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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// MeasurementDataChannelMeasurementData is the corresponding interface of MeasurementDataChannelMeasurementData
type MeasurementDataChannelMeasurementData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeasurementData
	// GetDeviceId returns DeviceId (property field)
	GetDeviceId() uint8
	// GetChannel returns Channel (property field)
	GetChannel() uint8
	// GetUnits returns Units (property field)
	GetUnits() MeasurementUnits
	// GetMultiplier returns Multiplier (property field)
	GetMultiplier() int8
	// GetMsb returns Msb (property field)
	GetMsb() uint8
	// GetLsb returns Lsb (property field)
	GetLsb() uint8
	// GetRawValue returns RawValue (virtual field)
	GetRawValue() uint16
	// GetValue returns Value (virtual field)
	GetValue() float64
}

// MeasurementDataChannelMeasurementDataExactly can be used when we want exactly this type and not a type which fulfills MeasurementDataChannelMeasurementData.
// This is useful for switch cases.
type MeasurementDataChannelMeasurementDataExactly interface {
	MeasurementDataChannelMeasurementData
	isMeasurementDataChannelMeasurementData() bool
}

// _MeasurementDataChannelMeasurementData is the data-structure of this message
type _MeasurementDataChannelMeasurementData struct {
	*_MeasurementData
	DeviceId   uint8
	Channel    uint8
	Units      MeasurementUnits
	Multiplier int8
	Msb        uint8
	Lsb        uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeasurementDataChannelMeasurementData) InitializeParent(parent MeasurementData, commandTypeContainer MeasurementCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_MeasurementDataChannelMeasurementData) GetParent() MeasurementData {
	return m._MeasurementData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeasurementDataChannelMeasurementData) GetDeviceId() uint8 {
	return m.DeviceId
}

func (m *_MeasurementDataChannelMeasurementData) GetChannel() uint8 {
	return m.Channel
}

func (m *_MeasurementDataChannelMeasurementData) GetUnits() MeasurementUnits {
	return m.Units
}

func (m *_MeasurementDataChannelMeasurementData) GetMultiplier() int8 {
	return m.Multiplier
}

func (m *_MeasurementDataChannelMeasurementData) GetMsb() uint8 {
	return m.Msb
}

func (m *_MeasurementDataChannelMeasurementData) GetLsb() uint8 {
	return m.Lsb
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MeasurementDataChannelMeasurementData) GetRawValue() uint16 {
	ctx := context.Background()
	_ = ctx
	return uint16(m.GetMsb()<<uint16(8) | m.GetLsb())
}

