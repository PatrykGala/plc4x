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

// ExtensionObjectEncodingMask is the corresponding interface of ExtensionObjectEncodingMask
type ExtensionObjectEncodingMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTypeIdSpecified returns TypeIdSpecified (property field)
	GetTypeIdSpecified() bool
	// GetXmlbody returns Xmlbody (property field)
	GetXmlbody() bool
	// GetBinaryBody returns BinaryBody (property field)
	GetBinaryBody() bool
}

// ExtensionObjectEncodingMaskExactly can be used when we want exactly this type and not a type which fulfills ExtensionObjectEncodingMask.
// This is useful for switch cases.
type ExtensionObjectEncodingMaskExactly interface {
	ExtensionObjectEncodingMask
	isExtensionObjectEncodingMask() bool
}

// _ExtensionObjectEncodingMask is the data-structure of this message
type _ExtensionObjectEncodingMask struct {
	TypeIdSpecified bool
	Xmlbody         bool
	BinaryBody      bool
	// Reserved Fields
	reservedField0 *int8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensionObjectEncodingMask) GetTypeIdSpecified() bool {
	return m.TypeIdSpecified
}

func (m *_ExtensionObjectEncodingMask) GetXmlbody() bool {
	return m.Xmlbody
}

func (m *_ExtensionObjectEncodingMask) GetBinaryBody() bool {
	return m.BinaryBody
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewExtensionObjectEncodingMask factory function for _ExtensionObjectEncodingMask
func NewExtensionObjectEncodingMask(typeIdSpecified bool, xmlbody bool, binaryBody bool) *_ExtensionObjectEncodingMask {
	return &_ExtensionObjectEncodingMask{TypeIdSpecified: typeIdSpecified, Xmlbody: xmlbody, BinaryBody: binaryBody}
}

// Deprecated: use the interface for direct cast
func CastExtensionObjectEncodingMask(structType any) ExtensionObjectEncodingMask {
	if casted, ok := structType.(ExtensionObjectEncodingMask); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensionObjectEncodingMask); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensionObjectEncodingMask) GetTypeName() string {
	return "ExtensionObjectEncodingMask"
}

func (m *_ExtensionObjectEncodingMask) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 5

	// Simple field (typeIdSpecified)
	lengthInBits += 1

	// Simple field (xmlbody)
	lengthInBits += 1

	// Simple field (binaryBody)
	lengthInBits += 1

	return lengthInBits
}

func (m *_ExtensionObjectEncodingMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ExtensionObjectEncodingMaskParse(ctx context.Context, theBytes []byte) (ExtensionObjectEncodingMask, error) {
	return ExtensionObjectEncodingMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ExtensionObjectEncodingMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ExtensionObjectEncodingMask, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ExtensionObjectEncodingMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensionObjectEncodingMask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *int8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadInt8("reserved", 5)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ExtensionObjectEncodingMask")
		}
		if reserved != int8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": int8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (typeIdSpecified)
	_typeIdSpecified, _typeIdSpecifiedErr := readBuffer.ReadBit("typeIdSpecified")
	if _typeIdSpecifiedErr != nil {
		return nil, errors.Wrap(_typeIdSpecifiedErr, "Error parsing 'typeIdSpecified' field of ExtensionObjectEncodingMask")
	}
	typeIdSpecified := _typeIdSpecified

	// Simple Field (xmlbody)
	_xmlbody, _xmlbodyErr := readBuffer.ReadBit("xmlbody")
	if _xmlbodyErr != nil {
		return nil, errors.Wrap(_xmlbodyErr, "Error parsing 'xmlbody' field of ExtensionObjectEncodingMask")
	}
	xmlbody := _xmlbody

	// Simple Field (binaryBody)
	_binaryBody, _binaryBodyErr := readBuffer.ReadBit("binaryBody")
	if _binaryBodyErr != nil {
		return nil, errors.Wrap(_binaryBodyErr, "Error parsing 'binaryBody' field of ExtensionObjectEncodingMask")
	}
	binaryBody := _binaryBody

	if closeErr := readBuffer.CloseContext("ExtensionObjectEncodingMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensionObjectEncodingMask")
	}

	// Create the instance
	return &_ExtensionObjectEncodingMask{
		TypeIdSpecified: typeIdSpecified,
		Xmlbody:         xmlbody,
		BinaryBody:      binaryBody,
		reservedField0:  reservedField0,
	}, nil
}

func (m *_ExtensionObjectEncodingMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ExtensionObjectEncodingMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ExtensionObjectEncodingMask"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExtensionObjectEncodingMask")
	}

	// Reserved Field (reserved)
	{
		var reserved int8 = int8(0x00)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": int8(0x00),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteInt8("reserved", 5, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (typeIdSpecified)
	typeIdSpecified := bool(m.GetTypeIdSpecified())
	_typeIdSpecifiedErr := writeBuffer.WriteBit("typeIdSpecified", (typeIdSpecified))
	if _typeIdSpecifiedErr != nil {
		return errors.Wrap(_typeIdSpecifiedErr, "Error serializing 'typeIdSpecified' field")
	}

	// Simple Field (xmlbody)
	xmlbody := bool(m.GetXmlbody())
	_xmlbodyErr := writeBuffer.WriteBit("xmlbody", (xmlbody))
	if _xmlbodyErr != nil {
		return errors.Wrap(_xmlbodyErr, "Error serializing 'xmlbody' field")
	}

	// Simple Field (binaryBody)
	binaryBody := bool(m.GetBinaryBody())
	_binaryBodyErr := writeBuffer.WriteBit("binaryBody", (binaryBody))
	if _binaryBodyErr != nil {
		return errors.Wrap(_binaryBodyErr, "Error serializing 'binaryBody' field")
	}

	if popErr := writeBuffer.PopContext("ExtensionObjectEncodingMask"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExtensionObjectEncodingMask")
	}
	return nil
}

func (m *_ExtensionObjectEncodingMask) isExtensionObjectEncodingMask() bool {
	return true
}

func (m *_ExtensionObjectEncodingMask) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}