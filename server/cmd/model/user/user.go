package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	ds "github.com/hojin-kr/go-http-game-server/cmd/ds"
	utils "github.com/hojin-kr/go-http-game-server/cmd/utils"
	_ "github.com/lib/pq"
)

type User struct {
	Id      int    `json:"id"`
	UUID    string `json:"uuid"`
	Status  int    `json:"status"`
	Created int    `json:"created"`
	Updated int    `json:"updated"`
}

type Recovery struct {
	UserID  int    `json:"user_id"`
	Code    string `json:"code"`
	Expired int    `json:"expired"`
}

type RecoveryRequest struct {
	Code string `json:"code"`
}

// const user status
const (
	USER_STATUS_ACTIVE   = 1
	USER_STATUS_INACTIVE = 2
)

func Init() User {
	ds := ds.GetClient()
	uuid := uuid.New()
	timestamp := time.Now().Unix()
	user := User{
		UUID:    uuid.String(),
		Status:  1,
		Created: int(timestamp),
		Updated: int(timestamp),
	}
	// insert
	_, err := ds.Exec("INSERT INTO switch.user (uuid, status, created, updated) VALUES ($1, $2, $3, $4)", user.UUID, user.Status, user.Created, user.Updated)
	if err != nil {
		panic(err)
	}
	row := ds.QueryRow("SELECT id, uuid, status, created, updated FROM switch.user WHERE uuid = $1", user.UUID)
	err = row.Scan(&user.Id, &user.UUID, &user.Status, &user.Created, &user.Updated)
	if err != nil {
		panic(err)
	}
	log.Println("user created" + user.UUID)
	return user
}

// GetByUUID if not exist create
func GetByUUID(uuid string) User {
	ds := ds.GetClient()
	var user User
	// check uuid format
	if !utils.IsValidatedUUID(uuid) {
		user = Init()
		return user
	}
	row := ds.QueryRow("SELECT id, uuid, status, created, updated FROM switch.user WHERE uuid = $1", uuid)
	err := row.Scan(&user.Id, &user.UUID, &user.Status, &user.Created, &user.Updated)
	if err != nil {
		user = Init()
	}
	return user
}

// GetRecoveryCodeByUserID
func GetRecoveryCodeByUserID(userId int) Recovery {
	ds := ds.GetClient()
	var recovery Recovery
	row := ds.QueryRow("SELECT user_id, code, expired FROM switch.recovery WHERE user_id = $1", userId)
	err := row.Scan(&recovery.UserID, &recovery.Code, &recovery.Expired)
	if expired := recovery.Expired; expired < int(time.Now().Unix()) || err != nil {
		recovery = InsertRecoveryCodeByUserID(userId, utils.GenerateRandomString(6), int(time.Now().Unix())+120)
	}
	return recovery
}

// InsertRecoveryCodeByUserID
func InsertRecoveryCodeByUserID(userId int, code string, expired int) Recovery {
	ds := ds.GetClient()
	recovery := Recovery{}
	err := ds.QueryRow("INSERT INTO switch.recovery (user_id, code, expired) VALUES ($1, $2, $3) RETURNING user_id, code, expired", userId, code, expired).Scan(&recovery.UserID, &recovery.Code, &recovery.Expired)

	if err != nil {
		panic(err)
	}
	return recovery
}

// RecoveryUser
func RecoveryUser(code string) User {
	// code의 user_id를 찾아서 새로운 uuid로 변경하고 user를 리턴한다.
	ds := ds.GetClient()
	user := User{}
	updated := int(time.Now().Unix())
	// get user_id
	var userId int
	err := ds.QueryRow("SELECT user_id FROM switch.recovery WHERE code = $1", code).Scan(&userId)

	if err != nil {
		panic(err)
	}

	// update uuid
	uuid := uuid.New()
	err = ds.QueryRow("UPDATE switch.user SET uuid = $1, updated = $2 WHERE id = $3 RETURNING id, uuid, status, created, updated", uuid.String(), updated, userId).Scan(&user.Id, &user.UUID, &user.Status, &user.Created, &user.Updated)

	if err != nil {
		panic(err)
	}

	// delete recovery code
	DeleteRecoveryCodeByUserID(userId)

	return user
}

// UpdateUUID
func UpdateUUID(userId int, uuid string) User {
	ds := ds.GetClient()
	user := User{}
	updated := int(time.Now().Unix())
	log.Println(uuid)
	err := ds.QueryRow("UPDATE switch.user SET uuid = $1, updated = $2 WHERE id = $3 RETURNING id, uuid, status, created, updated", uuid, updated, userId).Scan(&user.Id, &user.UUID, &user.Status, &user.Created, &user.Updated)

	if err != nil {
		panic(err)
	}
	return user
}

// DeleteRecoveryCodeByUserID
func DeleteRecoveryCodeByUserID(userId int) {
	ds := ds.GetClient()
	_, err := ds.Exec("DELETE FROM switch.recovery WHERE user_id = $1", userId)
	if err != nil {
		panic(err)
	}
}
