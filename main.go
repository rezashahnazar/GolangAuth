package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rezashahnazar/GolangAuth/controllers"
	"github.com/rezashahnazar/GolangAuth/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.Run()

}
