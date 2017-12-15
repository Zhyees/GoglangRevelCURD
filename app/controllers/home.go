package controllers

import "github.com/revel/revel"

type HomeCtrl struct {
	BaseCtrl
}

func (h HomeCtrl) Index() revel.Result{
	return h.Render()
}
