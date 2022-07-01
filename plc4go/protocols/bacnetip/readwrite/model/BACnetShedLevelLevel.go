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

// BACnetShedLevelLevel is the corresponding interface of BACnetShedLevelLevel
type BACnetShedLevelLevel interface {
	utils.LengthAware
	utils.Serializable
	BACnetShedLevel
	// GetLevel returns Level (property field)
	GetLevel() BACnetContextTagUnsignedInteger
}

// BACnetShedLevelLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetShedLevelLevel.
// This is useful for switch cases.
type BACnetShedLevelLevelExactly interface {
	BACnetShedLevelLevel
	isBACnetShedLevelLevel() bool
}

// _BACnetShedLevelLevel is the data-structure of this message
type _BACnetShedLevelLevel struct {
	*_BACnetShedLevel
	Level BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetShedLevelLevel) InitializeParent(parent BACnetShedLevel, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetShedLevelLevel) GetParent() BACnetShedLevel {
	return m._BACnetShedLevel
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelLevel) GetLevel() BACnetContextTagUnsignedInteger {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetShedLevelLevel factory function for _BACnetShedLevelLevel
func NewBACnetShedLevelLevel(level BACnetContextTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetShedLevelLevel {
	_result := &_BACnetShedLevelLevel{
		Level:            level,
		_BACnetShedLevel: NewBACnetShedLevel(peekedTagHeader),
	}
	_result._BACnetShedLevel._BACnetShedLevelChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelLevel(structType interface{}) BACnetShedLevelLevel {
	if casted, ok := structType.(BACnetShedLevelLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelLevel) GetTypeName() string {
	return "BACnetShedLevelLevel"
}

func (m *_BACnetShedLevelLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetShedLevelLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (level)
	lengthInBits += m.Level.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetShedLevelLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetShedLevelLevelParse(readBuffer utils.ReadBuffer) (BACnetShedLevelLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevelLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (level)
	if pullErr := readBuffer.PullContext("level"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for level")
	}
	_level, _levelErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field of BACnetShedLevelLevel")
	}
	level := _level.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for level")
	}

	if closeErr := readBuffer.CloseContext("BACnetShedLevelLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelLevel")
	}

	// Create a partially initialized instance
	_child := &_BACnetShedLevelLevel{
		Level:            level,
		_BACnetShedLevel: &_BACnetShedLevel{},
	}
	_child._BACnetShedLevel._BACnetShedLevelChildRequirements = _child
	return _child, nil
}

func (m *_BACnetShedLevelLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelLevel")
		}

		// Simple Field (level)
		if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for level")
		}
		_levelErr := writeBuffer.WriteSerializable(m.GetLevel())
		if popErr := writeBuffer.PopContext("level"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for level")
		}
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("BACnetShedLevelLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelLevel")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetShedLevelLevel) isBACnetShedLevelLevel() bool {
	return true
}

func (m *_BACnetShedLevelLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
