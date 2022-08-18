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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageSubscribeDigitalPinValue is the corresponding interface of FirmataMessageSubscribeDigitalPinValue
type FirmataMessageSubscribeDigitalPinValue interface {
	utils.LengthAware
	utils.Serializable
	FirmataMessage
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetEnable returns Enable (property field)
	GetEnable() bool
}

// FirmataMessageSubscribeDigitalPinValueExactly can be used when we want exactly this type and not a type which fulfills FirmataMessageSubscribeDigitalPinValue.
// This is useful for switch cases.
type FirmataMessageSubscribeDigitalPinValueExactly interface {
	FirmataMessageSubscribeDigitalPinValue
	isFirmataMessageSubscribeDigitalPinValue() bool
}

// _FirmataMessageSubscribeDigitalPinValue is the data-structure of this message
type _FirmataMessageSubscribeDigitalPinValue struct {
	*_FirmataMessage
	Pin    uint8
	Enable bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageSubscribeDigitalPinValue) GetMessageType() uint8 {
	return 0xD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageSubscribeDigitalPinValue) InitializeParent(parent FirmataMessage) {}

func (m *_FirmataMessageSubscribeDigitalPinValue) GetParent() FirmataMessage {
	return m._FirmataMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageSubscribeDigitalPinValue) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataMessageSubscribeDigitalPinValue) GetEnable() bool {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataMessageSubscribeDigitalPinValue factory function for _FirmataMessageSubscribeDigitalPinValue
func NewFirmataMessageSubscribeDigitalPinValue(pin uint8, enable bool, response bool) *_FirmataMessageSubscribeDigitalPinValue {
	_result := &_FirmataMessageSubscribeDigitalPinValue{
		Pin:             pin,
		Enable:          enable,
		_FirmataMessage: NewFirmataMessage(response),
	}
	_result._FirmataMessage._FirmataMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataMessageSubscribeDigitalPinValue(structType interface{}) FirmataMessageSubscribeDigitalPinValue {
	if casted, ok := structType.(FirmataMessageSubscribeDigitalPinValue); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageSubscribeDigitalPinValue); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageSubscribeDigitalPinValue) GetTypeName() string {
	return "FirmataMessageSubscribeDigitalPinValue"
}

func (m *_FirmataMessageSubscribeDigitalPinValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_FirmataMessageSubscribeDigitalPinValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 4

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enable)
	lengthInBits += 1

	return lengthInBits
}

func (m *_FirmataMessageSubscribeDigitalPinValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataMessageSubscribeDigitalPinValueParse(readBuffer utils.ReadBuffer, response bool) (FirmataMessageSubscribeDigitalPinValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageSubscribeDigitalPinValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageSubscribeDigitalPinValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 4)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field of FirmataMessageSubscribeDigitalPinValue")
	}
	pin := _pin

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of FirmataMessageSubscribeDigitalPinValue")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (enable)
	_enable, _enableErr := readBuffer.ReadBit("enable")
	if _enableErr != nil {
		return nil, errors.Wrap(_enableErr, "Error parsing 'enable' field of FirmataMessageSubscribeDigitalPinValue")
	}
	enable := _enable

	if closeErr := readBuffer.CloseContext("FirmataMessageSubscribeDigitalPinValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageSubscribeDigitalPinValue")
	}

	// Create a partially initialized instance
	_child := &_FirmataMessageSubscribeDigitalPinValue{
		_FirmataMessage: &_FirmataMessage{
			Response: response,
		},
		Pin:            pin,
		Enable:         enable,
		reservedField0: reservedField0,
	}
	_child._FirmataMessage._FirmataMessageChildRequirements = _child
	return _child, nil
}

func (m *_FirmataMessageSubscribeDigitalPinValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageSubscribeDigitalPinValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageSubscribeDigitalPinValue")
		}

		// Simple Field (pin)
		pin := uint8(m.GetPin())
		_pinErr := writeBuffer.WriteUint8("pin", 4, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (enable)
		enable := bool(m.GetEnable())
		_enableErr := writeBuffer.WriteBit("enable", (enable))
		if _enableErr != nil {
			return errors.Wrap(_enableErr, "Error serializing 'enable' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageSubscribeDigitalPinValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageSubscribeDigitalPinValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_FirmataMessageSubscribeDigitalPinValue) isFirmataMessageSubscribeDigitalPinValue() bool {
	return true
}

func (m *_FirmataMessageSubscribeDigitalPinValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
