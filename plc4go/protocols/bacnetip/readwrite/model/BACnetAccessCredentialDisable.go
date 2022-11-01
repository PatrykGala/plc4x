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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAccessCredentialDisable is an enum
type BACnetAccessCredentialDisable uint16

type IBACnetAccessCredentialDisable interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetAccessCredentialDisable_NONE                     BACnetAccessCredentialDisable = 0
	BACnetAccessCredentialDisable_DISABLE                  BACnetAccessCredentialDisable = 1
	BACnetAccessCredentialDisable_DISABLE_MANUAL           BACnetAccessCredentialDisable = 2
	BACnetAccessCredentialDisable_DISABLE_LOCKOUT          BACnetAccessCredentialDisable = 3
	BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE BACnetAccessCredentialDisable = 0xFFFF
)

var BACnetAccessCredentialDisableValues []BACnetAccessCredentialDisable

func init() {
	_ = errors.New
	BACnetAccessCredentialDisableValues = []BACnetAccessCredentialDisable{
		BACnetAccessCredentialDisable_NONE,
		BACnetAccessCredentialDisable_DISABLE,
		BACnetAccessCredentialDisable_DISABLE_MANUAL,
		BACnetAccessCredentialDisable_DISABLE_LOCKOUT,
		BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessCredentialDisableByValue(value uint16) (enum BACnetAccessCredentialDisable, ok bool) {
	switch value {
	case 0:
		return BACnetAccessCredentialDisable_NONE, true
	case 0xFFFF:
		return BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetAccessCredentialDisable_DISABLE, true
	case 2:
		return BACnetAccessCredentialDisable_DISABLE_MANUAL, true
	case 3:
		return BACnetAccessCredentialDisable_DISABLE_LOCKOUT, true
	}
	return 0, false
}

func BACnetAccessCredentialDisableByName(value string) (enum BACnetAccessCredentialDisable, ok bool) {
	switch value {
	case "NONE":
		return BACnetAccessCredentialDisable_NONE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE, true
	case "DISABLE":
		return BACnetAccessCredentialDisable_DISABLE, true
	case "DISABLE_MANUAL":
		return BACnetAccessCredentialDisable_DISABLE_MANUAL, true
	case "DISABLE_LOCKOUT":
		return BACnetAccessCredentialDisable_DISABLE_LOCKOUT, true
	}
	return 0, false
}

func BACnetAccessCredentialDisableKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessCredentialDisableValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessCredentialDisable(structType interface{}) BACnetAccessCredentialDisable {
	castFunc := func(typ interface{}) BACnetAccessCredentialDisable {
		if sBACnetAccessCredentialDisable, ok := typ.(BACnetAccessCredentialDisable); ok {
			return sBACnetAccessCredentialDisable
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessCredentialDisable) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetAccessCredentialDisable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessCredentialDisableParse(readBuffer utils.ReadBuffer) (BACnetAccessCredentialDisable, error) {
	val, err := readBuffer.ReadUint16("BACnetAccessCredentialDisable", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessCredentialDisable")
	}
	if enum, ok := BACnetAccessCredentialDisableByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetAccessCredentialDisable(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessCredentialDisable) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetAccessCredentialDisable", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessCredentialDisable) PLC4XEnumName() string {
	switch e {
	case BACnetAccessCredentialDisable_NONE:
		return "NONE"
	case BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessCredentialDisable_DISABLE:
		return "DISABLE"
	case BACnetAccessCredentialDisable_DISABLE_MANUAL:
		return "DISABLE_MANUAL"
	case BACnetAccessCredentialDisable_DISABLE_LOCKOUT:
		return "DISABLE_LOCKOUT"
	}
	return ""
}

func (e BACnetAccessCredentialDisable) String() string {
	return e.PLC4XEnumName()
}