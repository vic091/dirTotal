package utils

import (
	"dirTotal/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	MAX_CONCURRENCY = 5
	limitChan       = make(chan struct{}, MAX_CONCURRENCY)
)

// 获取所有子目录数量
func GetDirInfo(path string) (dirTotal config.DirTotal, err error) {
	ch := make(chan config.DirTotal, 5)
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

func ListDirInfo(path string, wg *sync.WaitGroup, ch chan config.DirTotal, chErr chan error) {
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
			d := config.DirTotal{
				DirCount:  1,
				FileCount: 0,
				TotalSize: 0,
			}
			ch <- d
			dirs = append(dirs, filepath.Join(path, fi.Name()))
		} else {
			d := config.DirTotal{
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

// 获取所有子目录数量
func GetHttpDirInfo(path string) (dirTotal config.DirTotal, err error) {
	ch := make(chan config.DirTotal, 5)
	chErr := make(chan error, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go ListHttpDirInfo(path, &wg, ch, chErr)
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

func ListHttpDirInfo(path string, wg *sync.WaitGroup, ch chan config.DirTotal, chErr chan error) {
	resp, err := http.Get(config.HOST_IP + "dir?path=" + path)
	if err != nil {
		chErr <- err
		return
	}

	defer resp.Body.Close()
	defer wg.Done()
	body, _ := ioutil.ReadAll(resp.Body)
	resDirs := config.ResDirs{}
	err = json.Unmarshal(body, &resDirs)
	if err != nil {
		chErr <- err
		return
	}
	dirs := make([]string, 0)

	for _, v := range resDirs.Dirs {
		if v.IsDir {
			d := config.DirTotal{
				DirCount:  1,
				FileCount: 0,
				TotalSize: 0,
			}
			ch <- d
			name := strings.Replace(v.Name, *config.RootPath, "", 0)
			dirs = append(dirs, name)
		} else {
			d := config.DirTotal{
				DirCount:  0,
				FileCount: 1,
				TotalSize: v.Size,
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

// 获取目录信息
func DirInfo(path string) (dirInfo []config.Dirs, err error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileInfo, _ := f.Readdir(-1)
	//操作系统指定的路径分隔符
	separator := string(os.PathSeparator)
	for _, info := range fileInfo {
		dir := config.Dirs{}
		if info.IsDir() {
			dir = config.Dirs{
				Name:  path + separator + info.Name(),
				IsDir: true,
				Size:  0,
			}
		} else {
			dir = config.Dirs{
				Name:  path + separator + info.Name(),
				IsDir: false,
				Size:  info.Size(),
			}
		}
		dirInfo = append(dirInfo, dir)
	}

	return dirInfo, nil
}
