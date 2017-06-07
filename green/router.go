package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	// In der IFrame-Lösung müssen wir über die im docker-compose.yml konfigurierten Host-Ports gehen, da der Browser
	// das IFrame füllt. Hübsch häßlich...
	constRedURL  = "http://localhost:9091" //"http://red:8080"
	constBlueURL = "http://localhost:9092" //"http://blue:8080"
)

type MyRouter struct {
	FrameTemplate *template.Template
	WelcomePage   *template.Template
}

type RoutingParams struct {
	IFrameTarget string
}

func main() {
	var router = MyRouter{}
	router.loadTemplates()

	var r = router.router()

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

//
// ------------------------ Routing
//

func (self *MyRouter) router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", self.RootHandler)
	r.HandleFunc("/welcome.html", self.WelcomeHandler)
	r.HandleFunc("/red", self.RedHandler)
	r.HandleFunc("/blue", self.BlueHandler)

	return r
}

// Zeigt die Index-Seite
func (self *MyRouter) RootHandler(resp http.ResponseWriter, req *http.Request) {
	self.FrameTemplate.Execute(resp, RoutingParams{IFrameTarget: "/welcome.html"})
}

func (self *MyRouter) WelcomeHandler(resp http.ResponseWriter, req *http.Request) {
	self.WelcomePage.Execute(resp, nil)
}

func (self *MyRouter) RedHandler(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[ROUTER] iframe for %s\n", constRedURL)
	self.FrameTemplate.Execute(resp, RoutingParams{IFrameTarget: constRedURL})
}

func (self *MyRouter) BlueHandler(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[ROUTER] iframe for %s\n", constBlueURL)
	self.FrameTemplate.Execute(resp, RoutingParams{IFrameTarget: constBlueURL})
}

//
// ------------------------ Templates
//
func (self *MyRouter) loadTemplates() {
	self.loadFrameTemplate()
	self.loadWelcomePage()
}

func (self *MyRouter) loadFrameTemplate() {
	var templ *template.Template
	templ = template.Must(template.ParseFiles("frame.html"))

	self.FrameTemplate = templ
}

func (self *MyRouter) loadWelcomePage() {
	var templ *template.Template
	templ = template.Must(template.ParseFiles("welcome.html"))

	self.WelcomePage = templ
}
