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

// BACnetServiceAckReadPropertyMultiple is the corresponding interface of BACnetServiceAckReadPropertyMultiple
type BACnetServiceAckReadPropertyMultiple interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetData returns Data (property field)
	GetData() []BACnetReadAccessResult
}

// BACnetServiceAckReadPropertyMultipleExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckReadPropertyMultiple.
// This is useful for switch cases.
type BACnetServiceAckReadPropertyMultipleExactly interface {
	BACnetServiceAckReadPropertyMultiple
	isBACnetServiceAckReadPropertyMultiple() bool
}

// _BACnetServiceAckReadPropertyMultiple is the data-structure of this message
type _BACnetServiceAckReadPropertyMultiple struct {
	*_BACnetServiceAck
	Data []BACnetReadAccessResult

	// Arguments.
	ServiceAckPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckReadPropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckReadPropertyMultiple) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckReadPropertyMultiple) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckReadPropertyMultiple) GetData() []BACnetReadAccessResult {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckReadPropertyMultiple factory function for _BACnetServiceAckReadPropertyMultiple
func NewBACnetServiceAckReadPropertyMultiple(data []BACnetReadAccessResult, serviceAckLength uint16, serviceAckPayloadLength uint16) *_BACnetServiceAckReadPropertyMultiple {
	_result := &_BACnetServiceAckReadPropertyMultiple{
		Data:              data,
		_BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckReadPropertyMultiple(structType interface{}) BACnetServiceAckReadPropertyMultiple {
	if casted, ok := structType.(BACnetServiceAckReadPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadPropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckReadPropertyMultiple) GetTypeName() string {
	return "BACnetServiceAckReadPropertyMultiple"
}

func (m *_BACnetServiceAckReadPropertyMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckReadPropertyMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetServiceAckReadPropertyMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckReadPropertyMultipleParse(readBuffer utils.ReadBuffer, serviceAckLength uint16, serviceAckPayloadLength uint16) (BACnetServiceAckReadPropertyMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadPropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckReadPropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Length array
	var data []BACnetReadAccessResult
	{
		_dataLength := serviceAckPayloadLength
		_dataEndPos := positionAware.GetPos() + uint16(_dataLength)
		for positionAware.GetPos() < _dataEndPos {
			_item, _err := BACnetReadAccessResultParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field of BACnetServiceAckReadPropertyMultiple")
			}
			data = append(data, _item.(BACnetReadAccessResult))
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadPropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckReadPropertyMultiple")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckReadPropertyMultiple{
		Data: data,
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckReadPropertyMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadPropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckReadPropertyMultiple")
		}

		// Array Field (data)
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for data")
		}
		for _, _element := range m.GetData() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for data")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadPropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckReadPropertyMultiple")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetServiceAckReadPropertyMultiple) isBACnetServiceAckReadPropertyMultiple() bool {
	return true
}

func (m *_BACnetServiceAckReadPropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
