package model

import (
	"time"

	ds "github.com/hojin-kr/go-http-game-server/cmd/ds"
	_ "github.com/lib/pq"
)

type Social struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	TargetID int    `json:"target_id"`
	Type     string `json:"type"`
	Vars     string `json:"vars"`
	Created  int    `json:"created"`
	Updated  int    `json:"updated"`
}

type SocialRequest struct {
	UserID   int    `json:"user_id"`
	TargetID int    `json:"target_id"`
	Type     string `json:"type"`
	Vars     string `json:"vars"`
}

type SocialDeleteRequest struct {
	ID int `json:"id"`
}

// GetByTargetIDAndTypeLimitOffset
func GetByTargetIDAndTypeLimitOffset(targetId int, socialType string, limit int, offset int) []Social {
	ds := ds.GetClient()
	social := Social{}
	socials := []Social{}
	rows, err := ds.Query("SELECT id, user_id, target_id, type, vars, created, updated FROM switch.social WHERE target_id = $1 AND type = $2 LIMIT $3 OFFSET $4", targetId, socialType, limit, offset)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&social.ID, &social.UserID, &social.TargetID, &social.Type, &social.Vars, &social.Created, &social.Updated)
		if err != nil {
			panic(err)
		}
		socials = append(socials, social)
	}

	return socials
}

// GetByUserIDAndTypeLimitOffset
func GetByUserIDAndTypeLimitOffset(userId int, socialType string, limit int, offset int) []Social {
	ds := ds.GetClient()
	social := Social{}
	socials := []Social{}
	rows, err := ds.Query("SELECT id, user_id, target_id, type, vars, created, updated FROM switch.social WHERE user_id = $1 AND type = $2 LIMIT $3 OFFSET $4", userId, socialType, limit, offset)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&social.ID, &social.UserID, &social.TargetID, &social.Type, &social.Vars, &social.Created, &social.Updated)
		if err != nil {
			panic(err)
		}
		socials = append(socials, social)
	}

	return socials
}

// GetCountByTargetIDAndType
func GetCountByTargetIDAndType(targetId int, socialType string) int {
	ds := ds.GetClient()
	var count int
	err := ds.QueryRow("SELECT COUNT(*) FROM switch.social WHERE target_id = $1 AND type = $2", targetId, socialType).Scan(&count)

	if err != nil {
		panic(err)
	}

	return count
}

// Insert
func Insert(userId int, targetId int, socialType string, vars string) Social {
	ds := ds.GetClient()
	social := Social{}
	tm := time.Now().Unix()
	social.Created = int(tm)
	social.Updated = int(tm)
	err := ds.QueryRow("INSERT INTO switch.social (user_id, target_id, type, vars) VALUES ($1, $2, $3, $4) RETURNING id, user_id, target_id, type, vars, created, updated", userId, targetId, socialType, vars).Scan(&social.ID, &social.UserID, &social.TargetID, &social.Type, &social.Vars, &social.Created, &social.Updated)

	if err != nil {
		panic(err)
	}
	return social
}

// UpdateVarsByID
func UpdateVarsByID(id int, vars string) Social {
	ds := ds.GetClient()
	social := Social{}
	social.Updated = int(time.Now().Unix())
	err := ds.QueryRow("UPDATE switch.social SET vars = $1 WHERE id = $2 RETURNING id, user_id, target_id, type, vars, created, updated", vars, id).Scan(&social.ID, &social.UserID, &social.TargetID, &social.Type, &social.Vars, &social.Created, &social.Updated)

	if err != nil {
		panic(err)
	}
	return social
}

// DeleteByID
func DeleteByID(id int) {
	ds := ds.GetClient()
	_, err := ds.Exec("DELETE FROM switch.social WHERE id = $1", id)

	if err != nil {
		panic(err)
	}
}
