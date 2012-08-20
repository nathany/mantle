// go test github.com/nathany/mantle/gfx
package gfx

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
