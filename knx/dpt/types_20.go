package dpt

// DPT_20002 represents DPT 20.002 (G) / DPT_BuildingMode.
type DPT_20002 uint8

const (
	BuildingMode_BuildingInUse      DPT_20002 = 0
	BuildingMode_BuildingNotUsed    DPT_20002 = 1
	BuildingMode_BuildingProtection DPT_20002 = 2
	// 3..255: reserverd, shall not be used
)

func (d DPT_20002) Pack() []byte {
	return packU8(uint8(d))
}

func (d *DPT_20002) Unpack(data []byte) error {
	return unpackU8(data, (*uint8)(d))
}

func (d DPT_20002) Unit() string {
	return ""
}

func (d DPT_20002) IsValid() bool {
	return d <= BuildingMode_BuildingProtection
}

func (d DPT_20002) String() string {
	switch d {
	case BuildingMode_BuildingInUse:
		return "Building in use"
	case BuildingMode_BuildingNotUsed:
		return "Building not used"
	case BuildingMode_BuildingProtection:
		return "Building protection"
	default:
		return "reserved, shall not be used"
	}
}

// DPT_20003 represents DPT 20.003 (G) / DPT_OccMode.
type DPT_20003 uint8

const (
	OccMode_occupied     DPT_20003 = 0
	OccMode_standby      DPT_20003 = 1
	OccMode_not_occupied DPT_20003 = 2
	// 3..255: not used, reserverd
)

func (d DPT_20003) Pack() []byte {
	return packU8(uint8(d))
}

func (d *DPT_20003) Unpack(data []byte) error {
	return unpackU8(data, (*uint8)(d))
}

func (d DPT_20003) Unit() string {
	return ""
}

func (d DPT_20003) IsValid() bool {
	return d <= OccMode_not_occupied
}

func (d DPT_20003) String() string {
	switch d {
	case OccMode_occupied:
		return "occupied"
	case OccMode_standby:
		return "standby"
	case OccMode_not_occupied:
		return "not occupied"
	default:
		return "no used; reserved"
	}
}

// DPT_20014 represents DPT 20.014 (G) / DPT_Beaufort_Wind_Force_Scale.
type DPT_20014 uint8

// Wind Force Scale
const (
	WindForceScale_Calm           DPT_20014 = 0
	WindForceScale_LightAir       DPT_20014 = 1
	WindForceScale_LightBreeze    DPT_20014 = 2
	WindForceScale_GentleBreeze   DPT_20014 = 3
	WindForceScale_ModerateBreeze DPT_20014 = 4
	WindForceScale_FreshBreeze    DPT_20014 = 5
	WindForceScale_StrongBreeze   DPT_20014 = 6
	WindForceScale_NearGale       DPT_20014 = 7
	WindForceScale_FreshGale      DPT_20014 = 8
	WindForceScale_StrongGale     DPT_20014 = 9
	WindForceScale_WholeGale      DPT_20014 = 10
	WindForceScale_ViolentStorm   DPT_20014 = 11
	WindForceScale_Hurricane      DPT_20014 = 12
	// 13..255: reserverd, shall not be used
)

func (d DPT_20014) Pack() []byte {
	return packU8(uint8(d))
}

func (d *DPT_20014) Unpack(data []byte) error {
	return unpackU8(data, (*uint8)(d))
}

func (d DPT_20014) Unit() string {
	return ""
}

func (d DPT_20014) IsValid() bool {
	return d <= WindForceScale_Hurricane
}

func (d DPT_20014) String() string {
	switch d {
	case WindForceScale_Calm:
		return "calm (no wind)"
	case WindForceScale_LightAir:
		return "light air"
	case WindForceScale_LightBreeze:
		return "light breeze"
	case WindForceScale_GentleBreeze:
		return "gentle breeze"
	case WindForceScale_ModerateBreeze:
		return "moderate breeze"
	case WindForceScale_FreshBreeze:
		return "fresh breeze"
	case WindForceScale_StrongBreeze:
		return "strong breeze"
	case WindForceScale_NearGale:
		return "near gale / moderate gale"
	case WindForceScale_FreshGale:
		return "fresh gale"
	case WindForceScale_StrongGale:
		return "strong gale"
	case WindForceScale_WholeGale:
		return "whole gale / storm"
	case WindForceScale_ViolentStorm:
		return "violent storm"
	case WindForceScale_Hurricane:
		return "hurricane"
	default:
		return "reserved, shall not be used"
	}
}

// DPT_20102 represents DPT 20.102 (HVAC) / DPT_HVACMode.
type DPT_20102 uint8

// HVACMode_.. defines the possible states.
const (
	HVACMode_Auto               DPT_20102 = 0
	HVACMode_Comfort            DPT_20102 = 1
	HVACMode_Standby            DPT_20102 = 2
	HVACMode_Economy            DPT_20102 = 3
	HVACMode_BuildingProtection DPT_20102 = 4
	// 5...255, reserved
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

// HVACContrMode_... defines the possible states.
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
