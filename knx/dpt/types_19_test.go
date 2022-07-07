// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"testing"
	"time"
)

func TestDPT_19001(t *testing.T) {
	var src, dst DPT_19001

	// Default, simply fill in time
	src.Date = time.Now()
	buf := src.Pack()
	err := dst.Unpack(buf)
	if err != nil {
		t.Errorf("Unpacking error")
	}

	if src.Date.Format("2006/01/02 15:04:05") != dst.String() {
		t.Errorf("Unexpected unpacking. Expected [%s], received [%s].", src.Date.Format("2006/01/02 15:04:05"), dst.String())
	}

	// A non-existant date, which golang will fix!
	src.Date = time.Date(2022, 2, 30, 24, 0, 1, 0, time.UTC)
	buf = src.Pack()
	err = dst.Unpack(buf)
	if err != nil {
		t.Errorf("Unpacking error")
	}

	// Date with hour=24, but with non-zero minutes/seconds (byte 6, 7)!
	// 2022/06/07 24:01:01
	buf = []byte{0x00, 0x7a, 0x06, 0x08, 0x78, 0x01, 0x01, 0x40, 0x00}
	err = dst.Unpack(buf)
	if err == nil {
		t.Errorf("Unpacking error, when hour = 24, minutes = seconds = 0!")
	}

	// Test attribute manual and default settings
	src.Date = time.Now()
	src.SRC = true
	buf = src.Pack()
	err = dst.Unpack(buf)

	if err != nil {
		t.Errorf("Unpacking error")
	}
	if !dst.SRC {
		t.Errorf("Unpack SRC attribute error")
	}
	if src.Date.IsDST() != dst.SUTI {
		t.Errorf("Summertime setting wrong")
	}

	// With time only
	src.Date = time.Now()
	src.NY = true
	src.ND = true
	src.NDOW = true
	buf = src.Pack()
	err = dst.Unpack(buf)
	if err != nil {
		t.Errorf("Unpacking error")
	}
	if src.Date.Format("15:04:05") != dst.String() {
		t.Errorf("Error in time only")
	}

	// Date without time
	src = DPT_19001{}
	src.Date = time.Now()
	src.NT = true
	buf = src.Pack()
	err = dst.Unpack(buf)
	if err != nil {
		t.Errorf("Unpacking error")
	}
	if src.Date.Format("2006/01/02") != dst.String() {
		t.Errorf("Error in date only")
	}

	// Date with time only as 24:00:00
	// Can put anything into the Date, will be ignored
	src = DPT_19001{}
	src.Date = time.Now()
	src.SP24 = true
	buf = src.Pack()
	dst.Unpack(buf)
	if dst.String() != "24:00:00" {
		t.Errorf("Error in special time 24:00:00")
	}

	// Date with time only, as 00:00:00
	src = DPT_19001{}
	src.Date = time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)
	src.NY = true
	src.ND = true
	src.NDOW = true
	buf = src.Pack()
	dst.Unpack(buf)
	if dst.String() != "00:00:00" {
		t.Errorf("Error in special time 00:00:00")
	}
}
