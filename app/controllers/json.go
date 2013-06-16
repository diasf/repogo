package controllers

import (
	"github.com/robfig/revel"
	"github.com/diasf/repogo/app/models"
	"fmt"
)

type JSONController struct {
	*revel.Controller
}

func (c JSONController) RenderJsonError(status int, messageKey string, args ...interface{}) revel.Result {
	c.Response.Status = status
	return c.RenderJson(models.BuildError(fmt.Sprintf("%v", status), c.Message(messageKey, args...)))
}
