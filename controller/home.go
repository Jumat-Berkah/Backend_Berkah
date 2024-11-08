package controller

import (
	"net/http"

	"Backend_berkah/helper"
	"Backend_berkah/model"
)

func GetHome(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Response = helper.GetIPaddress()
	helper.WriteJSON(respw, http.StatusOK, resp)
}


func NotFound(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Response = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, resp)
}