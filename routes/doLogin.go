package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"sso/data"
	"sso/model"
	"sso/util"
	"log"
	"time"
	"strings"
	"fmt"
	"strconv"
	"encoding/json"
)

func DoLogin(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Login")
	//parse request parameters
	req.ParseForm()

	service := req.FormValue("service")
	username := req.FormValue("username")
	password := req.FormValue("password")
	client_id := req.FormValue("client_id")
	redirect_uri := req.FormValue("redirect_uri")

	fmt.Println(client_id)
	fmt.Printf("%+v",req.Form)

	fmt.Println(redirect_uri)
	fmt.Println(client_id)
	log.Println("parsed params:service:" + service + ",username:" + username + ",password:" + password)
	fmt.Println("parsed params:service:" + service + ",username:" + username + ",password:" + password)

	//TODO Handle User Auth

	tgt := data.GrantTicketGrantingTicket(username, "")
	if tgt == nil {
		log.Fatalln("error grant tgt")
		fmt.Println("error grant tgt")
		r.HTML(200, "error", "")
		return
	}
	log.Println("tgt granted:" + tgt.Tgt)
	fmt.Println("tgt granted:" + tgt.Tgt)

	cookie := http.Cookie{Name: "CASTGC", Value: tgt.Tgt, Path: "/", Domain: util.COOKIE_DOMAIN, MaxAge: util.TICKET_GRANTING_TICKET_TIME_TO_LIVE}
	http.SetCookie(res, &cookie)

	st := data.GrantServiceTicket(tgt.Tgt, service)
	log.Println("st granted:" + st.St)
	fmt.Println("st granted:" + st.St)
	acookie := http.Cookie{Name: "GPIGC", Value: username, Path: "/", Domain: util.COOKIE_DOMAIN, MaxAge: util.TICKET_GRANTING_TICKET_TIME_TO_LIVE}
	http.SetCookie(res, &acookie)
	if st == nil {
		log.Fatalln("error grant st")
		fmt.Println("error grant st")
		r.HTML(200, "error", "")
		return
	}
	data.AddSTToTGT(tgt, st)
	validUser := data.GetLoginDetails(username, "")


	fmt.Printf("%+v\n",validUser)
	//fmt.Printf("validUser: "+validUser.Username)



	match := false
	if (validUser != nil) {
		//fmt.Println("valid user object is nil")
		match = util.CheckPasswordHash(password, validUser.Password)
	}

			
	log.Println(match)
	fmt.Println(match)
	fmt.Println("redirecting to... service")
	fmt.Println(st.Service)
	fmt.Println("************")
	if (match) {
		log.Println("matched")
		fmt.Println(st.Service)
		s := strings.Split(st.Service, "=")
    	var2 := s[1]
		suid := strconv.Itoa(validUser.ID)
		getSubscription := data.GetSubscription(var2, suid)
		//strconv.Itoa(validUser.ID)
		if (getSubscription == false) {
			fmt.Println("yes")
			gresult := make(map[string]interface{})
			gresult["username"] = validUser.Username
			gresult["uid"] = validUser.ID

			out, err := json.Marshal(st)
		    if err != nil {
		        //panic (err)
		    }

		    fmt.Println("%+v\n", out)
		    fmt.Println("---------"+string(out)+"--------")

		    key := "OAUTH_CODE_" + validUser.Username
			rErr := data.Cli.Set(key, string(out), time.Millisecond*util.SERVICE_TICKET_TIME_TO_LIVE).Err()
			if rErr != nil {
				fmt.Println(err)
				log.Fatalln("error grant st")
				fmt.Println("error grant st")
				r.HTML(200, "error", "")
				return
			}

			fmt.Printf("p-----"+string(out)+"----p")
			//fmt.Printf(string(validUser.ID))
			uid := strconv.Itoa(validUser.ID)


			//r.HTML(200, "subscribe", gresult)
			r.Redirect("http://178.128.251.254:3000/gpiSubscribe?uid="+uid+"&username="+validUser.Username)	
			return
		}
		
		redirectToServiced(r, st, username)
	 	//r.HTML(200, "login", service)
	} else {
		fmt.Println("not match o")
	 	//r.HTML(200, "login", service)
	 	//http://178.128.251.254:3000/oauth/authorize?client_id=gpitest&response_type=code&redirect_uri=my-gpi.io/gpitest&scope=read
	 	r.Redirect("http://178.128.251.254:3000/oauth/authorize/?client_id="+client_id+"&response_type=code&redirect_uri="+redirect_uri+"&scope=read")
	 		//http://178.128.251.254:3000/gpiSubscribe?uid="+uid+"&username="+validUser.Username)	
			return
	}
	//redirectToServiced(r, st, username)
	return

}

func validateServiced(service string) bool {
	if service == "" {
		return true
	}
	reg := regexp.MustCompile(`^(https|http)://.*`)
	return reg.MatchString(service)
}

func redirectToServiced(r render.Render, st *model.ServiceTicket, name string) {
	needAnd := strings.Contains(st.Service, "?")
	sep := "?"
	if needAnd {
		sep = "&"
	}
	redirectUrl := st.Service + sep + "ticket=" + st.St + "&parsedobj="+name
	log.Println("redirect to serviceb:" + redirectUrl)
	fmt.Println("redirect to serviceb:" + redirectUrl)
	r.Redirect(redirectUrl)
	//r.Redirect("callRoute/"+redirectUrl)
}
