package main

import (
	".."
	"fmt"
	"io"
	"net/http"
	"strings"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, param httpserver.CUrlParam, server httpserver.CHttpServer, userdata interface{}) bool {
	s := userdata.(*CServer)
	s.CommonLogic()
	io.WriteString(w, "index")
	return true
}

func HandleHello(w http.ResponseWriter, r *http.Request, param httpserver.CUrlParam, server httpserver.CHttpServer, userdata interface{}) bool {
	io.WriteString(w, strings.Join([]string{"hello", *param.ByName("name")}, ":"))
	return true
}

func HandleHello2(w http.ResponseWriter, r *http.Request, param httpserver.CUrlParam, server httpserver.CHttpServer, userdata interface{}) bool {
	io.WriteString(w, strings.Join([]string{"hello", *param.ByName("name"), ", age", *param.ByName("age")}, ":"))
	return true
}

func HandleError(w http.ResponseWriter, r *http.Request, param httpserver.CUrlParam, server httpserver.CHttpServer, userdata interface{}) bool {
	return false
}

func HandlePound(w http.ResponseWriter, r *http.Request, param httpserver.CUrlParam, server httpserver.CHttpServer, userdata interface{}) bool {
	return true
}

type CServer struct {
	m_http httpserver.CHttpServer
}

func (this *CServer) CommonLogic() {
	fmt.Println("common logic")
}

func (this *CServer) Start() {
	// new
	this.m_http = httpserver.NewHttpServer()
	// resubscribe
	this.m_http.Subscribe("/", httpserver.GET, httpserver.NewRouterHandler(this, HandleIndex))
	this.m_http.Subscribe("/error", httpserver.GET, httpserver.NewRouterHandler(this, HandleError))
	this.m_http.Subscribe("/hello/:name", httpserver.GET, httpserver.NewRouterHandler(this, HandleHello))
	this.m_http.Subscribe("/hello/name/:name/age/:age", httpserver.GET, httpserver.NewRouterHandler(this, HandleHello2))
	this.m_http.Subscribe("/pound/:name/#", httpserver.GET, httpserver.NewRouterHandler(this, HandlePound))
	http.ListenAndServe(":59000", this.m_http)
}

func main() {
	server := CServer{}
	server.Start()
}
