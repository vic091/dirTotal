package config

var RootPath *string

const HOST_IP = "http://127.0.0.1:8080/"

type ResDirs struct {
	Path string `json:"path"`
	Dirs []Dirs `json:"dirs"`
}

type ResDirsInfo struct {
	DirTotal
	Path string `json:"path"`
}

type DirTotal struct {
	DirCount  int   `json:"dirCount"`
	FileCount int   `json:"fileCount"`
	TotalSize int64 `json:"totalSize"`
}

type Dirs struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Size  int64  `json:"size"`
}
