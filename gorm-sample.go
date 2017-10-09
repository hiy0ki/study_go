package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id   int64
	Name string `sql:"size:255"`
}

func main() {
	db, err := gorm.Open("mysql", "user:password@/database?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	user := User{
		Name: "hiyoki",
	}

	db.Create(&user)

	fmt.Println(user.Id)
	fmt.Println(db.HasTable(&User{}))
	fmt.Println(db.HasTable("users"))
}
