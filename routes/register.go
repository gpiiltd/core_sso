package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	// "regexp"
	// "sso/data"
	// "sso/model"
	// //"sso/util"
	// "log"
	// "strings"
	// "fmt"
)

func Register(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	service := req.FormValue("service")
	
	r.HTML(200, "register", service)
	return
}