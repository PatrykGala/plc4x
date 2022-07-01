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

// BACnetEventParameterChangeOfCharacterString is the corresponding interface of BACnetEventParameterChangeOfCharacterString
type BACnetEventParameterChangeOfCharacterString interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfCharacterStringExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfCharacterString.
// This is useful for switch cases.
type BACnetEventParameterChangeOfCharacterStringExactly interface {
	BACnetEventParameterChangeOfCharacterString
	isBACnetEventParameterChangeOfCharacterString() bool
}

// _BACnetEventParameterChangeOfCharacterString is the data-structure of this message
type _BACnetEventParameterChangeOfCharacterString struct {
	*_BACnetEventParameter
	OpeningTag        BACnetOpeningTag
	TimeDelay         BACnetContextTagUnsignedInteger
	ListOfAlarmValues BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
	ClosingTag        BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfCharacterString) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfCharacterString) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetListOfAlarmValues() BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfCharacterString factory function for _BACnetEventParameterChangeOfCharacterString
func NewBACnetEventParameterChangeOfCharacterString(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, listOfAlarmValues BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterChangeOfCharacterString {
	_result := &_BACnetEventParameterChangeOfCharacterString{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		ListOfAlarmValues:     listOfAlarmValues,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfCharacterString(structType interface{}) BACnetEventParameterChangeOfCharacterString {
	if casted, ok := structType.(BACnetEventParameterChangeOfCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetTypeName() string {
	return "BACnetEventParameterChangeOfCharacterString"
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (listOfAlarmValues)
	lengthInBits += m.ListOfAlarmValues.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfCharacterStringParse(readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(17)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfCharacterString")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field of BACnetEventParameterChangeOfCharacterString")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (listOfAlarmValues)
	if pullErr := readBuffer.PullContext("listOfAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfAlarmValues")
	}
	_listOfAlarmValues, _listOfAlarmValuesErr := BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParse(readBuffer, uint8(uint8(1)))
	if _listOfAlarmValuesErr != nil {
		return nil, errors.Wrap(_listOfAlarmValuesErr, "Error parsing 'listOfAlarmValues' field of BACnetEventParameterChangeOfCharacterString")
	}
	listOfAlarmValues := _listOfAlarmValues.(BACnetEventParameterChangeOfCharacterStringListOfAlarmValues)
	if closeErr := readBuffer.CloseContext("listOfAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfAlarmValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(17)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfCharacterString")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfCharacterString")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfCharacterString{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		ListOfAlarmValues:     listOfAlarmValues,
		ClosingTag:            closingTag,
		_BACnetEventParameter: &_BACnetEventParameter{},
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfCharacterString")
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

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := writeBuffer.WriteSerializable(m.GetTimeDelay())
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (listOfAlarmValues)
		if pushErr := writeBuffer.PushContext("listOfAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfAlarmValues")
		}
		_listOfAlarmValuesErr := writeBuffer.WriteSerializable(m.GetListOfAlarmValues())
		if popErr := writeBuffer.PopContext("listOfAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfAlarmValues")
		}
		if _listOfAlarmValuesErr != nil {
			return errors.Wrap(_listOfAlarmValuesErr, "Error serializing 'listOfAlarmValues' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfCharacterString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfCharacterString) isBACnetEventParameterChangeOfCharacterString() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
