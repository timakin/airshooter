package client_test

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *MySuite) SetUpTest(c *C) {
	s.dir = c.MkDir()
	// Use s.dir to prepare some data.
}

func (s *TestSuite) TestGetClient(c *C) {
	c.Assert(1, Equals, m.Client{})
}
