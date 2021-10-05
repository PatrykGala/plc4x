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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/values"
	api "github.com/apache/plc4x/plc4go/pkg/plc4go/values"
	"github.com/pkg/errors"
	"time"
)

// Code generated by code-generation. DO NOT EDIT.

func DataItemParse(readBuffer utils.ReadBuffer, dataFormatName string, stringLength int32) (api.PlcValue, error) {
	readBuffer.PullContext("DataItem")
	switch {
	case dataFormatName == "IEC61131_BOOL": // BOOL

		// Reserved Field (Just skip the bytes)
		if _, _err := readBuffer.ReadUint8("reserved", 7); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadBit("value")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBOOL(value), nil
	case dataFormatName == "IEC61131_BYTE": // BitString

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBitString(value), nil
	case dataFormatName == "IEC61131_WORD": // BitString

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBitString(value), nil
	case dataFormatName == "IEC61131_DWORD": // BitString

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBitString(value), nil
	case dataFormatName == "IEC61131_SINT": // SINT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSINT(value), nil
	case dataFormatName == "IEC61131_USINT": // USINT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUSINT(value), nil
	case dataFormatName == "IEC61131_INT": // INT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcINT(value), nil
	case dataFormatName == "IEC61131_UINT": // UINT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUINT(value), nil
	case dataFormatName == "IEC61131_DINT": // DINT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDINT(value), nil
	case dataFormatName == "IEC61131_UDINT": // UDINT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUDINT(value), nil
	case dataFormatName == "IEC61131_LINT": // LINT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLINT(value), nil
	case dataFormatName == "IEC61131_ULINT": // ULINT

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcULINT(value), nil
	case dataFormatName == "IEC61131_REAL": // REAL

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat32("value", true, 8, 23)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcREAL(value), nil
	case dataFormatName == "IEC61131_LREAL": // LREAL

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat64("value", true, 11, 52)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLREAL(value), nil
	case dataFormatName == "IEC61131_CHAR": // STRING

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32((8)))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataFormatName == "IEC61131_WCHAR": // STRING

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32((16)))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataFormatName == "IEC61131_STRING": // STRING

		// Manual Field (value)
		value, _valueErr := StaticHelperParseAmsString(readBuffer, stringLength, "TF-")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataFormatName == "IEC61131_WSTRING": // STRING

		// Manual Field (value)
		value, _valueErr := StaticHelperParseAmsString(readBuffer, stringLength, "TF-")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataFormatName == "IEC61131_TIME": // TIME

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcTIME(value), nil
	case dataFormatName == "IEC61131_LTIME": // LTIME

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLTIME(value), nil
	case dataFormatName == "IEC61131_DATE": // DATE

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDATE(value), nil
	case dataFormatName == "IEC61131_TIME_OF_DAY": // TIME_OF_DAY

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcTIME_OF_DAY(value), nil
	case dataFormatName == "IEC61131_DATE_AND_TIME": // DATE_AND_TIME

		// Simple Field (secondsSinceEpoch)
		secondsSinceEpoch, _secondsSinceEpochErr := readBuffer.ReadUint32("secondsSinceEpoch", 32)
		if _secondsSinceEpochErr != nil {
			return nil, errors.Wrap(_secondsSinceEpochErr, "Error parsing 'secondsSinceEpoch' field")
		}
		value := time.Unix(int64(secondsSinceEpoch), 0)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDATE_AND_TIME(value), nil
	}
	// TODO: add more info which type it is actually
	return nil, errors.New("unsupported type")
}

func DataItemSerialize(writeBuffer utils.WriteBuffer, value api.PlcValue, dataFormatName string, stringLength int32) error {
	writeBuffer.PushContext("DataItem")
	switch {
	case dataFormatName == "IEC61131_BOOL": // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := writeBuffer.WriteUint8("reserved", 7, uint8(0x00)); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := writeBuffer.WriteBit("value", value.GetBool()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_BYTE": // BitString

		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_WORD": // BitString

		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, value.GetUint16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_DWORD": // BitString

		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_SINT": // SINT

		// Simple Field (value)
		if _err := writeBuffer.WriteInt8("value", 8, value.GetInt8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_USINT": // USINT

		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_INT": // INT

		// Simple Field (value)
		if _err := writeBuffer.WriteInt16("value", 16, value.GetInt16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_UINT": // UINT

		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, value.GetUint16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_DINT": // DINT

		// Simple Field (value)
		if _err := writeBuffer.WriteInt32("value", 32, value.GetInt32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_UDINT": // UDINT

		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_LINT": // LINT

		// Simple Field (value)
		if _err := writeBuffer.WriteInt64("value", 64, value.GetInt64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_ULINT": // ULINT

		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, value.GetUint64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_REAL": // REAL

		// Simple Field (value)
		if _err := writeBuffer.WriteFloat32("value", 32, value.GetFloat32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_LREAL": // LREAL

		// Simple Field (value)
		if _err := writeBuffer.WriteFloat64("value", 64, value.GetFloat64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_CHAR": // STRING

		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint8((8)), "TF-", value.GetString()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_WCHAR": // STRING

		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint8((16)), "TF-", value.GetString()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_STRING": // STRING

		// Manual Field (value)
		_valueErr := StaticHelperSerializeAmsString(writeBuffer, value, stringLength, "TF-")
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_WSTRING": // STRING

		// Manual Field (value)
		_valueErr := StaticHelperSerializeAmsString(writeBuffer, value, stringLength, "TF-")
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_TIME": // TIME

		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_LTIME": // LTIME

		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, value.GetUint64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_DATE": // DATE

		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_TIME_OF_DAY": // TIME_OF_DAY

		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataFormatName == "IEC61131_DATE_AND_TIME": // DATE_AND_TIME

		// Simple Field (secondsSinceEpoch)
		if _err := writeBuffer.WriteUint32("secondsSinceEpoch", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'secondsSinceEpoch' field")
		}
	default:
		// TODO: add more info which type it is actually
		return errors.New("unsupported type")
	}
	writeBuffer.PopContext("DataItem")
	return nil
}
