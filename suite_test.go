package gl3

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
