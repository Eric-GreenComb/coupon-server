package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/Eric-GreenComb/coupon-server/bean"
)

// MariaDB 数据库相关配置
var MariaDB bean.DBConfig

// Server Server Config
var Server bean.ServerConfig

// JWT Server Config
var JWT bean.JWTConfig

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	Server.Port = strings.Split(viper.GetString("server.port"), ",")
	Server.Mode = viper.GetString("server.mode")
	Server.GormLogMode = viper.GetString("server.gorm.LogMode")
	Server.ViewLimit = viper.GetInt("server.view.limit")
	Server.Passphrase = viper.GetString("server.passphrase")
	Server.Key = viper.GetString("server.key")

	MariaDB.Dialect = viper.GetString("database.dialect")
	MariaDB.Database = viper.GetString("database.database")
	MariaDB.User = viper.GetString("database.user")
	MariaDB.Password = viper.GetString("database.password")
	MariaDB.Host = viper.GetString("database.host")
	MariaDB.Port = viper.GetInt("database.port")
	MariaDB.Charset = viper.GetString("database.charset")
	MariaDB.MaxIdleConns = viper.GetInt("database.maxIdleConns")
	MariaDB.MaxOpenConns = viper.GetInt("database.maxOpenConns")
	MariaDB.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		MariaDB.User, MariaDB.Password, MariaDB.Host, MariaDB.Port, MariaDB.Database, MariaDB.Charset)

	JWT.SigningKey = viper.GetString("jwt.signing_key")
	JWT.ExpireTime = time.Minute * time.Duration(viper.GetInt("jwt.expire"))
	JWT.Realm = viper.GetString("jwt.realm")
}
