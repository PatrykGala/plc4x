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

// BACnetLightingOperation is an enum
type BACnetLightingOperation uint16

type IBACnetLightingOperation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetLightingOperation_NONE                     BACnetLightingOperation = 0
	BACnetLightingOperation_FADE_TO                  BACnetLightingOperation = 1
	BACnetLightingOperation_RAMP_TO                  BACnetLightingOperation = 2
	BACnetLightingOperation_STEP_UP                  BACnetLightingOperation = 3
	BACnetLightingOperation_STEP_DOWN                BACnetLightingOperation = 4
	BACnetLightingOperation_STEP_ON                  BACnetLightingOperation = 5
	BACnetLightingOperation_STEP_OFF                 BACnetLightingOperation = 6
	BACnetLightingOperation_WARN                     BACnetLightingOperation = 7
	BACnetLightingOperation_WARN_OFF                 BACnetLightingOperation = 8
	BACnetLightingOperation_WARN_RELINQUISH          BACnetLightingOperation = 9
	BACnetLightingOperation_STOP                     BACnetLightingOperation = 10
	BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE BACnetLightingOperation = 0xFFFF
)

var BACnetLightingOperationValues []BACnetLightingOperation

func init() {
	_ = errors.New
	BACnetLightingOperationValues = []BACnetLightingOperation{
		BACnetLightingOperation_NONE,
		BACnetLightingOperation_FADE_TO,
		BACnetLightingOperation_RAMP_TO,
		BACnetLightingOperation_STEP_UP,
		BACnetLightingOperation_STEP_DOWN,
		BACnetLightingOperation_STEP_ON,
		BACnetLightingOperation_STEP_OFF,
		BACnetLightingOperation_WARN,
		BACnetLightingOperation_WARN_OFF,
		BACnetLightingOperation_WARN_RELINQUISH,
		BACnetLightingOperation_STOP,
		BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLightingOperationByValue(value uint16) (enum BACnetLightingOperation, ok bool) {
	switch value {
	case 0:
		return BACnetLightingOperation_NONE, true
	case 0xFFFF:
		return BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetLightingOperation_FADE_TO, true
	case 10:
		return BACnetLightingOperation_STOP, true
	case 2:
		return BACnetLightingOperation_RAMP_TO, true
	case 3:
		return BACnetLightingOperation_STEP_UP, true
	case 4:
		return BACnetLightingOperation_STEP_DOWN, true
	case 5:
		return BACnetLightingOperation_STEP_ON, true
	case 6:
		return BACnetLightingOperation_STEP_OFF, true
	case 7:
		return BACnetLightingOperation_WARN, true
	case 8:
		return BACnetLightingOperation_WARN_OFF, true
	case 9:
		return BACnetLightingOperation_WARN_RELINQUISH, true
	}
	return 0, false
}

func BACnetLightingOperationByName(value string) (enum BACnetLightingOperation, ok bool) {
	switch value {
	case "NONE":
		return BACnetLightingOperation_NONE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE, true
	case "FADE_TO":
		return BACnetLightingOperation_FADE_TO, true
	case "STOP":
		return BACnetLightingOperation_STOP, true
	case "RAMP_TO":
		return BACnetLightingOperation_RAMP_TO, true
	case "STEP_UP":
		return BACnetLightingOperation_STEP_UP, true
	case "STEP_DOWN":
		return BACnetLightingOperation_STEP_DOWN, true
	case "STEP_ON":
		return BACnetLightingOperation_STEP_ON, true
	case "STEP_OFF":
		return BACnetLightingOperation_STEP_OFF, true
	case "WARN":
		return BACnetLightingOperation_WARN, true
	case "WARN_OFF":
		return BACnetLightingOperation_WARN_OFF, true
	case "WARN_RELINQUISH":
		return BACnetLightingOperation_WARN_RELINQUISH, true
	}
	return 0, false
}

func BACnetLightingOperationKnows(value uint16) bool {
	for _, typeValue := range BACnetLightingOperationValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLightingOperation(structType any) BACnetLightingOperation {
	castFunc := func(typ any) BACnetLightingOperation {
		if sBACnetLightingOperation, ok := typ.(BACnetLightingOperation); ok {
			return sBACnetLightingOperation
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLightingOperation) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetLightingOperation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingOperationParse(ctx context.Context, theBytes []byte) (BACnetLightingOperation, error) {
	return BACnetLightingOperationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLightingOperationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingOperation, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetLightingOperation", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLightingOperation")
	}
	if enum, ok := BACnetLightingOperationByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetLightingOperation")
		return BACnetLightingOperation(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLightingOperation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLightingOperation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetLightingOperation", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLightingOperation) PLC4XEnumName() string {
	switch e {
	case BACnetLightingOperation_NONE:
		return "NONE"
	case BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLightingOperation_FADE_TO:
		return "FADE_TO"
	case BACnetLightingOperation_STOP:
		return "STOP"
	case BACnetLightingOperation_RAMP_TO:
		return "RAMP_TO"
	case BACnetLightingOperation_STEP_UP:
		return "STEP_UP"
	case BACnetLightingOperation_STEP_DOWN:
		return "STEP_DOWN"
	case BACnetLightingOperation_STEP_ON:
		return "STEP_ON"
	case BACnetLightingOperation_STEP_OFF:
		return "STEP_OFF"
	case BACnetLightingOperation_WARN:
		return "WARN"
	case BACnetLightingOperation_WARN_OFF:
		return "WARN_OFF"
	case BACnetLightingOperation_WARN_RELINQUISH:
		return "WARN_RELINQUISH"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetLightingOperation) String() string {
	return e.PLC4XEnumName()
}
