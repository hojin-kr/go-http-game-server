package model

import (
	ds "github.com/hojin-kr/go-http-game-server/cmd/ds"
	_ "github.com/lib/pq"
)

type Profile struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
}

// Get
func Get(userId int) Profile {
	ds := ds.GetClient()
	profile := Profile{}
	err := ds.QueryRow("SELECT user_id, nickname FROM switch.profile WHERE user_id = $1", userId).Scan(&profile.UserId, &profile.Nickname)

	if err != nil {
		// insert
		profile = Insert(userId, "")
	}
	return profile
}

// Update
func Update(userId int, nickname string) Profile {
	ds := ds.GetClient()
	profile := Profile{}
	err := ds.QueryRow("UPDATE switch.profile SET nickname = $1 WHERE user_id = $2 RETURNING user_id, nickname", nickname, userId).Scan(&profile.UserId, &profile.Nickname)

	if err != nil {
		// insert
		profile = Insert(userId, nickname)
	}
	return profile
}

// Insert
func Insert(userId int, nickname string) Profile {
	ds := ds.GetClient()
	profile := Profile{}
	err := ds.QueryRow("INSERT INTO switch.profile (user_id, nickname) VALUES ($1, $2) RETURNING user_id, nickname", userId, nickname).Scan(&profile.UserId, &profile.Nickname)

	if err != nil {
		panic(err)
	}
	return profile
}

// Check nickname exists
func CheckNicknameExists(nickname string) bool {
	ds := ds.GetClient()
	var exists bool
	err := ds.QueryRow("SELECT EXISTS(SELECT 1 FROM switch.profile WHERE nickname = $1)", nickname).Scan(&exists)

	if err != nil {
		panic(err)
	}
	return exists
}
