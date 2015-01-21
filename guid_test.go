package guid

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
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

var guid string
var err error

func (g *GuidSuite) BenchmarkGuid(c *C) {
	for i := 0; i < c.N; i++ {
		key := fmt.Sprintf("%018d", i)

		//c.ResetTimer()
		guid, err = Calculate(key)
	}
}
