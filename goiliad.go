package goiliad

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	accountURL = "https://www.iliad.it/account/"
	keys       = []string{"calls", "sms", "data", "mms"}
)

// CreateCredentials allows to set user e password into a login request
func CreateCredentials(u, p string) url.Values {
	data := url.Values{}
	data.Set("login-ident", u)
	data.Add("login-pwd", p)
	return data
}

// PerformRequest allows to send the login request
//TODO: Write test
func PerformRequest(r *http.Request) (*http.Response, error) {
	client := &http.Client{}
	rsp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// CreateRequest allows to create a new login request
func CreateRequest(d url.Values, c *http.Cookie) (*http.Request, error) {
	req, err := http.NewRequest("POST", accountURL, strings.NewReader(d.Encode()))
	if err != nil {
		//TODO: Write test
		return nil, err
	}

	req.AddCookie(c)
	req.Header.Add("authority", "www.iliad.it")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("origin", "https://www.iliad.it")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("dnt", "1")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	return req, nil
}

// ReadResponse allows to read the response and to return a string version of it
func ReadResponse(r *http.Response) (string, error) {
	responseBody := r.Body

	bodyByte, err := ioutil.ReadAll(responseBody)
	if err != nil {
		//TODO: Write test
		defer responseBody.Close()
		return "", err
	}
	return string(bodyByte), nil
}

// GetInitialCookie retrieve the sessionCookie which allows to perform login
func GetInitialCookie() (*http.Cookie, error) {
	rs, err := http.Get(accountURL)
	if rs != nil {
		defer rs.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	cookies := rs.Cookies()
	return cookies[0], nil
}

// PageScraper extracts value from the body html page
func PageScraper(html string) (map[string]string, error) {
	var resultMap map[string]string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		//TODO: Write test
		return resultMap, err
	}

	conso := doc.Find(".conso__text")

	if conso.Size() < 4 {
		return resultMap, fmt.Errorf("Size doesn't match, found %d but need 4", conso.Size())
	}

	resultMap = make(map[string]string)

	for i, key := range keys {
		node := conso.Eq(i)
		resultMap[key] = node.Find("span").First().Text()
	}

	return resultMap, nil
}
