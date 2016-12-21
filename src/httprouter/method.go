package httprouter

type Method struct {
}

// 设置当前中间件
func (m *Method) Use() *Method {
	return m
}

func (m *Method) Get() *Method {
	return m
}

func (m *Method) Head() *Method {
	return m
}

func (m *Method) Options() *Method {
	return m
}

func (m *Method) Post() *Method {
	return m
}

func (m *Method) Put() *Method {
	return m
}

func (m *Method) Patch() *Method {
	return m
}

func (m *Method) Delete() *Method {
	return m
}
