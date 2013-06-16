package controllers

import (
	"github.com/robfig/revel"
	"github.com/diasf/repogo/app/models"
)

type User struct {
	JSONController
}

func (c User) GetAll() revel.Result {
	var users []models.User
	var err error
	if users, err = models.GetUserDAO().GetAll(); err != nil {
		return c.RenderJsonError(500, "server.error")
	}
	return c.RenderJson(users)
}

func (c User) GetUser(id int) revel.Result {
	var user models.User
	var err error
	if user, err = models.GetUserDAO().Get(id); err != nil {
		return c.RenderJsonError(404, "user.notfound")
	}
	return c.RenderJson(user)
}
