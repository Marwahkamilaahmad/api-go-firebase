package handlers

import (
	"context"
	"net/http"
	"go-firebase/models"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func AddTodoHandler(c *gin.Context, client *firestore.Client) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	_, _, err := client.Collection("todos").Add(ctx, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Todo added successfully"})
}

func GetTodoHandler(c *gin.Context, client *firestore.Client) {
	id := c.Param("id")

	ctx := context.Background()
	doc, err := client.Collection("todos").Doc(id).Get(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var todo models.Todo
	doc.DataTo(&todo)
	c.JSON(http.StatusOK, todo)
}
