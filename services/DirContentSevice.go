package services

import (
	"os"
	"path/filepath"
	"time"

	"github.com/PickupModifiedFiles/models"
)

// FindModedFiles returns modified file's FileInfo
func FindModedFiles(past models.DirContent, current models.DirContent) []models.FileInfo {
	pastContentFiles := getFileNames(past.Contents)
	currentContentFiles := getFileNames(current.Contents)

	var matches []string
	for _, val := range currentContentFiles {
		if strArrayContains(pastContentFiles, val) {
			matches = append(matches, val)
		}
	}

	var result []models.FileInfo
	for _, val := range matches {
		pastItem := getFileInfo(past.Contents, val)
		currentItem := getFileInfo(current.Contents, val)

		if pastItem.ModTime != currentItem.ModTime {
			result = append(result, currentItem)
		}
	}

	return result
}

// FindNewFiles returns different file past from current.
func FindNewFiles(past models.DirContent, current models.DirContent) []models.FileInfo {
	pastContentFiles := getFileNames(past.Contents)
	currentContentFiles := getFileNames(current.Contents)

	var addedFileNames []string
	for _, val := range currentContentFiles {
		if !strArrayContains(pastContentFiles, val) {
			addedFileNames = append(addedFileNames, val)
		}
	}

	var result []models.FileInfo
	for _, val := range addedFileNames {
		fileinfo := getFileInfo(current.Contents, val)
		result = append(result, fileinfo)
	}

	return result
}

// genDirContent gens JSON data.
func genDirContent(paths []string) models.DirContent {
	cd, _ := os.Getwd()
	var contents []models.FileInfo
	for _, item := range paths {
		path := filepath.Join(cd, item)
		file, _ := os.Open(path)
		finfo, _ := file.Stat()
		defer file.Close()

		contents = append(contents, models.FileInfo{Name: item, ModTime: finfo.ModTime().String()})
	}

	dirCOntent := models.DirContent{Root: cd, LogDate: time.Now().String(), Contents: contents}
	return dirCOntent
}

// getFileNames returns file names array.
func getFileNames(s []models.FileInfo) []string {
	var result []string
	for _, fileinfo := range s {
		result = append(result, fileinfo.Name)
	}

	return result
}

// getFileInfo returns FileInfo is matches with name.
func getFileInfo(s []models.FileInfo, name string) models.FileInfo {
	for _, val := range s {
		if val.Name == name {
			return val
		}
	}

	panic("No match FileInfo")
}

// strArrayContains returns s conatains i or not.
func strArrayContains(s []string, i string) bool {
	for _, val := range s {
		if val == i {
			return true
		}
	}

	return false
}
