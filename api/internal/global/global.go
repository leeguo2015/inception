package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"inception/api/config"
	"time"
)

var (
	Config    *config.Config
	DB        *gorm.DB
	REDIS     *redis.Client
	Log       *logrus.Logger
	Cors      CORS // `mapstructure:"cors" json:"cors" yaml:"cors"`
	Captcha   *captcha
	JwtTimout = 30 * 24 * time.Hour
)

const (
	//ConfigFilePath = "configs/api.yaml"
	ConfigFilePath = "configs/inceptionApi.yaml"
)

func Parse(c *config.Config) error {
	var err error
	logrus.Printf("初始化日志库")
	Log = ParesLog(&c.Logger)
	Log.Info("初始化 mysql")
	DB, err = ParesGormMysql(&c.Mysql)
	if err != nil {
		return err
	}
	Log.Info("初始化 redis")
	REDIS, err = ParseRedis(&c.Redis)
	if err != nil {
		return err
	}
	Captcha = new(captcha)
	if err = ParseCaptcha(Captcha); err != nil {
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
	if err := Parse(Config); err != nil {
		logrus.Panic(err)
	}
	return Config
}
