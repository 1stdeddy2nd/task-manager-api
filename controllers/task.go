package controllers

import (
	"net/http"
	"task-manager/database"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	userID := c.MustGet("userID").(uint)

	database.DB.Where("user_id = ?", userID).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	var task models.Task
	userID := c.MustGet("userID").(uint)

	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var input struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.MustGet("userID").(uint)

	task := models.Task{
		Title:  input.Title,
		Status: "pending",
		UserID: userID,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	userID := c.MustGet("userID").(uint)

	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input struct {
		Title  *string `json:"title"`
		Status *string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Title != nil {
		task.Title = *input.Title
	}

	if input.Status != nil {
		if *input.Status != "pending" && *input.Status != "done" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be 'pending' or 'done'"})
			return
		}
		task.Status = *input.Status
	}

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var task models.Task

	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
