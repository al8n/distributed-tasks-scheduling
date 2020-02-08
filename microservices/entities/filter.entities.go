package entities

import "net/http"

type FilterHandle func(w http.ResponseWriter, r *http.Request) error

// 拦截结构体
type Filter struct {
	// 用来储存需要拦截的uri
	filterMap map[string]FilterHandle
}

// 初始化函数
func NewFilter() *Filter  {
	return &Filter{filterMap: make(map[string]FilterHandle)}
}

// 注册拦截器
func(f *Filter) RegisterFilterUri(uri string, handler FilterHandle) {
	f.filterMap[uri] = handler
}

// 根据uri获取对应的handle
func (f *Filter) GetFilterHandle( uri  string) FilterHandle  {
	return f.filterMap[uri]
}


type WebHandle func(w http.ResponseWriter, r *http.Request)

// 执行 拦截器  返回函数类型
func (f *Filter) Handle(webHandle WebHandle) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		for path, handle := range f.filterMap {
			if path == r.RequestURI {
				err := handle(w, r)
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				break
			}
		}
		webHandle(w,r)
	}
}