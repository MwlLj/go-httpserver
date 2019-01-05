package httpserver

type urlParam struct {
	m_params map[string]string
}

func (this *urlParam) init() {
	this.m_params = make(map[string]string)
}

func (this *urlParam) add(key *string, value *string) {
	this.m_params[*key] = *value
}

func (this *urlParam) set(params *map[string]string) {
	this.m_params = *params
}

func (this *urlParam) ByName(name string) *string {
	v, ok := this.m_params[name]
	if !ok {
		return nil
	}
	return &v
}
