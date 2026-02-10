// Copyright 2017 Ole Krüger.
// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"encoding/binary"
	"math"
)

func packB1[T ~bool](b T) []byte {
	if b {
		return []byte{1}
	}

	return []byte{0}
}

func unpackB1[T ~bool](data []byte, b *T) error {
	if len(data) != 1 {
		return ErrInvalidLength
	}

	*b = data[0]&0x1 == 0x1

	return nil
}

func packB1U3(c bool, v uint8) []byte {
	if c {
		return []byte{v&0x7 | 0x8}
	}

	return []byte{v & 0x7}
}

func unpackB1U3(data []byte, c *bool, v *uint8) error {
	if len(data) != 1 {
		return ErrInvalidLength
	}

	*c = data[0]&0x8 == 0x8
	*v = data[0] & 0x7

	return nil
}

func packB2(b1 bool, b0 bool) []byte {
	var b byte

	if b1 {
		b |= 0x2
	}

	if b0 {
		b |= 0x1
	}

	return []byte{b}
}

func unpackB2(data []byte, b1 *bool, b0 *bool) error {
	if len(data) != 1 {
		return ErrInvalidLength
	}

	if data[0] > 0x3 {
		return ErrBadReservedBits
	}

	*b1 = data[0]&0x2 == 0x2
	*b0 = data[0]&0x1 == 0x1

	return nil
}

func packB4(b3 bool, b2 bool, b1 bool, b0 bool) byte {
	var b byte

	if b3 {
		b |= 0x8
	}

	if b2 {
		b |= 0x4
	}

	if b1 {
		b |= 0x2
	}

	if b0 {
		b |= 0x1
	}

	return byte(b)
}

func unpackB4(data byte, b3 *bool, b2 *bool, b1 *bool, b0 *bool) error {
	if uint8(data) > 15 {
		return ErrBadReservedBits
	}

	*b3 = data&0x8 == 0x8
	*b2 = data&0x4 == 0x4
	*b1 = data&0x2 == 0x2
	*b0 = data&0x1 == 0x1

	return nil
}

func packU8[T ~uint8](i T) []byte {
	return []byte{0, uint8(i)}
}

func unpackU8[T ~uint8](data []byte, i *T) error {
	if len(data) != 2 {
		return ErrInvalidLength
	}

	*i = T(data[1])

	return nil
}

func pack3U8[T ~uint8](u2 T, u1 T, u0 T) []byte {
	return []byte{0, uint8(u2), uint8(u1), uint8(u0)}
}

func unpack3U8(data []byte, u2 *uint8, u1 *uint8, u0 *uint8) error {
	if len(data) != 4 {
		return ErrInvalidLength
	}

	*u2 = uint8(data[1])
	*u1 = uint8(data[2])
	*u0 = uint8(data[3])

	return nil
}

func packV8[T ~int8](i T) []byte {
	return []byte{0, byte(i)}
}

func unpackV8[T ~int8](data []byte, i *T) error {
	if len(data) != 2 {
		return ErrInvalidLength
	}

	*i = T(data[1])

	return nil
}

func packU16[T ~uint16](i T) []byte {
	buffer := make([]byte, 3)

	binary.BigEndian.PutUint16(buffer[1:], uint16(i))

	return buffer
}

func unpackU16[T ~uint16](data []byte, i *T) error {
	if len(data) != 3 {
		return ErrInvalidLength
	}

	*i = T(binary.BigEndian.Uint16(data[1:]))

	return nil
}

func packV16[T ~int16](i T) []byte {
	buffer := make([]byte, 3)

	binary.BigEndian.PutUint16(buffer[1:], uint16(i))

	return buffer
}

func unpackV16[T ~int16](data []byte, i *T) error {
	if len(data) != 3 {
		return ErrInvalidLength
	}

	*i = T(int16(binary.BigEndian.Uint16(data[1:])))

	return nil
}

func packU32[T ~uint32](i T) []byte {
	buffer := make([]byte, 5)

	binary.BigEndian.PutUint32(buffer[1:], uint32(i))

	return buffer
}

func unpackU32[T ~uint32](data []byte, i *T) error {
	if len(data) != 5 {
		return ErrInvalidLength
	}

	*i = T(binary.BigEndian.Uint32(data[1:]))

	return nil
}

func packV32[T ~int32](i T) []byte {
	buffer := make([]byte, 5)

	binary.BigEndian.PutUint32(buffer[1:], uint32(i))

	return buffer
}

func unpackV32[T ~int32](data []byte, i *T) error {
	if len(data) != 5 {
		return ErrInvalidLength
	}

	*i = T(binary.BigEndian.Uint32(data[1:]))

	return nil
}

func packV64[T ~int64](i T) []byte {
	buffer := make([]byte, 9)

	binary.BigEndian.PutUint64(buffer[1:], uint64(i))

	return buffer
}

func unpackV64[T ~int64](data []byte, i *T) error {
	if len(data) != 9 {
		return ErrInvalidLength
	}

	*i = T(binary.BigEndian.Uint64(data[1:]))

	return nil
}

func packF16[T ~float64](f T) []byte {
	buffer := make([]byte, 3)

	if f > 670433.28 {
		f = 670433.28
	} else if f < -671088.64 {
		f = -671088.64
	}

	signedMantissa := f * 100.0
	exp := 0

	for signedMantissa > 2047 || signedMantissa < -2048 {
		signedMantissa /= 2
		exp++
	}

	buffer[1] |= uint8(exp&0xF) << 3

	if signedMantissa < 0 {
		signedMantissa += 2048
		buffer[1] |= 0x1 << 7
	}

	mantissa := uint(signedMantissa)

	buffer[1] |= uint8(mantissa>>8) & 0x7
	buffer[2] |= uint8(mantissa)

	return buffer
}

func unpackF16[T ~float64](data []byte, f *T) error {
	if len(data) != 3 {
		return ErrInvalidLength
	}

	// This value denotes invalid data, only applies to DPT 9.00x!
	if data[1]&0x7F == 0x7F && data[2]&0xFF == 0xFF {
		return ErrInvalidData
	}

	m := int(data[1]&0x7)<<8 | int(data[2])

	e := (data[1] >> 3) & 0x0F
	if data[1]&0x80 == 0x80 {
		m -= 2048
	}

	*f = T(m<<e) / 100.0

	return nil
}

func packF32[T ~float32](f T) []byte {
	buffer := make([]byte, 5)

	binary.BigEndian.PutUint32(buffer[1:], math.Float32bits(float32(f)))

	return buffer
}

func unpackF32[T ~float32](data []byte, f *T) error {
	if len(data) != 5 {
		return ErrInvalidLength
	}

	*f = T(math.Float32frombits(binary.BigEndian.Uint32(data[1:])))

	return nil
}
