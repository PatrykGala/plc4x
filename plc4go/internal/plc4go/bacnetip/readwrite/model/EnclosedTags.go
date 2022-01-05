/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type EnclosedTags struct {
	OpeningTag *BACnetContextTag
	Data       []*BACnetTag
	ClosingTag *BACnetContextTag
}

// The corresponding interface
type IEnclosedTags interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewEnclosedTags(openingTag *BACnetContextTag, data []*BACnetTag, closingTag *BACnetContextTag) *EnclosedTags {
	return &EnclosedTags{OpeningTag: openingTag, Data: data, ClosingTag: closingTag}
}

func CastEnclosedTags(structType interface{}) *EnclosedTags {
	castFunc := func(typ interface{}) *EnclosedTags {
		if casted, ok := typ.(EnclosedTags); ok {
			return &casted
		}
		if casted, ok := typ.(*EnclosedTags); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *EnclosedTags) GetTypeName() string {
	return "EnclosedTags"
}

func (m *EnclosedTags) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *EnclosedTags) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (openingTag)
	if m.OpeningTag != nil {
		lengthInBits += (*m.OpeningTag).LengthInBits()
	}

	// Manual Array Field (data)
	data := m.Data
	lengthInBits += TagsLength(data) * 8

	// Optional Field (closingTag)
	if m.ClosingTag != nil {
		lengthInBits += (*m.ClosingTag).LengthInBits()
	}

	return lengthInBits
}

func (m *EnclosedTags) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func EnclosedTagsParse(readBuffer utils.ReadBuffer, openingTagNumber uint8, closingTagNumber uint8) (*EnclosedTags, error) {
	if pullErr := readBuffer.PullContext("EnclosedTags"); pullErr != nil {
		return nil, pullErr
	}

	// Optional Field (openingTag) (Can be skipped, if a given expression evaluates to false)
	var openingTag *BACnetContextTag = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, openingTagNumber, BACnetDataType_NULL)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'openingTag' field")
		case _err == utils.ParseAssertError:
			readBuffer.SetPos(currentPos)
		default:
			openingTag = CastBACnetContextTag(_val)
			if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Manual Array Field (data)
	// Terminated array
	_dataList := make([]*BACnetTag, 0)
	for !((bool)(OpeningClosingTerminate(readBuffer, (openingTag)))) {
		_dataList = append(_dataList, ((*BACnetTag)(ParseTags(readBuffer))))

	}
	data := _dataList
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (closingTag) (Can be skipped, if a given expression evaluates to false)
	var closingTag *BACnetContextTag = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, closingTagNumber, BACnetDataType_NULL)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'closingTag' field")
		case _err == utils.ParseAssertError:
			readBuffer.SetPos(currentPos)
		default:
			closingTag = CastBACnetContextTag(_val)
			if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("EnclosedTags"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewEnclosedTags(openingTag, data, closingTag), nil
}

func (m *EnclosedTags) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("EnclosedTags"); pushErr != nil {
		return pushErr
	}

	// Optional Field (openingTag) (Can be skipped, if the value is null)
	var openingTag *BACnetContextTag = nil
	if m.OpeningTag != nil {
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return pushErr
		}
		openingTag = m.OpeningTag
		_openingTagErr := openingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return popErr
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}
	}

	// Manual Array Field (data)
	if m.Data != nil {
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, m := range m.Data {
			WriteTags(writeBuffer, m)
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Optional Field (closingTag) (Can be skipped, if the value is null)
	var closingTag *BACnetContextTag = nil
	if m.ClosingTag != nil {
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return pushErr
		}
		closingTag = m.ClosingTag
		_closingTagErr := closingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return popErr
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}
	}

	if popErr := writeBuffer.PopContext("EnclosedTags"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *EnclosedTags) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
