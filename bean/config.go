package bean

import (
	"time"
)

// ServerConfig ServerConfig Struct
type ServerConfig struct {
	Port        []string
	Mode        string
	GormLogMode string
	ViewLimit   int
	Passphrase  string
	Key         string
}

// DBConfig DBConfig Struct
type DBConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}

// JWTConfig JWTConfig
type JWTConfig struct {
	SigningKey string
	ExpireTime time.Duration
	Realm      string
}
