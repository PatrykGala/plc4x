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

// BACnetConstructedDataOctetstringValueAll is the corresponding interface of BACnetConstructedDataOctetstringValueAll
type BACnetConstructedDataOctetstringValueAll interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataOctetstringValueAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOctetstringValueAll.
// This is useful for switch cases.
type BACnetConstructedDataOctetstringValueAllExactly interface {
	BACnetConstructedDataOctetstringValueAll
	isBACnetConstructedDataOctetstringValueAll() bool
}

// _BACnetConstructedDataOctetstringValueAll is the data-structure of this message
type _BACnetConstructedDataOctetstringValueAll struct {
	*_BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOctetstringValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_OCTETSTRING_VALUE
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOctetstringValueAll) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

// NewBACnetConstructedDataOctetstringValueAll factory function for _BACnetConstructedDataOctetstringValueAll
func NewBACnetConstructedDataOctetstringValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOctetstringValueAll {
	_result := &_BACnetConstructedDataOctetstringValueAll{
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOctetstringValueAll(structType interface{}) BACnetConstructedDataOctetstringValueAll {
	if casted, ok := structType.(BACnetConstructedDataOctetstringValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOctetstringValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetTypeName() string {
	return "BACnetConstructedDataOctetstringValueAll"
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOctetstringValueAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOctetstringValueAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOctetstringValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOctetstringValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{"All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOctetstringValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOctetstringValueAll")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOctetstringValueAll{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOctetstringValueAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOctetstringValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOctetstringValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOctetstringValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOctetstringValueAll")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOctetstringValueAll) isBACnetConstructedDataOctetstringValueAll() bool {
	return true
}

func (m *_BACnetConstructedDataOctetstringValueAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
