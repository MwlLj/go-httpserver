package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var _ = fmt.Println

type httpServerImpl struct {
	m_urlParse urlParsePro
	m_topicMap sync.Map
}

func (this *httpServerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	isFind, findTopic, params := this.m_urlParse.findMatch(&url)
	if !isFind {
		io.WriteString(w, "no subscribe")
		return
	}
	v, ok := this.m_topicMap.Load(*findTopic)
	if !ok {
		io.WriteString(w, "not found methodMap")
		return
	}
	methodMap := v.(sync.Map)
	method := r.Method
	v, ok = methodMap.Load(method)
	if !ok {
		io.WriteString(w, "method no subscribe")
		return
	}
	var param urlParam
	param.init()
	param.set(params)
	handler := v.(CRouterHandler)
	res := handler.handler(w, r, &param, this)
	if !res {
		io.WriteString(w, "interior error")
	}
}

func (this *httpServerImpl) Subscribe(topic string, method string, handler CRouterHandler) {
	this.m_urlParse.regisnterUrl(topic)
	methodMap := sync.Map{}
	methodMap.Store(method, handler)
	this.m_topicMap.Store(topic, methodMap)
}
