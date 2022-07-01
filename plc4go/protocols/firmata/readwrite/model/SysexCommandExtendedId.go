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

// SysexCommandExtendedId is the corresponding interface of SysexCommandExtendedId
type SysexCommandExtendedId interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetId returns Id (property field)
	GetId() []int8
}

// SysexCommandExtendedIdExactly can be used when we want exactly this type and not a type which fulfills SysexCommandExtendedId.
// This is useful for switch cases.
type SysexCommandExtendedIdExactly interface {
	SysexCommandExtendedId
	isSysexCommandExtendedId() bool
}

// _SysexCommandExtendedId is the data-structure of this message
type _SysexCommandExtendedId struct {
	*_SysexCommand
	Id []int8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandExtendedId) GetCommandType() uint8 {
	return 0x00
}

func (m *_SysexCommandExtendedId) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandExtendedId) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandExtendedId) GetParent() SysexCommand {
	return m._SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandExtendedId) GetId() []int8 {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandExtendedId factory function for _SysexCommandExtendedId
func NewSysexCommandExtendedId(id []int8) *_SysexCommandExtendedId {
	_result := &_SysexCommandExtendedId{
		Id:            id,
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandExtendedId(structType interface{}) SysexCommandExtendedId {
	if casted, ok := structType.(SysexCommandExtendedId); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandExtendedId); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandExtendedId) GetTypeName() string {
	return "SysexCommandExtendedId"
}

func (m *_SysexCommandExtendedId) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandExtendedId) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Id) > 0 {
		lengthInBits += 8 * uint16(len(m.Id))
	}

	return lengthInBits
}

func (m *_SysexCommandExtendedId) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandExtendedIdParse(readBuffer utils.ReadBuffer, response bool) (SysexCommandExtendedId, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandExtendedId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandExtendedId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (id)
	if pullErr := readBuffer.PullContext("id", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for id")
	}
	// Count array
	id := make([]int8, uint16(2))
	// This happens when the size is set conditional to 0
	if len(id) == 0 {
		id = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
			_item, _err := readBuffer.ReadInt8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'id' field of SysexCommandExtendedId")
			}
			id[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("id", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for id")
	}

	if closeErr := readBuffer.CloseContext("SysexCommandExtendedId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandExtendedId")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandExtendedId{
		Id:            id,
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandExtendedId) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandExtendedId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandExtendedId")
		}

		// Array Field (id)
		if pushErr := writeBuffer.PushContext("id", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for id")
		}
		for _, _element := range m.GetId() {
			_elementErr := writeBuffer.WriteInt8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'id' field")
			}
		}
		if popErr := writeBuffer.PopContext("id", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for id")
		}

		if popErr := writeBuffer.PopContext("SysexCommandExtendedId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandExtendedId")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SysexCommandExtendedId) isSysexCommandExtendedId() bool {
	return true
}

func (m *_SysexCommandExtendedId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
