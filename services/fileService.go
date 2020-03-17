package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PickupModifiedFiles/models"
)

// GenerateCurrentDirContent returns DirContent of current directory.
func GenerateCurrentDirContent(cd string) models.DirContent {
	list := readDir(cd)
	return genDirContent(list)
}

// GenerateTransferDir gens directory tree for transfer
func GenerateTransferDir(dirContent models.DirContent) {
	t := time.Now()
	date := t.String()[:10]

	transferDirName := date
	if finfo, err := os.Stat(transferDirName); os.IsNotExist(err) || !finfo.IsDir() {
		os.Mkdir(transferDirName, 0777)
	}

	contents := dirContent.Contents
	for _, content := range contents {
		spPath := strings.Split(content.Name, "\\")
		path := date
		for _, pathFactor := range spPath {
			path = filepath.Join(path, pathFactor)
			dir, _ := filepath.Split(path)

			os.Mkdir(dir, 0777)
		}
	}
}

// CopyToTransferDir copies files to transferdir
func CopyToTransferDir(contents []models.FileInfo, srcPath, transferDirName string) {
	for _, file := range contents {
		dir, name := filepath.Split(file.Name)
		fmt.Println(dir)
		fmt.Println(name)
		os.Link(filepath.Join(srcPath, file.Name), filepath.Join(transferDirName, file.Name))
	}
}

// readDir is to read directories.
func readDir(rootPath string) []string {
	finfos, err := ioutil.ReadDir(rootPath)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, finfo := range finfos {
		if finfo.IsDir() {
			paths = append(paths, readDir(filepath.Join(rootPath, finfo.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(rootPath, finfo.Name()))
	}

	return paths
}
