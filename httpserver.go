package httpserver

import (
	"net/http"
)

var (
	POST    string = "POST"
	PUT     string = "PUT"
	GET     string = "GET"
	DELETE  string = "DELETE"
	OPTIONS string = "OPTIONS"
	HEAD    string = "HEAD"
	TRACE   string = "TRACE"
	CONNECT string = "CONNECT"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer, userdata interface{}) bool

type CHttpServer interface {
	Subscribe(topic string, method string, handler CRouterHandler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type CUrlParam interface {
	ByName(name string) *string
}

type CRouterHandler interface {
	handler(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer) bool
}

func NewHttpServer() CHttpServer {
	return &httpServerImpl{}
}

func NewRouterHandler(userdata interface{}, handle HandleFunc) CRouterHandler {
	return &routerHandler{m_userData: userdata, m_handlerFunc: handle}
}
