package service

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpTest(c *C) {
}

func (s *TestSuite) TestGetClient(c *C) {
	c.Assert(10, Equals, 10)
}
