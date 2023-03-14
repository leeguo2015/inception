package global

import (
	"inception/api/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	// 跨域配置
	Cors CORS // `mapstructure:"cors" json:"cors" yaml:"cors"`
)

const (
	//ConfigFilePath = "configs/api.yaml"
	ConfigFilePath = "configs/inceptionApi.yaml"
)

func Init(c *config.Config) error {
	var err error
	Log = ParesLog(&c.Logger)
	DB, err = ParesGormMysql(&c.Mysql)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig() *config.Config {
	Config = config.New()
	if err := Config.Load(ConfigFilePath); err != nil {
		logrus.Panic(err)
	}
	logrus.Printf("config path:\t%s", ConfigFilePath)
	logrus.Printf("config content:\t%#v", *Config)
	if err := Init(Config); err != nil {
		logrus.Panic(err)
	}
	return Config
}
