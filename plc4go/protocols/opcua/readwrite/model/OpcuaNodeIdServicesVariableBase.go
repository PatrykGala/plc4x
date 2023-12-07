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

// OpcuaNodeIdServicesVariableBase is an enum
type OpcuaNodeIdServicesVariableBase int32

type IOpcuaNodeIdServicesVariableBase interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableBase_BaseEventType_EventId                          OpcuaNodeIdServicesVariableBase = 2042
	OpcuaNodeIdServicesVariableBase_BaseEventType_EventType                        OpcuaNodeIdServicesVariableBase = 2043
	OpcuaNodeIdServicesVariableBase_BaseEventType_SourceNode                       OpcuaNodeIdServicesVariableBase = 2044
	OpcuaNodeIdServicesVariableBase_BaseEventType_SourceName                       OpcuaNodeIdServicesVariableBase = 2045
	OpcuaNodeIdServicesVariableBase_BaseEventType_Time                             OpcuaNodeIdServicesVariableBase = 2046
	OpcuaNodeIdServicesVariableBase_BaseEventType_ReceiveTime                      OpcuaNodeIdServicesVariableBase = 2047
	OpcuaNodeIdServicesVariableBase_BaseEventType_Message                          OpcuaNodeIdServicesVariableBase = 2050
	OpcuaNodeIdServicesVariableBase_BaseEventType_Severity                         OpcuaNodeIdServicesVariableBase = 2051
	OpcuaNodeIdServicesVariableBase_BaseEventType_LocalTime                        OpcuaNodeIdServicesVariableBase = 3190
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventId               OpcuaNodeIdServicesVariableBase = 3671
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventType             OpcuaNodeIdServicesVariableBase = 3672
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceNode            OpcuaNodeIdServicesVariableBase = 3673
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceName            OpcuaNodeIdServicesVariableBase = 3674
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Time                  OpcuaNodeIdServicesVariableBase = 3675
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ReceiveTime           OpcuaNodeIdServicesVariableBase = 3676
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_LocalTime             OpcuaNodeIdServicesVariableBase = 3677
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Message               OpcuaNodeIdServicesVariableBase = 3678
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Severity              OpcuaNodeIdServicesVariableBase = 3679
	OpcuaNodeIdServicesVariableBase_BaseAnalogType_Definition                      OpcuaNodeIdServicesVariableBase = 17565
	OpcuaNodeIdServicesVariableBase_BaseAnalogType_ValuePrecision                  OpcuaNodeIdServicesVariableBase = 17566
	OpcuaNodeIdServicesVariableBase_BaseAnalogType_InstrumentRange                 OpcuaNodeIdServicesVariableBase = 17567
	OpcuaNodeIdServicesVariableBase_BaseAnalogType_EURange                         OpcuaNodeIdServicesVariableBase = 17568
	OpcuaNodeIdServicesVariableBase_BaseAnalogType_EngineeringUnits                OpcuaNodeIdServicesVariableBase = 17569
	OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassId                 OpcuaNodeIdServicesVariableBase = 31771
	OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassName               OpcuaNodeIdServicesVariableBase = 31772
	OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassId              OpcuaNodeIdServicesVariableBase = 31773
	OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassName            OpcuaNodeIdServicesVariableBase = 31774
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassId      OpcuaNodeIdServicesVariableBase = 31887
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassName    OpcuaNodeIdServicesVariableBase = 31888
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassId   OpcuaNodeIdServicesVariableBase = 31889
	OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassName OpcuaNodeIdServicesVariableBase = 31890
)

var OpcuaNodeIdServicesVariableBaseValues []OpcuaNodeIdServicesVariableBase

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableBaseValues = []OpcuaNodeIdServicesVariableBase{
		OpcuaNodeIdServicesVariableBase_BaseEventType_EventId,
		OpcuaNodeIdServicesVariableBase_BaseEventType_EventType,
		OpcuaNodeIdServicesVariableBase_BaseEventType_SourceNode,
		OpcuaNodeIdServicesVariableBase_BaseEventType_SourceName,
		OpcuaNodeIdServicesVariableBase_BaseEventType_Time,
		OpcuaNodeIdServicesVariableBase_BaseEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableBase_BaseEventType_Message,
		OpcuaNodeIdServicesVariableBase_BaseEventType_Severity,
		OpcuaNodeIdServicesVariableBase_BaseEventType_LocalTime,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventId,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventType,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceNode,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceName,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Time,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_LocalTime,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Message,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Severity,
		OpcuaNodeIdServicesVariableBase_BaseAnalogType_Definition,
		OpcuaNodeIdServicesVariableBase_BaseAnalogType_ValuePrecision,
		OpcuaNodeIdServicesVariableBase_BaseAnalogType_InstrumentRange,
		OpcuaNodeIdServicesVariableBase_BaseAnalogType_EURange,
		OpcuaNodeIdServicesVariableBase_BaseAnalogType_EngineeringUnits,
		OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassName,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassName,
	}
}

