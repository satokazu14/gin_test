package db

import (
	"gin_test/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	DBMS := "mysql"
	USER := "admin"
	PASS := ""
	//local環境のmysqlに接続するには「tcp(127.0.0.1:3306)」
	//ec2環境のmysqlに接続するには「tcp(database-1.cenyz2blrddy.ap-northeast-1.rds.amazonaws.com:3306)」
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "hal_motor"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	//autoMigtation()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigtation() {
	db.AutoMigrate(&entity.User{})
}
