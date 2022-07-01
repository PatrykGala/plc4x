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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataCommandTimeArray is the corresponding interface of BACnetConstructedDataCommandTimeArray
type BACnetConstructedDataCommandTimeArray interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetCommandTimeArray returns CommandTimeArray (property field)
	GetCommandTimeArray() []BACnetTimeStamp
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataCommandTimeArrayExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCommandTimeArray.
// This is useful for switch cases.
type BACnetConstructedDataCommandTimeArrayExactly interface {
	BACnetConstructedDataCommandTimeArray
	isBACnetConstructedDataCommandTimeArray() bool
}

// _BACnetConstructedDataCommandTimeArray is the data-structure of this message
type _BACnetConstructedDataCommandTimeArray struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	CommandTimeArray     []BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCommandTimeArray) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COMMAND_TIME_ARRAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCommandTimeArray) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataCommandTimeArray) GetCommandTimeArray() []BACnetTimeStamp {
	return m.CommandTimeArray
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCommandTimeArray factory function for _BACnetConstructedDataCommandTimeArray
func NewBACnetConstructedDataCommandTimeArray(numberOfDataElements BACnetApplicationTagUnsignedInteger, commandTimeArray []BACnetTimeStamp, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCommandTimeArray {
	_result := &_BACnetConstructedDataCommandTimeArray{
		NumberOfDataElements:   numberOfDataElements,
		CommandTimeArray:       commandTimeArray,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCommandTimeArray(structType interface{}) BACnetConstructedDataCommandTimeArray {
	if casted, ok := structType.(BACnetConstructedDataCommandTimeArray); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCommandTimeArray); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCommandTimeArray) GetTypeName() string {
	return "BACnetConstructedDataCommandTimeArray"
}

func (m *_BACnetConstructedDataCommandTimeArray) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCommandTimeArray) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.CommandTimeArray) > 0 {
		for _, element := range m.CommandTimeArray {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCommandTimeArray) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCommandTimeArrayParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCommandTimeArray, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCommandTimeArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCommandTimeArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataCommandTimeArray")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (commandTimeArray)
	if pullErr := readBuffer.PullContext("commandTimeArray", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTimeArray")
	}
	// Terminated array
	var commandTimeArray []BACnetTimeStamp
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetTimeStampParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'commandTimeArray' field of BACnetConstructedDataCommandTimeArray")
			}
			commandTimeArray = append(commandTimeArray, _item.(BACnetTimeStamp))

		}
	}
	if closeErr := readBuffer.CloseContext("commandTimeArray", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTimeArray")
	}

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(commandTimeArray)) == (16)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"commandTimeArray should have exactly 16 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCommandTimeArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCommandTimeArray")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCommandTimeArray{
		NumberOfDataElements: numberOfDataElements,
		CommandTimeArray:     commandTimeArray,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCommandTimeArray) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCommandTimeArray"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCommandTimeArray")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (commandTimeArray)
		if pushErr := writeBuffer.PushContext("commandTimeArray", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for commandTimeArray")
		}
		for _, _element := range m.GetCommandTimeArray() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'commandTimeArray' field")
			}
		}
		if popErr := writeBuffer.PopContext("commandTimeArray", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for commandTimeArray")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCommandTimeArray"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCommandTimeArray")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCommandTimeArray) isBACnetConstructedDataCommandTimeArray() bool {
	return true
}

func (m *_BACnetConstructedDataCommandTimeArray) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
