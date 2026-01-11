package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
	ID         int64
	Username   string
	Email      string
	Created_at string
	Money      int64
}

var users []User

func Init_db(dsn_var string) {
	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", dsn_var)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("[X] DB 연결 실패..." + err.Error())
	}

	log.Println("[!] DB 연결 성공")
}

func Get_user(user_id string) User {
	var new_user User

	err := DB.QueryRow(
		"SELECT id, username, email, created_at, money FROM users WHERE userid = ?",
		user_id,
	).Scan(&new_user.ID, &new_user.Username, &new_user.Email, &new_user.Created_at, &new_user.Money)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("유저 없음")
			return new_user
		}
		log.Fatal(err)
	}

	return new_user
}
