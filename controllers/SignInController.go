package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jwkim1993/bdgg-server/dto"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

type SignInController struct {
}

func (s SignInController) Init(g *echo.Group) {
	g.POST("", s.GenerateJwtTokenUsingOAuthToken)
}

func (SignInController) GenerateJwtTokenUsingOAuthToken(c echo.Context) error {
	msg := &dto.OAuthDTO{}
	if err := c.Bind(msg); err != nil {
		return err
	}
	log.Printf("received token: %s", msg.AccessToken)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jaewon Kim"
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &dto.OAuthDTO{AccessToken: t})
}
