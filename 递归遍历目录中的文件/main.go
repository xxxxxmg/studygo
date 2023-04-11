package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	dirPath = flag.String("dir", "./", "遍历目录中的所有文件及路径")
)

func init() {
	flag.Parse()
}

func GetAllDirFilePath(dirName string) ([]string, error) {

	dirName = strings.TrimSuffix(dirName, string(os.PathSeparator))

	infos, err := ioutil.ReadDir(dirName)

	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, len(infos))

	for _, info := range infos {
		path := dirName + string(os.PathSeparator) + info.Name()
		if info.IsDir() {
			tmp, err := GetAllDirFilePath(path)
			if err != nil {
				return nil, err
			}
			paths = append(paths, tmp...)
			continue
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func main() {

	fmt.Println(*dirPath)

	paths, _ := GetAllDirFilePath(*dirPath)
	for _, path := range paths {
		strs := strings.Split(path, "\\")
		fmt.Println(strs[len(strs)-1], "文件的路径为：", path)
	}
}
