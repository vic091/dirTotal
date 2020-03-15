package route

import (
	"dir/service"
	"net/http"
)

type MyHandler struct {
	RootPath string
}

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := service.RootPath{Path: mh.RootPath}
	if r.URL.Path == "/dir" {
		p.DirService(w, r)
		return
	} else if r.URL.Path == "/dir_info" {
		p.DirInfoService(w, r)
		return
	}

	w.Write([]byte("路由不存在"))
}
