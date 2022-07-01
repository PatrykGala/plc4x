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

// BACnetRejectReasonTagged is the corresponding interface of BACnetRejectReasonTagged
type BACnetRejectReasonTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() BACnetRejectReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetRejectReasonTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetRejectReasonTagged.
// This is useful for switch cases.
type BACnetRejectReasonTaggedExactly interface {
	BACnetRejectReasonTagged
	isBACnetRejectReasonTagged() bool
}

// _BACnetRejectReasonTagged is the data-structure of this message
type _BACnetRejectReasonTagged struct {
	Value            BACnetRejectReason
	ProprietaryValue uint32

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRejectReasonTagged) GetValue() BACnetRejectReason {
	return m.Value
}

func (m *_BACnetRejectReasonTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetRejectReasonTagged) GetIsProprietary() bool {
	return bool(bool((m.GetValue()) == (BACnetRejectReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRejectReasonTagged factory function for _BACnetRejectReasonTagged
func NewBACnetRejectReasonTagged(value BACnetRejectReason, proprietaryValue uint32, actualLength uint32) *_BACnetRejectReasonTagged {
	return &_BACnetRejectReasonTagged{Value: value, ProprietaryValue: proprietaryValue, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetRejectReasonTagged(structType interface{}) BACnetRejectReasonTagged {
	if casted, ok := structType.(BACnetRejectReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRejectReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRejectReasonTagged) GetTypeName() string {
	return "BACnetRejectReasonTagged"
}

func (m *_BACnetRejectReasonTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetRejectReasonTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32(int32(0)) }, func() interface{} { return int32(int32(int32(m.ActualLength) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32(int32(int32(m.ActualLength) * int32(int32(8)))) }, func() interface{} { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetRejectReasonTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRejectReasonTaggedParse(readBuffer utils.ReadBuffer, actualLength uint32) (BACnetRejectReasonTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRejectReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRejectReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Manual Field (value)
	_value, _valueErr := ReadEnumGeneric(readBuffer, actualLength, BACnetRejectReason_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetRejectReasonTagged")
	}
	value := _value.(BACnetRejectReason)

	// Virtual field
	_isProprietary := bool((value) == (BACnetRejectReason_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(readBuffer, actualLength, isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetRejectReasonTagged")
	}
	proprietaryValue := _proprietaryValue.(uint32)

	if closeErr := readBuffer.CloseContext("BACnetRejectReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRejectReasonTagged")
	}

	// Create the instance
	return NewBACnetRejectReasonTagged(value, proprietaryValue, actualLength), nil
}

func (m *_BACnetRejectReasonTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetRejectReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRejectReasonTagged")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRejectReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRejectReasonTagged")
	}
	return nil
}

func (m *_BACnetRejectReasonTagged) isBACnetRejectReasonTagged() bool {
	return true
}

func (m *_BACnetRejectReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
