package basic

import (
	"errors"
	"log"
	"os"
	"strings"
)

type IHashOperator interface {
	HandleNewHash(hash string)
	AddActions(action IResultAction)
}

type BasicHashOperator struct {
	Url     string
	Actions []IResultAction
}

func (x BasicHashOperator) HandleNewHash(hash string) {
	saved_hash := getSavedHash(x.Url)
	resultInfo := ResultInfo{Filepath: getFilenameByPagename(x.Url), Hash: hash, Url: x.Url}

	if hash != saved_hash {
		saveHash(hash, x.Url)

		if saved_hash == "" {
			for _, element := range x.Actions {
				element.OnFirstHash(resultInfo)
			}
		} else {
			for _, element := range x.Actions {
				element.OnHashChanged(resultInfo)
			}
		}

	} else {
		for _, element := range x.Actions {
			element.OnHashUnchanged(resultInfo)
		}
	}
}

func (x BasicHashOperator) AddActions(action IResultAction) {
	x.Actions = append(x.Actions, action)
}

func getFilenameByPagename(pagename string) string {
	filename := pagename + ".txt"
	filename = strings.ReplaceAll(filename, "\\", "_")
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, ":", "_")
	return filename
}

func getSavedHash(page_name string) string {
	filename := getFilenameByPagename(page_name)

	if !isFileExists(filename) {
		return ""
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func isFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

func saveHash(hash, pagename string) {
	filename := getFilenameByPagename(pagename)
	err := os.WriteFile(filename, []byte(hash), 0644)

	if err != nil {
		log.Fatal(err)
	}
}
