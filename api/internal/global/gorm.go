package global

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Dialect      string
	URL          string
	DB           *sql.DB
	DBName       string
	User         string
	Password     string
	Host         string
	Port         int
	MaxIdleConns int
	MaxOpenConns int
}

func NewConfig(dialect, url, dbname, user, password, host string, port, maxidle, maxopen int) *Mysql {
	return &Mysql{
		Dialect:      dialect,
		URL:          url,
		DBName:       dbname,
		User:         user,
		Password:     password,
		Host:         host,
		Port:         port,
		MaxIdleConns: maxidle,
		MaxOpenConns: maxopen,
	}
}

// func ConnectMySQL(cfg MySQLConfig) (*gorm.DB, error) {
// 	dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func GormMysql() *gorm.DB {
	m := Config.Mysql
	if m.DB == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
