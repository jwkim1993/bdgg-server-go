package controllers

import (
	"fmt"
	"github.com/jwkim1993/bdgg-server/common"
	"github.com/jwkim1993/bdgg-server/dto"
	"github.com/labstack/echo"
	"net/http"
)

type SampleController struct {
}

func (c SampleController) Init(g *echo.Group) {
	g.GET("", c.GetHelloMessage)
	g.GET("param/:id", c.GetHelloMessageWithPathParam, common.IsLoggedIn, common.IsAdmin)
	g.GET("param", c.GetHelloMessageWithQueryParam, common.IsLoggedIn)
	g.POST("", c.PostMessage)
}

func (SampleController) GetHelloMessage(c echo.Context) error {
	retObj := &dto.SampleDTO{Message: "hello world"}

	return c.JSON(http.StatusOK, retObj)
}

func (SampleController) GetHelloMessageWithQueryParam(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")

	retObj := &dto.SampleDTO{Message: fmt.Sprintf("name: %s, age: %s", name, age)}

	return c.JSON(http.StatusOK, retObj)
}

func (SampleController) GetHelloMessageWithPathParam(c echo.Context) error {
	id := c.Param("id")

	retObj := &dto.SampleDTO{Message: fmt.Sprintf("id: %s", id)}

	return c.JSON(http.StatusOK, retObj)
}

func (SampleController) PostMessage(c echo.Context) error {
	msg := &dto.SampleDTO{}
	if err := c.Bind(msg); err != nil {
		return err
	}

	retObj := &dto.SampleDTO{Message: fmt.Sprintf("message: %s", msg.Message)}

	return c.JSON(http.StatusOK, retObj)
}
