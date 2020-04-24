package routes

import (
	//"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	// "regexp"
	"sso/data"
	// "sso/model"
	// //"sso/util"
	// "log"
	// "strings"
	// "fmt"
)

type U_Data struct {
	UserID int
	Token string
	Username string
	Redirect string
}

func GetProfilePicture(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//params := make(map[string]interface{})

	//username := req.FormValue("username")
	userid := req.FormValue("id")
	redirectUri := req.FormValue("redirect")
	//userid := paramss["id"]

	validUser := data.GetLoginDetails("", userid)

	if (validUser != nil) {
		params := &U_Data{
			UserID:	validUser.ID,
			Token: "",
			Username: validUser.Username,
			Redirect: redirectUri,
		}
		r.HTML(200, "profilePic", params)
		return	
	} else {
		r.HTML(200, "logind", "")
		return
	}
	
	
}