package cmd

import (
	"testing"
)

var containsTests = []struct {
	cidr     string
	ip       string
	expected bool
}{
	{"192.168.0.0/16", "192.168.2.4", true},
	{"192.168.0.0/16", "192.168.2.5", true},
	{"192.168.0.0/24", "192.168.2.4", false},
	{"192.168.0.0/24", "192.168.2.5", false},
	{"192.168.0.0/16", "192.167.2.5", false},
	{"192.168.0.0/16", "192", false},
	{"192.168.0.", "192", false},
}

func TestContains(t *testing.T) {
	for _, test := range containsTests {
		if got, _ := Contains(test.cidr, test.ip); got != test.expected {
			t.Errorf("Output %t not equal to expected %t", got, test.expected)
		}
	}
}
