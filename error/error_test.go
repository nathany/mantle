// go test github.com/nathany/gl3/error
package error

import (
  . "launchpad.net/gocheck"
  "testing"
)

// Hook gocheck into the gotest runner
func TestAll(t *testing.T) {
  TestingT(t)
}

type S struct{}

var _ = Suite(&S{})

func (s *S) TestHelloWorld(c *C) {
  c.Check(42, Equals, "42")
}
