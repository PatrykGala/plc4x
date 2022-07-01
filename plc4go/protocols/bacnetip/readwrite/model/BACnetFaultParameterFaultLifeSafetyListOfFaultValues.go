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

// BACnetFaultParameterFaultLifeSafetyListOfFaultValues is the corresponding interface of BACnetFaultParameterFaultLifeSafetyListOfFaultValues
type BACnetFaultParameterFaultLifeSafetyListOfFaultValues interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListIfFaultValues returns ListIfFaultValues (property field)
	GetListIfFaultValues() []BACnetLifeSafetyStateTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultLifeSafetyListOfFaultValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultLifeSafetyListOfFaultValues.
// This is useful for switch cases.
type BACnetFaultParameterFaultLifeSafetyListOfFaultValuesExactly interface {
	BACnetFaultParameterFaultLifeSafetyListOfFaultValues
	isBACnetFaultParameterFaultLifeSafetyListOfFaultValues() bool
}

// _BACnetFaultParameterFaultLifeSafetyListOfFaultValues is the data-structure of this message
type _BACnetFaultParameterFaultLifeSafetyListOfFaultValues struct {
	OpeningTag        BACnetOpeningTag
	ListIfFaultValues []BACnetLifeSafetyStateTagged
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetListIfFaultValues() []BACnetLifeSafetyStateTagged {
	return m.ListIfFaultValues
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultLifeSafetyListOfFaultValues factory function for _BACnetFaultParameterFaultLifeSafetyListOfFaultValues
func NewBACnetFaultParameterFaultLifeSafetyListOfFaultValues(openingTag BACnetOpeningTag, listIfFaultValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	return &_BACnetFaultParameterFaultLifeSafetyListOfFaultValues{OpeningTag: openingTag, ListIfFaultValues: listIfFaultValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultLifeSafetyListOfFaultValues(structType interface{}) BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	if casted, ok := structType.(BACnetFaultParameterFaultLifeSafetyListOfFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultLifeSafetyListOfFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetTypeName() string {
	return "BACnetFaultParameterFaultLifeSafetyListOfFaultValues"
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListIfFaultValues) > 0 {
		for _, element := range m.ListIfFaultValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultLifeSafetyListOfFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listIfFaultValues)
	if pullErr := readBuffer.PullContext("listIfFaultValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listIfFaultValues")
	}
	// Terminated array
	var listIfFaultValues []BACnetLifeSafetyStateTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listIfFaultValues' field of BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
			}
			listIfFaultValues = append(listIfFaultValues, _item.(BACnetLifeSafetyStateTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("listIfFaultValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listIfFaultValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}

	// Create the instance
	return NewBACnetFaultParameterFaultLifeSafetyListOfFaultValues(openingTag, listIfFaultValues, closingTag, tagNumber), nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
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

	// Array Field (listIfFaultValues)
	if pushErr := writeBuffer.PushContext("listIfFaultValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listIfFaultValues")
	}
	for _, _element := range m.GetListIfFaultValues() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listIfFaultValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listIfFaultValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listIfFaultValues")
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

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}
	return nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) isBACnetFaultParameterFaultLifeSafetyListOfFaultValues() bool {
	return true
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
