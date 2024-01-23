package starter

import (
	"database/sql"
	"pandax/kit/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormlog "gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

var Db *gorm.DB

type DbGorm struct {
	Type         string
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
}

func (dg *DbGorm) GormInit() *gorm.DB {
	switch dg.Type {
	case "mysql":
		Db = dg.GormMysql()
	case "postgresql":
		Db = dg.GormPostgresql()
	}
	return Db
}
func (dg *DbGorm) GormMysql() *gorm.DB {

	logger.Log.Infof("连接mysql [%s]", dg.Dsn)
	mysqlConfig := mysql.Config{
		DSN:                       dg.Dsn, // DSN data source name
		DefaultStringSize:         191,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,  // 根据版本自动配置
	}
	ormConfig := &gorm.Config{Logger: gormlog.Default.LogMode(gormlog.Silent)}
	db, err := gorm.Open(mysql.New(mysqlConfig), ormConfig)
	if err != nil {
		logger.Log.Panicf("连接mysql失败! [%s]", err.Error())
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dg.MaxOpenConns)
	return db
}

func (dg *DbGorm) GormPostgresql() *gorm.DB {

	db, err := sql.Open("postgres", dg.Dsn)
	if err != nil {
		logger.Log.Panicf("连接postgresql失败! [%s]", err.Error())
	}
	ormConfig := &gorm.Config{}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), ormConfig)
	if err != nil {
		logger.Log.Panicf("连接postgresql失败! [%s]", err.Error())
	}
	sqlDB, err := gormDb.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(dg.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(dg.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return gormDb
}
