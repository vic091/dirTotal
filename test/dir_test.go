package test

import (
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
