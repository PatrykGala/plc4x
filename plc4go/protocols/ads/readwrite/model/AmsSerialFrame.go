/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AmsSerialFrame is the corresponding interface of AmsSerialFrame
type AmsSerialFrame interface {
	utils.LengthAware
	utils.Serializable
	// GetMagicCookie returns MagicCookie (property field)
	GetMagicCookie() uint16
	// GetTransmitterAddress returns TransmitterAddress (property field)
	GetTransmitterAddress() int8
	// GetReceiverAddress returns ReceiverAddress (property field)
	GetReceiverAddress() int8
	// GetFragmentNumber returns FragmentNumber (property field)
	GetFragmentNumber() int8
	// GetLength returns Length (property field)
	GetLength() int8
	// GetUserdata returns Userdata (property field)
	GetUserdata() AmsPacket
	// GetCrc returns Crc (property field)
	GetCrc() uint16
}

// AmsSerialFrameExactly can be used when we want exactly this type and not a type which fulfills AmsSerialFrame.
// This is useful for switch cases.
type AmsSerialFrameExactly interface {
	AmsSerialFrame
	isAmsSerialFrame() bool
}

// _AmsSerialFrame is the data-structure of this message
type _AmsSerialFrame struct {
	MagicCookie        uint16
	TransmitterAddress int8
	ReceiverAddress    int8
	FragmentNumber     int8
	Length             int8
	Userdata           AmsPacket
	Crc                uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsSerialFrame) GetMagicCookie() uint16 {
	return m.MagicCookie
}

func (m *_AmsSerialFrame) GetTransmitterAddress() int8 {
	return m.TransmitterAddress
}

func (m *_AmsSerialFrame) GetReceiverAddress() int8 {
	return m.ReceiverAddress
}

func (m *_AmsSerialFrame) GetFragmentNumber() int8 {
	return m.FragmentNumber
}

func (m *_AmsSerialFrame) GetLength() int8 {
	return m.Length
}

func (m *_AmsSerialFrame) GetUserdata() AmsPacket {
	return m.Userdata
}

func (m *_AmsSerialFrame) GetCrc() uint16 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAmsSerialFrame factory function for _AmsSerialFrame
func NewAmsSerialFrame(magicCookie uint16, transmitterAddress int8, receiverAddress int8, fragmentNumber int8, length int8, userdata AmsPacket, crc uint16) *_AmsSerialFrame {
	return &_AmsSerialFrame{MagicCookie: magicCookie, TransmitterAddress: transmitterAddress, ReceiverAddress: receiverAddress, FragmentNumber: fragmentNumber, Length: length, Userdata: userdata, Crc: crc}
}

// Deprecated: use the interface for direct cast
func CastAmsSerialFrame(structType interface{}) AmsSerialFrame {
	if casted, ok := structType.(AmsSerialFrame); ok {
		return casted
	}
	if casted, ok := structType.(*AmsSerialFrame); ok {
		return *casted
	}
	return nil
}

func (m *_AmsSerialFrame) GetTypeName() string {
	return "AmsSerialFrame"
}

