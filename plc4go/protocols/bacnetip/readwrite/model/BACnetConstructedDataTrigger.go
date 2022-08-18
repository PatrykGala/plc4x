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

// BACnetConstructedDataTrigger is the corresponding interface of BACnetConstructedDataTrigger
type BACnetConstructedDataTrigger interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTrigger returns Trigger (property field)
	GetTrigger() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataTriggerExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTrigger.
// This is useful for switch cases.
type BACnetConstructedDataTriggerExactly interface {
	BACnetConstructedDataTrigger
	isBACnetConstructedDataTrigger() bool
}

// _BACnetConstructedDataTrigger is the data-structure of this message
type _BACnetConstructedDataTrigger struct {
	*_BACnetConstructedData
	Trigger BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrigger) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTrigger) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRIGGER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrigger) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTrigger) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrigger) GetTrigger() BACnetApplicationTagBoolean {
	return m.Trigger
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTrigger) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetTrigger())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTrigger factory function for _BACnetConstructedDataTrigger
func NewBACnetConstructedDataTrigger(trigger BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrigger {
	_result := &_BACnetConstructedDataTrigger{
		Trigger:                trigger,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrigger(structType interface{}) BACnetConstructedDataTrigger {
	if casted, ok := structType.(BACnetConstructedDataTrigger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrigger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrigger) GetTypeName() string {
	return "BACnetConstructedDataTrigger"
}

func (m *_BACnetConstructedDataTrigger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTrigger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (trigger)
	lengthInBits += m.Trigger.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTrigger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTriggerParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrigger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrigger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrigger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (trigger)
	if pullErr := readBuffer.PullContext("trigger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for trigger")
	}
	_trigger, _triggerErr := BACnetApplicationTagParse(readBuffer)
	if _triggerErr != nil {
		return nil, errors.Wrap(_triggerErr, "Error parsing 'trigger' field of BACnetConstructedDataTrigger")
	}
	trigger := _trigger.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("trigger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for trigger")
	}

	// Virtual field
	_actualValue := trigger
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrigger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrigger")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTrigger{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Trigger: trigger,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTrigger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrigger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrigger")
		}

		// Simple Field (trigger)
		if pushErr := writeBuffer.PushContext("trigger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for trigger")
		}
		_triggerErr := writeBuffer.WriteSerializable(m.GetTrigger())
		if popErr := writeBuffer.PopContext("trigger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for trigger")
		}
		if _triggerErr != nil {
			return errors.Wrap(_triggerErr, "Error serializing 'trigger' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrigger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrigger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrigger) isBACnetConstructedDataTrigger() bool {
	return true
}

func (m *_BACnetConstructedDataTrigger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
