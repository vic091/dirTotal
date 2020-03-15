package service

import (
	"dir/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResDirs struct {
	Path string       `json:"path"`
	Dirs []utils.Dirs `json:"dirs"`
}

type ResDirsInfo struct {
	utils.DirTotal
	Path string `json:"path"`
}
type RootPath struct {
	Path string
}

// 获取目录详情
func (p RootPath) DirService(w http.ResponseWriter, r *http.Request) {
	path := p.Path + r.FormValue("path")
	data, err := utils.DirInfo(path)
	if err != nil {
		w.Write([]byte("访问失败"))
	}
	resData := ResDirs{
		Path: r.FormValue("path"),
		Dirs: data,
	}
	res, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("start http server fail:", err)
		w.Write([]byte("访问失败"))
	}
	w.Write(res)
}

// 统计目录信息
func (p RootPath) DirInfoService(w http.ResponseWriter, r *http.Request) {
	path := p.Path + r.FormValue("path")

	//data := make([]map[string]string, 0)
	info, err := utils.GetDirInfo(path)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("请求失败"))
		return
	}
	resData := ResDirsInfo{
		DirTotal: info,
		Path:     r.FormValue("path"),
	}
	res, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("start http server fail:", err)
	}

	w.Write(res)
}