func (m *_AmsSerialFrame) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AmsSerialFrame) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (magicCookie)
	lengthInBits += 16

	// Simple field (transmitterAddress)
	lengthInBits += 8

	// Simple field (receiverAddress)
	lengthInBits += 8

	// Simple field (fragmentNumber)
	lengthInBits += 8

	// Simple field (length)
	lengthInBits += 8

	// Simple field (userdata)
	lengthInBits += m.Userdata.GetLengthInBits()

	// Simple field (crc)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AmsSerialFrame) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AmsSerialFrameParse(readBuffer utils.ReadBuffer) (AmsSerialFrame, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsSerialFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsSerialFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (magicCookie)
	_magicCookie, _magicCookieErr := readBuffer.ReadUint16("magicCookie", 16)
	if _magicCookieErr != nil {
		return nil, errors.Wrap(_magicCookieErr, "Error parsing 'magicCookie' field of AmsSerialFrame")
	}
	magicCookie := _magicCookie

	// Simple Field (transmitterAddress)
	_transmitterAddress, _transmitterAddressErr := readBuffer.ReadInt8("transmitterAddress", 8)
	if _transmitterAddressErr != nil {
		return nil, errors.Wrap(_transmitterAddressErr, "Error parsing 'transmitterAddress' field of AmsSerialFrame")
	}
	transmitterAddress := _transmitterAddress

	// Simple Field (receiverAddress)
	_receiverAddress, _receiverAddressErr := readBuffer.ReadInt8("receiverAddress", 8)
	if _receiverAddressErr != nil {
		return nil, errors.Wrap(_receiverAddressErr, "Error parsing 'receiverAddress' field of AmsSerialFrame")
	}
	receiverAddress := _receiverAddress

	// Simple Field (fragmentNumber)
	_fragmentNumber, _fragmentNumberErr := readBuffer.ReadInt8("fragmentNumber", 8)
	if _fragmentNumberErr != nil {
		return nil, errors.Wrap(_fragmentNumberErr, "Error parsing 'fragmentNumber' field of AmsSerialFrame")
	}
	fragmentNumber := _fragmentNumber

	// Simple Field (length)
	_length, _lengthErr := readBuffer.ReadInt8("length", 8)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AmsSerialFrame")
	}
	length := _length

	// Simple Field (userdata)
	if pullErr := readBuffer.PullContext("userdata"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userdata")
	}
	_userdata, _userdataErr := AmsPacketParse(readBuffer)
	if _userdataErr != nil {
		return nil, errors.Wrap(_userdataErr, "Error parsing 'userdata' field of AmsSerialFrame")
	}
	userdata := _userdata.(AmsPacket)
	if closeErr := readBuffer.CloseContext("userdata"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userdata")
	}

	// Simple Field (crc)
	_crc, _crcErr := readBuffer.ReadUint16("crc", 16)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field of AmsSerialFrame")
	}
	crc := _crc

	if closeErr := readBuffer.CloseContext("AmsSerialFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsSerialFrame")
	}

	// Create the instance
	return NewAmsSerialFrame(magicCookie, transmitterAddress, receiverAddress, fragmentNumber, length, userdata, crc), nil
}

func (m *_AmsSerialFrame) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AmsSerialFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsSerialFrame")
	}

	// Simple Field (magicCookie)
	magicCookie := uint16(m.GetMagicCookie())
	_magicCookieErr := writeBuffer.WriteUint16("magicCookie", 16, (magicCookie))
	if _magicCookieErr != nil {
		return errors.Wrap(_magicCookieErr, "Error serializing 'magicCookie' field")
	}

	// Simple Field (transmitterAddress)
	transmitterAddress := int8(m.GetTransmitterAddress())
	_transmitterAddressErr := writeBuffer.WriteInt8("transmitterAddress", 8, (transmitterAddress))
	if _transmitterAddressErr != nil {
		return errors.Wrap(_transmitterAddressErr, "Error serializing 'transmitterAddress' field")
	}

	// Simple Field (receiverAddress)
	receiverAddress := int8(m.GetReceiverAddress())
	_receiverAddressErr := writeBuffer.WriteInt8("receiverAddress", 8, (receiverAddress))
	if _receiverAddressErr != nil {
		return errors.Wrap(_receiverAddressErr, "Error serializing 'receiverAddress' field")
	}

	// Simple Field (fragmentNumber)
	fragmentNumber := int8(m.GetFragmentNumber())
	_fragmentNumberErr := writeBuffer.WriteInt8("fragmentNumber", 8, (fragmentNumber))
	if _fragmentNumberErr != nil {
		return errors.Wrap(_fragmentNumberErr, "Error serializing 'fragmentNumber' field")
	}

	// Simple Field (length)
	length := int8(m.GetLength())
	_lengthErr := writeBuffer.WriteInt8("length", 8, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (userdata)
	if pushErr := writeBuffer.PushContext("userdata"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for userdata")
	}
	_userdataErr := writeBuffer.WriteSerializable(m.GetUserdata())
	if popErr := writeBuffer.PopContext("userdata"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for userdata")
	}
	if _userdataErr != nil {
		return errors.Wrap(_userdataErr, "Error serializing 'userdata' field")
	}

	// Simple Field (crc)
	crc := uint16(m.GetCrc())
	_crcErr := writeBuffer.WriteUint16("crc", 16, (crc))
	if _crcErr != nil {
		return errors.Wrap(_crcErr, "Error serializing 'crc' field")
	}

	if popErr := writeBuffer.PopContext("AmsSerialFrame"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsSerialFrame")
	}
	return nil
}

func (m *_AmsSerialFrame) isAmsSerialFrame() bool {
	return true
}

func (m *_AmsSerialFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
