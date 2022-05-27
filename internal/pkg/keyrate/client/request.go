package client

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client is an interface for testing a request object.
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// GetKeyRateXMLRequest Request
type GetKeyRateXMLRequest struct {
	XMLName  xml.Name `xml:"KeyRateXML"`
	Xmlns    string   `xml:"xmlns,attr"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

type SoapBody struct {
	XMLName xml.Name `xml:"soap12:Body"`
	Request interface{}
}

type SoapRoot struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Body    SoapBody
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
