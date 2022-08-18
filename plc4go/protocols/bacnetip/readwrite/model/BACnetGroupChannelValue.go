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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetGroupChannelValue is the corresponding interface of BACnetGroupChannelValue
type BACnetGroupChannelValue interface {
	utils.LengthAware
	utils.Serializable
	// GetChannel returns Channel (property field)
	GetChannel() BACnetContextTagUnsignedInteger
	// GetOverridingPriority returns OverridingPriority (property field)
	GetOverridingPriority() BACnetContextTagUnsignedInteger
	// GetValue returns Value (property field)
	GetValue() BACnetChannelValue
}

// BACnetGroupChannelValueExactly can be used when we want exactly this type and not a type which fulfills BACnetGroupChannelValue.
// This is useful for switch cases.
type BACnetGroupChannelValueExactly interface {
	BACnetGroupChannelValue
	isBACnetGroupChannelValue() bool
}

// _BACnetGroupChannelValue is the data-structure of this message
type _BACnetGroupChannelValue struct {
	Channel            BACnetContextTagUnsignedInteger
	OverridingPriority BACnetContextTagUnsignedInteger
	Value              BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetGroupChannelValue) GetChannel() BACnetContextTagUnsignedInteger {
	return m.Channel
}

func (m *_BACnetGroupChannelValue) GetOverridingPriority() BACnetContextTagUnsignedInteger {
	return m.OverridingPriority
}

func (m *_BACnetGroupChannelValue) GetValue() BACnetChannelValue {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetGroupChannelValue factory function for _BACnetGroupChannelValue
func NewBACnetGroupChannelValue(channel BACnetContextTagUnsignedInteger, overridingPriority BACnetContextTagUnsignedInteger, value BACnetChannelValue) *_BACnetGroupChannelValue {
	return &_BACnetGroupChannelValue{Channel: channel, OverridingPriority: overridingPriority, Value: value}
}

// Deprecated: use the interface for direct cast
func CastBACnetGroupChannelValue(structType interface{}) BACnetGroupChannelValue {
	if casted, ok := structType.(BACnetGroupChannelValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetGroupChannelValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetGroupChannelValue) GetTypeName() string {
	return "BACnetGroupChannelValue"
}

func (m *_BACnetGroupChannelValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetGroupChannelValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (channel)
	lengthInBits += m.Channel.GetLengthInBits()

	// Optional Field (overridingPriority)
	if m.OverridingPriority != nil {
		lengthInBits += m.OverridingPriority.GetLengthInBits()
	}

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetGroupChannelValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetGroupChannelValueParse(readBuffer utils.ReadBuffer) (BACnetGroupChannelValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetGroupChannelValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetGroupChannelValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (channel)
	if pullErr := readBuffer.PullContext("channel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channel")
	}
	_channel, _channelErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _channelErr != nil {
		return nil, errors.Wrap(_channelErr, "Error parsing 'channel' field of BACnetGroupChannelValue")
	}
	channel := _channel.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("channel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channel")
	}

	// Optional Field (overridingPriority) (Can be skipped, if a given expression evaluates to false)
	var overridingPriority BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("overridingPriority"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for overridingPriority")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'overridingPriority' field of BACnetGroupChannelValue")
		default:
			overridingPriority = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("overridingPriority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for overridingPriority")
			}
		}
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := BACnetChannelValueParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetGroupChannelValue")
	}
	value := _value.(BACnetChannelValue)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("BACnetGroupChannelValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetGroupChannelValue")
	}

	// Create the instance
	return &_BACnetGroupChannelValue{
		Channel:            channel,
		OverridingPriority: overridingPriority,
		Value:              value,
	}, nil
}

func (m *_BACnetGroupChannelValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetGroupChannelValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetGroupChannelValue")
	}

	// Simple Field (channel)
	if pushErr := writeBuffer.PushContext("channel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for channel")
	}
	_channelErr := writeBuffer.WriteSerializable(m.GetChannel())
	if popErr := writeBuffer.PopContext("channel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for channel")
	}
	if _channelErr != nil {
		return errors.Wrap(_channelErr, "Error serializing 'channel' field")
	}

	// Optional Field (overridingPriority) (Can be skipped, if the value is null)
	var overridingPriority BACnetContextTagUnsignedInteger = nil
	if m.GetOverridingPriority() != nil {
		if pushErr := writeBuffer.PushContext("overridingPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for overridingPriority")
		}
		overridingPriority = m.GetOverridingPriority()
		_overridingPriorityErr := writeBuffer.WriteSerializable(overridingPriority)
		if popErr := writeBuffer.PopContext("overridingPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for overridingPriority")
		}
		if _overridingPriorityErr != nil {
			return errors.Wrap(_overridingPriorityErr, "Error serializing 'overridingPriority' field")
		}
	}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for value")
	}
	_valueErr := writeBuffer.WriteSerializable(m.GetValue())
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for value")
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetGroupChannelValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetGroupChannelValue")
	}
	return nil
}

func (m *_BACnetGroupChannelValue) isBACnetGroupChannelValue() bool {
	return true
}

func (m *_BACnetGroupChannelValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
