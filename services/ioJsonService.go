package services

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PickupAndCopyModifiedFIles/models"
)

// OutputJSON outputs Json file
func OutputJSON(dirContent models.DirContent) {
	jsonBytes, err := json.Marshal(dirContent)
	if err != nil {
		panic(err)
	}

	logfile, err := os.Create(`DirLog.json`)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	logfile.WriteString(string(jsonBytes))
}

// ReadJSON read Json file
func ReadJSON() models.DirContent {
	jsonFile, err := ioutil.ReadFile("DirLog.json")
	if err != nil {
		panic(err)
	}

	var dirContent models.DirContent
	if err := json.Unmarshal(jsonFile, &dirContent); err != nil {
		panic(err)
	}

	return dirContent
}
