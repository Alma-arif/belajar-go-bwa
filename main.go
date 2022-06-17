package main

import (
	"log"
	"membuat-api-bwa/handler"
	"membuat-api-bwa/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/ayo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	// input := user.LoginInput{
	// 	Email:    "kapan@kapan.com",
	// 	Password: "haidiasd",
	// }

	// user, err := userService.Login(input)

	// if err != nil {
	// 	fmt.Println("terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// userbyemail, err := userRepository.FindByEmail("gin@mail.com")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(userbyemail.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1/")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
