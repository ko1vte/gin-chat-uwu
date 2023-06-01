package dao

import (
	"database/sql"
	"fmt"
	"gin-chat-uwu/database"
	"gin-chat-uwu/models"
	"log"
)

//key：models.User的其中一个
func SelectUserByKEY(key string, value string) (*models.User, error) {
	db, err := database.InitMysqlDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	selectStr := fmt.Sprintf("SELECT * FROM users WHERE %s = ?", key)
	fmt.Println(selectStr)
	stmt, err := db.Prepare(selectStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var selectUser models.User
	row := stmt.QueryRow(value)
	err = row.Scan(&selectUser.Id, &selectUser.Username, &selectUser.Password, &selectUser.CreatedAt, &selectUser.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //用户不存在
		}
		return nil, err //查询错误
	}
	return &selectUser, nil //查询到了用户
}

//插入元素
func AddUser(user *models.User) error {
	db, err := database.InitMysqlDB()
	if err != nil {
		return err
	}
	insertStr := "INSERT INTO users (Username, Password, CreatedAt, Name) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(insertStr)

	if err != nil {
		return err
	}
	defer stmt.Close()

	//执行预处理的语句
	result, err := stmt.Exec(user.Username, user.Password, user.CreatedAt, user.Name)
	log.Println(result)
	if err != nil {
		return err
	}
	log.Printf("用户 %s 注册成功，注册时间：%s", user.Username, user.CreatedAt.Time)
	return nil
}

func DeleUser(username string) error {
	db, err := database.InitMysqlDB()
	if err != nil {
		return err
	}
	deleStr := "DELETE FROM users WHERE Username = ?"
	stmt, err := db.Prepare(deleStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(username)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
