// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"time"
)

// DPT_19001 represents DPT 19.001 (G) / DPT_DateTime.
// Setting the field SP24 = true will produce 24:00:00 as a time only.
// Used to define time ranges 00:00:00 -> 24:00:00.
type DPT_19001 struct {
	Date time.Time
	SP24 bool
	F    bool
	WD   bool
	NWD  bool
	NY   bool
	ND   bool
	NDOW bool
	NT   bool
	SUTI bool
	CLQ  bool
	SRC  bool
}

func (d DPT_19001) Pack() []byte {
	var buf = make([]byte, 9)

	if d.SP24 {
		d.WD = false
		weekday := uint8(0)
		buf[4] = weekday<<5 | uint8(24)
	} else {
		year := d.Date.Year()
		if year < 1900 || year > 2155 {
			year = 1900
		}

		buf[1] = uint8(year - 1900)

		month := d.Date.Month()
		buf[2] = uint8(month & 0xF)

		day := d.Date.Day()
		buf[3] = uint8(day & 0x1F)

		weekday := uint8(d.Date.Weekday())
		if weekday == 0 {
			weekday = 7
		}

		d.WD = (weekday <= 5)
		hour := uint8(d.Date.Hour())
		buf[4] = (weekday << 5) | hour
		minutes := uint8(d.Date.Minute())
		buf[5] = minutes & 0x3F

		seconds := uint8(d.Date.Second())
		buf[6] = seconds & 0x3F
	}

	if d.F {
		buf[7] |= 0x80
	}

	if d.WD {
		buf[7] |= 0x40
	}

	if d.NWD {
		buf[7] |= 0x20
	}

	if d.NY {
		buf[7] |= 0x10
	}

	if d.ND {
		buf[7] |= 0x08
	}

	if d.NDOW {
		buf[7] |= 0x04
	}

	if d.NT {
		buf[7] |= 0x02
	}

	d.SUTI = d.Date.IsDST()
	if d.SUTI {
		buf[7] |= 0x01
	}

	if d.CLQ {
		buf[8] |= 0x80
	}

	if d.SRC {
		buf[8] |= 0x40
	}

	return buf
}

func (d *DPT_19001) Unpack(data []byte) error {
	if len(data) != 9 {
		return ErrInvalidLength
	}

	year := int(data[1]) + 1900

	month := time.Month(data[2] & 0xF)
	day := int(data[3] & 0x1F)

	// dayOfWeek is unused because we are relying on go to get correct weekday
	// dayOfWeek := uint8(data[4] >> 5)

	hour := int(data[4] & 0x1F)
	minutes := int(data[5] & 0x3F)
	seconds := int(data[6] & 0x3F)

	if hour == 24 {
		if minutes == 0 && seconds == 0 {
			d.SP24 = true
		} else {
			return ErrOutOfRange
		}
	} else {
		d.SP24 = false
	}

	d.Date = time.Date(year, month, day, hour, minutes, seconds, 0, time.UTC)

	d.F = (data[7]&0x80 == 0x80)
	d.WD = (data[7]&0x40 == 0x40)
	d.NWD = (data[7]&0x20 == 0x20)
	d.NY = (data[7]&0x10 == 0x10)
	d.ND = (data[7]&0x8 == 0x8)
	d.NDOW = (data[7]&0x4 == 0x4)
	d.NT = (data[7]&0x2 == 0x2)
	d.SUTI = (data[7]&0x1 == 0x1)

	d.CLQ = (data[8]&0x80 == 0x80)
	d.SRC = (data[8]&0x40 == 0x40)

	return nil
}

func (d DPT_19001) Unit() string {
	return ""
}

func (d DPT_19001) String() string {
	if d.SP24 {
		return "24:00:00"
	}

	if d.NT {
		return d.Date.Format("2006/01/02")
	}

	if d.NY && d.ND && d.NDOW {
		return d.Date.Format("15:04:05")
	}

	return d.Date.Format("2006/01/02 15:04:05")
}
