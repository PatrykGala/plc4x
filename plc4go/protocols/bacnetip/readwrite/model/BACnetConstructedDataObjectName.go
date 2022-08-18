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

// BACnetConstructedDataObjectName is the corresponding interface of BACnetConstructedDataObjectName
type BACnetConstructedDataObjectName interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetObjectName returns ObjectName (property field)
	GetObjectName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataObjectNameExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataObjectName.
// This is useful for switch cases.
type BACnetConstructedDataObjectNameExactly interface {
	BACnetConstructedDataObjectName
	isBACnetConstructedDataObjectName() bool
}

// _BACnetConstructedDataObjectName is the data-structure of this message
type _BACnetConstructedDataObjectName struct {
	*_BACnetConstructedData
	ObjectName BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataObjectName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataObjectName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OBJECT_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataObjectName) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataObjectName) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataObjectName) GetObjectName() BACnetApplicationTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataObjectName) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetObjectName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataObjectName factory function for _BACnetConstructedDataObjectName
func NewBACnetConstructedDataObjectName(objectName BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataObjectName {
	_result := &_BACnetConstructedDataObjectName{
		ObjectName:             objectName,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataObjectName(structType interface{}) BACnetConstructedDataObjectName {
	if casted, ok := structType.(BACnetConstructedDataObjectName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataObjectName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataObjectName) GetTypeName() string {
	return "BACnetConstructedDataObjectName"
}

func (m *_BACnetConstructedDataObjectName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataObjectName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataObjectName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataObjectNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataObjectName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataObjectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataObjectName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectName)
	if pullErr := readBuffer.PullContext("objectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectName")
	}
	_objectName, _objectNameErr := BACnetApplicationTagParse(readBuffer)
	if _objectNameErr != nil {
		return nil, errors.Wrap(_objectNameErr, "Error parsing 'objectName' field of BACnetConstructedDataObjectName")
	}
	objectName := _objectName.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("objectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectName")
	}

	// Virtual field
	_actualValue := objectName
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataObjectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataObjectName")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataObjectName{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ObjectName: objectName,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataObjectName) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataObjectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataObjectName")
		}

		// Simple Field (objectName)
		if pushErr := writeBuffer.PushContext("objectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectName")
		}
		_objectNameErr := writeBuffer.WriteSerializable(m.GetObjectName())
		if popErr := writeBuffer.PopContext("objectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectName")
		}
		if _objectNameErr != nil {
			return errors.Wrap(_objectNameErr, "Error serializing 'objectName' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataObjectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataObjectName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataObjectName) isBACnetConstructedDataObjectName() bool {
	return true
}

func (m *_BACnetConstructedDataObjectName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
