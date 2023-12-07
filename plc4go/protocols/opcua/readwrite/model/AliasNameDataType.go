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

// AliasNameDataType is the corresponding interface of AliasNameDataType
type AliasNameDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAliasName returns AliasName (property field)
	GetAliasName() QualifiedName
	// GetNoOfReferencedNodes returns NoOfReferencedNodes (property field)
	GetNoOfReferencedNodes() int32
	// GetReferencedNodes returns ReferencedNodes (property field)
	GetReferencedNodes() []ExpandedNodeId
}

// AliasNameDataTypeExactly can be used when we want exactly this type and not a type which fulfills AliasNameDataType.
// This is useful for switch cases.
type AliasNameDataTypeExactly interface {
	AliasNameDataType
	isAliasNameDataType() bool
}

// _AliasNameDataType is the data-structure of this message
type _AliasNameDataType struct {
	*_ExtensionObjectDefinition
	AliasName           QualifiedName
	NoOfReferencedNodes int32
	ReferencedNodes     []ExpandedNodeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AliasNameDataType) GetIdentifier() string {
	return "23470"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AliasNameDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_AliasNameDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AliasNameDataType) GetAliasName() QualifiedName {
	return m.AliasName
}

func (m *_AliasNameDataType) GetNoOfReferencedNodes() int32 {
	return m.NoOfReferencedNodes
}

func (m *_AliasNameDataType) GetReferencedNodes() []ExpandedNodeId {
	return m.ReferencedNodes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAliasNameDataType factory function for _AliasNameDataType
func NewAliasNameDataType(aliasName QualifiedName, noOfReferencedNodes int32, referencedNodes []ExpandedNodeId) *_AliasNameDataType {
	_result := &_AliasNameDataType{
		AliasName:                  aliasName,
		NoOfReferencedNodes:        noOfReferencedNodes,
		ReferencedNodes:            referencedNodes,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAliasNameDataType(structType any) AliasNameDataType {
	if casted, ok := structType.(AliasNameDataType); ok {
		return casted
	}
	if casted, ok := structType.(*AliasNameDataType); ok {
		return *casted
	}
	return nil
}

func (m *_AliasNameDataType) GetTypeName() string {
	return "AliasNameDataType"
}

func (m *_AliasNameDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (aliasName)
	lengthInBits += m.AliasName.GetLengthInBits(ctx)

	// Simple field (noOfReferencedNodes)
	lengthInBits += 32

	// Array field
	if len(m.ReferencedNodes) > 0 {
		for _curItem, element := range m.ReferencedNodes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencedNodes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AliasNameDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AliasNameDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (AliasNameDataType, error) {
	return AliasNameDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func AliasNameDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (AliasNameDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AliasNameDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AliasNameDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (aliasName)
	if pullErr := readBuffer.PullContext("aliasName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for aliasName")
	}
	_aliasName, _aliasNameErr := QualifiedNameParseWithBuffer(ctx, readBuffer)
	if _aliasNameErr != nil {
		return nil, errors.Wrap(_aliasNameErr, "Error parsing 'aliasName' field of AliasNameDataType")
	}
	aliasName := _aliasName.(QualifiedName)
	if closeErr := readBuffer.CloseContext("aliasName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for aliasName")
	}

	// Simple Field (noOfReferencedNodes)
	_noOfReferencedNodes, _noOfReferencedNodesErr := readBuffer.ReadInt32("noOfReferencedNodes", 32)
	if _noOfReferencedNodesErr != nil {
		return nil, errors.Wrap(_noOfReferencedNodesErr, "Error parsing 'noOfReferencedNodes' field of AliasNameDataType")
	}
	noOfReferencedNodes := _noOfReferencedNodes

	// Array field (referencedNodes)
	if pullErr := readBuffer.PullContext("referencedNodes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referencedNodes")
	}
	// Count array
	referencedNodes := make([]ExpandedNodeId, utils.Max(noOfReferencedNodes, 0))
	// This happens when the size is set conditional to 0
	if len(referencedNodes) == 0 {
		referencedNodes = nil
	}
	{
		_numItems := uint16(utils.Max(noOfReferencedNodes, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExpandedNodeIdParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'referencedNodes' field of AliasNameDataType")
			}
			referencedNodes[_curItem] = _item.(ExpandedNodeId)
		}
	}
	if closeErr := readBuffer.CloseContext("referencedNodes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referencedNodes")
	}

	if closeErr := readBuffer.CloseContext("AliasNameDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AliasNameDataType")
	}

	// Create a partially initialized instance
	_child := &_AliasNameDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		AliasName:                  aliasName,
		NoOfReferencedNodes:        noOfReferencedNodes,
		ReferencedNodes:            referencedNodes,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_AliasNameDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AliasNameDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AliasNameDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AliasNameDataType")
		}

		// Simple Field (aliasName)
		if pushErr := writeBuffer.PushContext("aliasName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for aliasName")
		}
		_aliasNameErr := writeBuffer.WriteSerializable(ctx, m.GetAliasName())
		if popErr := writeBuffer.PopContext("aliasName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for aliasName")
		}
		if _aliasNameErr != nil {
			return errors.Wrap(_aliasNameErr, "Error serializing 'aliasName' field")
		}

		// Simple Field (noOfReferencedNodes)
		noOfReferencedNodes := int32(m.GetNoOfReferencedNodes())
		_noOfReferencedNodesErr := writeBuffer.WriteInt32("noOfReferencedNodes", 32, int32((noOfReferencedNodes)))
		if _noOfReferencedNodesErr != nil {
			return errors.Wrap(_noOfReferencedNodesErr, "Error serializing 'noOfReferencedNodes' field")
		}

		// Array Field (referencedNodes)
		if pushErr := writeBuffer.PushContext("referencedNodes", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referencedNodes")
		}
		for _curItem, _element := range m.GetReferencedNodes() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetReferencedNodes()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'referencedNodes' field")
			}
		}
		if popErr := writeBuffer.PopContext("referencedNodes", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referencedNodes")
		}

		if popErr := writeBuffer.PopContext("AliasNameDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AliasNameDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AliasNameDataType) isAliasNameDataType() bool {
	return true
}

func (m *_AliasNameDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
