package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/PrincessFluffyButt937/GoTrace/internal/structure"
)

func (cfg *apiConfig) Scanfolder(rootFilePath string) error {
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

	//Parses contents of folder (file information)
	//Full logic yet to be implemented

	for {
		files, err := openDir.ReadDir(cfg.parseLimit)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Folder parsing failed: %s", err.Error())
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		fmt.Println("------")
	}

	return nil
}

func parseStringToTime(timeString string) (time.Time, error) {
	//to be implemented
	return time.Now(), nil
}

type projectData struct {
	Project  string
	Revision string
}

func parseStringToProjectRevision(projectRev string) (projectData, error) {
	//to be implemented
	return projectData{}, nil
}

// don't forget to change signature
func ParseXML(XMLFilePath string) (structure.TraceabilityData, error) {
	//to be implemented
	return structure.TraceabilityData{}, nil
}
