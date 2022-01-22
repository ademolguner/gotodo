package main

import (
	"github.com/ademolguner/gotodo/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Echo framework için yeni bir instance oluşturulur
	e := echo.New()

	// Middleware tanımlamaları yapılır
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoint'lerimiz oluşturulur.
	e.GET("/", services.Hello)

	e.GET("/todos", services.AllTodos)
	e.POST("/todos", services.CreateTodo)

	e.GET("/todos/:id", services.GetTodo)
	e.PUT("/todos/:id/complete", services.UpdeteTodoIsComplete)
	e.PUT("/todos/:id/uncomplete", services.UpdeteTodoIsUncomplete)
	e.DELETE("/todos/:id", services.DeleteTodo)

	// 3200 portundan API'ı ayağa kaldıralım
	e.Logger.Fatal(e.Start(":3200"))
}

func main() {

}
