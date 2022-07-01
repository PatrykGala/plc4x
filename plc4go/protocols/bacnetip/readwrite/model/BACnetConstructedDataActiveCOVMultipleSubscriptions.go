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

// BACnetConstructedDataActiveCOVMultipleSubscriptions is the corresponding interface of BACnetConstructedDataActiveCOVMultipleSubscriptions
type BACnetConstructedDataActiveCOVMultipleSubscriptions interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetActiveCOVMultipleSubscriptions returns ActiveCOVMultipleSubscriptions (property field)
	GetActiveCOVMultipleSubscriptions() []BACnetCOVMultipleSubscription
}

// BACnetConstructedDataActiveCOVMultipleSubscriptionsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataActiveCOVMultipleSubscriptions.
// This is useful for switch cases.
type BACnetConstructedDataActiveCOVMultipleSubscriptionsExactly interface {
	BACnetConstructedDataActiveCOVMultipleSubscriptions
	isBACnetConstructedDataActiveCOVMultipleSubscriptions() bool
}

// _BACnetConstructedDataActiveCOVMultipleSubscriptions is the data-structure of this message
type _BACnetConstructedDataActiveCOVMultipleSubscriptions struct {
	*_BACnetConstructedData
	ActiveCOVMultipleSubscriptions []BACnetCOVMultipleSubscription
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_COV_MULTIPLE_SUBSCRIPTIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetActiveCOVMultipleSubscriptions() []BACnetCOVMultipleSubscription {
	return m.ActiveCOVMultipleSubscriptions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActiveCOVMultipleSubscriptions factory function for _BACnetConstructedDataActiveCOVMultipleSubscriptions
func NewBACnetConstructedDataActiveCOVMultipleSubscriptions(activeCOVMultipleSubscriptions []BACnetCOVMultipleSubscription, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActiveCOVMultipleSubscriptions {
	_result := &_BACnetConstructedDataActiveCOVMultipleSubscriptions{
		ActiveCOVMultipleSubscriptions: activeCOVMultipleSubscriptions,
		_BACnetConstructedData:         NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActiveCOVMultipleSubscriptions(structType interface{}) BACnetConstructedDataActiveCOVMultipleSubscriptions {
	if casted, ok := structType.(BACnetConstructedDataActiveCOVMultipleSubscriptions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveCOVMultipleSubscriptions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetTypeName() string {
	return "BACnetConstructedDataActiveCOVMultipleSubscriptions"
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.ActiveCOVMultipleSubscriptions) > 0 {
		for _, element := range m.ActiveCOVMultipleSubscriptions {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataActiveCOVMultipleSubscriptionsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActiveCOVMultipleSubscriptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActiveCOVMultipleSubscriptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (activeCOVMultipleSubscriptions)
	if pullErr := readBuffer.PullContext("activeCOVMultipleSubscriptions", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for activeCOVMultipleSubscriptions")
	}
	// Terminated array
	var activeCOVMultipleSubscriptions []BACnetCOVMultipleSubscription
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetCOVMultipleSubscriptionParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'activeCOVMultipleSubscriptions' field of BACnetConstructedDataActiveCOVMultipleSubscriptions")
			}
			activeCOVMultipleSubscriptions = append(activeCOVMultipleSubscriptions, _item.(BACnetCOVMultipleSubscription))

		}
	}
	if closeErr := readBuffer.CloseContext("activeCOVMultipleSubscriptions", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for activeCOVMultipleSubscriptions")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActiveCOVMultipleSubscriptions")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataActiveCOVMultipleSubscriptions{
		ActiveCOVMultipleSubscriptions: activeCOVMultipleSubscriptions,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActiveCOVMultipleSubscriptions")
		}

		// Array Field (activeCOVMultipleSubscriptions)
		if pushErr := writeBuffer.PushContext("activeCOVMultipleSubscriptions", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for activeCOVMultipleSubscriptions")
		}
		for _, _element := range m.GetActiveCOVMultipleSubscriptions() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'activeCOVMultipleSubscriptions' field")
			}
		}
		if popErr := writeBuffer.PopContext("activeCOVMultipleSubscriptions", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for activeCOVMultipleSubscriptions")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActiveCOVMultipleSubscriptions")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) isBACnetConstructedDataActiveCOVMultipleSubscriptions() bool {
	return true
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
