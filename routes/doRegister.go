package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	//"regexp"
	"sso/model"
	"sso/data"
	"sso/util"
	"log"
	//"strings"
	"fmt"
)

func DoRegister(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in register")
	fmt.Printf("in register")
	founderrors := false
	//parse request parameters
	req.ParseForm()
	//fmt.Printf("%+v\n",)
	fmt.Printf("______________________________")
	var ThisNewUserDetails model.NewGPIUser
	service := req.FormValue("service")
	fmt.Println("")
	fmt.Println(service)
	ThisNewUserDetails.Lastname = req.FormValue("lastname") //     string  `json:"lastname"`
	ThisNewUserDetails.Firstname = req.FormValue("othernames") //string  `json:"firstname"`
	ThisNewUserDetails.Dob = "" //      string `json:"dob"`
	ThisNewUserDetails.Sex = "" //      string `json:"sex"`
	ThisNewUserDetails.Email = req.FormValue("email") //        string  `json:"email"`
	ThisNewUserDetails.Address = req.FormValue("address") // string  `json:"address"`
	ThisNewUserDetails.City = req.FormValue("city") //      string `json:"city"`
	ThisNewUserDetails.Username = req.FormValue("username") //      string `json:"username"`
	ThisNewUserDetails.Password = req.FormValue("password") //	string `json:"username"`
	ThisNewUserDetails.RPassword = req.FormValue("rpassword") //	string `json:"username"`

	fmt.Println(ThisNewUserDetails.RPassword)

	if ThisNewUserDetails.Password != ThisNewUserDetails.RPassword {
		//founderrors = true
		fmt.Println("password mismatch")
		fmt.Println(ThisNewUserDetails.Password)
		fmt.Println(ThisNewUserDetails.RPassword)
	}


	thisUser := data.NewUserDetails(ThisNewUserDetails)

	//nameExists := data.UserNameExist()
	if data.UserNameExist(thisUser.Username) == true {
		founderrors = true
		fmt.Println("username taken")
	}

	if data.UserEmailExist(ThisNewUserDetails.Email) == true {
		founderrors = true
		fmt.Println("email taken")
		r.Redirect("http://178.128.251.254:5000/register?service="+service+"&info=eMail Taken. Please use another email address!")
		//http://178.128.251.254:5000/register?service=http://178.128.251.254/gpitest
		return
	}

	
	if founderrors == false {
		fmt.Printf("writing to db")
	hash, _ := util.HashPassword(thisUser.Password) // ignore error for the sake of simplicity
	companyid := "1"
	_, _ = data.Conn.Exec("insert into users ( myemail,username,password,fname,lname,company,address,city)" +
		" values ( '" + ThisNewUserDetails.Email + "', '" + ThisNewUserDetails.Username + "', '" + hash + "', '" +
		ThisNewUserDetails.Firstname + "', '" + ThisNewUserDetails.Lastname + "', '" + companyid +"', '" + ThisNewUserDetails.Address + "', '" +  ThisNewUserDetails.City + "')")

	}


	//myUserInsertId, err := lastIdVar.LastInsertId()
	if founderrors == false {
		fmt.Printf("writing permission")
		//founderrors = false
		mypermission := "1"
		_, _ = data.Conn.Exec("insert into user_permission_matches ( user_id,permission_id )" + 
			" values ( '" + "1" + "', '" + mypermission + "')")

	}

	
	//fmt.Println("redirect to url:" + service+"/registration_not_confirmed")
	if founderrors == true {
		//redirectUrl = "register"
		//fmt.Println("redirect to url:" + service+"/registration_not_confirmed")
		//r.Redirect(service+"/registration_not_confirmed/?info=Username&nbsp;Taken")
		dbuild := make(map[string]interface{})
		dbuild["service"] = service
		dbuild["information"] = "Username taken! Please try again using a different name."
		dbuild["status"] = false
		//r.HTML(200, "register", dbuild)
		r.Redirect("http://178.128.251.254:5000/register?service="+service+"&info=Username Taken. Please choose a new name and try again!")
		//http://178.128.251.254:5000/register?service=http://178.128.251.254/gpitest
		return
	} else {
		redirectUrl := service+"/registration_confirmation"
	//redirectUrl = "subscriptions"
	//r.HTML(200, "subscribe", service)
	//return

	log.Println("redirect to url:" + redirectUrl)
	r.Redirect(redirectUrl)
	return
	}
	
}

func DoRegisterEnt(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in register")
	fmt.Printf("in register")
	founderrors := false
	//parse request parameters
	req.ParseForm()
	fmt.Printf("%+v\n", req.Form)
	fmt.Printf("______________________________")
	//req.ParseForm()
	var ThisNewUserDetails model.NewGPIUser
	service := req.FormValue("service")
	companyid := req.FormValue("companyid")
	ThisNewUserDetails.Lastname = req.FormValue("lname") //     string  `json:"lastname"`
	ThisNewUserDetails.Firstname = req.FormValue("fname") //string  `json:"firstname"`
	ThisNewUserDetails.Dob = "" //      string `json:"dob"`
	ThisNewUserDetails.Sex = "" //      string `json:"sex"`
	ThisNewUserDetails.Email = req.FormValue("email") //        string  `json:"email"`
	ThisNewUserDetails.Address = "" // req.FormValue("address") // string  `json:"address"`
	ThisNewUserDetails.City = req.FormValue("city") //      string `json:"city"`
	ThisNewUserDetails.Username = req.FormValue("username") //      string `json:"username"`
	ThisNewUserDetails.Password = req.FormValue("password") //	string `json:"username"`
	ThisNewUserDetails.RPassword = req.FormValue("rpassword") //	string `json:"username"`
	result := make(map[string]interface{})

	if ThisNewUserDetails.Password != ThisNewUserDetails.RPassword {
		//founderrors = true
	}


	thisUser := data.NewUserDetails(ThisNewUserDetails)

	//nameExists := data.UserNameExist()
	if data.UserNameExist(thisUser.Username) == true {
		founderrors = true		
	}

	hash, _ := util.HashPassword(thisUser.Password) // ignore error for the sake of simplicity
	//companyid := companyid
	_, _ = data.Conn.Exec("insert into users ( myemail,username,password,fname,lname,company,address,city)" +
		" values ( '" + ThisNewUserDetails.Email + "', '" + ThisNewUserDetails.Username + "', '" + hash + "', '" +
		ThisNewUserDetails.Firstname + "', '" + ThisNewUserDetails.Lastname + "', '" + companyid +"', '" + ThisNewUserDetails.Address + "', '" +  ThisNewUserDetails.City + "')")

	//myUserInsertId, err := lastIdVar.LastInsertId()
	err := false
	if err != false {
		founderrors = true
	} else {
		founderrors = false
		mypermission := "1"
		_, _ = data.Conn.Exec("insert into user_permission_matches ( user_id,permission_id )" + 
			" values ( '" + "1" + "', '" + mypermission + "')")

	}

	redirectUrl := service+"/registration_confirmation"

	if founderrors {
		redirectUrl = "register"
		log.Println("redirect to url:" + redirectUrl)
		//r.Redirect(redirectUrl)
		r.JSON(401, "{\"status\":\"error\"}")
		return
	}

	//redirectUrl = "subscriptions"
	//r.HTML(200, "subscribe", service)
	//return
	result["status"] = "OK"
	result["desription"] = "Registration Complated Successfully"
	log.Println("redirect to url:" + redirectUrl)
	//r.Redirect(redirectUrl)
	r.JSON(200, result)
	return
}