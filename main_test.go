package main

import (
	"fmt"
	"testing"
)

func TestProjecDataParse(t *testing.T) {
	//test data
	data1 := projectDataInput{
		fileName:    "1513028976-PB30577001I-1-NO-PCB-BARCODE2545-20240405063324.XML",
		programName: "M1468_PB3057700_I_BOT",
		panelName:   "1.4.PB3057700",
	}
	r_data1 := projectData{
		Project:  "PB3057700",
		Revision: "I",
		Side:     "BOT",
	}

	data2 := projectDataInput{
		fileName:    "I am random invalid string",
		programName: "I_am_Invalid_too",
		panelName:   "1.4.PB3057700",
	}
	r_data2 := projectData{
		Project:  "PB3057700",
		Revision: "N/A",
		Side:     "N/A",
	}

	data3 := projectDataInput{
		fileName:    "1513028976-PB30577001I-1-NO-PCB-BARCODE2545-20240405063324.XML",
		programName: "M1468_PB123_I_BOT",
		panelName:   "me.not.Working",
	}
	r_data3 := projectData{
		Project:  "PB3057700",
		Revision: "I",
		Side:     "N/A",
	}

	data4 := projectDataInput{
		fileName:    "Hello-PBthere-master-kenobi.XML",
		programName: "M1468_PB3057700_I_NOT",
		panelName:   "1.2.3",
	}
	r_data4 := projectData{
		Project:  "PB3057700",
		Revision: "I",
		Side:     "N/A",
	}

	data5 := projectDataInput{
		fileName:    "Hello-PBthere-master-kenobi.XML",
		programName: "M1468_PB305700_I_TOP", //should not work
		panelName:   "1.2.3",
	}
	r_data5 := projectData{}
	data6 := projectDataInput{
		fileName:    "1513028976-PB30577001I-1-NO-PCB-BARCODE2545-20240405063324.XML",
		programName: "M1468_PB3057700_I_TOP",
		panelName:   "1.4.PB3057700",
	}
	r_data6 := projectData{
		Project:  "PB3057700",
		Revision: "I",
		Side:     "TOP",
	}

	//testing

	outData1, err := parseStringToProjectRevision(data1)
	if err != nil {
		fmt.Println("outData1 fail")
		fmt.Println("unexpected error raised")
		t.Fail()
	}
	if outData1 != r_data1 {
		fmt.Println("outData1 fail")
		fmt.Printf("Exp: %s, %s, %s\n", r_data1.Project, r_data1.Revision, r_data1.Side)
		fmt.Printf("Got: %s, %s, %s\n", outData1.Project, outData1.Revision, outData1.Side)
		t.Fail()
	}

	outData2, err := parseStringToProjectRevision(data2)
	if err != nil {
		fmt.Println("outData2 fail")
		fmt.Println("unexpected error raised")
		t.Fail()
	}
	if outData2 != r_data2 {
		fmt.Println("outData2 fail")
		fmt.Printf("Exp: %s, %s, %s\n", r_data2.Project, r_data2.Revision, r_data2.Side)
		fmt.Printf("Got: %s, %s, %s\n", outData2.Project, outData2.Revision, outData2.Side)
		t.Fail()
	}
	outData3, err := parseStringToProjectRevision(data3)
	if err != nil {
		fmt.Println("outData3 fail")
		fmt.Println("unexpected error raised")
		t.Fail()
	}
	if outData3 != r_data3 {
		fmt.Println("outData3 fail")
		fmt.Printf("Exp: %s, %s, %s\n", r_data3.Project, r_data3.Revision, r_data3.Side)
		fmt.Printf("Got: %s, %s, %s\n", outData3.Project, outData3.Revision, outData3.Side)
		t.Fail()
	}

	outData4, err := parseStringToProjectRevision(data4)
	if err != nil {
		fmt.Println("outData4 fail")
		fmt.Println("unexpected error raised")
		t.Fail()
	}
	if outData4 != r_data4 {
		fmt.Println("outData4 fail")
		fmt.Printf("Exp: %s, %s, %s\n", r_data4.Project, r_data4.Revision, r_data4.Side)
		fmt.Printf("Got: %s, %s, %s\n", outData4.Project, outData4.Revision, outData4.Side)
		t.Fail()
	}

	outData5, err := parseStringToProjectRevision(data5)
	if err == nil {
		fmt.Println("outData5 fail")
		fmt.Println("error should be raised")
		t.Fail()
	}
	if outData5 != r_data5 {
		fmt.Println("outData5 fail")
		fmt.Println("Exp: empty data")
		fmt.Printf("Got: %s, %s, %s\n", outData5.Project, outData5.Revision, outData5.Side)
		t.Fail()
	}

	outData6, err := parseStringToProjectRevision(data6)
	if err != nil {
		fmt.Println("outData6 fail")
		fmt.Println("unexpected error raised")
		t.Fail()
	}
	if outData6 != r_data6 {
		fmt.Println("outData6 fail")
		fmt.Printf("Exp: %s, %s, %s\n", r_data6.Project, r_data6.Revision, r_data6.Side)
		fmt.Printf("Got: %s, %s, %s\n", outData6.Project, outData6.Revision, outData6.Side)
		t.Fail()
	}
}
