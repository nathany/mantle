// go test github.com/nathany/mantle/gfx
package gfx_test

import (
  "github.com/nathany/mantle/gfx"
  . "launchpad.net/gocheck"
  "testing"
)

// Hook gocheck into the gotest runner
func TestAll(t *testing.T) {
  TestingT(t)
}

type S struct{}

var _ = Suite(&S{})

func (s *S) TestErrorCodeToString(c *C) {
  str := gfx.ErrorCodeToString(0x500)
  c.Check(str, Equals, "Invalid Enum (0x500)")
}
