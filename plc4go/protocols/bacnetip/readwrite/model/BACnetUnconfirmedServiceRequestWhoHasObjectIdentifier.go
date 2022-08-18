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

// BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier is the corresponding interface of BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
type BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequestWhoHasObject
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
}

// BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierExactly interface {
	BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
	isBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier() bool
}

// _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier struct {
	*_BACnetUnconfirmedServiceRequestWhoHasObject
	ObjectIdentifier BACnetContextTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) InitializeParent(parent BACnetUnconfirmedServiceRequestWhoHasObject, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetParent() BACnetUnconfirmedServiceRequestWhoHasObject {
	return m._BACnetUnconfirmedServiceRequestWhoHasObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier factory function for _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
func NewBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier, peekedTagHeader BACnetTagHeader) *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier {
	_result := &_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier{
		ObjectIdentifier: objectIdentifier,
		_BACnetUnconfirmedServiceRequestWhoHasObject: NewBACnetUnconfirmedServiceRequestWhoHasObject(peekedTagHeader),
	}
	_result._BACnetUnconfirmedServiceRequestWhoHasObject._BACnetUnconfirmedServiceRequestWhoHasObjectChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier(structType interface{}) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierParse(readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier{
		_BACnetUnconfirmedServiceRequestWhoHasObject: &_BACnetUnconfirmedServiceRequestWhoHasObject{},
		ObjectIdentifier: objectIdentifier,
	}
	_child._BACnetUnconfirmedServiceRequestWhoHasObject._BACnetUnconfirmedServiceRequestWhoHasObjectChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) isBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
