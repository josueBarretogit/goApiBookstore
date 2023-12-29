package main

import (
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/initializers"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Id: "1", Item: "clean room", Completed: false},
	{Id: "1", Item: "clean dirty", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var todo Todo
	if err := context.BindJSON(&todo); err != nil {
		return
	}
	todos = append(todos, todo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func toggleTodoEstado(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	//Db := initializers.ConnectToDB()

	r := gin.Default()

	r.POST("/post/create", controllers.PostCreate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
