package dpt

import (
	"fmt"
)

// DPT_232600 represents DPT 232.600 (G) / DPT_Colour_RGB.
// Colour RGB - RGB value 4x(0..100%) / U8 U8 U8
type DPT_232600 struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (d DPT_232600) Pack() []byte {

	return pack3U8(d.Red, d.Green, d.Blue)
}

func (d *DPT_232600) Unpack(data []byte) error {

	if len(data) != 4 {
		return ErrInvalidLength
	}

	err := unpack3U8(data, &d.Red, &d.Green, &d.Blue)

	if err != nil {
		return ErrInvalidLength
	}

	return nil
}

func (d DPT_232600) Unit() string {
	return ""
}

func (d DPT_232600) String() string {
	return fmt.Sprintf("%d %d %d", d.Red, d.Green, d.Blue)
}
