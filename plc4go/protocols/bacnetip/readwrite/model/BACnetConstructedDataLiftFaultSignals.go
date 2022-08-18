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

// BACnetConstructedDataLiftFaultSignals is the corresponding interface of BACnetConstructedDataLiftFaultSignals
type BACnetConstructedDataLiftFaultSignals interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultSignals returns FaultSignals (property field)
	GetFaultSignals() []BACnetLiftFaultTagged
}

// BACnetConstructedDataLiftFaultSignalsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLiftFaultSignals.
// This is useful for switch cases.
type BACnetConstructedDataLiftFaultSignalsExactly interface {
	BACnetConstructedDataLiftFaultSignals
	isBACnetConstructedDataLiftFaultSignals() bool
}

// _BACnetConstructedDataLiftFaultSignals is the data-structure of this message
type _BACnetConstructedDataLiftFaultSignals struct {
	*_BACnetConstructedData
	FaultSignals []BACnetLiftFaultTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLiftFaultSignals) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFT
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_SIGNALS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLiftFaultSignals) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLiftFaultSignals) GetFaultSignals() []BACnetLiftFaultTagged {
	return m.FaultSignals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLiftFaultSignals factory function for _BACnetConstructedDataLiftFaultSignals
func NewBACnetConstructedDataLiftFaultSignals(faultSignals []BACnetLiftFaultTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLiftFaultSignals {
	_result := &_BACnetConstructedDataLiftFaultSignals{
		FaultSignals:           faultSignals,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLiftFaultSignals(structType interface{}) BACnetConstructedDataLiftFaultSignals {
	if casted, ok := structType.(BACnetConstructedDataLiftFaultSignals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLiftFaultSignals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetTypeName() string {
	return "BACnetConstructedDataLiftFaultSignals"
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.FaultSignals) > 0 {
		for _, element := range m.FaultSignals {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLiftFaultSignalsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLiftFaultSignals, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLiftFaultSignals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLiftFaultSignals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (faultSignals)
	if pullErr := readBuffer.PullContext("faultSignals", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultSignals")
	}
	// Terminated array
	var faultSignals []BACnetLiftFaultTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLiftFaultTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'faultSignals' field of BACnetConstructedDataLiftFaultSignals")
			}
			faultSignals = append(faultSignals, _item.(BACnetLiftFaultTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("faultSignals", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultSignals")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLiftFaultSignals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLiftFaultSignals")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLiftFaultSignals{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultSignals: faultSignals,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLiftFaultSignals) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLiftFaultSignals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLiftFaultSignals")
		}

		// Array Field (faultSignals)
		if pushErr := writeBuffer.PushContext("faultSignals", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultSignals")
		}
		for _, _element := range m.GetFaultSignals() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'faultSignals' field")
			}
		}
		if popErr := writeBuffer.PopContext("faultSignals", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultSignals")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLiftFaultSignals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLiftFaultSignals")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLiftFaultSignals) isBACnetConstructedDataLiftFaultSignals() bool {
	return true
}

func (m *_BACnetConstructedDataLiftFaultSignals) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
