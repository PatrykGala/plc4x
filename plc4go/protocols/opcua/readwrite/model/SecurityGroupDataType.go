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

// SecurityGroupDataType is the corresponding interface of SecurityGroupDataType
type SecurityGroupDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetNoOfSecurityGroupFolder returns NoOfSecurityGroupFolder (property field)
	GetNoOfSecurityGroupFolder() int32
	// GetSecurityGroupFolder returns SecurityGroupFolder (property field)
	GetSecurityGroupFolder() []PascalString
	// GetKeyLifetime returns KeyLifetime (property field)
	GetKeyLifetime() float64
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetMaxFutureKeyCount returns MaxFutureKeyCount (property field)
	GetMaxFutureKeyCount() uint32
	// GetMaxPastKeyCount returns MaxPastKeyCount (property field)
	GetMaxPastKeyCount() uint32
	// GetSecurityGroupId returns SecurityGroupId (property field)
	GetSecurityGroupId() PascalString
	// GetNoOfRolePermissions returns NoOfRolePermissions (property field)
	GetNoOfRolePermissions() int32
	// GetRolePermissions returns RolePermissions (property field)
	GetRolePermissions() []ExtensionObjectDefinition
	// GetNoOfGroupProperties returns NoOfGroupProperties (property field)
	GetNoOfGroupProperties() int32
	// GetGroupProperties returns GroupProperties (property field)
	GetGroupProperties() []ExtensionObjectDefinition
}

// SecurityGroupDataTypeExactly can be used when we want exactly this type and not a type which fulfills SecurityGroupDataType.
// This is useful for switch cases.
type SecurityGroupDataTypeExactly interface {
	SecurityGroupDataType
	isSecurityGroupDataType() bool
}

// _SecurityGroupDataType is the data-structure of this message
type _SecurityGroupDataType struct {
	*_ExtensionObjectDefinition
	Name                    PascalString
	NoOfSecurityGroupFolder int32
	SecurityGroupFolder     []PascalString
	KeyLifetime             float64
	SecurityPolicyUri       PascalString
	MaxFutureKeyCount       uint32
	MaxPastKeyCount         uint32
	SecurityGroupId         PascalString
	NoOfRolePermissions     int32
	RolePermissions         []ExtensionObjectDefinition
	NoOfGroupProperties     int32
	GroupProperties         []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SecurityGroupDataType) GetIdentifier() string {
	return "23603"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityGroupDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SecurityGroupDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityGroupDataType) GetName() PascalString {
	return m.Name
}

func (m *_SecurityGroupDataType) GetNoOfSecurityGroupFolder() int32 {
	return m.NoOfSecurityGroupFolder
}

func (m *_SecurityGroupDataType) GetSecurityGroupFolder() []PascalString {
	return m.SecurityGroupFolder
}

func (m *_SecurityGroupDataType) GetKeyLifetime() float64 {
	return m.KeyLifetime
}

func (m *_SecurityGroupDataType) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_SecurityGroupDataType) GetMaxFutureKeyCount() uint32 {
	return m.MaxFutureKeyCount
}

func (m *_SecurityGroupDataType) GetMaxPastKeyCount() uint32 {
	return m.MaxPastKeyCount
}

func (m *_SecurityGroupDataType) GetSecurityGroupId() PascalString {
	return m.SecurityGroupId
}

func (m *_SecurityGroupDataType) GetNoOfRolePermissions() int32 {
	return m.NoOfRolePermissions
}

func (m *_SecurityGroupDataType) GetRolePermissions() []ExtensionObjectDefinition {
	return m.RolePermissions
}

func (m *_SecurityGroupDataType) GetNoOfGroupProperties() int32 {
	return m.NoOfGroupProperties
}

