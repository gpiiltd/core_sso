package main

import (
	//"net/http"
	"os"
	"log"
	//"io"
	"github.com/martini-contrib/cors"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"sso/routes"
	"sso/util"
)

func init() {
	file, err := os.Create(util.LOG_FILE)
	if err != nil {
		log.Println("error create logFile")
		return
	} else {
		log.SetOutput(file)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("sso:")

}

func main() {
	/*http.HandleFunc("/test.png", serveimage)
	func serveimage(w http.ResponseWriter, r *http.Request) {
     http.ServeFile(w, r, "test.png")
	}*/
	//fs := http.FileServer(http.Dir("templates"))
	//http.Handle("/templates", http.StripPrefix("/templates/", fs))
  	//http.Handle("/", fs)
  	//m.Get("/templates",fs)
  	//m.Get("/templates", http.StripPrefix("/templates/", fs))
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		//IndentJSON: true,                       // Output human readable JSON
		//IndentXML:  true,                       // Output human readable XML
	}))
	m.Use(martini.Static("templates"))

	m.Use(cors.Allow(&cors.Options{
	    AllowOrigins:     []string{"*"},
	    AllowMethods:     []string{"GET, POST"},
	    AllowHeaders:     []string{"Origin, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"},
	    ExposeHeaders:    []string{"Content-Length, Content-Type"},
	    AllowCredentials: false,
  	}))
	//m.Use(martini.Static("publicam"))

	//m.Get("/", func() string {
    //return "Hello world!"
  //})
  //http.Handle("/", m)*/
  	//fs := http.FileServer(http.Dir("template"))
	//m.Handlers(http.Handle("/templates", http.StripPrefix("/templates/", fs)))
	/*m.Get("/template", func (w http.ResponseWriter, r *http.Request) {
     http.Handle("/template", http.StripPrefix("/template/", http.FileServer(http.Dir("template")))) //http.FileServer(http.Dir("template")) //http.ServeFile(w, r, "template")
	})*/ 
	///sso/:client/:id

	// h1 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello\n")
	// }

	// m.Get("/", h1)
	m.Post("/profilePicture", routes.ProfilePicture)
	m.Post("/doRegisterEnt", routes.DoRegisterEnt)
	m.Get("/getProfilePicture", routes.GetProfilePicture)
	m.Get("/callRoute/:client/:id", routes.CallRoute)
	m.Get("/subscribe", routes.Subscribe)
	m.Get("/register", routes.Register)
	m.Post("/doRegister", routes.DoRegister)
	m.Get("/login", routes.Login)
	m.Post("/doLogin", routes.DoLogin)
	m.Get("/subsribeToService", routes.SubscribeToService)
	m.Get("/serviceValidate", routes.ServiceValidate)
	//m.Run()
	m.RunOnAddr("178.128.251.254:5000")
	
	//log.Println("Listening...")
  	//http.ListenAndServe(":5000", nil)
}