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
	"math"
)

// Code generated by code-generation. DO NOT EDIT.

// S7VarPayloadDataItem is the corresponding interface of S7VarPayloadDataItem
type S7VarPayloadDataItem interface {
	utils.LengthAware
	utils.Serializable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetData returns Data (property field)
	GetData() []byte
}

// S7VarPayloadDataItemExactly can be used when we want exactly this type and not a type which fulfills S7VarPayloadDataItem.
// This is useful for switch cases.
type S7VarPayloadDataItemExactly interface {
	S7VarPayloadDataItem
	isS7VarPayloadDataItem() bool
}

// _S7VarPayloadDataItem is the data-structure of this message
type _S7VarPayloadDataItem struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	Data          []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7VarPayloadDataItem) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_S7VarPayloadDataItem) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_S7VarPayloadDataItem) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7VarPayloadDataItem factory function for _S7VarPayloadDataItem
func NewS7VarPayloadDataItem(returnCode DataTransportErrorCode, transportSize DataTransportSize, data []byte) *_S7VarPayloadDataItem {
	return &_S7VarPayloadDataItem{ReturnCode: returnCode, TransportSize: transportSize, Data: data}
}

// Deprecated: use the interface for direct cast
func CastS7VarPayloadDataItem(structType interface{}) S7VarPayloadDataItem {
	if casted, ok := structType.(S7VarPayloadDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarPayloadDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarPayloadDataItem) GetTypeName() string {
	return "S7VarPayloadDataItem"
}

func (m *_S7VarPayloadDataItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7VarPayloadDataItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Implicit Field (dataLength)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	// Padding Field (padding)
	_timesPadding := uint8(int32(int32(len(m.GetData()))) % int32(int32(2)))
	for ; _timesPadding > 0; _timesPadding-- {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_S7VarPayloadDataItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7VarPayloadDataItemParse(readBuffer utils.ReadBuffer) (S7VarPayloadDataItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarPayloadDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarPayloadDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnCode")
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParse(readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field of S7VarPayloadDataItem")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnCode")
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
	_transportSize, _transportSizeErr := DataTransportSizeParse(readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field of S7VarPayloadDataItem")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	// Implicit Field (dataLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := readBuffer.ReadUint16("dataLength", 16)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field of S7VarPayloadDataItem")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf(transportSize.SizeInBits(), func() interface{} { return uint16(math.Ceil(float64(dataLength) / float64(float64(8.0)))) }, func() interface{} { return uint16(dataLength) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of S7VarPayloadDataItem")
	}

	// Padding Field (padding)
	{
		if pullErr := readBuffer.PullContext("padding", utils.WithRenderAsList(true)); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for padding")
		}
		_timesPadding := (int32(int32(len(data))) % int32(int32(2)))
		for ; (readBuffer.HasMore(8)) && (_timesPadding > 0); _timesPadding-- {
			// Just read the padding data and ignore it
			_, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'padding' field of S7VarPayloadDataItem")
			}
		}
		if closeErr := readBuffer.CloseContext("padding", utils.WithRenderAsList(true)); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for padding")
		}
	}

	if closeErr := readBuffer.CloseContext("S7VarPayloadDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarPayloadDataItem")
	}

	// Create the instance
	return NewS7VarPayloadDataItem(returnCode, transportSize, data), nil
}

func (m *_S7VarPayloadDataItem) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7VarPayloadDataItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7VarPayloadDataItem")
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnCode")
	}
	_returnCodeErr := writeBuffer.WriteSerializable(m.GetReturnCode())
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnCode")
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transportSize")
	}
	_transportSizeErr := writeBuffer.WriteSerializable(m.GetTransportSize())
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transportSize")
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint16(uint16(uint16(len(m.GetData()))) * uint16(uint16(utils.InlineIf(bool(bool((m.GetTransportSize()) == (DataTransportSize_BIT))), func() interface{} { return uint16(uint16(1)) }, func() interface{} {
		return uint16(uint16(utils.InlineIf(m.GetTransportSize().SizeInBits(), func() interface{} { return uint16(uint16(8)) }, func() interface{} { return uint16(uint16(1)) }).(uint16)))
	}).(uint16))))
	_dataLengthErr := writeBuffer.WriteUint16("dataLength", 16, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	// Padding Field (padding)
	{
		if pushErr := writeBuffer.PushContext("padding", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for padding")
		}
		_timesPadding := uint8(int32(int32(len(m.GetData()))) % int32(int32(2)))
		for ; _timesPadding > 0; _timesPadding-- {
			_paddingValue := uint8(0x00)
			_paddingErr := writeBuffer.WriteUint8("", 8, (_paddingValue))
			if _paddingErr != nil {
				return errors.Wrap(_paddingErr, "Error serializing 'padding' field")
			}
		}
		if popErr := writeBuffer.PopContext("padding", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for padding")
		}
	}

	if popErr := writeBuffer.PopContext("S7VarPayloadDataItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7VarPayloadDataItem")
	}
	return nil
}

func (m *_S7VarPayloadDataItem) isS7VarPayloadDataItem() bool {
	return true
}

func (m *_S7VarPayloadDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
