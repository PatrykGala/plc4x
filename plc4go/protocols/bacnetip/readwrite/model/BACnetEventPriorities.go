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

// BACnetEventPriorities is the corresponding interface of BACnetEventPriorities
type BACnetEventPriorities interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetToOffnormal returns ToOffnormal (property field)
	GetToOffnormal() BACnetApplicationTagUnsignedInteger
	// GetToFault returns ToFault (property field)
	GetToFault() BACnetApplicationTagUnsignedInteger
	// GetToNormal returns ToNormal (property field)
	GetToNormal() BACnetApplicationTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventPrioritiesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventPriorities.
// This is useful for switch cases.
type BACnetEventPrioritiesExactly interface {
	BACnetEventPriorities
	isBACnetEventPriorities() bool
}

// _BACnetEventPriorities is the data-structure of this message
type _BACnetEventPriorities struct {
	OpeningTag  BACnetOpeningTag
	ToOffnormal BACnetApplicationTagUnsignedInteger
	ToFault     BACnetApplicationTagUnsignedInteger
	ToNormal    BACnetApplicationTagUnsignedInteger
	ClosingTag  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventPriorities) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventPriorities) GetToOffnormal() BACnetApplicationTagUnsignedInteger {
	return m.ToOffnormal
}

func (m *_BACnetEventPriorities) GetToFault() BACnetApplicationTagUnsignedInteger {
	return m.ToFault
}

func (m *_BACnetEventPriorities) GetToNormal() BACnetApplicationTagUnsignedInteger {
	return m.ToNormal
}

func (m *_BACnetEventPriorities) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventPriorities factory function for _BACnetEventPriorities
func NewBACnetEventPriorities(openingTag BACnetOpeningTag, toOffnormal BACnetApplicationTagUnsignedInteger, toFault BACnetApplicationTagUnsignedInteger, toNormal BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventPriorities {
	return &_BACnetEventPriorities{OpeningTag: openingTag, ToOffnormal: toOffnormal, ToFault: toFault, ToNormal: toNormal, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventPriorities(structType interface{}) BACnetEventPriorities {
	if casted, ok := structType.(BACnetEventPriorities); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventPriorities); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventPriorities) GetTypeName() string {
	return "BACnetEventPriorities"
}

func (m *_BACnetEventPriorities) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventPriorities) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits()

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits()

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventPriorities) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventPrioritiesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventPriorities, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventPriorities"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventPriorities")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventPriorities")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (toOffnormal)
	if pullErr := readBuffer.PullContext("toOffnormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toOffnormal")
	}
	_toOffnormal, _toOffnormalErr := BACnetApplicationTagParse(readBuffer)
	if _toOffnormalErr != nil {
		return nil, errors.Wrap(_toOffnormalErr, "Error parsing 'toOffnormal' field of BACnetEventPriorities")
	}
	toOffnormal := _toOffnormal.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("toOffnormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toOffnormal")
	}

	// Simple Field (toFault)
	if pullErr := readBuffer.PullContext("toFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toFault")
	}
	_toFault, _toFaultErr := BACnetApplicationTagParse(readBuffer)
	if _toFaultErr != nil {
		return nil, errors.Wrap(_toFaultErr, "Error parsing 'toFault' field of BACnetEventPriorities")
	}
	toFault := _toFault.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("toFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toFault")
	}

	// Simple Field (toNormal)
	if pullErr := readBuffer.PullContext("toNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toNormal")
	}
	_toNormal, _toNormalErr := BACnetApplicationTagParse(readBuffer)
	if _toNormalErr != nil {
		return nil, errors.Wrap(_toNormalErr, "Error parsing 'toNormal' field of BACnetEventPriorities")
	}
	toNormal := _toNormal.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("toNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toNormal")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventPriorities")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventPriorities"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventPriorities")
	}

	// Create the instance
	return NewBACnetEventPriorities(openingTag, toOffnormal, toFault, toNormal, closingTag, tagNumber), nil
}

func (m *_BACnetEventPriorities) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventPriorities"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventPriorities")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (toOffnormal)
	if pushErr := writeBuffer.PushContext("toOffnormal"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toOffnormal")
	}
	_toOffnormalErr := writeBuffer.WriteSerializable(m.GetToOffnormal())
	if popErr := writeBuffer.PopContext("toOffnormal"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toOffnormal")
	}
	if _toOffnormalErr != nil {
		return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
	}

	// Simple Field (toFault)
	if pushErr := writeBuffer.PushContext("toFault"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toFault")
	}
	_toFaultErr := writeBuffer.WriteSerializable(m.GetToFault())
	if popErr := writeBuffer.PopContext("toFault"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toFault")
	}
	if _toFaultErr != nil {
		return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
	}

	// Simple Field (toNormal)
	if pushErr := writeBuffer.PushContext("toNormal"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toNormal")
	}
	_toNormalErr := writeBuffer.WriteSerializable(m.GetToNormal())
	if popErr := writeBuffer.PopContext("toNormal"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toNormal")
	}
	if _toNormalErr != nil {
		return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventPriorities"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventPriorities")
	}
	return nil
}

func (m *_BACnetEventPriorities) isBACnetEventPriorities() bool {
	return true
}

func (m *_BACnetEventPriorities) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
