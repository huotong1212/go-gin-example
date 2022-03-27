package models

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

/**
初始化数据库
*/
func Setup() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbSettings := setting.DatabaseSetting
	dbType = dbSettings.Type
	dbName = dbSettings.Name
	user = dbSettings.User
	password = dbSettings.Password
	host = dbSettings.Host
	tablePrefix = dbSettings.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

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

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallBack)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createdTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createdTimeField.IsBlank {
				createdTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 根据入参获取设置了字面值的参数，例如本文中是 gorm:update_column ，它会去查找含这个字面值的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// 设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// scope.Get("gorm:delete_option") 检查是否手动指定了 delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		// 获取我们约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(), // 返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),                 // scope.AddToVars 该方法可以添加值作为 SQL 的参数，也可用于防范 SQL 注入
				addExtraSpaceIfExist(scope.CombinedConditionSql()), //返回组合好的条件 SQL，看一下方法原型很明了
				addExtraSpaceIfExist(extraOption),
			)).Exec()
			fmt.Printf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(), // 返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),                 // scope.AddToVars 该方法可以添加值作为 SQL 的参数，也可用于防范 SQL 注入
				addExtraSpaceIfExist(scope.CombinedConditionSql()), //返回组合好的条件 SQL，看一下方法原型很明了
				addExtraSpaceIfExist(extraOption),
			)
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}

	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
