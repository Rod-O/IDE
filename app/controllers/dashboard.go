package controllers

import "github.com/revel/revel"

type Dashboard struct {
	*revel.Controller
}

func (c Dashboard) Index() revel.Result {
 	return	c.Render()
}

func (c Dashboard) FileManager() revel.Result {
	return c.Render()
}
