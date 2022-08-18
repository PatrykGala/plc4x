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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataDoorStatus is the corresponding interface of BACnetConstructedDataDoorStatus
type BACnetConstructedDataDoorStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoorStatus returns DoorStatus (property field)
	GetDoorStatus() BACnetDoorStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorStatusTagged
}

// BACnetConstructedDataDoorStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDoorStatus.
// This is useful for switch cases.
type BACnetConstructedDataDoorStatusExactly interface {
	BACnetConstructedDataDoorStatus
	isBACnetConstructedDataDoorStatus() bool
}

// _BACnetConstructedDataDoorStatus is the data-structure of this message
type _BACnetConstructedDataDoorStatus struct {
	*_BACnetConstructedData
	DoorStatus BACnetDoorStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorStatus) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDoorStatus) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetDoorStatus() BACnetDoorStatusTagged {
	return m.DoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetActualValue() BACnetDoorStatusTagged {
	return CastBACnetDoorStatusTagged(m.GetDoorStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorStatus factory function for _BACnetConstructedDataDoorStatus
func NewBACnetConstructedDataDoorStatus(doorStatus BACnetDoorStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorStatus {
	_result := &_BACnetConstructedDataDoorStatus{
		DoorStatus:             doorStatus,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorStatus(structType interface{}) BACnetConstructedDataDoorStatus {
	if casted, ok := structType.(BACnetConstructedDataDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorStatus) GetTypeName() string {
	return "BACnetConstructedDataDoorStatus"
}

func (m *_BACnetConstructedDataDoorStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDoorStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorStatus)
	lengthInBits += m.DoorStatus.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDoorStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorStatus)
	if pullErr := readBuffer.PullContext("doorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorStatus")
	}
	_doorStatus, _doorStatusErr := BACnetDoorStatusTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _doorStatusErr != nil {
		return nil, errors.Wrap(_doorStatusErr, "Error parsing 'doorStatus' field of BACnetConstructedDataDoorStatus")
	}
	doorStatus := _doorStatus.(BACnetDoorStatusTagged)
	if closeErr := readBuffer.CloseContext("doorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorStatus")
	}

	// Virtual field
	_actualValue := doorStatus
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDoorStatus{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DoorStatus: doorStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDoorStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorStatus")
		}

		// Simple Field (doorStatus)
		if pushErr := writeBuffer.PushContext("doorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorStatus")
		}
		_doorStatusErr := writeBuffer.WriteSerializable(m.GetDoorStatus())
		if popErr := writeBuffer.PopContext("doorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorStatus")
		}
		if _doorStatusErr != nil {
			return errors.Wrap(_doorStatusErr, "Error serializing 'doorStatus' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorStatus) isBACnetConstructedDataDoorStatus() bool {
	return true
}

func (m *_BACnetConstructedDataDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
