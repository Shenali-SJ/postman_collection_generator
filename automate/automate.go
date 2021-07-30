package automate

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
)

type Item struct {
	Item []APICall `json:"item"`
}

type APICall struct {
	Name string       `json:"name"`
	Request *Request  `json:"request"`
	Response []string `json:"response"`
}

type Request struct {
	Method string   `json:"method"`
	Header []Header `json:"header"`
	Body   Body     `json:"body,omitempty"`
	Url    URL      `json:"url"`
}

type Header struct {
	Key string `json:"key,omitempty"`
	Value []string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
}

type Body struct {
	Mode string `json:"mode,omitempty"`
	Raw string `json:"raw"`
}

type URL struct {
	Raw string      `json:"raw"`
	Protocol string `json:"protocol"`
	Host string     `json:"host"`
	Port string     `json:"port,omitempty"`
	Query []Query   `json:"query,omitempty"`
	Path []string   `json:"path,omitempty"`
}

type Query struct {
	Key string `json:"key"`
	Value []string `json:"value"`
}

// CreateCollection generates a postman collection by extracting the HTTP request data from the request
func CreateCollection(c echo.Context, runningPort string, fileName string) {
	name := c.Request().URL.RawQuery
	method := c.Request().Method
	protocol := c.Request().URL.Scheme
	host := c.Request().Host

	header := c.Request().Header
	var headers []Header
	for key, value := range header {
		headerValue := Header{
			Key:   key,
			Value: value,
		}
		headers  = append(headers, headerValue)
	}

	var body string

	port := runningPort
	path := c.Request().URL.Path
	pathParams := []string{path}

	queryMap := c.Request().URL.Query()
	var query []Query
	for key, value := range queryMap {
		queryValue := Query{
			Key:   key,
			Value: value,
		}
		query = append(query, queryValue)
	}

	response := []string{}

	if method == "POST" || method == "PUT" || method == "PATCH" {

		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request().Body)
		body = buf.String()

	}

	apiCall := APICall{
		Name:     name,
		Request:  &Request{
			Method: method,
			Header: headers,
			Body: Body{
				Raw:  body,
			},
			Url:    URL{
				Raw:      name,
				Protocol: protocol,
				Host:     host,
				Port:     port,
				Query: query,
				Path:  pathParams,
			},
		},
		Response: response,
	}

	apiCalls := []APICall{apiCall}
	item := Item{Item: apiCalls}

	// writes data to the JSON file
	file, _ := json.MarshalIndent(item, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)

}



