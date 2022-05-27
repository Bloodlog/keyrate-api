package client

import (
	"encoding/xml"
	"time"
)

// Envelope Response
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    *Body    `xml:"Body"`
}

type Body struct {
	Response Response `xml:"KeyRateXMLResponse"`
}

type Response struct {
	Result Result `xml:"KeyRateXMLResult"`
}
type Result struct {
	Rows []Rows `xml:"KeyRate"`
}

type Rows struct {
	KeyRates []KeyRates `xml:"KR"`
}

type KeyRates struct {
	Date string `xml:"DT" json:"date"`
	Rate string `xml:"Rate" json:"rate"`
}

type Clienter interface {
	Get(fromDate time.Time, toDate time.Time) ([]KeyRates, error)
}

func Get(fromDate time.Time, toDate time.Time) ([]KeyRates, error) {
	data, error := KeyRateByDate(fromDate, toDate)
	return data.Body.Response.Result.Rows[0].KeyRates, error
}

func KeyRateByDate(fromDate time.Time, toDate time.Time) (Envelope, error) {
	var request = GetKeyRateXMLRequest{}
	request.Xmlns = "http://web.cbr.ru/"
	layout := "2006-01-02"
	request.FromDate = fromDate.Format(layout)
	request.ToDate = toDate.Format(layout)

	rawXmlData := SoapCall("https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx", request)

	var data Envelope

	return data, xml.Unmarshal([]byte(rawXmlData), &data)
}
