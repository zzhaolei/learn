package main

import (
	"database/sql"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Code      string
	Price     uint
	CreatedAt sql.NullString
}

func main() {
}
