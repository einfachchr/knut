package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	// In der IFrame-Lösung müssen wir über die im docker-compose.yml konfigurierten Host-Ports gehen, da der Browser
	// das IFrame füllt. Hübsch häßlich...
	constRedURL  = "http://red:8080"  //"http://localhost:9091" //"http://red:8080"
	constBlueURL = "http://blue:8080" //"http://localhost:9092" //"http://blue:8080"
)

type MyRouter struct {
	FrameTemplate *template.Template
	WelcomePage   *template.Template
}

type RoutingParams struct {
	Content string
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
	var welcome bytes.Buffer
	var err = self.WelcomePage.Execute(&welcome, nil)
	if err != nil {
		panic(err)
	}
	self.FrameTemplate.Execute(resp, RoutingParams{Content: welcome.String()})
}

func (self *MyRouter) WelcomeHandler(resp http.ResponseWriter, req *http.Request) {
	self.WelcomePage.Execute(resp, nil)
}

func (self *MyRouter) RedHandler(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[RedHandler] request for %s", constRedURL)
	var content = self.getContent(req, constRedURL)
	log.Printf("[RedHandler] got: %s", content)

	self.FrameTemplate.Execute(resp, RoutingParams{Content: content})
}

func (self *MyRouter) BlueHandler(resp http.ResponseWriter, req *http.Request) {
	log.Printf("[RedHandler] request for %s", constBlueURL)
	var content = self.getContent(req, constBlueURL)
	log.Printf("[RedHandler] got: %s", content)

	self.FrameTemplate.Execute(resp, RoutingParams{Content: content})
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

//
// ------------------------ GET-Requests ins Backend
//
func (self *MyRouter) getContent(req *http.Request, target string) string {
	fmt.Printf("[GET] request-url: %#v\n", req.URL)
	fmt.Printf("[GET] request-header: %#v\n", req.Header)
	resp, err := http.Get(target)
	if err != nil {
		return fmt.Sprintf("got error while GET: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("got error while reading response body: %s", err)
	}

	return string(body)
}
