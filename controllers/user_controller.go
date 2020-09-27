package controllers

import (
	"github.com/jwkim1993/bdgg-server/db"
	"github.com/jwkim1993/bdgg-server/dto"
	"github.com/jwkim1993/bdgg-server/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type UserController struct {
}

func (u UserController) Init(e *echo.Group) {
	e.GET("", u.GetUserListByUserName)
}

func (u UserController) GetUserListByUserName(c echo.Context) error {
	username := c.QueryParam("username")
	log.Info("username: ", username)

	if len(username) < 2 {
		c.JSON(http.StatusBadRequest, dto.CustomError{
			Error:   "Invalid Parameter",
			Message: "Length must be at least 2",
		})
	}
	dbm := db.DBManager()

	var members []models.Member
	result := dbm.Where("name like ?", "%"+username+"%").Find(&members)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, dto.CustomError{
			Error:   "InterServerError",
			Message: result.Error.Error(),
		})
	}
	log.Info("result: ", result.RowsAffected)

	var response dto.MemberList
	response.SetMembers(members)

	return c.JSON(http.StatusOK, response)
}
