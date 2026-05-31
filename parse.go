package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
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

//Time parsing start

func parseStringToTime(timeString string) (time.Time, error) {
	//example	2024040414291701
	//format	yyyyMMDDhhmmssxx (year, month, day, hour, minute, second, miliseconds?)
	layout := "20060102150405"
	if len(timeString) != 16 {
		return time.Time{}, fmt.Errorf("Unkonwn time format:\nexp: \"yyyMMDDhhmmssxx (year, month, day, hour, minute, second, miliseconds)\"\ngot: %s", timeString)
	}
	timeStringSnip := timeString[:14]
	parsedTime, err := time.Parse(layout, timeStringSnip)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

//Time parsing end

//Project Data extraction start

type projectDataInput struct {
	fileName    string
	programName string
	panelName   string
	//different option in case one method fails.
}

type projectData struct {
	Project  string
	Revision string
	Side     string
}

func parseStringToProjectRevision(dataInput projectDataInput) (projectData, error) {
	data := projectData{}
	//data extraction from program name:
	programSplit := strings.Split(dataInput.programName, "_")
	if len(programSplit) == 4 {
		project := programSplit[1]
		revision := programSplit[2]
		side := programSplit[3]
		if (len(project) == 9 && strings.HasPrefix(strings.ToLower(project), "pb")) &&
			len(revision) == 1 {
			data.Project = project
			data.Revision = revision
			if len(side) == 3 && (strings.ToLower(side) == "top" || strings.ToLower(side) == "bot") {
				data.Side = side
			} else {
				data.Side = "N/A"
			}
			return data, nil
		}
	}
	//data extration from file name:
	filenameSplit := strings.Split(dataInput.fileName, "-")
	if len(filenameSplit) == 7 {
		projectAndRev := filenameSplit[1]
		if len(projectAndRev) == 10 && strings.HasPrefix(strings.ToLower(projectAndRev), "pb") {
			project := projectAndRev[:9]
			revison := projectAndRev[9:]
			data.Project = project
			data.Revision = revison
			data.Side = "N/A"
			return data, nil
		}
	}
	//data extraction from panel name:
	panelNameSplit := strings.Split(dataInput.panelName, ".")
	if len(panelNameSplit) == 3 {
		project := panelNameSplit[2]
		if len(project) == 9 && strings.HasPrefix(strings.ToLower(project), "pb") {
			data.Project = project
			data.Revision = "N/A"
			data.Side = "N/A"
			return data, nil
		}
	}
	//no compatible data found
	return projectData{}, fmt.Errorf("No compatible project data can be extracted:\n FileName: %s\nProgName: %s\nPaneName: %s\n", dataInput.fileName, dataInput.programName, dataInput.panelName)
}

//Project Data extraction end

func parseXML(XMLFilePath string) (structure.TraceabilityData, error) {
	file, err := os.Open(XMLFilePath)
	if err != nil {
		return structure.TraceabilityData{}, err
	}
	defer file.Close()
	data := structure.TraceabilityData{}

	file_body, err := io.ReadAll(file)
	if err != nil {
		return structure.TraceabilityData{}, err
	}

	if err := xml.Unmarshal(file_body, &data); err != nil {
		return structure.TraceabilityData{}, err
	}
	return data, nil
}

func formatTraceData(inputData structure.TraceabilityData) (structure.FormatedXMLdata, error) {
	//to be implemented
	return structure.FormatedXMLdata{}, nil
}
