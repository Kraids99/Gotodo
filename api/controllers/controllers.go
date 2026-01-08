package controllers

import (
	"gotodo/database"
	"gotodo/models"
	"strings"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

var todos = []models.Gotodo{}
var nextID = 1

// c = nama variabel
// *gin.Context = tipe data
// := buat variabel + assign

// Create
func CreateGotodo(c *gin.Context) {
	var newTodo models.Gotodo

	// ambil body JSON → masuk ke struct todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// simpan ke database
	if err := database.DB.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create todo",
		})
		return
	}

	// todos = append(todos, newTodo)

	// kirim balik data
	c.JSON(http.StatusCreated, newTodo)
}

// Read
func GetGotodo(c *gin.Context) {
	// c.JSON(http.StatusOK, todos)

	var todo []models.Gotodo

	// ambil semua data dari table todos
	if err := database.DB.Find(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch todos",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Update
func UpdateGotodo(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var todo models.Gotodo

	// cari todo berdasarkan ID
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	var updateData struct {
		Title string `json:"title"`
		Completed bool `json:"completed"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if strings.TrimSpace(updateData.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title cannot be empty",
		})
		return
	}

	// update field
	if updateData.Title != "" {
		todo.Title = updateData.Title
	}
	todo.Completed = updateData.Completed

	database.DB.Save(&todo)

	c.JSON(http.StatusOK, todo)

	// for i, todo := range todos {
	// 	if todo.ID == id {
	// 		todos[i].Title = updateData.Title
	// 		todos[i].Completed = updateData.Completed
	// 		c.JSON(http.StatusOK, todos[i])
	// 		return
	// 	}
	// }
}

// delete
func DeleteGotodo(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	// for i, todo := range todos {
	// 	if todo.ID == id {
	// 		todos = append(todos[:i], todos[i+1:]...)

	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "todo deleted",
	// 		})
	// 		return
	// 	}
	// }

	if err := database.DB.Delete(&models.Gotodo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})

}
