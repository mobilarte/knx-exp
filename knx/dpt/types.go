// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"errors"
	"fmt"
)

// A DatapointValue is a value of a datapoint.
type DatapointValue interface {
	// Pack the datapoint to a byte array.
	Pack() []byte

	// Unpack a the datapoint value from a byte array.
	Unpack(data []byte) error
}

// DatapointMeta gives meta information about a datapoint
type DatapointMeta interface {
	// Unit returns the unit of this datapoint type or empty string if it doesn't have a unit.
	Unit() string

	// fmt.Stringer provides a string representation of the datapoint.
	fmt.Stringer
}

// Datapoint represents a datapoint with both its value and metadata.
type Datapoint interface {
	DatapointValue
	DatapointMeta
}

var (
	// ErrInvalidLength is returned when the application data has unexpected length.
	ErrInvalidLength = errors.New("application data has invalid length")
	// ErrInvalidData is returned when a bit or content denotes invalid data per KNX specification.
	ErrInvalidData = errors.New("application data is noted as invalid")
	// ErrOutOfRange is returned when the unpacked value is out of range.
	ErrOutOfRange = errors.New("application data is out of range")
	// ErrBadReservedBits is returned when reserved bits are populated. E.g. if bit number 5 of a r4B4 field is populated.
	ErrBadReservedBits = errors.New("reserved bits in the application data have been populated")
)
