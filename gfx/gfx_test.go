// go test github.com/nathany/mantle/gfx
package gfx_test

import (
    "github.com/nathany/mantle/core/gl"
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

func (s *S) TestGetError(c *C) {
    fakerc := new(FakeContext)

    err := gfx.GetErrors(fakerc)
    c.Check(err, IsNil)
}

type FakeContext struct{}

func (_ FakeContext) GetError() gl.ErrorCode {
    return gl.NoError
}
