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

// KnxNetIpTunneling is the corresponding interface of KnxNetIpTunneling
type KnxNetIpTunneling interface {
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetIpTunnelingExactly can be used when we want exactly this type and not a type which fulfills KnxNetIpTunneling.
// This is useful for switch cases.
type KnxNetIpTunnelingExactly interface {
	KnxNetIpTunneling
	isKnxNetIpTunneling() bool
}

// _KnxNetIpTunneling is the data-structure of this message
type _KnxNetIpTunneling struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpTunneling) GetServiceType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpTunneling) InitializeParent(parent ServiceId) {}

func (m *_KnxNetIpTunneling) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpTunneling) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpTunneling factory function for _KnxNetIpTunneling
func NewKnxNetIpTunneling(version uint8) *_KnxNetIpTunneling {
	_result := &_KnxNetIpTunneling{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpTunneling(structType interface{}) KnxNetIpTunneling {
	if casted, ok := structType.(KnxNetIpTunneling); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpTunneling); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpTunneling) GetTypeName() string {
	return "KnxNetIpTunneling"
}

func (m *_KnxNetIpTunneling) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxNetIpTunneling) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpTunneling) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpTunnelingParse(readBuffer utils.ReadBuffer) (KnxNetIpTunneling, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpTunneling"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpTunneling")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of KnxNetIpTunneling")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetIpTunneling"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpTunneling")
	}

	// Create a partially initialized instance
	_child := &_KnxNetIpTunneling{
		Version:    version,
		_ServiceId: &_ServiceId{},
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetIpTunneling) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpTunneling"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpTunneling")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpTunneling"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpTunneling")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_KnxNetIpTunneling) isKnxNetIpTunneling() bool {
	return true
}

func (m *_KnxNetIpTunneling) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
