// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_13001 represents DPT 13.001 (G) / DPT_Value_4_Count.
type DPT_13001 int32

func (d DPT_13001) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13001) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13001) Unit() string {
	return "counter pulses"
}

func (d DPT_13001) String() string {
	return fmt.Sprintf("%d counter pulses", int32(d))
}

// DPT_13002 represents DPT 13.002 (G) / DPT_FlowRate_m3/h.
type DPT_13002 int32

func (d DPT_13002) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13002) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13002) Unit() string {
	return "m³/h"
}

func (d DPT_13002) String() string {
	return fmt.Sprintf("%d m³/h", int32(d))
}

// DPT_13010 represents DPT 13.010 (G) / DPT_ActiveEnergy.
type DPT_13010 int32

func (d DPT_13010) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13010) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13010) Unit() string {
	return "Wh"
}

func (d DPT_13010) String() string {
	return fmt.Sprintf("%d Wh", int32(d))
}

// DPT_13011 represents DPT 13.011 (G) / DPT_ApparantEnergy.
type DPT_13011 int32

func (d DPT_13011) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13011) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13011) Unit() string {
	return "VAh"
}

func (d DPT_13011) String() string {
	return fmt.Sprintf("%d VAh", int32(d))
}

// DPT_13012 represents DPT 13.012 (G) / DPT_ReactiveEnergy.
type DPT_13012 int32

func (d DPT_13012) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13012) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13012) Unit() string {
	return "VARh"
}

func (d DPT_13012) String() string {
	return fmt.Sprintf("%d VARh", int32(d))
}

// DPT_13013 represents DPT 13.013 (G) / DPT_ActiveEnergy_kWh.
type DPT_13013 int32

func (d DPT_13013) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13013) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13013) Unit() string {
	return "kWh"
}

func (d DPT_13013) String() string {
	return fmt.Sprintf("%d kWh", int32(d))
}

// DPT_13014 represents DPT 13.014 (G) / DPT_Apparant_Energy_kVAh.
type DPT_13014 int32

func (d DPT_13014) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13014) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13014) Unit() string {
	return "kVAh"
}

func (d DPT_13014) String() string {
	return fmt.Sprintf("%d kVAh", int32(d))
}

// DPT_13015 represents DPT 13.015 (G) / DPT_ReactiveEnergy_kVARh.
type DPT_13015 int32

func (d DPT_13015) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13015) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13015) Unit() string {
	return "kVARh"
}

func (d DPT_13015) String() string {
	return fmt.Sprintf("%d kVARh", int32(d))
}

// DPT_13016 represents DPT 13.016 (G) / DPT_ActiveEnergy_MWh.
type DPT_13016 int32

func (d DPT_13016) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13016) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13016) Unit() string {
	return "MWh"
}

func (d DPT_13016) String() string {
	return fmt.Sprintf("%d MWh", int32(d))
}

// DPT_13100 represents DPT 13.100 (G) / DPT_LongDeltaTimeSec.
type DPT_13100 int32

func (d DPT_13100) Pack() []byte {
	return packV32(d)
}

func (d *DPT_13100) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_13100) Unit() string {
	return "s"
}

func (d DPT_13100) String() string {
	return fmt.Sprintf("%d s", int32(d))
}

// DPT_131200 represents DPT 13.1200 (G) / DPT_DeltaVolumeLiquid_Litre.
type DPT_131200 int32

func (d DPT_131200) Pack() []byte {
	return packV32(d)
}

func (d *DPT_131200) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_131200) Unit() string {
	return "litre"
}

func (d DPT_131200) String() string {
	return fmt.Sprintf("%d litre", int32(d))
}

// DPT_131201 represents DPT 13.1200 (G) / DPT_DeltaVolume_m3.
type DPT_131201 int32

func (d DPT_131201) Pack() []byte {
	return packV32(d)
}

func (d *DPT_131201) Unpack(data []byte) error {
	return unpackV32(data, d)
}

func (d DPT_131201) Unit() string {
	return "m³"
}

func (d DPT_131201) String() string {
	return fmt.Sprintf("%d  m³", int32(d))
}
