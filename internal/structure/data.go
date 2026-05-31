package structure

import (
	"encoding/xml"
	"time"
)

//raw XML data structure

type TraceabilityData struct {
	XMLName      xml.Name `xml:"TraceabilityData"`  //
	Text         string   `xml:",chardata"`         //unimportant??
	Version      string   `xml:"version,attr"`      //traciability version data?
	Line         string   `xml:"line,attr"`         //linka
	Station      string   `xml:"station,attr"`      //specigic mounting station
	Job          string   `xml:"job,attr"`          //program name - eg. "Valor\Line 1\Multi\M1468_PB3057700_I_BOT"
	BoardID      string   `xml:"boardID,attr"`      //serial number
	BoardName    string   `xml:"boardName,attr"`    //ommit
	DateBegin    string   `xml:"dateBegin,attr"`    //prod start? 	2024040414275301
	DateComplete string   `xml:"dateComplete,attr"` //prod end?		2024040414291701
	ErrorLabel   string   `xml:"errorLabel,attr"`   //??
	Setup        string   `xml:"setup,attr"`        //setup plan - multi / solo
	Panel        struct {
		Text string `xml:",chardata"` //unimportant??
		Name string `xml:"name,attr"` //project name? eg "1.3.PB3057700"
		Omit string `xml:"omit,attr"` //txt boolean
		Ref  []struct {
			Text   string   `xml:",chardata"` //unimportant??
			ID     string   `xml:"id,attr"`   //ID which references to HU / Charge
			RefDes []string `xml:"RefDes"`    //refdes belonging to ID/HU
		} `xml:"ref"`
	} `xml:"panel"`
	Charge []struct {
		Text     string `xml:",chardata"`     //unimportant??
		ID       string `xml:"id,attr"`       //references ID from pannel - list of refdes
		Comp     string `xml:"comp,attr"`     //CGCK PN?
		Barc1    string `xml:"barc1,attr"`    //CGCK PN? identical with comp
		Barc2    string `xml:"barc2,attr"`    //component lotcode
		Barc3    string `xml:"barc3,attr"`    //unkown number value eg. 101500, 102500, 102000, 105000
		Barc4    string `xml:"barc4,attr"`    //empty
		Barc5    string `xml:"barc5,attr"`    //empty
		Barc6    string `xml:"barc6,attr"`    //HU
		ContSize string `xml:"contSize,attr"` //number 0?
		PackForm string `xml:"packForm,attr"` //CGCK component package marking eg. "Valor\DIS0603R.45-2", "Valor\SOT23-2", "Valor\MELFSTD-4"
		TableID  string `xml:"tableID,attr"`  //unkown table reference eg. "DZS00000377", "BXW00000127, "DZS00000377"
		Operator string `xml:"operator,attr"` //unfortunately empty
		Loc      string `xml:"loc,attr"`      //unkown number
		Track    string `xml:"track,attr"`    //unkown number
		Div      string `xml:"div,attr"`      //unkown number
		MTCid    string `xml:"MTCid,attr"`    //unkown number
		MTCloc   string `xml:"MTCloc,attr"`   //unkown number
		MTCtower string `xml:"MTCtower,attr"` //unkown number
		MTCtray  string `xml:"MTCtray,attr"`  //unkown number
		MTCdiv   string `xml:"MTCdiv,attr"`   //unkown number
		Changed  string `xml:"changed,attr"`  //txt boolean eg. "fAlSe" yes... capital A
	} `xml:"charge"`
}

// sn, hu, ref-list, time-finish, PN, lotcode, project, project rev

//trace: ID, SN, HU, ref-list, time-finish
//board: ID, SN (unique), project, project rev
//comp: ID, HU (unique), PN, lotcode

type FormatedXMLdata struct {
	SerialNumber string
	Project      string
	Revision     string
	Component    []struct {
		PartNumber   string
		HandlingUnit string
		LotCode      string
		PlacedIn     time.Time
		RefereceList []string
	}
}
