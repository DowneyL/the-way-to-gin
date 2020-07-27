package models

import (
	"fmt"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db  *gorm.DB
	wdb *gorm.DB
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		section                                           *ini.Section
		dbType, dbName, user, password, host, tablePrefix string
	)
	section, err = setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database':%v", err)
	}
	dbType = section.Key("TYPE").MustString("mysql")
	dbName = section.Key("NAME").MustString("")
	user = section.Key("USER").MustString("root")
	password = section.Key("PASSWORD").MustString("")
	host = section.Key("HOST").MustString("127.0.0.1")
	tablePrefix = section.Key("TABLE_PREFIX").MustString("")

	info := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(dbType, info)
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// write db
	wdb = db
}

func CloseDB() {
	defer db.Close()
}

func CloseWriteDB() {
	defer wdb.Close()
}
