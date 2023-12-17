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

// OpcuaNodeIdServicesVariableSemantic is an enum
type OpcuaNodeIdServicesVariableSemantic int32

type IOpcuaNodeIdServicesVariableSemantic interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes OpcuaNodeIdServicesVariableSemantic = 2739
)

var OpcuaNodeIdServicesVariableSemanticValues []OpcuaNodeIdServicesVariableSemantic

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableSemanticValues = []OpcuaNodeIdServicesVariableSemantic{
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes,
	}
}

func OpcuaNodeIdServicesVariableSemanticByValue(value int32) (enum OpcuaNodeIdServicesVariableSemantic, ok bool) {
	switch value {
	case 2739:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSemanticByName(value string) (enum OpcuaNodeIdServicesVariableSemantic, ok bool) {
	switch value {
	case "SemanticChangeEventType_Changes":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSemanticKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableSemanticValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableSemantic(structType any) OpcuaNodeIdServicesVariableSemantic {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableSemantic {
		if sOpcuaNodeIdServicesVariableSemantic, ok := typ.(OpcuaNodeIdServicesVariableSemantic); ok {
			return sOpcuaNodeIdServicesVariableSemantic
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableSemantic) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableSemantic) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableSemanticParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableSemantic, error) {
	return OpcuaNodeIdServicesVariableSemanticParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableSemanticParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableSemantic, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableSemantic", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableSemantic")
	}
	if enum, ok := OpcuaNodeIdServicesVariableSemanticByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableSemantic")
		return OpcuaNodeIdServicesVariableSemantic(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableSemantic) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableSemantic) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableSemantic", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableSemantic) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes:
		return "SemanticChangeEventType_Changes"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableSemantic) String() string {
	return e.PLC4XEnumName()
}
