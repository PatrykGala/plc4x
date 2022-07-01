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

// BACnetEventParameterChangeOfValue is the corresponding interface of BACnetEventParameterChangeOfValue
type BACnetEventParameterChangeOfValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetCovCriteria returns CovCriteria (property field)
	GetCovCriteria() BACnetEventParameterChangeOfValueCivCriteria
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfValueExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfValue.
// This is useful for switch cases.
type BACnetEventParameterChangeOfValueExactly interface {
	BACnetEventParameterChangeOfValue
	isBACnetEventParameterChangeOfValue() bool
}

// _BACnetEventParameterChangeOfValue is the data-structure of this message
type _BACnetEventParameterChangeOfValue struct {
	*_BACnetEventParameter
	OpeningTag  BACnetOpeningTag
	TimeDelay   BACnetContextTagUnsignedInteger
	CovCriteria BACnetEventParameterChangeOfValueCivCriteria
	ClosingTag  BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfValue) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterChangeOfValue) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfValue) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfValue) GetCovCriteria() BACnetEventParameterChangeOfValueCivCriteria {
	return m.CovCriteria
}

func (m *_BACnetEventParameterChangeOfValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfValue factory function for _BACnetEventParameterChangeOfValue
func NewBACnetEventParameterChangeOfValue(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, covCriteria BACnetEventParameterChangeOfValueCivCriteria, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterChangeOfValue {
	_result := &_BACnetEventParameterChangeOfValue{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		CovCriteria:           covCriteria,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfValue(structType interface{}) BACnetEventParameterChangeOfValue {
	if casted, ok := structType.(BACnetEventParameterChangeOfValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfValue) GetTypeName() string {
	return "BACnetEventParameterChangeOfValue"
}

func (m *_BACnetEventParameterChangeOfValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (covCriteria)
	lengthInBits += m.CovCriteria.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfValueParse(readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(2)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfValue")
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
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field of BACnetEventParameterChangeOfValue")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (covCriteria)
	if pullErr := readBuffer.PullContext("covCriteria"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for covCriteria")
	}
	_covCriteria, _covCriteriaErr := BACnetEventParameterChangeOfValueCivCriteriaParse(readBuffer, uint8(uint8(1)))
	if _covCriteriaErr != nil {
		return nil, errors.Wrap(_covCriteriaErr, "Error parsing 'covCriteria' field of BACnetEventParameterChangeOfValue")
	}
	covCriteria := _covCriteria.(BACnetEventParameterChangeOfValueCivCriteria)
	if closeErr := readBuffer.CloseContext("covCriteria"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for covCriteria")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(2)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfValue")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfValue{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		CovCriteria:           covCriteria,
		ClosingTag:            closingTag,
		_BACnetEventParameter: &_BACnetEventParameter{},
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfValue")
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

		// Simple Field (covCriteria)
		if pushErr := writeBuffer.PushContext("covCriteria"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for covCriteria")
		}
		_covCriteriaErr := writeBuffer.WriteSerializable(m.GetCovCriteria())
		if popErr := writeBuffer.PopContext("covCriteria"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for covCriteria")
		}
		if _covCriteriaErr != nil {
			return errors.Wrap(_covCriteriaErr, "Error serializing 'covCriteria' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfValue) isBACnetEventParameterChangeOfValue() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
