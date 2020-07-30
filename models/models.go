package models

import (
	"fmt"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var (
	db  *gorm.DB
	wdb *gorm.DB
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
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
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// write db
	wdb = db
}

func CloseDB() {
	defer db.Close()
}

func CloseWriteDB() {
	defer wdb.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifiedTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifiedTimeField.IsBlank {
				_ = modifiedTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	var extraOption string
	if str, ok := scope.Get("gorm:delete_option"); ok {
		extraOption = fmt.Sprint(str)
	}

	deleteOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
	if !scope.Search.Unscoped && hasDeletedOnField {
		scope.Raw(fmt.Sprintf(
			"UPDATE %v SET %v=%v%v%v",
			scope.QuotedTableName(),
			scope.Quote(deleteOnField.DBName),
			scope.AddToVars(time.Now().Unix()),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
			)).Exec()
	} else {
		scope.Raw(fmt.Sprintf(
			"DELETE FROM %v%v%v",
			scope.QuotedTableName(),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}