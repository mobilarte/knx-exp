package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 20.xxx (N8)
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

/*
	knxValue := []byte{0, 4}
	dptValue := DPT_20102(4)

	var tmpDPT DPT_20102
	assert.NoError(t, tmpDPT.Unpack(knxValue))
	assert.Equal(t, dptValue, tmpDPT)

	assert.Equal(t, knxValue, dptValue.Pack())

	assert.Equal(t, "Building Protection", dptValue.String())
}
*/
