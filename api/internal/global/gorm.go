package global

import (
	"errors"
	"inception/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ParesGormMysql(m *config.Mysql) (*gorm.DB, error) {
	if m.DB == "" {
		return nil, errors.New("database configuration is required")
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConn)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		return db, nil
	}
}
