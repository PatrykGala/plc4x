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

// BACnetFaultParameterFaultCharacterStringListOfFaultValues is the corresponding interface of BACnetFaultParameterFaultCharacterStringListOfFaultValues
type BACnetFaultParameterFaultCharacterStringListOfFaultValues interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() []BACnetApplicationTagCharacterString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultCharacterStringListOfFaultValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultCharacterStringListOfFaultValues.
// This is useful for switch cases.
type BACnetFaultParameterFaultCharacterStringListOfFaultValuesExactly interface {
	BACnetFaultParameterFaultCharacterStringListOfFaultValues
	isBACnetFaultParameterFaultCharacterStringListOfFaultValues() bool
}

// _BACnetFaultParameterFaultCharacterStringListOfFaultValues is the data-structure of this message
type _BACnetFaultParameterFaultCharacterStringListOfFaultValues struct {
	OpeningTag        BACnetOpeningTag
	ListOfFaultValues []BACnetApplicationTagCharacterString
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetListOfFaultValues() []BACnetApplicationTagCharacterString {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultCharacterStringListOfFaultValues factory function for _BACnetFaultParameterFaultCharacterStringListOfFaultValues
func NewBACnetFaultParameterFaultCharacterStringListOfFaultValues(openingTag BACnetOpeningTag, listOfFaultValues []BACnetApplicationTagCharacterString, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultCharacterStringListOfFaultValues {
	return &_BACnetFaultParameterFaultCharacterStringListOfFaultValues{OpeningTag: openingTag, ListOfFaultValues: listOfFaultValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultCharacterStringListOfFaultValues(structType interface{}) BACnetFaultParameterFaultCharacterStringListOfFaultValues {
	if casted, ok := structType.(BACnetFaultParameterFaultCharacterStringListOfFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultCharacterStringListOfFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetTypeName() string {
	return "BACnetFaultParameterFaultCharacterStringListOfFaultValues"
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfFaultValues) > 0 {
		for _, element := range m.ListOfFaultValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultCharacterStringListOfFaultValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultCharacterStringListOfFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfFaultValues)
	if pullErr := readBuffer.PullContext("listOfFaultValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfFaultValues")
	}
	// Terminated array
	var listOfFaultValues []BACnetApplicationTagCharacterString
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfFaultValues' field of BACnetFaultParameterFaultCharacterStringListOfFaultValues")
			}
			listOfFaultValues = append(listOfFaultValues, _item.(BACnetApplicationTagCharacterString))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfFaultValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfFaultValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}

	// Create the instance
	return &_BACnetFaultParameterFaultCharacterStringListOfFaultValues{
		TagNumber:         tagNumber,
		OpeningTag:        openingTag,
		ListOfFaultValues: listOfFaultValues,
		ClosingTag:        closingTag,
	}, nil
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
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

	// Array Field (listOfFaultValues)
	if pushErr := writeBuffer.PushContext("listOfFaultValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfFaultValues")
	}
	for _, _element := range m.GetListOfFaultValues() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfFaultValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfFaultValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfFaultValues")
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

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) isBACnetFaultParameterFaultCharacterStringListOfFaultValues() bool {
	return true
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
