package httpserver

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer, userdata interface{}) bool {
	io.WriteString(w, "index")
	return true
}

func HandleIndex2(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer, userdata interface{}) bool {
	io.WriteString(w, "index2")
	return true
}

func HandleHello(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer, userdata interface{}) bool {
	io.WriteString(w, strings.Join([]string{"hello", *param.ByName("name")}, ":"))
	return true
}

func HandleHello2(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer, userdata interface{}) bool {
	io.WriteString(w, strings.Join([]string{"hello", *param.ByName("name"), "age", *param.ByName("age")}, ":"))
	return true
}

func TestHttpServer(t *testing.T) {
	server := NewHttpServer()
	server.Subscribe("/", GET, NewRouterHandler(nil, HandleIndex))
	server.Subscribe("/index", GET, NewRouterHandler(nil, HandleIndex2))
	server.Subscribe("/hello/:name", GET, NewRouterHandler(nil, HandleHello))
	server.Subscribe("/hello/name/:name/age/:age", GET, NewRouterHandler(nil, HandleHello2))
	http.ListenAndServe(":59000", server)
}
