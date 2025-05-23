// Copyright 2025 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.
// Described in 03_08_07 KNXnetIP Remote Configuration and Diagnosis v01.01.02 AS.pdf

package knxnet

import (
	"errors"
	"fmt"
	"net"

	"github.com/mobilarte/knx-exp/knx/cemi"
	"github.com/mobilarte/knx-exp/knx/util"
)

const (
	friendlyNameMaxLen = 30
)

// DescriptionType describes the type of a DeviceInformationBlock.
type DescriptionType uint8

const (
	// DescriptionTypeDeviceInfo describes Device information e.g. KNX medium.
	DescriptionTypeDeviceInfo DescriptionType = 0x01

	// DescriptionTypeSupportedServiceFamilies describes Service families supported by the device.
	DescriptionTypeSupportedServiceFamilies DescriptionType = 0x02

	// DescriptionTypeIPConfig describes IP configuration.
	DescriptionTypeIPConfig DescriptionType = 0x03

	// DescriptionTypeIPCurrentConfig describes current IP configuration.
	DescriptionTypeIPCurrentConfig DescriptionType = 0x04

	// DescriptionTypeKNXAddresses describes KNX addresses.
	DescriptionTypeKNXAddresses DescriptionType = 0x05

	// DescriptionTypeManufacturerData describes a DIB structure for further data defined by device manufacturer.
	DescriptionTypeManufacturerData DescriptionType = 0xfe
)

// KNXMedium describes the KNX medium type.
type KNXMedium uint8

const (
	// KNXMediumTP1 is the TP1 medium
	KNXMediumTP1 KNXMedium = 0x02
	// KNXMediumPL110 is the PL110 medium
	KNXMediumPL110 KNXMedium = 0x04
	// KNXMediumRF is the RF medium
	KNXMediumRF KNXMedium = 0x10
	// KNXMediumIP is the IP medium
	KNXMediumIP KNXMedium = 0x20
)

// ProjectInstallationIdentifier describes a KNX project installation identifier.
type ProjectInstallationIdentifier uint16

// DeviceStatus describes the device status.
type DeviceStatus uint8

// DeviceSerialNumber desribes the serial number of a device.
type DeviceSerialNumber [6]byte

// DeviceInformationBlock contains information about a device.
type DeviceInformationBlock struct {
	Type                    DescriptionType
	Medium                  KNXMedium
	Status                  DeviceStatus
	Source                  cemi.IndividualAddr
	ProjectIdentifier       ProjectInstallationIdentifier
	SerialNumber            DeviceSerialNumber
	RoutingMulticastAddress Address
	HardwareAddr            net.HardwareAddr
	FriendlyName            string
}

// Size returns the packed size.
func (DeviceInformationBlock) Size() uint {
	return 54
}

// Pack assembles the device information structure in the given buffer.
func (dib *DeviceInformationBlock) Pack(buffer []byte) {
	buf := make([]byte, friendlyNameMaxLen)
	util.PackString(buf, friendlyNameMaxLen, dib.FriendlyName)

	util.PackSome(
		buffer,
		uint8(dib.Size()), uint8(dib.Type),
		uint8(dib.Medium), uint8(dib.Status),
		uint16(dib.Source),
		uint16(dib.ProjectIdentifier),
		dib.SerialNumber[:],
		dib.RoutingMulticastAddress[:],
		[]byte(dib.HardwareAddr),
		buf,
	)
}

// Unpack parses the given data in order to initialize the structure.
func (dib *DeviceInformationBlock) Unpack(data []byte) (n uint, err error) {
	var length uint8

	dib.HardwareAddr = make([]byte, 6)
	if n, err = util.UnpackSome(
		data,
		&length, (*uint8)(&dib.Type),
		(*uint8)(&dib.Medium), (*uint8)(&dib.Status),
		(*uint16)(&dib.Source),
		(*uint16)(&dib.ProjectIdentifier),
		dib.SerialNumber[:],
		dib.RoutingMulticastAddress[:],
		[]byte(dib.HardwareAddr),
	); err != nil {
		return
	}

	nn, err := util.UnpackString(data[n:], friendlyNameMaxLen, &dib.FriendlyName)
	if err != nil {
		return n, err
	}
	n += nn

	if length != uint8(dib.Size()) {
		return n, errors.New("device info structure length is invalid")
	}

	return
}

// SupportedServicesDIB contains information about the supported services of a device.
type SupportedServicesDIB struct {
	Type     DescriptionType
	Families []ServiceFamily
}

// Size returns the packed size.
func (sdib SupportedServicesDIB) Size() uint {
	size := uint(2)
	for _, f := range sdib.Families {
		size += f.Size()
	}

	return size
}

// Pack assembles the supported services structure in the given buffer.
func (sdib *SupportedServicesDIB) Pack(buffer []byte) {
	util.PackSome(
		buffer,
		uint8(sdib.Size()), uint8(sdib.Type),
	)

	offset := uint(2)
	for _, f := range sdib.Families {
		f.Pack(buffer[offset:])
		offset += f.Size()
	}
}

