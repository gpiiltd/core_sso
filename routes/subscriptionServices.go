package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	//"sso/data"
	"sso/model"
	//"sso/util"
	"log"
	"strings"
	"fmt"
)

func Subscribe(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in subscribe")
	fmt.Println("in subscribe")
	//parse request parameters
	req.ParseForm()
	service := "gpitest" //"Default"

	log.Println("subscribe service:" + service)
	fmt.Println("Subscribe service:" + service)

	
	
	fmt.Println("redirect, service is nill")
	log.Println("redirect to subscribe page")
	r.HTML(200, "subscribe", service)
	return

}

func SubscribeToService(r render.Render, res http.ResponseWriter, req *http.Request) {
	// log.Println("in Login")
	// fmt.Println("in login")
	// //parse request parameters
	// req.ParseForm()
	// service := req.FormValue("service")

	// log.Println("login service:" + service)
	// fmt.Println("login service:" + service)

	// var ticket string
	// cookies := req.Cookies()
	// for _, cookie := range cookies {
	// 	if cookie.Name == "CASTGC" {
	// 		ticket = cookie.Value
	// 		log.Println("find tgt in cookies:" + ticket)
	// 	}
	// }
	// if ticket == "" {
	// 	fmt.Println("TICKET IS EMPTYk")
	// 	log.Println("TICKET IS EMPTY")
	// 	if validateService(service) {
	// 		fmt.Println("service is valid")
	// 		log.Println("service is valid")
	// 		log.Println("service validate, redirect to login page")
	// 		fmt.Println("service validate, redirect to login page")
	// 		r.HTML(200, "login", service)
	// 		return
	// 	}
	// }
	// fmt.Println("login ticket found! "+ticket)
	// tgt := data.FindTGT(ticket)
	
	// if tgt != nil {
	// 	if service == "" {
	// 		log.Println("tgt found and valid, no service found, redirect to default success page")
	// 		fmt.Println("tgt found and valid, no service found, redirect to default success page")
	// 		r.HTML(200, "success", "")
	// 		return
	// 	}
	// 	if validateService(service) {
	// 		log.Println("tgt found and valid, service:" + service)
	// 		fmt.Println("tgt found and valid, service:" + service)
	// 		st := data.GrantServiceTicket(tgt.Tgt, service)
	// 		log.Println("st granted, st:" + st.St)
	// 		fmt.Println("st granted, st:" + st.St)
	// 		data.AddSTToTGT(tgt, st)
	// 		redirectToService(r, st,"")
	// 		return
	// 	}
	// 	log.Println("service invalid")
	// }
	// fmt.Println("redirect to login page tgt is nill")
	// log.Println("redirect to login page")
	r.HTML(200, "logibbn", "service")
	return

}


func nvalidateService(service string) bool {
	if service == "" {
		return true
	}
	reg := regexp.MustCompile(`^(https|http)://.*`)
	return reg.MatchString(service)
}

func nredirectToService(r render.Render, st *model.ServiceTicket, name string) {
	needAnd := strings.Contains(st.Service, "?")
	sep := "?"
	if needAnd {
		sep = "&"
	}
	redirectUrl := st.Service + sep + "ticket=" + st.St + "&parsedobj="+name
	log.Println("redirect to service:" + redirectUrl)
	fmt.Println("redirect to service:" + redirectUrl)
	r.Redirect(redirectUrl)
}
