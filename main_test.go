package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateCredentials(t *testing.T) {
	data := createCredentials("user", "password")

	if data.Get("login-ident") != "user" || data.Get("login-pwd") != "password" {
		t.Errorf("User and password not setted on the Values structure")
	}
}

func TestGetInitialCookie(t *testing.T) {
	got, err := getInitialCookie()
	if err != nil {
		t.Error(err)
	}

	expected := "ACCOUNT_SESSID"

	if got.Name != expected {
		t.Errorf("got %s, expected %s", got, expected)
	}
}

func TestCreateRequest(t *testing.T) {
	dataFake := url.Values{}
	dataFake.Set("test1", "test1")
	dataFake.Add("test2", "test2")

	req, err := createRequest(dataFake, &http.Cookie{})
	if err != nil {
		t.Error(err)
	}

	got := req.Header.Get("authority")
	expected := "www.iliad.it"

	if got != expected {
		t.Errorf("got %s, expected %s", got, expected)
	}
}

func TestReadResponse(t *testing.T) {

	expected := "Fake Result"

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>"+expected+"</body></html>")
	}

	req := httptest.NewRequest("POST", "http://test.com/account", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	hresponse := w.Result()

	got, err := readResponse(hresponse)

	if err != nil {
		t.Error(err)
	}

	if ok := strings.Contains(got, expected); !ok {
		t.Errorf("Got %s, expected %s", got, expected)
	}

}
