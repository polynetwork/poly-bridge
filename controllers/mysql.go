package controllers

import (
	"github.com/astaxie/beego"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func newDB() *gorm.DB{
	user := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	scheme := beego.AppConfig.String("mysqldb")
	db, err := gorm.Open(mysql.Open(user + ":" +password + "@tcp(" + url + ")/" + scheme + "?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
