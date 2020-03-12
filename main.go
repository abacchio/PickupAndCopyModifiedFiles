package main

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/PickupAndCopyModifiedFIles/models"
	"github.com/PickupAndCopyModifiedFIles/services"
)

//
type Config struct {
	Logfile string   `toml:"logfilename"`
	Ignore  []string `toml:"ignorefile"`
}

func init() {

	var config Config
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

}

func main() {
	current := services.GenerateCurrentDirContent(".")

	var past models.DirContent
	if _, err := os.Open("DirLog.json"); err == nil {
		fmt.Println("******************** Read previous data. ********************")
		past = services.ReadJSON()
		fmt.Println("END")
		fmt.Println("******************** Start compare. *************************")
		if past.Root[3:] != current.Root[3:] {
			fmt.Println("Root directory has changed, it's unauthorized operation.\nRegenerate JSON file.")
			services.OutputJSON(current)
			return
		}

		addedFiles := services.FindNewFiles(past, current)
		fmt.Println("New added files are :")
		if addedFiles == nil {
			fmt.Println("Nothing")
		} else {
			for _, i := range addedFiles {
				fmt.Println(i)
			}
		}

		modifiedFiles := services.FindModedFiles(past, current)
		fmt.Println("Modified files are :")
		if modifiedFiles == nil {
			fmt.Println("Nothing")
		} else {
			for _, i := range modifiedFiles {
				fmt.Println(i)
			}
		}

		fmt.Println("END")

		fmt.Println("******************** Copy files to TransferDir. ********************")
		cd, _ := os.Getwd()
		var transferContent []models.FileInfo
		for _, modifiedFile := range modifiedFiles {
			transferContent = append(addedFiles, modifiedFile)
		}

		transferDirContent := models.DirContent{Root: cd, LogDate: time.Now().String(), Contents: transferContent}
		services.GenerateTransferDir(transferDirContent)

		t := time.Now()
		date := t.String()[:10]
		services.CopyToTransferDir(addedFiles, cd, date)
		services.CopyToTransferDir(modifiedFiles, cd, "date")
		fmt.Println("END")

	} else {
		fmt.Println("******** No previous log data, generate new log data. *********")
		fmt.Println("END")
	}

	services.OutputJSON(current)
}
