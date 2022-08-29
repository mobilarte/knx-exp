package dpt

// DPT_20102 represents DPT 20.102 (HVAC) / DPT_HVACMode.
type DPT_20102 uint8

const (
	Auto               DPT_20102 = 0
	Comfort            DPT_20102 = 1
	Standby            DPT_20102 = 2
	Economy            DPT_20102 = 3
	BuildingProtection DPT_20102 = 4
)

func (d DPT_20102) Pack() []byte {
	return packU8(uint8(d))
}

func (d *DPT_20102) Unpack(data []byte) error {
	return unpackU8(data, (*uint8)(d))
}

func (d DPT_20102) Unit() string {
	return ""
}

func (d DPT_20102) IsValid() bool {
	return d <= BuildingProtection
}

func (d DPT_20102) String() string {
	switch d {
	case Auto:
		return "Auto"
	case Comfort:
		return "Comfort"
	case Standby:
		return "Standby"
	case Economy:
		return "Economy"
	case BuildingProtection:
		return "Building Protection"
	default:
		return "reserved"
	}
}
