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

// BACnetConstructedDataRequestedShedLevel is the corresponding interface of BACnetConstructedDataRequestedShedLevel
type BACnetConstructedDataRequestedShedLevel interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRequestedShedLevel returns RequestedShedLevel (property field)
	GetRequestedShedLevel() BACnetShedLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetShedLevel
}

// BACnetConstructedDataRequestedShedLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRequestedShedLevel.
// This is useful for switch cases.
type BACnetConstructedDataRequestedShedLevelExactly interface {
	BACnetConstructedDataRequestedShedLevel
	isBACnetConstructedDataRequestedShedLevel() bool
}

// _BACnetConstructedDataRequestedShedLevel is the data-structure of this message
type _BACnetConstructedDataRequestedShedLevel struct {
	*_BACnetConstructedData
	RequestedShedLevel BACnetShedLevel
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REQUESTED_SHED_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) GetRequestedShedLevel() BACnetShedLevel {
	return m.RequestedShedLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) GetActualValue() BACnetShedLevel {
	return CastBACnetShedLevel(m.GetRequestedShedLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRequestedShedLevel factory function for _BACnetConstructedDataRequestedShedLevel
func NewBACnetConstructedDataRequestedShedLevel(requestedShedLevel BACnetShedLevel, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRequestedShedLevel {
	_result := &_BACnetConstructedDataRequestedShedLevel{
		RequestedShedLevel:     requestedShedLevel,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRequestedShedLevel(structType interface{}) BACnetConstructedDataRequestedShedLevel {
	if casted, ok := structType.(BACnetConstructedDataRequestedShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRequestedShedLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetTypeName() string {
	return "BACnetConstructedDataRequestedShedLevel"
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (requestedShedLevel)
	lengthInBits += m.RequestedShedLevel.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRequestedShedLevelParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRequestedShedLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRequestedShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRequestedShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestedShedLevel)
	if pullErr := readBuffer.PullContext("requestedShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestedShedLevel")
	}
	_requestedShedLevel, _requestedShedLevelErr := BACnetShedLevelParse(readBuffer)
	if _requestedShedLevelErr != nil {
		return nil, errors.Wrap(_requestedShedLevelErr, "Error parsing 'requestedShedLevel' field of BACnetConstructedDataRequestedShedLevel")
	}
	requestedShedLevel := _requestedShedLevel.(BACnetShedLevel)
	if closeErr := readBuffer.CloseContext("requestedShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestedShedLevel")
	}

	// Virtual field
	_actualValue := requestedShedLevel
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRequestedShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRequestedShedLevel")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRequestedShedLevel{
		RequestedShedLevel: requestedShedLevel,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRequestedShedLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRequestedShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRequestedShedLevel")
		}

		// Simple Field (requestedShedLevel)
		if pushErr := writeBuffer.PushContext("requestedShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestedShedLevel")
		}
		_requestedShedLevelErr := writeBuffer.WriteSerializable(m.GetRequestedShedLevel())
		if popErr := writeBuffer.PopContext("requestedShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestedShedLevel")
		}
		if _requestedShedLevelErr != nil {
			return errors.Wrap(_requestedShedLevelErr, "Error serializing 'requestedShedLevel' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRequestedShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRequestedShedLevel")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRequestedShedLevel) isBACnetConstructedDataRequestedShedLevel() bool {
	return true
}

func (m *_BACnetConstructedDataRequestedShedLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
