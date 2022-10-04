package main

import (
	"Practical_1/initializers"

	"Practical_1/controllers"

	"github.com/gin-gonic/gin"
)
func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	
	r := gin.Default()
	r.GET("/posts", controllers.PeopleAll)
	r.GET("/posts/:id", controllers.PeopleIndex)

	r.POST("/posts", controllers.PostsCreate)

	r.PUT("/posts/:id", controllers.PeopleUpdate)

	r.DELETE("/posts/:id", controllers.PeopleDelete)
	r.Run()
}