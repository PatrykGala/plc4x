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

// BACnetConstructedDataLightingCommand is the corresponding interface of BACnetConstructedDataLightingCommand
type BACnetConstructedDataLightingCommand interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLightingCommand returns LightingCommand (property field)
	GetLightingCommand() BACnetLightingCommand
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLightingCommand
}

// BACnetConstructedDataLightingCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLightingCommand.
// This is useful for switch cases.
type BACnetConstructedDataLightingCommandExactly interface {
	BACnetConstructedDataLightingCommand
	isBACnetConstructedDataLightingCommand() bool
}

// _BACnetConstructedDataLightingCommand is the data-structure of this message
type _BACnetConstructedDataLightingCommand struct {
	*_BACnetConstructedData
	LightingCommand BACnetLightingCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLightingCommand) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingCommand) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLightingCommand) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetLightingCommand() BACnetLightingCommand {
	return m.LightingCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetActualValue() BACnetLightingCommand {
	return CastBACnetLightingCommand(m.GetLightingCommand())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingCommand factory function for _BACnetConstructedDataLightingCommand
func NewBACnetConstructedDataLightingCommand(lightingCommand BACnetLightingCommand, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingCommand {
	_result := &_BACnetConstructedDataLightingCommand{
		LightingCommand:        lightingCommand,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingCommand(structType interface{}) BACnetConstructedDataLightingCommand {
	if casted, ok := structType.(BACnetConstructedDataLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingCommand) GetTypeName() string {
	return "BACnetConstructedDataLightingCommand"
}

func (m *_BACnetConstructedDataLightingCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLightingCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lightingCommand)
	lengthInBits += m.LightingCommand.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLightingCommandParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightingCommand)
	if pullErr := readBuffer.PullContext("lightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightingCommand")
	}
	_lightingCommand, _lightingCommandErr := BACnetLightingCommandParse(readBuffer)
	if _lightingCommandErr != nil {
		return nil, errors.Wrap(_lightingCommandErr, "Error parsing 'lightingCommand' field of BACnetConstructedDataLightingCommand")
	}
	lightingCommand := _lightingCommand.(BACnetLightingCommand)
	if closeErr := readBuffer.CloseContext("lightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightingCommand")
	}

	// Virtual field
	_actualValue := lightingCommand
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommand")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLightingCommand{
		LightingCommand: lightingCommand,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLightingCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommand")
		}

		// Simple Field (lightingCommand)
		if pushErr := writeBuffer.PushContext("lightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lightingCommand")
		}
		_lightingCommandErr := writeBuffer.WriteSerializable(m.GetLightingCommand())
		if popErr := writeBuffer.PopContext("lightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lightingCommand")
		}
		if _lightingCommandErr != nil {
			return errors.Wrap(_lightingCommandErr, "Error serializing 'lightingCommand' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommand")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingCommand) isBACnetConstructedDataLightingCommand() bool {
	return true
}

func (m *_BACnetConstructedDataLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
