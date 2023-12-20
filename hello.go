package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	id        string `json:"id"`
	item      string `json:"item"`
	completed bool   `json:"completed"`
}

var todos = []todo{
	{id: "1", item: "clean room", completed: false},
	{id: "1", item: "clean dirty", completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)

	router.Run("localhost:8080")
}
