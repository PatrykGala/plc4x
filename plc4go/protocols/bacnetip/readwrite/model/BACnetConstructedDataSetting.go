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

// BACnetConstructedDataSetting is the corresponding interface of BACnetConstructedDataSetting
type BACnetConstructedDataSetting interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSetting returns Setting (property field)
	GetSetting() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataSettingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSetting.
// This is useful for switch cases.
type BACnetConstructedDataSettingExactly interface {
	BACnetConstructedDataSetting
	isBACnetConstructedDataSetting() bool
}

// _BACnetConstructedDataSetting is the data-structure of this message
type _BACnetConstructedDataSetting struct {
	*_BACnetConstructedData
	Setting BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSetting) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSetting) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SETTING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSetting) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSetting) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSetting) GetSetting() BACnetApplicationTagUnsignedInteger {
	return m.Setting
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSetting) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetSetting())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSetting factory function for _BACnetConstructedDataSetting
func NewBACnetConstructedDataSetting(setting BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSetting {
	_result := &_BACnetConstructedDataSetting{
		Setting:                setting,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSetting(structType interface{}) BACnetConstructedDataSetting {
	if casted, ok := structType.(BACnetConstructedDataSetting); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSetting); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSetting) GetTypeName() string {
	return "BACnetConstructedDataSetting"
}

func (m *_BACnetConstructedDataSetting) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSetting) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (setting)
	lengthInBits += m.Setting.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSetting) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSettingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSetting, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSetting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSetting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (setting)
	if pullErr := readBuffer.PullContext("setting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for setting")
	}
	_setting, _settingErr := BACnetApplicationTagParse(readBuffer)
	if _settingErr != nil {
		return nil, errors.Wrap(_settingErr, "Error parsing 'setting' field of BACnetConstructedDataSetting")
	}
	setting := _setting.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("setting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for setting")
	}

	// Virtual field
	_actualValue := setting
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSetting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSetting")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSetting{
		Setting: setting,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSetting) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSetting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSetting")
		}

		// Simple Field (setting)
		if pushErr := writeBuffer.PushContext("setting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for setting")
		}
		_settingErr := writeBuffer.WriteSerializable(m.GetSetting())
		if popErr := writeBuffer.PopContext("setting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for setting")
		}
		if _settingErr != nil {
			return errors.Wrap(_settingErr, "Error serializing 'setting' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSetting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSetting")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSetting) isBACnetConstructedDataSetting() bool {
	return true
}

func (m *_BACnetConstructedDataSetting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
