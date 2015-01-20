package guid

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&GuidSuite{})

type GuidSuite struct{}

// TODO add some more etkeys to test
var TestGuids = map[string]string{
	"991203671241403262": "2CF3C58D435FEE819AE0826BF7A5FEFE",
}

func (g *GuidSuite) TestGuid(c *C) {
	for k, g := range TestGuids {
		guid, err := Calculate(k)

		c.Assert(guid, HasLen, 32)
		c.Assert(err, IsNil)
		c.Assert(guid, Equals, g)
	}
}
