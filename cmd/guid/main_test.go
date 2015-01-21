package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/twcclan/go-guid"
	. "gopkg.in/check.v1"
)

// TODO add some more etkeys to test
var TestGuids = map[string]string{
	"991203671241403262": "2CF3C58D435FEE819AE0826BF7A5FEFE",
}

var _ = Suite(&HttpSuite{})

type HttpSuite struct {
	ts *httptest.Server
}

func (h *HttpSuite) SetUpSuite(c *C) {
	h.ts = httptest.NewServer(getMux())
}

func (h *HttpSuite) TearDownSuite(c *C) {
	h.ts = httptest.NewServer(getMux())
}

func Test(t *testing.T) { TestingT(t) }

func (h *HttpSuite) TestCalculateGuid(c *C) {
	for k, g := range TestGuids {
		res, err := http.Post(h.ts.URL, "text/plain", strings.NewReader(k))
		c.Assert(err, IsNil, Commentf("Failed requesting guid for key %s: %s", k, err))

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		c.Assert(err, IsNil, Commentf("Failed reading response body: %s", err))
		c.Assert(g, Equals, string(body))
	}
}

func (h *HttpSuite) testInvalidEtKey(key string, c *C) {
	res, err := http.Post(h.ts.URL, "text/plain", strings.NewReader(key))

	c.Assert(err, IsNil)
	c.Assert(res.StatusCode, Equals, http.StatusBadRequest)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	c.Assert(err, IsNil)
	c.Assert(string(body), Equals, fmt.Sprintf("Failed to calculate guid: %s", guid.ErrInvalidEtKey))
}

func (h *HttpSuite) TestBadRequests(c *C) {
	// test too short
	h.testInvalidEtKey("too short", c)

	// test too long
	h.testInvalidEtKey("way too long way too long way too long", c)
}
