package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	Name  string
	Age   int
	Email string
}

func seeds(db *gorm.DB) error {

	users := Users{Name: "taro", Age: 20, Email: "taro@taro.com"}
	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	users2 := Users{Name: "jiro", Age: 20, Email: "jiro@jiro.com"}
	if err := db.Create(&users2).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	users3 := Users{Name: "saburo", Age: 20, Email: "saburo@saburo.com"}
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
