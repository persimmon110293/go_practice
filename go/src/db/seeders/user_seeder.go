package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	name       string
	age        int
	email      string
	created_at time.Time
	updated_at time.Time
}

func seeds(db *gorm.DB) error {

	users := Users{name: "taro", age: 20, email: "taro@taro.com", created_at: time.Now(), updated_at: time.Now()}
	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	users2 := Users{name: "jiro", age: 20, email: "jiro@jiro.com", created_at: time.Now(), updated_at: time.Now()}
	if err := db.Create(&users2).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	users3 := Users{name: "saburo", age: 20, email: "saburo@saburo.com", created_at: time.Now(), updated_at: time.Now()}
	if err := db.Create(&users3).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}

func openConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(mysql:3306)/practice?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func main() {
	db := openConnection()
	defer db.Close()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
