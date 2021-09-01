package ganLaoLi

import (
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type GanLaoLi struct {
	router map[string]HandleFunc
}

func New()  *GanLaoLi{
	return &GanLaoLi{router: make(map[string]HandleFunc)}
}

func (*GanLaoLi) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	// 找路由, 业务处理
}

func (g *GanLaoLi) Run (addr string) error {
	return http.ListenAndServe(addr, g)
}
