package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CalResult struct {
	XMLName xml.Name
	Body    struct {
			XMLName     xml.Name
			Addresponse struct {
				XMLName xml.Name
				Return  []string `xml:"AddResult"`
			} `xml:"AddResponse"`
	}
}

func main() {
	// wsdl service url
	// url := fmt.Sprintf("%s",
	// 	"http://www.dneonline.com/calculator.asmx",
	// )
	url := "http://www.dneonline.com/calculator.asmx"

	// payload
	payload := []byte(strings.TrimSpace(`
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:tem="http://tempuri.org/">
   <soap:Header/>
   <soap:Body>
      <tem:Add>
         <tem:intA>12</tem:intA>
         <tem:intB>4</tem:intB>
      </tem:Add>
   </soap:Body>
</soap:Envelope>`,
	))

	httpMethod := "POST"

	// soap action
	soapAction := "urn:Add"

	// authorization credentials
	// username := "admin"
	// password := "admin"

	log.Println("-> Preparing the request")

	// prepare the request
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}

	// set the content type header, as well as the oter required headers
	req.Header.Set("Content-type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)
	// req.SetBasicAuth(username, password)

	// prepare the client request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	log.Println("-> Dispatching the request")

	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}
	log.Println("-> Retrieving and parsing the response")
	
	// read and parse the response body
	result := new(CalResult)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}
	fmt.Println(result)
	log.Println("-> Everything is good, printing users data")

	// print the users data
	addresult := result.Body.Addresponse.Return
	fmt.Println(strings.Join(addresult, ", "))
}