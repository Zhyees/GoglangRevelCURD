package controllers

import (
	"encoding/json"
	"io"
	"github.com/revel/revel"
)

const(
	SuperAdminRoleId = 1
	OperationalAdminRoleId = 2
)

type BaseCtrlIFace interface {
	parseData(v interface{})
}

type BaseCtrl struct {
	GorpController
	Model DefaultModel
}

type DefaultModel struct {
	ViewCode string
	ViewTitle string
}

func (b BaseCtrl) GetDefaultModel() DefaultModel{
	return b.Model
}

func (b BaseCtrl) parseData(r io.Reader, v interface{}) (error) {
	err := json.NewDecoder(r).Decode(v)
	return err
}

func (b BaseCtrl) getConfigValue(param string) string{
	p, found := revel.Config.String(param)
	revel.WARN.Printf("%#v", p)
	if found{
		return p
	}

	return ""
}

func (b BaseCtrl) toJson(v interface{}) (string, error) {
	resByte, err := json.Marshal(v)
	return string(resByte), err
}

func (b BaseCtrl) fromJson(jsonToken string, v interface{}) (error) {
	err := json.Unmarshal([]byte(jsonToken), v)
	return err
}


