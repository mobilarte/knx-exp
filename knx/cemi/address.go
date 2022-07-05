// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package cemi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// IndividualAddr is an individual address for a KNX device.
type IndividualAddr uint16

// NewIndividualAddr3 generates an individual address from "a.b.c".
// a is the area address [0..15], b is the line address [0..15] and
// c is the device address [0..255].
func NewIndividualAddr3(a, b, c uint8) IndividualAddr {
	return IndividualAddr(a&0xF)<<12 | IndividualAddr(b&0xF)<<8 | IndividualAddr(c)
}

// NewIndividualAddr2 generates an individual address from "a.b".
// a is the subnetwork address [0..255], b is the device address [0.255].
func NewIndividualAddr2(a, b uint8) IndividualAddr {
	return IndividualAddr(a)<<8 | IndividualAddr(b)
}

// NewIndividualAddrString parses the given string to an individual address.
// Supported formats are
// %d.%d.%d ([0..15], [0..15], [0..255]),
// %d.%d ([0..255], [0..255]) and
// %d ([0..65535]). Ranges are checked.
func NewIndividualAddrString(addr string) (IndividualAddr, error) {
	var nums []int

	numstrings := strings.Split(addr, ".")

	for _, s := range numstrings {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		nums = append(nums, i)
	}

	switch len(nums) {
	case 3:
		if nums[0] < 0 || nums[0] > 15 ||
			nums[1] < 0 || nums[1] > 15 ||
			nums[2] < 0 || nums[2] > 255 {
			return 0, errors.New("invalid area or main line number")
		}
		return NewIndividualAddr3(uint8(nums[0]), uint8(nums[1]), uint8(nums[2])), nil
	case 2:
		if nums[0] < 0 || nums[0] > 255 || nums[1] < 0 || nums[1] > 255 {
			return 0, errors.New("invalid subnetwork or device number")
		}
		return NewIndividualAddr2(uint8(nums[0]), uint8(nums[1])), nil
	case 1:
		if nums[0] < 0 || nums[0] > 65535 {
			return 0, errors.New("invalid area or main line number")
		}
		return IndividualAddr(nums[0]), nil
	}

	return 0, errors.New("string cannot be parsed to an individual address")
}

// String generates a string representation "a.b.c" with
// a = area address = 4 bits, b = line address = 4 bits,
// c = device address = 1 byte.
func (addr IndividualAddr) String() string {
	return fmt.Sprintf("%d.%d.%d", uint8(addr>>12)&0xF, uint8(addr>>8)&0xF, uint8(addr))
}

// GroupAddr is an address for a KNX group object. Group address
// zero is a broadcast.
type GroupAddr uint16

// NewGroupAddr3 generates a group address from "a/b/c".
// a = 5 bits, b = 3 bits, c = 1 byte.
func NewGroupAddr3(a, b, c uint8) GroupAddr {
	return GroupAddr(a&0x1F)<<11 | GroupAddr(b&0x7)<<8 | GroupAddr(c)
}

// NewGroupAddr2 generates a group address from "a/b".
// a = 5 bits, b = 11 bits.
func NewGroupAddr2(a uint8, b uint16) GroupAddr {
	return GroupAddr(a)<<8 | GroupAddr(b&0x7FF)
}

// NewGroupAddrString parses the given string to a group address.
// Supported formats are:
// %d/%d/%d ([0..31], [0..7], [0..255]),
// %d/%d ([0..31], [0..2047]) and
// %d ([0..65535]). Ranges are checked.
func NewGroupAddrString(addr string) (GroupAddr, error) {
	var nums []int

	numstrings := strings.Split(addr, "/")

	for _, s := range numstrings {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		nums = append(nums, i)
	}

	switch len(nums) {
	case 3:
		if nums[0] < 0 || nums[0] > 31 ||
			nums[1] < 0 || nums[1] > 7 ||
			nums[2] < 0 || nums[2] > 255 {
			return 0, errors.New("invalid main, middle group address")
		}
		if nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {
			return 0, errors.New("invalid group address 0/0/0")
		}
		return NewGroupAddr3(uint8(nums[0]), uint8(nums[1]), uint8(nums[2])), nil
	case 2:
		if nums[0] < 0 || nums[0] > 31 ||
			nums[1] < 0 || nums[1] > 2047 {
			return 0, errors.New("invalid short group number")
		}
		if nums[0] == 0 && nums[1] == 0 {
			return 0, errors.New("invalid group address 0/0")
		}
		return NewGroupAddr2(uint8(nums[0]), uint16(nums[1])), nil
	case 1:
		if nums[0] < 0 || nums[0] > 65535 {
			return 0, errors.New("invalid free group number")
		}
		return GroupAddr(nums[0]), nil
	}

	return 0, errors.New("string cannot be parsed to a group address")
}

// String generates a string representation with middle group "a/b/c" where
// a = 5 bits, b = 3 bits, c = 1 byte.
func (addr GroupAddr) String() string {
	return fmt.Sprintf("%d/%d/%d", uint8(addr>>11)&0x1F, uint8(addr>>8)&0x7, uint8(addr))
}
