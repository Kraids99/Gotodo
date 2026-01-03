package routes

// Huruf BESAR = public (exported)
// Huruf kecil = private (internal package)

// InitRoutes --> bisa dipanggil dari main.go
// initRoutes --> TIDAK bisa dipanggil dari main.go

import (
	"github.com/gin-gonic/gin"
	"gotodo/controllers"
)

func InitRoutes(r *gin.Engine) {

	// test server
	r.GET("/ping", controllers.Ping)

	// create todo
	r.POST("/gotodo", controllers.CreateGotodo)

	// read todo
	r.GET("/gotodo", controllers.GetGotodo)

	// update todo
	r.PUT("/gotodo/:id", controllers.UpdateGotodo)

	r.DELETE("/gotodo/:id", controllers.DeleteGotodo)
}

// func InitRoutes(r *gin.Engine) {
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// }
