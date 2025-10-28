// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_29010 represents DPT 29.010 (G) / DPT_ActiveEnergy_V64.
type DPT_29010 int64

func (d DPT_29010) Pack() []byte {
	return packV64(d)
}

func (d *DPT_29010) Unpack(data []byte) error {
	return unpackV64(data, d)
}

func (d DPT_29010) Unit() string {
	return "Wh"
}

func (d DPT_29010) String() string {
	return fmt.Sprintf("%d Wh", int64(d))
}

// DPT_29011 represents DPT 29.011 (G) / DPT_ApparentEnergy_V64.
type DPT_29011 int64

func (d DPT_29011) Pack() []byte {
	return packV64(d)
}

func (d *DPT_29011) Unpack(data []byte) error {
	return unpackV64(data, d)
}

func (d DPT_29011) Unit() string {
	return "VAh"
}

func (d DPT_29011) String() string {
	return fmt.Sprintf("%d VAh", int64(d))
}

// DPT_29012 represents DPT 29.012 (G) / DPT_ReactiveEnergy_V64.
type DPT_29012 int64

func (d DPT_29012) Pack() []byte {
	return packV64(d)
}

func (d *DPT_29012) Unpack(data []byte) error {
	return unpackV64(data, d)
}

func (d DPT_29012) Unit() string {
	return "VARh"
}

func (d DPT_29012) String() string {
	return fmt.Sprintf("%d VARh", int64(d))
}
