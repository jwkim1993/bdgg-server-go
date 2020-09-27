package controllers

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/jwkim1993/bdgg-server/dto"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type FCMController struct {
	app *firebase.App
}

func (f FCMController) Init(g *echo.Group, app *firebase.App) {
	f.app = app
	g.GET("", f.sendMessage)
}

func (f FCMController) sendMessage(c echo.Context) error {
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := f.app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	//registrationToken := "fNScPvO1RDuyJJJkBLtqSI:APA91bGl8_mf7aquheUnOftMjLqsFgAb2XVre2weHq-qDvPifKpfpbP8mF6lQRYHneSdenbsHun_CCh1zTD8rxmFpmV-xPPpeU0QFyW07vj-a9WHx9UT-If_O7T2ngCTHoyJINecDKcF"
	registrationToken := "cvloJc1tTdu5NC2qXIVwXy:APA91bFw_xCFdAYt_1rO8dX70C9AE4Y95157A4nqJ0LN2n2SXvloKw0_l3y7rhPFvCvqKwsg5tPfoHiYROUmD25c_LjgbWN2-ks8f2cUQncKrF1m2-t6DStkmRwofZOzU4ehxzoOrfm0"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "hello",
			Body:  "World",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	retObj := dto.SampleDTO{Message: fmt.Sprintf("Successfully sent message: %s", response)}

	return c.JSON(http.StatusOK, retObj)
}
