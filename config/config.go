package config

import (
	"fmt"
	"gin_exercise/controller"
	"gin_exercise/help"
	"gin_exercise/mydb"

	"github.com/cihub/seelog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GlobalConfig *Config

type Config struct {
	ListenPort int
	User       string
	Password   string
	Host       string
	Port       string
}

func init() {
	GlobalConfig = GetConfig()
}

func GetConfig() *Config {
	return &Config{
		ListenPort: 8080,
		//ssh服务器用户名，密码，ip地址，端口号
		User:     "ycx",
		Password: "20231105",
		Host:     "211.71.76.205",
		Port:     "22",
	}
}

func Initlog() {
	help.SetupLogger()
	seelog.Info(fmt.Sprintf("Begin Seelog"))
}

func InitUserdatabase() {
	//数据库账号root 密码123456 地址127.0.0.1:3306
	dsn := "root:123456@tcp(127.0.0.1:3306)/userInfo?charset=utf8&parseTime=True&loc=Local"
	userDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		seelog.Error(err)
	} else {
		mydb.UserDB = userDB
	}
	p := mydb.UserDB.Migrator().HasTable(&controller.User{}) //检测是否存在users表单
	if !p {
		seelog.Info(fmt.Sprintf("create table users"))
		mydb.UserDB.Migrator().CreateTable(&controller.User{}) //不存在则创建
	} else {
		seelog.Info(fmt.Sprintf("table users already exists"))
	}
}
