// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package utils_test

import (
	"github.com/juju/testing"
	gc "gopkg.in/check.v1"

	"github.com/juju/utils"
)

var _ = gc.Suite(&sizeSuite{})

type sizeSuite struct {
	testing.IsolationSuite
}

func (*sizeSuite) TestParseSize(c *gc.C) {
	type test struct {
		in  string
		out uint64
		err string
	}
	tests := []test{{
		in:  "",
		err: `expected a non-negative number with optional multiplier suffix \(M/G/T/P\), got ""`,
	}, {
		in:  "-1",
		err: `expected a non-negative number with optional multiplier suffix \(M/G/T/P\), got "-1"`,
	}, {
		in:  "1MZ",
		err: `invalid multiplier suffix "MZ"`,
	}, {
		in:  "0",
		out: 0,
	}, {
		in:  "123",
		out: 123,
	}, {
		in:  "1M",
		out: 1,
	}, {
		in:  "0.5G",
		out: 512,
	}, {
		in:  "0.5GB",
		out: 512,
	}, {
		in:  "0.5GiB",
		out: 512,
	}, {
		in:  "0.5T",
		out: 524288,
	}, {
		in:  "0.5P",
		out: 536870912,
	}}
	for i, test := range tests {
		c.Logf("test %d: %+v", i, test)
		size, err := utils.ParseSize(test.in)
		if test.err != "" {
			c.Assert(err, gc.NotNil)
			c.Assert(err, gc.ErrorMatches, test.err)
		} else {
			c.Assert(err, gc.IsNil)
			c.Assert(size, gc.Equals, test.out)
		}
	}
}
