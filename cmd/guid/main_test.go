package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TODO add some more etkeys to test
var TestGuids = map[string]string{
	"991203671241403262": "2CF3C58D435FEE819AE0826BF7A5FEFE",
}

func TestCalculateGuid(t *testing.T) {
	ts := httptest.NewServer(getMux())
	defer ts.Close()

	for k, g := range TestGuids {
		res, err := http.Post(ts.URL, "text/plain", strings.NewReader(k))
		if err != nil {
			t.Fatalf("Failed requesting guid for key %s: %s", k, err)
		}

		defer res.Body.Close()
		if body, err := ioutil.ReadAll(res.Body); err != nil {
			t.Fatalf("Failed reading response body: %s", err)
		} else {
			if g != string(body) {
				t.Fatalf("Expected %s, received %s", g, string(body))
			}

			t.Logf("Requested guid for key %s: %s", k, g)
		}
	}
}
