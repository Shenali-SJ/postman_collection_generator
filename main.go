package main

import (
	"automaticPostmanCollection/automate"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name string `json:"name" xml:"name"`
	Age int `json:"age" xml:"age"`
}

var port string

const getHomeFileName = "basePath.json"
const getUserFileName = "user.json"
const createUserFileName = "createUser.json"

func main() {
	e := echo.New()

	e.GET("/", getHome)
	e.GET("/user", getUser)
	e.POST("/user/create", createUser)
	e.Logger.Fatal(e.Start(":1323"))

	port = e.Listener.Addr().String()
}


func getHome(c echo.Context) error {
	// Generates postman collection based on the specific end point.
	automate.CreateCollection(c, port, getHomeFileName)
	return c.String(http.StatusOK, "Welcome")
}

func getUser(c echo.Context) error {
	name := c.QueryParam("name")

	// Generates postman collection based on the specific end point.
	automate.CreateCollection(c, port, getUserFileName)

	return c.String(http.StatusOK, "Hi " + name)
}

func createUser(c echo.Context) error {
	user := new(User)

	if err := c.Bind(user); err != nil {
		fmt.Println(err)
		return err   //TODO: http error - 400
	}

	// Generates postman collection based on the specific end point.
	automate.CreateCollection(c, port, createUserFileName, user)

	return c.JSON(http.StatusOK, user)
}