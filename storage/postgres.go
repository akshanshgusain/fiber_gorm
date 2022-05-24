package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := "host=localhost user=hello_fastapi password=hello_fastapi dbname=hello_fastapi_dev port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := fmt.Sprintf("host:%s port:%s user:%s password:%s dbname=%s sslmode:%s",
	//	config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, err
}
