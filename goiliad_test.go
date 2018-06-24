package goiliad

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateCredentials(t *testing.T) {
	data := CreateCredentials("user", "password")

	if data.Get("login-ident") != "user" || data.Get("login-pwd") != "password" {
		t.Errorf("User and password not setted on the Values structure")
	}
}

func TestGetInitialCookie(t *testing.T) {
	got, err := GetInitialCookie()
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

	req, err := CreateRequest(dataFake, &http.Cookie{})
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

	got, err := ReadResponse(hresponse)

	if err != nil {
		t.Error(err)
	}

	if ok := strings.Contains(got, expected); !ok {
		t.Errorf("Got %s, expected %s", got, expected)
	}

}

func TestPageScraper(t *testing.T) {
	htmlHappyTest := `
<!doctype html>
<html lang="it">
  <head>
    <meta charset="utf-8">
    <title>Benvenuto in iliad</title>
  </head>
  <body>
    <div class="remodal-bg">
      <div id="container" canvas="container">
        <div id="page-container">
	  <div class="page-container main">
	    <div class="page-content">
	      <div class="page p-conso">
	        <h2>
		  <span class="bold">Offerta iliad</span> - Credito : <b class="red">0.00€</b>
		  <div class="toggle-conso">
		    <a href="#" data-target="local" class="selected">In Italia</a>
		    <a href="#" data-target="roaming">Estero</a>
		  </div>
		</h2>
		<div class="conso-infos conso-local">
		  <div class="grid-l conso__grid">
		    <div class="grid-c w-4 w-tablet-4">
		      <div class="conso__content">
		        <div class="conso__text">
			Chiamate: <span class="red">0s</span><br>
			Consumi voce: <span class="red">0.00€</span>
		      </div>
		    </div>
		  </div>
		  <div class="grid-c w-4 w-tablet-4">
		    <div class="conso__content">
		      <div class="conso__text"><span class="red">0 SMS</span><br>
		      SMS extra: <span class="red">0.00€</span>
		    </div>
		  </div>
		</div>
		<div class="grid-l conso__grid">
		  <div class="grid-c w-4 w-tablet-4">
		    <div class="conso__content">
		      <div class="conso__text">
		        <span class="red">62,76MB</span> / 30GB<br>
			Consumi Dati: <span class="red">0.00€</span>
		      </div>
		    </div>
		  <div class="grid-c w-4 w-tablet-4">
		    <div class="conso__content">
		      <div class="conso__text">
		        <span class="red">0 MMS<br></span>
			Consumi MMS: <span class="red">0.00€</span>
		      </div>
		    </div>
		  </div>
		</div>
	      </div>
            </div>
	  </div>
	</div>
      </div>
    </div>
  </div>
  </body>
</html>`

	t.Run("Happy Path", func(t *testing.T) {
		got, err := PageScraper(htmlHappyTest)
		if err != nil {
			t.Error(err)
		}

		checkKeys(got, t)

		checkValues(got, t)

	})

	htmlWrongNumberOfDatas := `
<!doctype html>
<html lang="it">
  <head>
    <meta charset="utf-8">
    <title>Benvenuto in iliad</title>
  </head>
  <body>
    <div class="remodal-bg">
      <div id="container" canvas="container">
        <div id="page-container">
	  <div class="page-container main">
	    <div class="page-content">
	      <div class="page p-conso">
	        <h2>
		  <span class="bold">Offerta iliad</span> - Credito : <b class="red">0.00€</b>
		  <div class="toggle-conso">
		    <a href="#" data-target="local" class="selected">In Italia</a>
		    <a href="#" data-target="roaming">Estero</a>
		  </div>
		</h2>
		<div class="conso-infos conso-local">
		  <div class="grid-l conso__grid">
		    <div class="grid-c w-4 w-tablet-4">
		      <div class="conso__content">
		        <div class="conso__text">
			Chiamate: <span class="red">0s</span><br>
			Consumi voce: <span class="red">0.00€</span>
		      </div>
		    </div>
		  </div>
		<div class="grid-l conso__grid">
		  <div class="grid-c w-4 w-tablet-4">
		    <div class="conso__content">
		      <div class="conso__text">
		        <span class="red">62,76MB</span> / 30GB<br>
			Consumi Dati: <span class="red">0.00€</span>
		      </div>
		    </div>
		  <div class="grid-c w-4 w-tablet-4">
		    <div class="conso__content">
		      <div class="conso__text">
		        <span class="red">0 MMS<br></span>
			Consumi MMS: <span class="red">0.00€</span>
		      </div>
		    </div>
		  </div>
		</div>
	      </div>
            </div>
	  </div>
	</div>
      </div>
    </div>
  </div>
  </body>
</html>`

	t.Run("Wrong Number of datas", func(t *testing.T) {
		_, err := PageScraper(htmlWrongNumberOfDatas)
		if err == nil {
			t.Error(err)
		}
	})

}

func checkKeys(g map[string]string, t *testing.T) {
	t.Helper()

	for _, key := range keys {
		if _, ok := g[key]; !ok {
			t.Errorf("No %s key are found on the result map", key)

		}
	}
}

func checkValues(g map[string]string, t *testing.T) {
	t.Helper()

	var fakeResults = []string{"0s", "0 SMS", "62,76MB", "0 MMS"}

	for i, key := range keys {
		if v := g[key]; v != fakeResults[i] {
			t.Errorf("The %s value is not %s but %s", key, fakeResults[i], v)
		}
	}
}
