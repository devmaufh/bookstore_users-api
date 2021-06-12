package app

import "github.com/devmaufh/bookstore_users-api/controllers/user"

func mapUrls() {
	router.GET("/user", user.GetUsers)
	router.GET("/user/:user_id", user.GetUser)
	router.GET("/user/search", user.SearchUser)
	router.POST("/user", user.CreateUser)
	router.DELETE("/user/:user_id", user.DeleteUser)
}
