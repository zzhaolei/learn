package config

import (
	"gorm.io/gorm"
)

type Config struct {
	DB   *gorm.DB
	Salt string
}

func New(db *gorm.DB) Config {
	return Config{
		DB:   db,
		Salt: "salt", // must be use env
	}
}
