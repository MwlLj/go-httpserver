package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var _ = fmt.Println

type httpServer struct {
	m_routerMap sync.Map
	m_urlParse  urlParse
}

func (this *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.String()
	var param urlParam
	ok, v := this.m_urlParse.findMatch(&topic, &this.m_routerMap, &param)
	if !ok {
		io.WriteString(w, "no subscribe")
		return
	}
	methodMap := v.(sync.Map)
	method := r.Method
	v, ok = methodMap.Load(method)
	if !ok {
		io.WriteString(w, "method no subscribe")
		return
	}
	handler := v.(*routerHandler)
	res := handler.handler(w, r, &param, this)
	if !res {
		io.WriteString(w, "interior error")
	}
}

func (this *httpServer) Subscribe(topic string, method string, handler CRouterHandler) {
	methodMap := sync.Map{}
	methodMap.Store(method, handler)
	this.m_routerMap.Store(topic, methodMap)
}
