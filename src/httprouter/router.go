package httprouter

import "net/http"

type Router struct {

}

func New() *Router {
	return &Router{
	}
}

// 设置路径
func (r *Router) Path(path string) *Method {
	m := Method{}
	return m
}

// 设置前缀
func (r *Router) Prefix(prefix string) *Router {
	return r
}

// 设置中间件
func (r *Router) Use() *Router {
	return r
}

// HTTP 中间件
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	print("hello")
}