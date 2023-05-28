package models

import "github.com/go-sql-driver/mysql"

type User struct {
	Id        int            `json:"id"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	CreatedAt mysql.NullTime `json:"created_at"`
	Name      string         `json:"name"`
}
