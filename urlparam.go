package httpserver

import (
	"sync"
)

type urlParam struct {
	m_params sync.Map
}

func (this *urlParam) add(key *string, value *string) {
	this.m_params.Store(*key, *value)
}

func (this *urlParam) ByName(name string) *string {
	v, ok := this.m_params.Load(name)
	if !ok {
		return nil
	}
	value := v.(string)
	return &value
}
