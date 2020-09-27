package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/jwkim1993/bdgg-server/controllers"
	"github.com/jwkim1993/bdgg-server/db"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
	"google.golang.org/api/option"
	"log"
)

const (
	ServerPort string = ":8080"
)

func main() {
	// db initialize
	if err := db.InitDB(); err != nil {
		fmt.Println("error has occured")
		return
	}

	//firebase initialize
	opt := option.WithCredentialsFile("/Users/jaewonkim/Project/bdgg-server/src/main/resources/keys/bdgg-firebase-adminsdk.json")
	config := &firebase.Config{ProjectID: "bdgg-d3822"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	e := echo.New()

	//e.Use(common.IsLoggedIn)

	//controllers.SampleController{}.Init(e.Group("/"))
	controllers.FCMController{}.Init(e.Group("/send"), app)
	controllers.SignInController{}.Init(e.Group("/signup"))
	controllers.UserController{}.Init(e.Group("/users"))

	if err := e.Start(ServerPort); err != nil {
		log.Println(err)
	}
}
