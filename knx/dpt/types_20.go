package dpt

// DPT_20102 represents DPT 20.102 (HVAC) / DPT_HVACMode.
type DPT_20102 uint8

const (
	HVACMode_Auto               DPT_20102 = 0
	HVACMode_Comfort            DPT_20102 = 1
	HVACMode_Standby            DPT_20102 = 2
	HVACMode_Economy            DPT_20102 = 3
	HVACMode_BuildingProtection DPT_20102 = 4
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
	return d <= HVACMode_BuildingProtection
}

func (d DPT_20102) String() string {
	switch d {
	case HVACMode_Auto:
		return "Auto"
	case HVACMode_Comfort:
		return "Comfort"
	case HVACMode_Standby:
		return "Standby"
	case HVACMode_Economy:
		return "Economy"
	case HVACMode_BuildingProtection:
		return "Building Protection"
	default:
		return "reserved"
	}
}

// DPT_20105 represents DPT 20.105 (HVAC) / DPT_HVACContrMode.
type DPT_20105 uint8

const (
	HVACContrMode_Auto                    DPT_20105 = 0
	HVACContrMode_Heat                    DPT_20105 = 1
	HVACContrMode_Morning_Pump            DPT_20105 = 2
	HVACContrMode_Cool                    DPT_20105 = 3
	HVACContrMode_Night_Purge             DPT_20105 = 4
	HVACContrMode_Precool                 DPT_20105 = 5
	HVACContrMode_Off                     DPT_20105 = 6
	HVACContrMode_Test                    DPT_20105 = 7
	HVACContrMode_Emergency_Heat          DPT_20105 = 8
	HVACContrMode_Fan_only                DPT_20105 = 9
	HVACContrMode_Free_Cool               DPT_20105 = 10
	HVACContrMode_Ice                     DPT_20105 = 11
	HVACContrMode_Maximum_Heating_Mode    DPT_20105 = 12
	HVACContrMode_Economic_Heat_Cool_Mode DPT_20105 = 13
	HVACContrMode_Dehumidification        DPT_20105 = 14
	HVACContrMode_Calibration_Mode        DPT_20105 = 15
	HVACContrMode_Emergency_Cool_Mode     DPT_20105 = 16
	HVACContrMode_Emergency_Steam_Mode    DPT_20105 = 17
	HVACContrMode_NoDem                   DPT_20105 = 20
)

func (d DPT_20105) Pack() []byte {
	return packU8(uint8(d))
}

func (d *DPT_20105) Unpack(data []byte) error {
	return unpackU8(data, (*uint8)(d))
}

func (d DPT_20105) Unit() string {
	return ""
}

func (d DPT_20105) IsValid() bool {
	return d <= HVACContrMode_NoDem && uint8(d) != 18 && uint8(d) != 19
}

func (d DPT_20105) String() string {
	switch d {
	case HVACContrMode_Auto:
		return "Auto"
	case HVACContrMode_Heat:
		return "Heat"
	case HVACContrMode_Morning_Pump:
		return "Morning Pump"
	case HVACContrMode_Cool:
		return "Cool"
	case HVACContrMode_Night_Purge:
		return "Night Purge"
	case HVACContrMode_Precool:
		return "Precool"
	case HVACContrMode_Off:
		return "Off"
	case HVACContrMode_Test:
		return "Test"
	case HVACContrMode_Emergency_Heat:
		return "Emergency Heat"
	case HVACContrMode_Fan_only:
		return "Fan only"
	case HVACContrMode_Free_Cool:
		return "Free Cool"
	case HVACContrMode_Ice:
		return "Ice"
	case HVACContrMode_Maximum_Heating_Mode:
		return "Maximum Heating Mode"
	case HVACContrMode_Economic_Heat_Cool_Mode:
		return "Economic Heat/Cool Mode"
	case HVACContrMode_Dehumidification:
		return "Dehumidification"
	case HVACContrMode_Calibration_Mode:
		return "Calibration Mode"
	case HVACContrMode_Emergency_Cool_Mode:
		return "Emergency Cool Mode"
	case HVACContrMode_Emergency_Steam_Mode:
		return "Emergency Steam Mode"
	case HVACContrMode_NoDem:
		return "NoDem"
	default:
		return "reserved"
	}
}