func (m *_SecurityGroupDataType) GetGroupProperties() []ExtensionObjectDefinition {
	return m.GroupProperties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityGroupDataType factory function for _SecurityGroupDataType
func NewSecurityGroupDataType(name PascalString, noOfSecurityGroupFolder int32, securityGroupFolder []PascalString, keyLifetime float64, securityPolicyUri PascalString, maxFutureKeyCount uint32, maxPastKeyCount uint32, securityGroupId PascalString, noOfRolePermissions int32, rolePermissions []ExtensionObjectDefinition, noOfGroupProperties int32, groupProperties []ExtensionObjectDefinition) *_SecurityGroupDataType {
	_result := &_SecurityGroupDataType{
		Name:                       name,
		NoOfSecurityGroupFolder:    noOfSecurityGroupFolder,
		SecurityGroupFolder:        securityGroupFolder,
		KeyLifetime:                keyLifetime,
		SecurityPolicyUri:          securityPolicyUri,
		MaxFutureKeyCount:          maxFutureKeyCount,
		MaxPastKeyCount:            maxPastKeyCount,
		SecurityGroupId:            securityGroupId,
		NoOfRolePermissions:        noOfRolePermissions,
		RolePermissions:            rolePermissions,
		NoOfGroupProperties:        noOfGroupProperties,
		GroupProperties:            groupProperties,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityGroupDataType(structType any) SecurityGroupDataType {
	if casted, ok := structType.(SecurityGroupDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityGroupDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityGroupDataType) GetTypeName() string {
	return "SecurityGroupDataType"
}

func (m *_SecurityGroupDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (noOfSecurityGroupFolder)
	lengthInBits += 32

	// Array field
	if len(m.SecurityGroupFolder) > 0 {
		for _curItem, element := range m.SecurityGroupFolder {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SecurityGroupFolder), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (keyLifetime)
	lengthInBits += 64

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (maxFutureKeyCount)
	lengthInBits += 32

	// Simple field (maxPastKeyCount)
	lengthInBits += 32

	// Simple field (securityGroupId)
	lengthInBits += m.SecurityGroupId.GetLengthInBits(ctx)

	// Simple field (noOfRolePermissions)
	lengthInBits += 32

	// Array field
	if len(m.RolePermissions) > 0 {
		for _curItem, element := range m.RolePermissions {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.RolePermissions), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfGroupProperties)
	lengthInBits += 32

	// Array field
	if len(m.GroupProperties) > 0 {
		for _curItem, element := range m.GroupProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GroupProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SecurityGroupDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityGroupDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (SecurityGroupDataType, error) {
	return SecurityGroupDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SecurityGroupDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SecurityGroupDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SecurityGroupDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityGroupDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of SecurityGroupDataType")
	}
	name := _name.(PascalString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	// Simple Field (noOfSecurityGroupFolder)
	_noOfSecurityGroupFolder, _noOfSecurityGroupFolderErr := readBuffer.ReadInt32("noOfSecurityGroupFolder", 32)
	if _noOfSecurityGroupFolderErr != nil {
		return nil, errors.Wrap(_noOfSecurityGroupFolderErr, "Error parsing 'noOfSecurityGroupFolder' field of SecurityGroupDataType")
	}
	noOfSecurityGroupFolder := _noOfSecurityGroupFolder

	// Array field (securityGroupFolder)
	if pullErr := readBuffer.PullContext("securityGroupFolder", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityGroupFolder")
	}
	// Count array
	securityGroupFolder := make([]PascalString, utils.Max(noOfSecurityGroupFolder, 0))
	// This happens when the size is set conditional to 0
	if len(securityGroupFolder) == 0 {
		securityGroupFolder = nil
	}
	{
		_numItems := uint16(utils.Max(noOfSecurityGroupFolder, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'securityGroupFolder' field of SecurityGroupDataType")
			}
			securityGroupFolder[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("securityGroupFolder", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityGroupFolder")
	}

	// Simple Field (keyLifetime)
	_keyLifetime, _keyLifetimeErr := readBuffer.ReadFloat64("keyLifetime", 64)
	if _keyLifetimeErr != nil {
		return nil, errors.Wrap(_keyLifetimeErr, "Error parsing 'keyLifetime' field of SecurityGroupDataType")
	}
	keyLifetime := _keyLifetime

	// Simple Field (securityPolicyUri)
	if pullErr := readBuffer.PullContext("securityPolicyUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityPolicyUri")
	}
	_securityPolicyUri, _securityPolicyUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _securityPolicyUriErr != nil {
		return nil, errors.Wrap(_securityPolicyUriErr, "Error parsing 'securityPolicyUri' field of SecurityGroupDataType")
	}
	securityPolicyUri := _securityPolicyUri.(PascalString)
	if closeErr := readBuffer.CloseContext("securityPolicyUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityPolicyUri")
	}

	// Simple Field (maxFutureKeyCount)
	_maxFutureKeyCount, _maxFutureKeyCountErr := readBuffer.ReadUint32("maxFutureKeyCount", 32)
	if _maxFutureKeyCountErr != nil {
		return nil, errors.Wrap(_maxFutureKeyCountErr, "Error parsing 'maxFutureKeyCount' field of SecurityGroupDataType")
	}
	maxFutureKeyCount := _maxFutureKeyCount

	// Simple Field (maxPastKeyCount)
	_maxPastKeyCount, _maxPastKeyCountErr := readBuffer.ReadUint32("maxPastKeyCount", 32)
	if _maxPastKeyCountErr != nil {
		return nil, errors.Wrap(_maxPastKeyCountErr, "Error parsing 'maxPastKeyCount' field of SecurityGroupDataType")
	}
	maxPastKeyCount := _maxPastKeyCount

	// Simple Field (securityGroupId)
	if pullErr := readBuffer.PullContext("securityGroupId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityGroupId")
	}
	_securityGroupId, _securityGroupIdErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _securityGroupIdErr != nil {
		return nil, errors.Wrap(_securityGroupIdErr, "Error parsing 'securityGroupId' field of SecurityGroupDataType")
	}
	securityGroupId := _securityGroupId.(PascalString)
	if closeErr := readBuffer.CloseContext("securityGroupId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityGroupId")
	}

	// Simple Field (noOfRolePermissions)
	_noOfRolePermissions, _noOfRolePermissionsErr := readBuffer.ReadInt32("noOfRolePermissions", 32)
	if _noOfRolePermissionsErr != nil {
		return nil, errors.Wrap(_noOfRolePermissionsErr, "Error parsing 'noOfRolePermissions' field of SecurityGroupDataType")
	}
	noOfRolePermissions := _noOfRolePermissions

	// Array field (rolePermissions)
	if pullErr := readBuffer.PullContext("rolePermissions", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for rolePermissions")
	}
	// Count array
	rolePermissions := make([]ExtensionObjectDefinition, utils.Max(noOfRolePermissions, 0))
	// This happens when the size is set conditional to 0
	if len(rolePermissions) == 0 {
		rolePermissions = nil
	}
	{
		_numItems := uint16(utils.Max(noOfRolePermissions, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "98")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'rolePermissions' field of SecurityGroupDataType")
			}
			rolePermissions[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("rolePermissions", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for rolePermissions")
	}

	// Simple Field (noOfGroupProperties)
	_noOfGroupProperties, _noOfGroupPropertiesErr := readBuffer.ReadInt32("noOfGroupProperties", 32)
	if _noOfGroupPropertiesErr != nil {
		return nil, errors.Wrap(_noOfGroupPropertiesErr, "Error parsing 'noOfGroupProperties' field of SecurityGroupDataType")
	}
	noOfGroupProperties := _noOfGroupProperties

	// Array field (groupProperties)
	if pullErr := readBuffer.PullContext("groupProperties", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for groupProperties")
	}
	// Count array
	groupProperties := make([]ExtensionObjectDefinition, utils.Max(noOfGroupProperties, 0))
	// This happens when the size is set conditional to 0
	if len(groupProperties) == 0 {
		groupProperties = nil
	}
	{
		_numItems := uint16(utils.Max(noOfGroupProperties, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "14535")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'groupProperties' field of SecurityGroupDataType")
			}
			groupProperties[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("groupProperties", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for groupProperties")
	}

	if closeErr := readBuffer.CloseContext("SecurityGroupDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityGroupDataType")
	}

	// Create a partially initialized instance
	_child := &_SecurityGroupDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Name:                       name,
		NoOfSecurityGroupFolder:    noOfSecurityGroupFolder,
		SecurityGroupFolder:        securityGroupFolder,
		KeyLifetime:                keyLifetime,
		SecurityPolicyUri:          securityPolicyUri,
		MaxFutureKeyCount:          maxFutureKeyCount,
		MaxPastKeyCount:            maxPastKeyCount,
		SecurityGroupId:            securityGroupId,
		NoOfRolePermissions:        noOfRolePermissions,
		RolePermissions:            rolePermissions,
		NoOfGroupProperties:        noOfGroupProperties,
		GroupProperties:            groupProperties,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SecurityGroupDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityGroupDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityGroupDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityGroupDataType")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
		}

		// Simple Field (noOfSecurityGroupFolder)
		noOfSecurityGroupFolder := int32(m.GetNoOfSecurityGroupFolder())
		_noOfSecurityGroupFolderErr := writeBuffer.WriteInt32("noOfSecurityGroupFolder", 32, (noOfSecurityGroupFolder))
		if _noOfSecurityGroupFolderErr != nil {
			return errors.Wrap(_noOfSecurityGroupFolderErr, "Error serializing 'noOfSecurityGroupFolder' field")
		}

		// Array Field (securityGroupFolder)
		if pushErr := writeBuffer.PushContext("securityGroupFolder", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityGroupFolder")
		}
		for _curItem, _element := range m.GetSecurityGroupFolder() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetSecurityGroupFolder()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'securityGroupFolder' field")
			}
		}
		if popErr := writeBuffer.PopContext("securityGroupFolder", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityGroupFolder")
		}

		// Simple Field (keyLifetime)
		keyLifetime := float64(m.GetKeyLifetime())
		_keyLifetimeErr := writeBuffer.WriteFloat64("keyLifetime", 64, (keyLifetime))
		if _keyLifetimeErr != nil {
			return errors.Wrap(_keyLifetimeErr, "Error serializing 'keyLifetime' field")
		}

		// Simple Field (securityPolicyUri)
		if pushErr := writeBuffer.PushContext("securityPolicyUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityPolicyUri")
		}
		_securityPolicyUriErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityPolicyUri())
		if popErr := writeBuffer.PopContext("securityPolicyUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityPolicyUri")
		}
		if _securityPolicyUriErr != nil {
			return errors.Wrap(_securityPolicyUriErr, "Error serializing 'securityPolicyUri' field")
		}

		// Simple Field (maxFutureKeyCount)
		maxFutureKeyCount := uint32(m.GetMaxFutureKeyCount())
		_maxFutureKeyCountErr := writeBuffer.WriteUint32("maxFutureKeyCount", 32, (maxFutureKeyCount))
		if _maxFutureKeyCountErr != nil {
			return errors.Wrap(_maxFutureKeyCountErr, "Error serializing 'maxFutureKeyCount' field")
		}

		// Simple Field (maxPastKeyCount)
		maxPastKeyCount := uint32(m.GetMaxPastKeyCount())
		_maxPastKeyCountErr := writeBuffer.WriteUint32("maxPastKeyCount", 32, (maxPastKeyCount))
		if _maxPastKeyCountErr != nil {
			return errors.Wrap(_maxPastKeyCountErr, "Error serializing 'maxPastKeyCount' field")
		}

		// Simple Field (securityGroupId)
		if pushErr := writeBuffer.PushContext("securityGroupId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityGroupId")
		}
		_securityGroupIdErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityGroupId())
		if popErr := writeBuffer.PopContext("securityGroupId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityGroupId")
		}
		if _securityGroupIdErr != nil {
			return errors.Wrap(_securityGroupIdErr, "Error serializing 'securityGroupId' field")
		}

		// Simple Field (noOfRolePermissions)
		noOfRolePermissions := int32(m.GetNoOfRolePermissions())
		_noOfRolePermissionsErr := writeBuffer.WriteInt32("noOfRolePermissions", 32, (noOfRolePermissions))
		if _noOfRolePermissionsErr != nil {
			return errors.Wrap(_noOfRolePermissionsErr, "Error serializing 'noOfRolePermissions' field")
		}

		// Array Field (rolePermissions)
		if pushErr := writeBuffer.PushContext("rolePermissions", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for rolePermissions")
		}
		for _curItem, _element := range m.GetRolePermissions() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetRolePermissions()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'rolePermissions' field")
			}
		}
		if popErr := writeBuffer.PopContext("rolePermissions", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for rolePermissions")
		}

		// Simple Field (noOfGroupProperties)
		noOfGroupProperties := int32(m.GetNoOfGroupProperties())
		_noOfGroupPropertiesErr := writeBuffer.WriteInt32("noOfGroupProperties", 32, (noOfGroupProperties))
		if _noOfGroupPropertiesErr != nil {
			return errors.Wrap(_noOfGroupPropertiesErr, "Error serializing 'noOfGroupProperties' field")
		}

		// Array Field (groupProperties)
		if pushErr := writeBuffer.PushContext("groupProperties", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for groupProperties")
		}
		for _curItem, _element := range m.GetGroupProperties() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetGroupProperties()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'groupProperties' field")
			}
		}
		if popErr := writeBuffer.PopContext("groupProperties", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for groupProperties")
		}

		if popErr := writeBuffer.PopContext("SecurityGroupDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityGroupDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityGroupDataType) isSecurityGroupDataType() bool {
	return true
}

func (m *_SecurityGroupDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}