// Unpack parses the given data in order to initialize the structure.
func (sdib *SupportedServicesDIB) Unpack(data []byte) (n uint, err error) {
	var length uint8
	if n, err = util.UnpackSome(
		data,
		&length, (*uint8)(&sdib.Type),
	); err != nil {
		return
	}

	for n < uint(length) {
		f := ServiceFamily{}
		nn, err := f.Unpack(data[n:])
		if err != nil {
			return n, errors.New("unable to unpack service family")
		}

		n += nn
		sdib.Families = append(sdib.Families, f)
	}

	if length != uint8(sdib.Size()) {
		return n, errors.New("invalid length for Supported Services structure")
	}

	return
}

type IpAssignment uint8
type IpCapabilities uint8

// IpConfigDIB contains information about the IP configuration.
type IpConfigDIB struct {
	Type           DescriptionType
	IpAddress      Address
	SubnetMask     Address
	DefaultGateway Address
	IPCapabilities IpCapabilities
	IPAssignment   IpAssignment
}

// Size returns the packed size.
func (IpConfigDIB) Size() uint {
	return 16
}

// Pack assembles the device information structure in the given buffer.
func (ipconf *IpConfigDIB) Pack(buffer []byte) {
	util.PackSome(
		buffer,
		uint8(ipconf.Size()), uint8(ipconf.Type),
		ipconf.IpAddress[:], ipconf.SubnetMask[:], ipconf.DefaultGateway[:],
		uint8(ipconf.IPCapabilities),
		uint8(ipconf.IPAssignment),
	)
}

// Unpack parses the given data in order to initialize the structure.
func (ipconf *IpConfigDIB) Unpack(data []byte) (n uint, err error) {
	var length uint8

	if n, err = util.UnpackSome(
		data,
		&length, (*uint8)(&ipconf.Type),
		ipconf.IpAddress[:], ipconf.SubnetMask[:], ipconf.DefaultGateway[:],
		(*uint8)(&ipconf.IPCapabilities), (*uint8)(&ipconf.IPAssignment),
	); err != nil {
		fmt.Println(err)
		return
	}

	if length != uint8(ipconf.Size()) {
		return n, errors.New("ipconfigdib structure length is invalid")
	}

	return
}

// CurConfigDIB contains information about the current IP configuration.
type CurConfigDIB struct {
	Type           DescriptionType
	IpAddress      Address
	SubnetMask     Address
	DefaultGateway Address
	DHCPServer     Address
	IPAssignment   IpAssignment
	Reserved       byte
}

// Size returns the packed size.
func (CurConfigDIB) Size() uint {
	return 20
}

// Pack assembles the device information structure in the given buffer.
func (curconf *CurConfigDIB) Pack(buffer []byte) {

	util.PackSome(
		buffer,
		uint8(curconf.Size()), uint8(curconf.Type),
		curconf.IpAddress[:], curconf.SubnetMask[:],
		curconf.DefaultGateway[:], curconf.DHCPServer[:],
		uint8(curconf.IPAssignment), curconf.Reserved,
	)
}

// Unpack parses the given data in order to initialize the structure.
func (curconf *CurConfigDIB) Unpack(data []byte) (n uint, err error) {
	var length uint8

	if n, err = util.UnpackSome(
		data,
		&length, (*uint8)(&curconf.Type),
		curconf.IpAddress[:], curconf.SubnetMask[:],
		curconf.DefaultGateway[:], curconf.DHCPServer[:],
		(*uint8)(&curconf.IPAssignment), &curconf.Reserved,
	); err != nil {
		fmt.Println(err)
		return
	}

	if length != uint8(curconf.Size()) {
		return n, errors.New("curconfigdib structure length is invalid")
	}

	return
}

// KnxAddrDIB contains information about the current IP configuration.
type KnxAddrDIB struct {
	Type         DescriptionType
	KNXAddresses []cemi.IndividualAddr
}

// Size returns the packed size. This is the size, type + all
// the individual addresses.
func (knxaddr *KnxAddrDIB) Size() uint {
	return uint(2 + len(knxaddr.KNXAddresses)*2)
}

// Pack assembles the device information structure in the given buffer.
func (knxaddr *KnxAddrDIB) Pack(buffer []byte) {
	util.PackSome(
		buffer, uint8(knxaddr.Size()), uint8(knxaddr.Type),
	)

	offset := uint(2)
	for _, e := range knxaddr.KNXAddresses {
		fmt.Println(e)
		util.PackSome(buffer[offset:], uint16(e))
		offset += 2
	}
}

// Unpack parses the given data in order to initialize the structure.
func (knxaddr *KnxAddrDIB) Unpack(data []byte) (n uint, err error) {
	var length uint8

	if n, err = util.UnpackSome(data, &length, (*uint8)(&knxaddr.Type)); err != nil {
		return
	}

	for n < uint(length) {
		var k cemi.IndividualAddr
		nn, err := util.UnpackSome(data[n:], (*uint16)(&k))
		if err != nil {
			return n, errors.New("unable to unpack individual address")
		}
		n += nn
		knxaddr.KNXAddresses = append(knxaddr.KNXAddresses, k)
	}

	if length != uint8(knxaddr.Size()) {
		return n, errors.New("knxaddr structure length is invalid")
	}

	return
}

