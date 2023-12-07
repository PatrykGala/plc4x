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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// PascalByteString is the corresponding interface of PascalByteString
type PascalByteString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetStringLength returns StringLength (property field)
	GetStringLength() int32
	// GetStringValue returns StringValue (property field)
	GetStringValue() []byte
}

// PascalByteStringExactly can be used when we want exactly this type and not a type which fulfills PascalByteString.
// This is useful for switch cases.
type PascalByteStringExactly interface {
	PascalByteString
	isPascalByteString() bool
}

// _PascalByteString is the data-structure of this message
type _PascalByteString struct {
	StringLength int32
	StringValue  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PascalByteString) GetStringLength() int32 {
	return m.StringLength
}

func (m *_PascalByteString) GetStringValue() []byte {
	return m.StringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPascalByteString factory function for _PascalByteString
func NewPascalByteString(stringLength int32, stringValue []byte) *_PascalByteString {
	return &_PascalByteString{StringLength: stringLength, StringValue: stringValue}
}

// Deprecated: use the interface for direct cast
func CastPascalByteString(structType any) PascalByteString {
	if casted, ok := structType.(PascalByteString); ok {
		return casted
	}
	if casted, ok := structType.(*PascalByteString); ok {
		return *casted
	}
	return nil
}

func (m *_PascalByteString) GetTypeName() string {
	return "PascalByteString"
}

func (m *_PascalByteString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (stringLength)
	lengthInBits += 32

	// Array field
	if len(m.StringValue) > 0 {
		lengthInBits += 8 * uint16(len(m.StringValue))
	}

	return lengthInBits
}

func (m *_PascalByteString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PascalByteStringParse(ctx context.Context, theBytes []byte) (PascalByteString, error) {
	return PascalByteStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PascalByteStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PascalByteString, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PascalByteString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PascalByteString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (stringLength)
	_stringLength, _stringLengthErr := readBuffer.ReadInt32("stringLength", 32)
	if _stringLengthErr != nil {
		return nil, errors.Wrap(_stringLengthErr, "Error parsing 'stringLength' field of PascalByteString")
	}
	stringLength := _stringLength
	// Byte Array field (stringValue)
	numberOfBytesstringValue := int(utils.InlineIf(bool((stringLength) == (-(1))), func() any { return uint16(uint16(0)) }, func() any { return uint16(stringLength) }).(uint16))
	stringValue, _readArrayErr := readBuffer.ReadByteArray("stringValue", numberOfBytesstringValue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'stringValue' field of PascalByteString")
	}

	if closeErr := readBuffer.CloseContext("PascalByteString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PascalByteString")
	}

	// Create the instance
	return &_PascalByteString{
		StringLength: stringLength,
		StringValue:  stringValue,
	}, nil
}

func (m *_PascalByteString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PascalByteString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PascalByteString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PascalByteString")
	}

	// Simple Field (stringLength)
	stringLength := int32(m.GetStringLength())
	_stringLengthErr := writeBuffer.WriteInt32("stringLength", 32, int32((stringLength)))
	if _stringLengthErr != nil {
		return errors.Wrap(_stringLengthErr, "Error serializing 'stringLength' field")
	}

	// Array Field (stringValue)
	// Byte Array field (stringValue)
	if err := writeBuffer.WriteByteArray("stringValue", m.GetStringValue()); err != nil {
		return errors.Wrap(err, "Error serializing 'stringValue' field")
	}

	if popErr := writeBuffer.PopContext("PascalByteString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PascalByteString")
	}
	return nil
}

func (m *_PascalByteString) isPascalByteString() bool {
	return true
}

func (m *_PascalByteString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
