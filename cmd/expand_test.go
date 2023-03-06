package cmd

import (
	"testing"
)

type addExpandTest struct {
	cidr     string
	expected string
}

var addExpandTests = []addExpandTest{
	{"192.168.0.0/32", "{\"cidr\":\"192.168.0.0/32\",\"ips\":[[\"192.168.0.0\"]]}"},
}

func TestExpand(t *testing.T) {
	for _, test := range addExpandTests {
		if got, _ := Expand(test.cidr); got != test.expected {
			t.Errorf("Output %q not equal to expected %q", got, test.expected)
		}
	}
}
