package core

import (
	"blog/config"
	"blog/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// mysql初始化配置
var Db *gorm.DB

func MysqlInit() error {
	var err error
	var dbconfig = config.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbconfig.Username, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.Db, dbconfig.Charset)
	//连接数据库
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	if Db.Error != nil {
		return err
	}
	sqlDb, err := Db.DB()
	sqlDb.SetMaxIdleConns(dbconfig.MaxIdle)
	sqlDb.SetMaxOpenConns(dbconfig.MaxOpen)
	global.Log.Infof("[mysql数据库连接成功]")
	return nil
}
