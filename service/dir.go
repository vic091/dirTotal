package service

import (
	"dirTotal/config"
	"dirTotal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

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
	resData := config.ResDirs{
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
	resData := config.ResDirsInfo{
		DirTotal: info,
		Path:     r.FormValue("path"),
	}
	res, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("start http server fail:", err)
	}

	w.Write(res)
}

func (p RootPath) DirHttpInfoService(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")

	//data := make([]map[string]string, 0)
	info, err := utils.GetHttpDirInfo(path)
	if err != nil {
		w.Write([]byte("请求失败"))
		return
	}
	resData := config.ResDirsInfo{
		DirTotal: info,
		Path:     r.FormValue("path"),
	}
	res, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("start http server fail:", err)
	}

	w.Write(res)
}