func OpcuaNodeIdServicesVariableBaseByValue(value int32) (enum OpcuaNodeIdServicesVariableBase, ok bool) {
	switch value {
	case 17565:
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_Definition, true
	case 17566:
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_ValuePrecision, true
	case 17567:
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_InstrumentRange, true
	case 17568:
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_EURange, true
	case 17569:
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_EngineeringUnits, true
	case 2042:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_EventId, true
	case 2043:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_EventType, true
	case 2044:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_SourceNode, true
	case 2045:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_SourceName, true
	case 2046:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_Time, true
	case 2047:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ReceiveTime, true
	case 2050:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_Message, true
	case 2051:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_Severity, true
	case 31771:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassId, true
	case 31772:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassName, true
	case 31773:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassId, true
	case 31774:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassName, true
	case 31887:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassId, true
	case 31888:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassName, true
	case 31889:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassId, true
	case 31890:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassName, true
	case 3190:
		return OpcuaNodeIdServicesVariableBase_BaseEventType_LocalTime, true
	case 3671:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventId, true
	case 3672:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventType, true
	case 3673:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceNode, true
	case 3674:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceName, true
	case 3675:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Time, true
	case 3676:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ReceiveTime, true
	case 3677:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_LocalTime, true
	case 3678:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Message, true
	case 3679:
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBaseByName(value string) (enum OpcuaNodeIdServicesVariableBase, ok bool) {
	switch value {
	case "BaseAnalogType_Definition":
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_Definition, true
	case "BaseAnalogType_ValuePrecision":
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_ValuePrecision, true
	case "BaseAnalogType_InstrumentRange":
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_InstrumentRange, true
	case "BaseAnalogType_EURange":
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_EURange, true
	case "BaseAnalogType_EngineeringUnits":
		return OpcuaNodeIdServicesVariableBase_BaseAnalogType_EngineeringUnits, true
	case "BaseEventType_EventId":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_EventId, true
	case "BaseEventType_EventType":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_EventType, true
	case "BaseEventType_SourceNode":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_SourceNode, true
	case "BaseEventType_SourceName":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_SourceName, true
	case "BaseEventType_Time":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_Time, true
	case "BaseEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ReceiveTime, true
	case "BaseEventType_Message":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_Message, true
	case "BaseEventType_Severity":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_Severity, true
	case "BaseEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassId, true
	case "BaseEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassName, true
	case "BaseEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassId, true
	case "BaseEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassName, true
	case "BaseModelChangeEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassId, true
	case "BaseModelChangeEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassName, true
	case "BaseModelChangeEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassId, true
	case "BaseModelChangeEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassName, true
	case "BaseEventType_LocalTime":
		return OpcuaNodeIdServicesVariableBase_BaseEventType_LocalTime, true
	case "BaseModelChangeEventType_EventId":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventId, true
	case "BaseModelChangeEventType_EventType":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventType, true
	case "BaseModelChangeEventType_SourceNode":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceNode, true
	case "BaseModelChangeEventType_SourceName":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceName, true
	case "BaseModelChangeEventType_Time":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Time, true
	case "BaseModelChangeEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ReceiveTime, true
	case "BaseModelChangeEventType_LocalTime":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_LocalTime, true
	case "BaseModelChangeEventType_Message":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Message, true
	case "BaseModelChangeEventType_Severity":
		return OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBaseKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableBaseValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableBase(structType any) OpcuaNodeIdServicesVariableBase {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableBase {
		if sOpcuaNodeIdServicesVariableBase, ok := typ.(OpcuaNodeIdServicesVariableBase); ok {
			return sOpcuaNodeIdServicesVariableBase
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableBase) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableBase) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableBaseParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableBase, error) {
	return OpcuaNodeIdServicesVariableBaseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableBaseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableBase, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableBase", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableBase")
	}
	if enum, ok := OpcuaNodeIdServicesVariableBaseByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableBase")
		return OpcuaNodeIdServicesVariableBase(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableBase) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableBase) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableBase", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableBase) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableBase_BaseAnalogType_Definition:
		return "BaseAnalogType_Definition"
	case OpcuaNodeIdServicesVariableBase_BaseAnalogType_ValuePrecision:
		return "BaseAnalogType_ValuePrecision"
	case OpcuaNodeIdServicesVariableBase_BaseAnalogType_InstrumentRange:
		return "BaseAnalogType_InstrumentRange"
	case OpcuaNodeIdServicesVariableBase_BaseAnalogType_EURange:
		return "BaseAnalogType_EURange"
	case OpcuaNodeIdServicesVariableBase_BaseAnalogType_EngineeringUnits:
		return "BaseAnalogType_EngineeringUnits"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_EventId:
		return "BaseEventType_EventId"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_EventType:
		return "BaseEventType_EventType"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_SourceNode:
		return "BaseEventType_SourceNode"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_SourceName:
		return "BaseEventType_SourceName"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_Time:
		return "BaseEventType_Time"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_ReceiveTime:
		return "BaseEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_Message:
		return "BaseEventType_Message"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_Severity:
		return "BaseEventType_Severity"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassId:
		return "BaseEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionClassName:
		return "BaseEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassId:
		return "BaseEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_ConditionSubClassName:
		return "BaseEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassId:
		return "BaseModelChangeEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionClassName:
		return "BaseModelChangeEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassId:
		return "BaseModelChangeEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ConditionSubClassName:
		return "BaseModelChangeEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableBase_BaseEventType_LocalTime:
		return "BaseEventType_LocalTime"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventId:
		return "BaseModelChangeEventType_EventId"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_EventType:
		return "BaseModelChangeEventType_EventType"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceNode:
		return "BaseModelChangeEventType_SourceNode"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_SourceName:
		return "BaseModelChangeEventType_SourceName"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Time:
		return "BaseModelChangeEventType_Time"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_ReceiveTime:
		return "BaseModelChangeEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_LocalTime:
		return "BaseModelChangeEventType_LocalTime"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Message:
		return "BaseModelChangeEventType_Message"
	case OpcuaNodeIdServicesVariableBase_BaseModelChangeEventType_Severity:
		return "BaseModelChangeEventType_Severity"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableBase) String() string {
	return e.PLC4XEnumName()
}
