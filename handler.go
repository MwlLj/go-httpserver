package httpserver

import (
	"net/http"
)

type routerHandler struct {
	m_userData    interface{}
	m_handlerFunc HandleFunc
}

func (this *routerHandler) userData() interface{} {
	return this.m_userData
}

func (this *routerHandler) handlerFunc() HandleFunc {
	return this.m_handlerFunc
}

func (this *routerHandler) handler(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer) bool {
	return this.m_handlerFunc(w, r, param, server, this.m_userData)
}
