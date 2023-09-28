package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&Person{}); err != nil {
		log.Fatal(err)
	}

	// 创建
	db.Create(&Person{Code: "abc", Price: 8})

	var p Person
	db.First(&p, "code = ?", "abc")

	db.Model(&p).Update("Price", 18)
	// db.Delete(&p, p.ID)
	fmt.Println(p)
}
