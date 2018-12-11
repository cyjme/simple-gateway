package app

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	mysqlArgs := Config.DB.User + ":" + Config.DB.Password + "@tcp(" + Config.DB.Host + ":" + Config.DB.Port + ")/" + Config.DB.Name + "?charset=utf8&parseTime=True&loc=Local"
	//Db, err = gorm.Open("mysql", "root:123456789@/chat?charset=utf8&parseTime=True&loc=Local")
	log.Println("mysqlArgs", mysqlArgs)
	DB, err = gorm.Open("mysql", mysqlArgs)

	// 全局禁用表名复数
	DB.SingularTable(true)
	DB.DB().SetConnMaxLifetime(60 * time.Second)

	//defer DB.Close()

	if err != nil {
		panic(err)
		//Db.DropTableIfExists(
		//	&model.User{})
		//
	}

	// DB.CreateTable(
	// 	&model.Route{},
	// 	&model.RouteGroup{},
	// )

	fmt.Println("DB connect success!!!", mysqlArgs)
	DB.LogMode(true)
}
