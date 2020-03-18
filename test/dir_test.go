package test

import (
	"dirTotal/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

type DirTotal struct {
	path      string
	DirCount  int
	FileCount int
	TotalSize int64
}

var RootPath = "/Users/admin/www/learn/go/src/dir/data"

func TestDir(t *testing.T) {
	time.Sleep(2 * time.Second)
}

// 目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func TestHttp(t *testing.T) {
	path := "numpy"
	resp, err := http.Get("http://127.0.0.1:8080/dir?path=" + path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	dirs := config.ResDirs{}
	json.Unmarshal(body, &dirs)
	fmt.Println(dirs)
	for k, v := range dirs.Dirs {
		fmt.Println(k, "======", v)
	}

}
