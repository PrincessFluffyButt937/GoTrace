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

func TestTimeParse(t *testing.T) {
	sample1 := "2024040506332203"
	exp1 := "time.Date(2024, time.April, 5, 6, 33, 22, 0, time.UTC)"

	parsedTime1, err := parseStringToTime(sample1)
	if err != nil {
		println(err.Error())
		t.Fail()
	}
	if parsedTime1.GoString() != exp1 {
		t.Fail()
		fmt.Printf("TestTimeParse, parsedTime1 mismatch:\nexp: %s\ngot: %s\n", exp1, parsedTime1)
	}

	sample2 := "2021050506332403"
	exp2 := "time.Date(2021, time.May, 5, 6, 33, 24, 0, time.UTC)"

	parsedTime2, err := parseStringToTime(sample2)
	if err != nil {
		println(err.Error())
		t.Fail()
	}
	if parsedTime2.GoString() != exp2 {
		t.Fail()
		fmt.Printf("TestTimeParse, parsedTime2 mismatch:\nexp: %s\ngot: %s\n", exp2, parsedTime2)
	}

	sample3 := "2021"

	if _, err := parseStringToTime(sample3); err == nil {
		t.Fail()
		fmt.Println("TestTimeParse, sample3: error should have been raised")
	}

	sample4 := "20Hello506332203"

	if _, err := parseStringToTime(sample4); err == nil {
		t.Fail()
		fmt.Println("TestTimeParse, sample4: error should have been raised")
	}
}

func TestSortKEyGeneration(t *testing.T) {
	ref1 := "R1"
	exp1 := "000001"

	ref2 := "CD146"
	exp2 := "000146"

	ref3 := "TX123456"
	exp3 := "123456"

	key1 := refSortKey(ref1)
	t.Attr(exp1, key1)

	key2 := refSortKey(ref2)
	t.Attr(exp2, key2)

	key3 := refSortKey(ref3)
	t.Attr(exp3, key3)
}

func TestRefToString(t *testing.T) {
	rList1 := []string{"R10", "R2", "R12", "R3"}
	exp1 := "R2, R3, R10, R12"

	rList2 := []string{"R3"}
	exp2 := "R3"

	rList3 := []string{"R1", "R2", "R12", "R3", "R26", "R123456"}
	exp3 := "R1, R2, R3, R12, R26, R123456"

	rList4 := []string{"R1", "R2", "R12", "R3", "R26", "R123456"}
	exp4 := "R1, R2, R12, R3, R26, R123456"

	rSring1 := refListToString(rList1, true)
	if rSring1 != exp1 {
		t.Fail()
		fmt.Printf("TestRefToString / exp1 - rString1 mismatch:\nexp: %s\ngot: %s\n", exp1, rSring1)
	}

	rSring2 := refListToString(rList2, true)
	if rSring2 != exp2 {
		t.Fail()
		fmt.Printf("TestRefToString / exp2 - rString2 mismatch:\nexp: %s\ngot: %s\n", exp2, rSring2)
	}

	rSring3 := refListToString(rList3, true)
	if rSring3 != exp3 {
		t.Fail()
		fmt.Printf("TestRefToString / exp3 - rString3 mismatch:\nexp: %s\ngot: %s\n", exp3, rSring3)
	}

	rSring4 := refListToString(rList4, false)
	if rSring4 != exp4 {
		t.Fail()
		fmt.Printf("TestRefToString / exp4 - rString4 mismatch:\nexp: %s\ngot: %s\n", exp4, rSring4)
	}
}
