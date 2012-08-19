package gl3

import (
  . "launchpad.net/gocheck"
  "testing"
)

// Hook gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type SuiteErrors struct{}

var _ = Suite(&SuiteErrors{})

func (s *SuiteErrors) TestHelloWorld(c *C) {
  c.Check(42, Equals, "42")
}
