package controllers

import (
	_ "models"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {

	c.DoInit()

	c.TplName = "index.html"
}
