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

// BACnetConstructedDataAccumulatorLowLimit is the corresponding interface of BACnetConstructedDataAccumulatorLowLimit
type BACnetConstructedDataAccumulatorLowLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAccumulatorLowLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccumulatorLowLimit.
// This is useful for switch cases.
type BACnetConstructedDataAccumulatorLowLimitExactly interface {
	BACnetConstructedDataAccumulatorLowLimit
	isBACnetConstructedDataAccumulatorLowLimit() bool
}

// _BACnetConstructedDataAccumulatorLowLimit is the data-structure of this message
type _BACnetConstructedDataAccumulatorLowLimit struct {
	*_BACnetConstructedData
	LowLimit BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetLowLimit() BACnetApplicationTagUnsignedInteger {
	return m.LowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccumulatorLowLimit factory function for _BACnetConstructedDataAccumulatorLowLimit
func NewBACnetConstructedDataAccumulatorLowLimit(lowLimit BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorLowLimit {
	_result := &_BACnetConstructedDataAccumulatorLowLimit{
		LowLimit:               lowLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorLowLimit(structType interface{}) BACnetConstructedDataAccumulatorLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorLowLimit"
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccumulatorLowLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccumulatorLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lowLimit)
	if pullErr := readBuffer.PullContext("lowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowLimit")
	}
	_lowLimit, _lowLimitErr := BACnetApplicationTagParse(readBuffer)
	if _lowLimitErr != nil {
		return nil, errors.Wrap(_lowLimitErr, "Error parsing 'lowLimit' field of BACnetConstructedDataAccumulatorLowLimit")
	}
	lowLimit := _lowLimit.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("lowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowLimit")
	}

	// Virtual field
	_actualValue := lowLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorLowLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccumulatorLowLimit{
		LowLimit: lowLimit,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorLowLimit")
		}

		// Simple Field (lowLimit)
		if pushErr := writeBuffer.PushContext("lowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowLimit")
		}
		_lowLimitErr := writeBuffer.WriteSerializable(m.GetLowLimit())
		if popErr := writeBuffer.PopContext("lowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowLimit")
		}
		if _lowLimitErr != nil {
			return errors.Wrap(_lowLimitErr, "Error serializing 'lowLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorLowLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) isBACnetConstructedDataAccumulatorLowLimit() bool {
	return true
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
