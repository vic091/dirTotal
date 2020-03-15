package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	MAX_CONCURRENCY = 5
	limitChan       = make(chan struct{}, MAX_CONCURRENCY)
)

type DirTotal struct {
	DirCount  int   `json:"dirCount"`
	FileCount int   `json:"fileCount"`
	TotalSize int64 `json:"totalSize"`
}

// 获取所有子目录数量
func GetDirInfo(path string) (dirTotal DirTotal, err error) {
	ch := make(chan DirTotal, 5)
	chErr := make(chan error, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go ListDirInfo(path, &wg, ch, chErr)
	go func() {
		wg.Wait()
		close(ch)
		close(chErr)
	}()
	var ok bool
Loop:
	for {
		select {
		case d, ok := <-ch:
			if ok {
				dirTotal.TotalSize += d.TotalSize
				dirTotal.DirCount += d.DirCount
				dirTotal.FileCount += d.FileCount
			} else {
				break Loop
			}
		case err, ok = <-chErr:
			if ok {
				break Loop
			}
		}
	}

	return
}

func ListDirInfo(path string, wg *sync.WaitGroup, ch chan DirTotal, chErr chan error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		chErr <- err
		return
	}
	defer f.Close()
	defer wg.Done()
	fileInfo, _ := f.Readdir(-1)
	//操作系统指定的路径分隔符
	separator := string(os.PathSeparator)
	dirs := make([]string, 0)
	for _, fii := range fileInfo {
		fi, err := os.Stat(path + separator + fii.Name())
		if err != nil {
			continue
		}

		if fi.IsDir() {
			d := DirTotal{
				DirCount:  1,
				FileCount: 0,
				TotalSize: 0,
			}
			ch <- d
			dirs = append(dirs, filepath.Join(path, fi.Name()))
		} else {
			d := DirTotal{
				DirCount:  0,
				FileCount: 1,
				TotalSize: fi.Size(),
			}
			ch <- d
		}
	}
	for _, fi := range dirs {
		wg.Add(1)
		limitChan <- struct{}{}
		go ListDirInfo(fi, wg, ch, chErr)
		<-limitChan
	}
}

type Dirs struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Size  int64  `json:"size"`
}

// 获取目录信息
func DirInfo(path string) (dirInfo []Dirs, err error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer f.Close()
	fileInfo, _ := f.Readdir(-1)
	//操作系统指定的路径分隔符
	separator := string(os.PathSeparator)
	for _, info := range fileInfo {
		dir := Dirs{}
		if info.IsDir() {
			dir = Dirs{
				Name:  path + separator + info.Name(),
				IsDir: true,
				Size:  0,
			}
		} else {
			dir = Dirs{
				Name:  path + separator + info.Name(),
				IsDir: false,
				Size:  info.Size(),
			}
		}
		dirInfo = append(dirInfo, dir)
	}

	return dirInfo, nil
}
