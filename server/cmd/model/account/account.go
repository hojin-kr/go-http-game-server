package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	ds "github.com/hojin-kr/go-http-game-server/cmd/ds"
	utils "github.com/hojin-kr/go-http-game-server/cmd/utils"
	_ "github.com/lib/pq"
)

type Account struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func Init() Account {
	ds := ds.GetClient()
	uuid := uuid.New()
	// timestamp to string
	timestamp := time.Now().Format("20060102150405")
	account := Account{
		Id:    uuid.String(),
		Token: "token" + timestamp,
	}
	// insert
	_, err := ds.Exec("INSERT INTO account (id, token) VALUES ($1, $2)", account.Id, account.Token)
	if err != nil {
		panic(err)
	}
	row := ds.QueryRow("SELECT id, token FROM account WHERE token = $1", account.Token)
	err = row.Scan(&account.Id, &account.Token)
	if err != nil {
		panic(err)
	}
	log.Println("user created" + account.Id)
	return account
}

// GetByUUID if not exist create
func GetByUUID(uuid string) Account {
	ds := ds.GetClient()
	var account Account
	// check uuid format
	if !utils.IsValidatedUUID(uuid) {
		account = Init()
		return account
	}
	row := ds.QueryRow("SELECT id, token FROM account WHERE id = $1", uuid)
	err := row.Scan(&account.Id, &account.Token)
	if err != nil {
		account = Init()
	}
	return account
}
