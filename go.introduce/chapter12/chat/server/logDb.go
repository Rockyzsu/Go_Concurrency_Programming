package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB
var dbErr error

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "go_chat"
	CHARSET   = "utf8"
)

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	//打开连接 如果写成MysqlDb, err := sql.Open("mysql", dbDSN)那么MysqlDb可能是局部的变量,全局MysqlDb是nil
	MysqlDb, dbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if dbErr != nil {
		fmt.Println("Mysql Open:", dbErr)
	}
	MysqlDb.SetMaxOpenConns(50)
	MysqlDb.SetMaxIdleConns(10)
	MysqlDb.SetConnMaxLifetime(30 * time.Second)
	//Ping确定连接是否可用
	if dbErr = MysqlDb.Ping(); dbErr != nil {
		panic("Mysql数据库连接失败: " + dbErr.Error())
	}
}

//LogToDb 插入数据到数据库
func LogToDb(msg string, address string) int64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("LogToDb:", err)
		}
	}()
	//可以进行增\删\改操作(query string, args ...interface{})
	//MysqlDb可能是nil
	ret, err := MysqlDb.Exec("insert into chat_logs (message,address) values(?,?)", msg, address)
	if err != nil {
		fmt.Println("LogToDb错误:", err)
		return -1
	}
	rows, _ := ret.RowsAffected()
	return rows
}
