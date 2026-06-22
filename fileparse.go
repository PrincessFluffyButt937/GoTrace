package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type UserInput struct {
	data map[string]bool
}

func (cfg *apiConfig) SerialNumberScan(rootFilePath string, input *UserInput) error {
	//valid directory check
	dir, err := os.Stat(rootFilePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("Filepath does not exist:\n%s\n", err.Error())
	}
	if !dir.IsDir() {
		return fmt.Errorf("Filepath is not a directory:\n%s\n", err.Error())
	}
	//open directory
	openDir, err := os.Open(rootFilePath)
	if err != nil {
		return err
	}
	defer openDir.Close()

	for {
		files, err := openDir.ReadDir(cfg.parseLimit)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Folder parsing failed: %s", err.Error())
		}
		for _, file := range files {
			FileName := file.Name()
			tempFilePath := filepath.Join(rootFilePath, FileName)
			rawXML, err := parseXML(tempFilePath)
			if err != nil {
				//loggging
				continue
			}
			//SN check against user input
			if _, ok := input.data[rawXML.BoardID]; !ok {
				continue
			}
			input.data[rawXML.BoardID] = true

			xmlData, err := formatTraceData(rawXML, FileName)
			if err != nil {
				//logging
				continue
			}
			if err := cfg.saveToDatabase(xmlData); err != nil {
				//logging
				continue
			}
		}
	}
	return nil
}
