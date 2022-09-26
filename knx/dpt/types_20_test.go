package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 20.102 (N8)
func TestDPT_20102(t *testing.T) {
	type DPT20 struct {
		Dpv DatapointValue
		Idx uint8
		Val string
	}

	var types_20 = []DPT20{
		{new(DPT_20102), 0, "Auto"},
		{new(DPT_20102), 1, "Comfort"},
		{new(DPT_20102), 2, "Standby"},
		{new(DPT_20102), 3, "Economy"},
		{new(DPT_20102), 4, "Building Protection"},
		{new(DPT_20102), 5, "reserved"},
	}
	for _, e := range types_20 {
		err := e.Dpv.Unpack(packU8(e.Idx))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.Val {
			t.Errorf("%#v has wrong value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.Dpv)
		}
	}
}

// Test DPT 20.105 (N8)
func TestDPT_20105(t *testing.T) {
	type DPT20 struct {
		Dpv DatapointValue
		Idx uint8
		Val string
	}

	var types_20 = []DPT20{
		{new(DPT_20105), 0, "Auto"},
		{new(DPT_20105), 1, "Heat"},
		{new(DPT_20105), 2, "Morning Pump"},
		{new(DPT_20105), 3, "Cool"},
		{new(DPT_20105), 4, "Night Purge"},
		{new(DPT_20105), 5, "Precool"},
		{new(DPT_20105), 6, "Off"},
		{new(DPT_20105), 7, "Test"},
		{new(DPT_20105), 8, "Emergency Heat"},
		{new(DPT_20105), 9, "Fan only"},
		{new(DPT_20105), 10, "Free Cool"},
		{new(DPT_20105), 11, "Ice"},
		{new(DPT_20105), 12, "Maximum Heating Mode"},
		{new(DPT_20105), 13, "Economic Heat/Cool Mode"},
		{new(DPT_20105), 14, "Dehumidification"},
		{new(DPT_20105), 15, "Calibration Mode"},
		{new(DPT_20105), 16, "Emergency Cool Mode"},
		{new(DPT_20105), 17, "Emergency Steam Mode"},
		{new(DPT_20105), 18, "reserved"},
		{new(DPT_20105), 19, "reserved"},
		{new(DPT_20105), 20, "NoDem"},
		{new(DPT_20105), 21, "reserved"},
	}

	for _, e := range types_20 {
		err := e.Dpv.Unpack(packU8(e.Idx))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.Val {
			t.Errorf("%#v has wrong value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.Dpv)
		}
	}
}
