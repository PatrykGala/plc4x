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

// BACnetConstructedDataExtendedTimeEnable is the corresponding interface of BACnetConstructedDataExtendedTimeEnable
type BACnetConstructedDataExtendedTimeEnable interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetExtendedTimeEnable returns ExtendedTimeEnable (property field)
	GetExtendedTimeEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataExtendedTimeEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataExtendedTimeEnable.
// This is useful for switch cases.
type BACnetConstructedDataExtendedTimeEnableExactly interface {
	BACnetConstructedDataExtendedTimeEnable
	isBACnetConstructedDataExtendedTimeEnable() bool
}

// _BACnetConstructedDataExtendedTimeEnable is the data-structure of this message
type _BACnetConstructedDataExtendedTimeEnable struct {
	*_BACnetConstructedData
	ExtendedTimeEnable BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXTENDED_TIME_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) GetExtendedTimeEnable() BACnetApplicationTagBoolean {
	return m.ExtendedTimeEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetExtendedTimeEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataExtendedTimeEnable factory function for _BACnetConstructedDataExtendedTimeEnable
func NewBACnetConstructedDataExtendedTimeEnable(extendedTimeEnable BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExtendedTimeEnable {
	_result := &_BACnetConstructedDataExtendedTimeEnable{
		ExtendedTimeEnable:     extendedTimeEnable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExtendedTimeEnable(structType interface{}) BACnetConstructedDataExtendedTimeEnable {
	if casted, ok := structType.(BACnetConstructedDataExtendedTimeEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExtendedTimeEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetTypeName() string {
	return "BACnetConstructedDataExtendedTimeEnable"
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (extendedTimeEnable)
	lengthInBits += m.ExtendedTimeEnable.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataExtendedTimeEnableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataExtendedTimeEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExtendedTimeEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExtendedTimeEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (extendedTimeEnable)
	if pullErr := readBuffer.PullContext("extendedTimeEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for extendedTimeEnable")
	}
	_extendedTimeEnable, _extendedTimeEnableErr := BACnetApplicationTagParse(readBuffer)
	if _extendedTimeEnableErr != nil {
		return nil, errors.Wrap(_extendedTimeEnableErr, "Error parsing 'extendedTimeEnable' field of BACnetConstructedDataExtendedTimeEnable")
	}
	extendedTimeEnable := _extendedTimeEnable.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("extendedTimeEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for extendedTimeEnable")
	}

	// Virtual field
	_actualValue := extendedTimeEnable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExtendedTimeEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExtendedTimeEnable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataExtendedTimeEnable{
		ExtendedTimeEnable: extendedTimeEnable,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataExtendedTimeEnable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExtendedTimeEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExtendedTimeEnable")
		}

		// Simple Field (extendedTimeEnable)
		if pushErr := writeBuffer.PushContext("extendedTimeEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for extendedTimeEnable")
		}
		_extendedTimeEnableErr := writeBuffer.WriteSerializable(m.GetExtendedTimeEnable())
		if popErr := writeBuffer.PopContext("extendedTimeEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for extendedTimeEnable")
		}
		if _extendedTimeEnableErr != nil {
			return errors.Wrap(_extendedTimeEnableErr, "Error serializing 'extendedTimeEnable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExtendedTimeEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExtendedTimeEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExtendedTimeEnable) isBACnetConstructedDataExtendedTimeEnable() bool {
	return true
}

func (m *_BACnetConstructedDataExtendedTimeEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
