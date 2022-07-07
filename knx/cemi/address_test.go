// Copyright (c) 2022 mobilarte.
// Licensed under the MIT license which can be found in the LICENSE file.

package cemi

import (
	"testing"
)

// Test Individual Address
func Test_IndividualAddress(t *testing.T) {
	type Addr struct {
		Src   string
		Valid bool
	}

	var addrs = []Addr{
		{"1.2.3", true},
		{"1.3.255", true},
		{"1.3.0", true},
		{"75.235", true},
		{"65535", true},
		{"15.15.255", true},
		{"15.15.0", true},
		{"13057", true},
		{"1..0", false},
		{"15.15.", false},
		{" . .15", false},
		{"18.15.240", false},
		{"1.3.450", false},
		{"1.450", false},
		{"-2", false},
		{"400", false},
		{"-11.0.0", false},
	}

	for _, a := range addrs {
		_, err := NewIndividualAddrString(a.Src)
		if err != nil && a.Valid {
			t.Errorf("%#v has error %s.", a.Src, err)
		}
	}
}

func Test_GroupAddress(t *testing.T) {
	type Addr struct {
		Src   string
		Valid bool
	}

	var addrs = []Addr{
		{"1/2/3", true},
		{"31/7/255", true},
		{"31/2040", true},
		{"65535", true},
		{"84/230", false},
		{"0/0/0", false},
		{"0/0", false},
	}
	for _, a := range addrs {
		_, err := NewGroupAddrString(a.Src)
		if err != nil && a.Valid {
			t.Errorf("%#v has error %s.", a.Src, err)
		}
	}
}
