package main

import (
	"douying/config"
	"douying/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := config.DBConnectString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect mysql!")
	}

	db.AutoMigrate(&model.User{})

	db.Create(&model.User{Name: "xgf", Password: "1234"})
	fmt.Println("ok")

	/* var user model.User
	db.First(&user, 1)
	if user.Name != "xgf" {
		fmt.Println("Fail")
	}
	db.Model(&user).Update("password", "4567")
	if user.Password == "1234" {
		fmt.Println("Fail")
	} */
}
