package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// UsersRoutes
	users := e.Group("/users")
	users.GET("", getUsers)
	users.GET("/:id", getUser)
	users.POST("", createUser)
	e.Logger.Fatal(e.Start(":8081"))

}

type UserModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []UserModel{
	{ID: 1, Name: "Michael"},
	{ID: 2, Name: "Jack"},
	{ID: 3, Name: "Jhon"},
	{ID: 4, Name: "Mickey"},
	{ID: 5, Name: "Minny"},
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.String(http.StatusInternalServerError, "不正なユーザーIDです.")
	}

	for _, v := range users {
		if v.ID == userId {
			return c.JSON(http.StatusOK, v)
		}
	}
	return c.JSON(http.StatusNotFound, "対象のユーザーが見つかりません")
}

type UserRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func createUser(c echo.Context) error {
	var user UserRequest
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "不正なユーザー名です")
	}

	userModel := UserModel{ID: user.Id, Name: user.Name}
	users = append(users, userModel)

	return c.JSON(http.StatusCreated, users)
}
