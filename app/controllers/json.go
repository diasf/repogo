package controllers

import (
	"fmt"
	"github.com/diasf/repogo/app/models"
	"github.com/robfig/revel"
)

type JSONController struct {
	*revel.Controller
}

func (c JSONController) RenderJsonError(status int, messageKey string, args ...interface{}) revel.Result {
	c.Response.Status = status
	return c.RenderJson(models.BuildError(fmt.Sprintf("%v", status), c.Message(messageKey, args...)))
}
