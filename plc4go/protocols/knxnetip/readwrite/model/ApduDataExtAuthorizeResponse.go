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

// ApduDataExtAuthorizeResponse is the corresponding interface of ApduDataExtAuthorizeResponse
type ApduDataExtAuthorizeResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetLevel returns Level (property field)
	GetLevel() uint8
}

// ApduDataExtAuthorizeResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtAuthorizeResponse.
// This is useful for switch cases.
type ApduDataExtAuthorizeResponseExactly interface {
	ApduDataExtAuthorizeResponse
	isApduDataExtAuthorizeResponse() bool
}

// _ApduDataExtAuthorizeResponse is the data-structure of this message
type _ApduDataExtAuthorizeResponse struct {
	*_ApduDataExt
	Level uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtAuthorizeResponse) GetExtApciType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtAuthorizeResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtAuthorizeResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtAuthorizeResponse) GetLevel() uint8 {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtAuthorizeResponse factory function for _ApduDataExtAuthorizeResponse
func NewApduDataExtAuthorizeResponse(level uint8, length uint8) *_ApduDataExtAuthorizeResponse {
	_result := &_ApduDataExtAuthorizeResponse{
		Level:        level,
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtAuthorizeResponse(structType interface{}) ApduDataExtAuthorizeResponse {
	if casted, ok := structType.(ApduDataExtAuthorizeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtAuthorizeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtAuthorizeResponse) GetTypeName() string {
	return "ApduDataExtAuthorizeResponse"
}

func (m *_ApduDataExtAuthorizeResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtAuthorizeResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (level)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ApduDataExtAuthorizeResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtAuthorizeResponseParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtAuthorizeResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtAuthorizeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtAuthorizeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (level)
	_level, _levelErr := readBuffer.ReadUint8("level", 8)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field of ApduDataExtAuthorizeResponse")
	}
	level := _level

	if closeErr := readBuffer.CloseContext("ApduDataExtAuthorizeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtAuthorizeResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtAuthorizeResponse{
		Level: level,
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtAuthorizeResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtAuthorizeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtAuthorizeResponse")
		}

		// Simple Field (level)
		level := uint8(m.GetLevel())
		_levelErr := writeBuffer.WriteUint8("level", 8, (level))
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtAuthorizeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtAuthorizeResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtAuthorizeResponse) isApduDataExtAuthorizeResponse() bool {
	return true
}

func (m *_ApduDataExtAuthorizeResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
