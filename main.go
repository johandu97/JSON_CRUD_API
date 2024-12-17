package main

import (
	"JSON_CRUD_API/controllers"
	"JSON_CRUD_API/initializers"

	_ "JSON_CRUD_API/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

// @title CRUD POST API
// @version 1.0.0
// @description This is a sample server for managing blog posts.
// @termsOfService http://swagger.io/terms/
// @contact.name Zohan
// @contact.email johandu97@gmail.com
// @host localhost:3000
// @BasePath /
func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", controllers.Ping)
	r.POST("/createtable", controllers.CreateTable)
	r.POST("/createpost", controllers.CreatePost)
	r.GET("/viewpost/:ID", controllers.ViewPost)
	r.GET("/viewposts", controllers.ViewPosts)
	r.PUT("/updatepost/:ID", controllers.UpdatePost)
	r.DELETE("/deletepost/:ID", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
