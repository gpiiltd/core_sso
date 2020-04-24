package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"sso/data"
	"sso/model"
	//"sso/util"
	"log"
	"strings"
	"fmt"
)

func Login(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Loginn")
	fmt.Println("in loginn")
	//parse request parameters
	req.ParseForm()
	service := req.FormValue("service")
	//clientcallback := req.FormValue("clientcallback")
	//service := req.FormValue("service")

	log.Println("login service:" + service)
	fmt.Println("login service:" + service)

	var ticket string
	cookies := req.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "CASTGC" {
			ticket = cookie.Value
			log.Println("find tgt in cookies:" + ticket)
		}
	}
	if ticket == "" {
		fmt.Println("TICKET IS EMPTYk")
		log.Println("TICKET IS EMPTY")
		if validateServicel(service) {
			fmt.Println("service is valid")
			log.Println("service is valid")
			log.Println("service validate, redirect to login page")
			fmt.Println("service validate, redirect to login page")
			r.HTML(200, "login", service)
			return
		}
	}
	fmt.Println("login ticket found! "+ticket)
	tgt := data.FindTGT(ticket)
	
	if tgt != nil {
		if service == "" {
			log.Println("tgt found and valid, no service found, redirect to default success page")
			fmt.Println("tgt found and valid, no service found, redirect to default success page")
			r.HTML(200, "logincc", "gpitest")
			return
		}
		if validateServicel(service) {
			log.Println("tgt found and valid, servince:" + service)
			fmt.Println("tgt found and valid, servince:" + service)
			st := data.GrantServiceTicket(tgt.Tgt, service)
			log.Println("st granteed, st:" + st.St)
			fmt.Println("st granteed, st:" + st.St)
			data.AddSTToTGT(tgt, st)
			redirectToServicel(r, st,"")
			return
		}
		log.Println("service invalid")
	}
	fmt.Println("redirect to login page tgt is nill")
	log.Println("redirect to login page")
	fmt.Println("fixing clientcallback issue")
	log.Println("fixing clientcallback issue")
	r.HTML(200, "login", service)
	return

}

func validateServicel(service string) bool {
	if service == "" {
		return true
	}
	reg := regexp.MustCompile(`^(https|http)://.*`)
	return reg.MatchString(service)
}

func redirectToServicel(r render.Render, st *model.ServiceTicket, name string) {
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
