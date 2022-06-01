package models

import "project/freelance/danspro/simpleProject/config/db"

type User struct {
	Id       int
	Name     string
	Password string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUser(username, password string) User {
	var (
		user User
	)

	tx := db.MysqlConn().Begin()

	tx.Where("name = ? AND password = ?", username, password).Find(&user)

	tx.Commit()

	return user
}
