package app

import "github.com/Aaron-Yu1/Go-Test/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