// ServiceFamilyType describes a KNXnet service family type.
type ServiceFamilyType uint8

const (
	// ServiceFamilyTypeIPCore is the KNXnet/IP Core family type.
	ServiceFamilyTypeIPCore ServiceFamilyType = 0x02
	// ServiceFamilyTypeIPDeviceManagement is the KNXnet/IP Device Management family type.
	ServiceFamilyTypeIPDeviceManagement ServiceFamilyType = 0x03
	// ServiceFamilyTypeIPTunnelling is the KNXnet/IP Tunnelling family type.
	ServiceFamilyTypeIPTunnelling ServiceFamilyType = 0x04
	// ServiceFamilyTypeIPRouting is the KNXnet/IP Routing family type.
	ServiceFamilyTypeIPRouting ServiceFamilyType = 0x05
	// ServiceFamilyTypeIPRemoteLogging is the KNXnet/IP Remote Logging family type.
	ServiceFamilyTypeIPRemoteLogging ServiceFamilyType = 0x06
	// ServiceFamilyTypeIPRemoteConfigurationAndDiagnosis is the KNXnet/IP Remote
	// Configuration and Diagnosis family type.
	ServiceFamilyTypeIPRemoteConfigurationAndDiagnosis ServiceFamilyType = 0x07
	// ServiceFamilyTypeIPObjectServer is the KNXnet/IP Object Server family type.
	ServiceFamilyTypeIPObjectServer ServiceFamilyType = 0x08
)

// ServiceFamily describes a KNXnet service supported by a device.
type ServiceFamily struct {
	Type    ServiceFamilyType
	Version uint8
}

// Size returns the packed size.
func (ServiceFamily) Size() uint {
	return 2
}

// Pack assembles the service family structure in the given buffer.
func (f *ServiceFamily) Pack(buffer []byte) {
	util.PackSome(
		buffer,
		uint8(f.Type), f.Version,
	)
}

// Unpack parses the given data in order to initialize the structure.
func (f *ServiceFamily) Unpack(data []byte) (n uint, err error) {
	return util.UnpackSome(data, (*uint8)(&f.Type), &f.Version)
}

// DescriptionBlock is returned by a Search Request, a Description Request
// or a DiagnosticRequest. Not all DIBs may be returned.
type DescriptionBlock struct {
	DeviceHardware    DeviceInformationBlock
	SupportedServices SupportedServicesDIB
	IpConfig          IpConfigDIB
	CurConfig         CurConfigDIB
	KnxAddr           KnxAddrDIB
	UnknownBlocks     []UnknownDescriptionBlock
}

// Unpack parses the given service payload in order to initialize the Description Block.
// It can cope with not in sequence and unknown Device Information Blocks (DIB).
func (di *DescriptionBlock) Unpack(data []byte) (n uint, err error) {
	var length uint8
	var ty DescriptionType

	n = 0
	for n < uint(len(data)) {
		// DIBs should always have a length and a type.
		_, err := util.UnpackSome(data[n:], &length, (*uint8)(&ty))
		if err != nil {
			return 0, err
		}

		switch ty {
		case DescriptionTypeDeviceInfo:
			_, err = di.DeviceHardware.Unpack(data[n : n+uint(length)])
			if err != nil {
				return 0, err
			}
			n += uint(length)

		case DescriptionTypeSupportedServiceFamilies:
			_, err = di.SupportedServices.Unpack(data[n : n+uint(length)])
			if err != nil {
				return 0, err
			}
			n += uint(length)

		case DescriptionTypeIPConfig:
			_, err = di.IpConfig.Unpack(data[n : n+uint(length)])
			if err != nil {
				return 0, err
			}
			n += uint(length)

		case DescriptionTypeIPCurrentConfig:
			_, err = di.CurConfig.Unpack(data[n : n+uint(length)])
			if err != nil {
				return 0, err
			}
			n += uint(length)

		case DescriptionTypeKNXAddresses:
			_, err = di.KnxAddr.Unpack(data[n : n+uint(length)])
			if err != nil {
				return 0, err
			}
			n += uint(length)

		case DescriptionTypeManufacturerData:
			u := UnknownDescriptionBlock{Type: ty}

			// known DIBs without data will be silently ignored.
			if length > 2 {
				_, err = u.Unpack(data[n+2 : n+uint(length)-2])
				if err != nil {
					return 0, err
				}
				di.UnknownBlocks = append(di.UnknownBlocks, u)
				util.Log(di, "DIB not parsed: 0x%02x", ty)
			}
			n += uint(length)

		default:
			util.Log(di, "Found unsupported DIB with code: 0x%02x", ty)
			n += uint(length)
		}
	}

	return n, err
}

// UnknownDescriptionBlock is a placeholder for unknown DIBs.
type UnknownDescriptionBlock struct {
	Type DescriptionType
	Data []byte
}

// Unpack Unknown Description Blocks into a buffer.
func (u *UnknownDescriptionBlock) Unpack(data []byte) (n uint, err error) {
	u.Data = make([]byte, len(data))
	return util.UnpackSome(data, u.Data)
}
