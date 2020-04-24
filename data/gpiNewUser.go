package data

import (
	"sso/model"
	//"sso/util"
	//"strings"
	"fmt"
	"log"
)

func NewUserDetails(details model.NewGPIUser) *model.NewGPIUser {
	//defer Conn.Close()
	detail := new(model.NewGPIUser)

	detail.Lastname = details.Lastname
	detail.Firstname = details.Firstname
	detail.Dob = details.Dob
	detail.Sex = details.Sex
	detail.Email = details.Email
	detail.Address = details.Address
	detail.City = details.City
	detail.Username = details.Username
	detail.Password = details.Password
	detail.RPassword = details.RPassword

	return detail

}

func UserNameExist(username string) bool {
	//defer Conn.Close()
	row := Conn.QueryRow("SELECT id FROM users WHERE username = '" + username + "'")
	var thisNewUser__id int
	
	row.Scan(&thisNewUser__id)

	if thisNewUser__id > 0 {
		return true
	}
	return false
}

func UserEmailExist(username string) bool {
	//defer Conn.Close()
	row := Conn.QueryRow("SELECT id FROM users WHERE myemail = '" + username + "'")
	var thisNewUser__id int
	
	row.Scan(&thisNewUser__id)

	if thisNewUser__id > 0 {
		return true
	}
	return false
}

func GetLoginDetails(user_name string, userid string) *model.GPIUser {
	//defer Conn.Close()
	row := Conn.QueryRow("select id, myemail, username, fname, lname, password, dob, gender, address, city FROM users where (username='" + user_name + "' or id = '"+userid+"')")
	var id int
	var myemail string
	var username string
	var fname string
	var lname string
	var dob string
	var gender string
	var password string
	var address string
	var city string

	row.Scan(&id, &myemail, &username, &fname, &lname, &password, &dob, &gender, &address, &city)

	if username != "" {
		var u = new(model.GPIUser)
		u.ID = id
		u.Lastname = lname
		u.Firstname = fname
		u.Dob = dob
		u.Sex = gender
		u.Email = myemail
		u.Address = address
		u.City = city
		u.Username = username
		u.Password = password
		fmt.Println("Looged in: "+username)
		log.Println("Looged in: "+username)
		return u
	}
	
	return nil

}