func (m *_MeasurementDataChannelMeasurementData) GetValue() float64 {
	ctx := context.Background()
	_ = ctx
	return float64(float64(float64(m.GetRawValue())*float64(m.GetMultiplier())) * float64(float64(10)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeasurementDataChannelMeasurementData factory function for _MeasurementDataChannelMeasurementData
func NewMeasurementDataChannelMeasurementData(deviceId uint8, channel uint8, units MeasurementUnits, multiplier int8, msb uint8, lsb uint8, commandTypeContainer MeasurementCommandTypeContainer) *_MeasurementDataChannelMeasurementData {
	_result := &_MeasurementDataChannelMeasurementData{
		DeviceId:         deviceId,
		Channel:          channel,
		Units:            units,
		Multiplier:       multiplier,
		Msb:              msb,
		Lsb:              lsb,
		_MeasurementData: NewMeasurementData(commandTypeContainer),
	}
	_result._MeasurementData._MeasurementDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeasurementDataChannelMeasurementData(structType any) MeasurementDataChannelMeasurementData {
	if casted, ok := structType.(MeasurementDataChannelMeasurementData); ok {
		return casted
	}
	if casted, ok := structType.(*MeasurementDataChannelMeasurementData); ok {
		return *casted
	}
	return nil
}

func (m *_MeasurementDataChannelMeasurementData) GetTypeName() string {
	return "MeasurementDataChannelMeasurementData"
}

func (m *_MeasurementDataChannelMeasurementData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (deviceId)
	lengthInBits += 8

	// Simple field (channel)
	lengthInBits += 8

	// Simple field (units)
	lengthInBits += 8

	// Simple field (multiplier)
	lengthInBits += 8

	// Simple field (msb)
	lengthInBits += 8

	// Simple field (lsb)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MeasurementDataChannelMeasurementData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeasurementDataChannelMeasurementDataParse(ctx context.Context, theBytes []byte) (MeasurementDataChannelMeasurementData, error) {
	return MeasurementDataChannelMeasurementDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MeasurementDataChannelMeasurementDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeasurementDataChannelMeasurementData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MeasurementDataChannelMeasurementData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeasurementDataChannelMeasurementData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceId)
	_deviceId, _deviceIdErr := readBuffer.ReadUint8("deviceId", 8)
	if _deviceIdErr != nil {
		return nil, errors.Wrap(_deviceIdErr, "Error parsing 'deviceId' field of MeasurementDataChannelMeasurementData")
	}
	deviceId := _deviceId

	// Simple Field (channel)
	_channel, _channelErr := readBuffer.ReadUint8("channel", 8)
	if _channelErr != nil {
		return nil, errors.Wrap(_channelErr, "Error parsing 'channel' field of MeasurementDataChannelMeasurementData")
	}
	channel := _channel

	// Simple Field (units)
	if pullErr := readBuffer.PullContext("units"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for units")
	}
	_units, _unitsErr := MeasurementUnitsParseWithBuffer(ctx, readBuffer)
	if _unitsErr != nil {
		return nil, errors.Wrap(_unitsErr, "Error parsing 'units' field of MeasurementDataChannelMeasurementData")
	}
	units := _units
	if closeErr := readBuffer.CloseContext("units"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for units")
	}

	// Simple Field (multiplier)
	_multiplier, _multiplierErr := readBuffer.ReadInt8("multiplier", 8)
	if _multiplierErr != nil {
		return nil, errors.Wrap(_multiplierErr, "Error parsing 'multiplier' field of MeasurementDataChannelMeasurementData")
	}
	multiplier := _multiplier

	// Simple Field (msb)
	_msb, _msbErr := readBuffer.ReadUint8("msb", 8)
	if _msbErr != nil {
		return nil, errors.Wrap(_msbErr, "Error parsing 'msb' field of MeasurementDataChannelMeasurementData")
	}
	msb := _msb

	// Simple Field (lsb)
	_lsb, _lsbErr := readBuffer.ReadUint8("lsb", 8)
	if _lsbErr != nil {
		return nil, errors.Wrap(_lsbErr, "Error parsing 'lsb' field of MeasurementDataChannelMeasurementData")
	}
	lsb := _lsb

	// Virtual field
	_rawValue := msb<<uint16(8) | lsb
	rawValue := uint16(_rawValue)
	_ = rawValue

	// Virtual field
	_value := float64(float64(rawValue)*float64(multiplier)) * float64(float64(10))
	value := float64(_value)
	_ = value

	if closeErr := readBuffer.CloseContext("MeasurementDataChannelMeasurementData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeasurementDataChannelMeasurementData")
	}

	// Create a partially initialized instance
	_child := &_MeasurementDataChannelMeasurementData{
		_MeasurementData: &_MeasurementData{},
		DeviceId:         deviceId,
		Channel:          channel,
		Units:            units,
		Multiplier:       multiplier,
		Msb:              msb,
		Lsb:              lsb,
	}
	_child._MeasurementData._MeasurementDataChildRequirements = _child
	return _child, nil
}

func (m *_MeasurementDataChannelMeasurementData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeasurementDataChannelMeasurementData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeasurementDataChannelMeasurementData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeasurementDataChannelMeasurementData")
		}

		// Simple Field (deviceId)
		deviceId := uint8(m.GetDeviceId())
		_deviceIdErr := writeBuffer.WriteUint8("deviceId", 8, uint8((deviceId)))
		if _deviceIdErr != nil {
			return errors.Wrap(_deviceIdErr, "Error serializing 'deviceId' field")
		}

		// Simple Field (channel)
		channel := uint8(m.GetChannel())
		_channelErr := writeBuffer.WriteUint8("channel", 8, uint8((channel)))
		if _channelErr != nil {
			return errors.Wrap(_channelErr, "Error serializing 'channel' field")
		}

		// Simple Field (units)
		if pushErr := writeBuffer.PushContext("units"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for units")
		}
		_unitsErr := writeBuffer.WriteSerializable(ctx, m.GetUnits())
		if popErr := writeBuffer.PopContext("units"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for units")
		}
		if _unitsErr != nil {
			return errors.Wrap(_unitsErr, "Error serializing 'units' field")
		}

		// Simple Field (multiplier)
		multiplier := int8(m.GetMultiplier())
		_multiplierErr := writeBuffer.WriteInt8("multiplier", 8, int8((multiplier)))
		if _multiplierErr != nil {
			return errors.Wrap(_multiplierErr, "Error serializing 'multiplier' field")
		}

		// Simple Field (msb)
		msb := uint8(m.GetMsb())
		_msbErr := writeBuffer.WriteUint8("msb", 8, uint8((msb)))
		if _msbErr != nil {
			return errors.Wrap(_msbErr, "Error serializing 'msb' field")
		}

		// Simple Field (lsb)
		lsb := uint8(m.GetLsb())
		_lsbErr := writeBuffer.WriteUint8("lsb", 8, uint8((lsb)))
		if _lsbErr != nil {
			return errors.Wrap(_lsbErr, "Error serializing 'lsb' field")
		}
		// Virtual field
		rawValue := m.GetRawValue()
		_ = rawValue
		if _rawValueErr := writeBuffer.WriteVirtual(ctx, "rawValue", m.GetRawValue()); _rawValueErr != nil {
			return errors.Wrap(_rawValueErr, "Error serializing 'rawValue' field")
		}
		// Virtual field
		value := m.GetValue()
		_ = value
		if _valueErr := writeBuffer.WriteVirtual(ctx, "value", m.GetValue()); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("MeasurementDataChannelMeasurementData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeasurementDataChannelMeasurementData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeasurementDataChannelMeasurementData) isMeasurementDataChannelMeasurementData() bool {
	return true
}

func (m *_MeasurementDataChannelMeasurementData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
