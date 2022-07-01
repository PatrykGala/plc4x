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

// BACnetPrescale is the corresponding interface of BACnetPrescale
type BACnetPrescale interface {
	utils.LengthAware
	utils.Serializable
	// GetMultiplier returns Multiplier (property field)
	GetMultiplier() BACnetContextTagUnsignedInteger
	// GetModuloDivide returns ModuloDivide (property field)
	GetModuloDivide() BACnetContextTagUnsignedInteger
}

// BACnetPrescaleExactly can be used when we want exactly this type and not a type which fulfills BACnetPrescale.
// This is useful for switch cases.
type BACnetPrescaleExactly interface {
	BACnetPrescale
	isBACnetPrescale() bool
}

// _BACnetPrescale is the data-structure of this message
type _BACnetPrescale struct {
	Multiplier   BACnetContextTagUnsignedInteger
	ModuloDivide BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPrescale) GetMultiplier() BACnetContextTagUnsignedInteger {
	return m.Multiplier
}

func (m *_BACnetPrescale) GetModuloDivide() BACnetContextTagUnsignedInteger {
	return m.ModuloDivide
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPrescale factory function for _BACnetPrescale
func NewBACnetPrescale(multiplier BACnetContextTagUnsignedInteger, moduloDivide BACnetContextTagUnsignedInteger) *_BACnetPrescale {
	return &_BACnetPrescale{Multiplier: multiplier, ModuloDivide: moduloDivide}
}

// Deprecated: use the interface for direct cast
func CastBACnetPrescale(structType interface{}) BACnetPrescale {
	if casted, ok := structType.(BACnetPrescale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPrescale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPrescale) GetTypeName() string {
	return "BACnetPrescale"
}

func (m *_BACnetPrescale) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPrescale) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (multiplier)
	lengthInBits += m.Multiplier.GetLengthInBits()

	// Simple field (moduloDivide)
	lengthInBits += m.ModuloDivide.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPrescale) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPrescaleParse(readBuffer utils.ReadBuffer) (BACnetPrescale, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPrescale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPrescale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (multiplier)
	if pullErr := readBuffer.PullContext("multiplier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for multiplier")
	}
	_multiplier, _multiplierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _multiplierErr != nil {
		return nil, errors.Wrap(_multiplierErr, "Error parsing 'multiplier' field of BACnetPrescale")
	}
	multiplier := _multiplier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("multiplier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for multiplier")
	}

	// Simple Field (moduloDivide)
	if pullErr := readBuffer.PullContext("moduloDivide"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for moduloDivide")
	}
	_moduloDivide, _moduloDivideErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _moduloDivideErr != nil {
		return nil, errors.Wrap(_moduloDivideErr, "Error parsing 'moduloDivide' field of BACnetPrescale")
	}
	moduloDivide := _moduloDivide.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("moduloDivide"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for moduloDivide")
	}

	if closeErr := readBuffer.CloseContext("BACnetPrescale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPrescale")
	}

	// Create the instance
	return NewBACnetPrescale(multiplier, moduloDivide), nil
}

func (m *_BACnetPrescale) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPrescale"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPrescale")
	}

	// Simple Field (multiplier)
	if pushErr := writeBuffer.PushContext("multiplier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for multiplier")
	}
	_multiplierErr := writeBuffer.WriteSerializable(m.GetMultiplier())
	if popErr := writeBuffer.PopContext("multiplier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for multiplier")
	}
	if _multiplierErr != nil {
		return errors.Wrap(_multiplierErr, "Error serializing 'multiplier' field")
	}

	// Simple Field (moduloDivide)
	if pushErr := writeBuffer.PushContext("moduloDivide"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for moduloDivide")
	}
	_moduloDivideErr := writeBuffer.WriteSerializable(m.GetModuloDivide())
	if popErr := writeBuffer.PopContext("moduloDivide"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for moduloDivide")
	}
	if _moduloDivideErr != nil {
		return errors.Wrap(_moduloDivideErr, "Error serializing 'moduloDivide' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPrescale"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPrescale")
	}
	return nil
}

func (m *_BACnetPrescale) isBACnetPrescale() bool {
	return true
}

func (m *_BACnetPrescale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
