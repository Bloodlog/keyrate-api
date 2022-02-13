package client

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Envelope Response
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body *Body `xml:"Body"`
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

// GetKeyRateXMLRequest Request
type GetKeyRateXMLRequest struct {
	XMLName xml.Name `xml:"KeyRateXML"`
	Xmlns string `xml:"xmlns,attr"`
	FromDate string `xml:"fromDate"`
	ToDate string `xml:"ToDate"`
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

type SoapBody struct {
	XMLName xml.Name `xml:"soap12:Body"`
	Request interface{}
}

type SoapRoot struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	Xsi    string `xml:"xmlns:xsi,attr"`
	Xsd    string `xml:"xmlns:xsd,attr"`
	Soap12 string `xml:"xmlns:soap12,attr"`
	Body   SoapBody
}

func SoapCall(service string, request interface{}) string {
	var root = SoapRoot{}
	root.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	root.Xsd = "http://www.w3.org/2001/XMLSchema"
	root.Soap12 = "http://www.w3.org/2003/05/soap-envelope"
	root.Body = SoapBody{}
	root.Body.Request = request

	out, _ := xml.MarshalIndent(&root, " ", "  ")

	body := string(out)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	response, err := client.Post(service, "application/soap+xml", bytes.NewBufferString(body))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	return strings.TrimSpace(string(content))
}
