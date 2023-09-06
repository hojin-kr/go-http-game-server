package model

import (
	ds "github.com/hojin-kr/go-http-game-server/cmd/ds"
	_ "github.com/lib/pq"
)

type Etc struct {
	UserID int    `json:"user_id"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

// Get
func Get(userId int, key string) Etc {
	ds := ds.GetClient()
	etc := Etc{}
	err := ds.QueryRow("SELECT user_id, key, value FROM switch.etc WHERE user_id = $1 AND key = $2", userId, key).Scan(&etc.UserID, &etc.Key, &etc.Value)

	if err != nil {
		// insert
		etc = Insert(userId, key, "")
	}

	return etc
}

// GetAll
func GetAll(userId int) []Etc {
	ds := ds.GetClient()
	etc := Etc{}
	etcs := []Etc{}
	rows, err := ds.Query("SELECT user_id, key, value FROM switch.etc WHERE user_id = $1", userId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&etc.UserID, &etc.Key, &etc.Value)
		if err != nil {
			panic(err)
		}
		etcs = append(etcs, etc)
	}

	return etcs
}

// Update
func Update(userId int, key string, value string) Etc {
	ds := ds.GetClient()
	etc := Etc{}

	err := ds.QueryRow("UPDATE switch.etc SET value = $1 WHERE user_id = $2 AND key = $3 RETURNING user_id, key, value", value, userId, key).Scan(&etc.UserID, &etc.Key, &etc.Value)

	if err != nil {
		// insert
		etc = Insert(userId, key, value)
	}

	return etc
}

// Insert
func Insert(userId int, key string, value string) Etc {
	ds := ds.GetClient()
	etc := Etc{}

	err := ds.QueryRow("INSERT INTO switch.etc (user_id, key, value) VALUES ($1, $2, $3) RETURNING user_id, key, value", userId, key, value).Scan(&etc.UserID, &etc.Key, &etc.Value)

	if err != nil {
		panic(err)
	}

	return etc
}
