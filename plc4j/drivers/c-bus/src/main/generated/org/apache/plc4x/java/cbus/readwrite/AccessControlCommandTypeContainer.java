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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum AccessControlCommandTypeContainer {
  AccessControlCommandCloseAccessPoint(
      (short) 0x02,
      (short) 2,
      AccessControlCommandType.CLOSE_ACCESS_POINT,
      AccessControlCategory.SYSTEM_REQUEST),
  AccessControlCommandLockAccessPoint(
      (short) 0x0A,
      (short) 2,
      AccessControlCommandType.LOCK_ACCESS_POINT,
      AccessControlCategory.SYSTEM_REQUEST),
  AccessControlCommandAccessPointLeftOpen(
      (short) 0x12,
      (short) 2,
      AccessControlCommandType.ACCESS_POINT_LEFT_OPEN,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandAccessPointForcedOpen(
      (short) 0x1A,
      (short) 2,
      AccessControlCommandType.ACCESS_POINT_FORCED_OPEN,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandAccessPointClosed(
      (short) 0x22,
      (short) 2,
      AccessControlCommandType.ACCESS_POINT_CLOSED,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandRequestToExit(
      (short) 0x32,
      (short) 2,
      AccessControlCommandType.REQUEST_TO_EXIT,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_0Bytes(
      (short) 0xA0,
      (short) 0,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_1Bytes(
      (short) 0xA1,
      (short) 1,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_2Bytes(
      (short) 0xA2,
      (short) 2,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_3Bytes(
      (short) 0xA3,
      (short) 3,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_4Bytes(
      (short) 0xA4,
      (short) 4,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_5Bytes(
      (short) 0xA5,
      (short) 5,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_6Bytes(
      (short) 0xA6,
      (short) 6,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_7Bytes(
      (short) 0xA7,
      (short) 7,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_8Bytes(
      (short) 0xA8,
      (short) 8,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_9Bytes(
      (short) 0xA9,
      (short) 9,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_10Bytes(
      (short) 0xAA,
      (short) 10,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_11Bytes(
      (short) 0xAB,
      (short) 11,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_12Bytes(
      (short) 0xAC,
      (short) 12,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_13Bytes(
      (short) 0xAD,
      (short) 13,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_14Bytes(
      (short) 0xAE,
      (short) 14,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_15Bytes(
      (short) 0xAF,
      (short) 15,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_16Bytes(
      (short) 0xB0,
      (short) 16,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_17Bytes(
      (short) 0xB1,
      (short) 17,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_18Bytes(
      (short) 0xB2,
      (short) 18,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_19Bytes(
      (short) 0xB3,
      (short) 19,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_20Bytes(
      (short) 0xB4,
      (short) 20,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_21Bytes(
      (short) 0xB5,
      (short) 21,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_22Bytes(
      (short) 0xB6,
      (short) 22,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_23Bytes(
      (short) 0xB7,
      (short) 23,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_24Bytes(
      (short) 0xB8,
      (short) 24,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_25Bytes(
      (short) 0xB9,
      (short) 25,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_26Bytes(
      (short) 0xBA,
      (short) 26,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_27Bytes(
      (short) 0xBB,
      (short) 27,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_28Bytes(
      (short) 0xBC,
      (short) 28,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_29Bytes(
      (short) 0xBD,
      (short) 29,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_30Bytes(
      (short) 0xBE,
      (short) 30,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_31Bytes(
      (short) 0xBF,
      (short) 31,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_0Bytes(
      (short) 0xC0,
      (short) 0,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_1Bytes(
      (short) 0xC1,
      (short) 1,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_2Bytes(
      (short) 0xC2,
      (short) 2,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_3Bytes(
      (short) 0xC3,
      (short) 3,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_4Bytes(
      (short) 0xC4,
      (short) 4,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_5Bytes(
      (short) 0xC5,
      (short) 5,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_6Bytes(
      (short) 0xC6,
      (short) 6,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_7Bytes(
      (short) 0xC7,
      (short) 7,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_8Bytes(
      (short) 0xC8,
      (short) 8,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_9Bytes(
      (short) 0xC9,
      (short) 9,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_10Bytes(
      (short) 0xCA,
      (short) 10,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_11Bytes(
      (short) 0xCB,
      (short) 11,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_12Bytes(
      (short) 0xCC,
      (short) 12,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_13Bytes(
      (short) 0xCD,
      (short) 13,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_14Bytes(
      (short) 0xCE,
      (short) 14,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_15Bytes(
      (short) 0xCF,
      (short) 15,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_16Bytes(
      (short) 0xD0,
      (short) 16,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_17Bytes(
      (short) 0xD1,
      (short) 17,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_18Bytes(
      (short) 0xD2,
      (short) 18,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_19Bytes(
      (short) 0xD3,
      (short) 19,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_20Bytes(
      (short) 0xD4,
      (short) 20,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_21Bytes(
      (short) 0xD5,
      (short) 21,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_22Bytes(
      (short) 0xD6,
      (short) 22,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_23Bytes(
      (short) 0xD7,
      (short) 23,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_24Bytes(
      (short) 0xD8,
      (short) 24,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_25Bytes(
      (short) 0xD9,
      (short) 25,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_26Bytes(
      (short) 0xDA,
      (short) 26,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_27Bytes(
      (short) 0xDB,
      (short) 27,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_28Bytes(
      (short) 0xDC,
      (short) 28,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_29Bytes(
      (short) 0xDD,
      (short) 29,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_30Bytes(
      (short) 0xDE,
      (short) 30,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_31Bytes(
      (short) 0xDF,
      (short) 31,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY);
  private static final Map<Short, AccessControlCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (AccessControlCommandTypeContainer value : AccessControlCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final short numBytes;
  private final AccessControlCommandType commandType;
  private final AccessControlCategory category;

  AccessControlCommandTypeContainer(
      short value,
      short numBytes,
      AccessControlCommandType commandType,
      AccessControlCategory category) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
    this.category = category;
  }

  public short getValue() {
    return value;
  }

  public short getNumBytes() {
    return numBytes;
  }

  public static AccessControlCommandTypeContainer firstEnumForFieldNumBytes(short fieldValue) {
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AccessControlCommandTypeContainer> enumsForFieldNumBytes(short fieldValue) {
    List<AccessControlCommandTypeContainer> _values = new ArrayList<>();
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public AccessControlCommandType getCommandType() {
    return commandType;
  }

  public static AccessControlCommandTypeContainer firstEnumForFieldCommandType(
      AccessControlCommandType fieldValue) {
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AccessControlCommandTypeContainer> enumsForFieldCommandType(
      AccessControlCommandType fieldValue) {
    List<AccessControlCommandTypeContainer> _values = new ArrayList<>();
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public AccessControlCategory getCategory() {
    return category;
  }

  public static AccessControlCommandTypeContainer firstEnumForFieldCategory(
      AccessControlCategory fieldValue) {
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCategory() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AccessControlCommandTypeContainer> enumsForFieldCategory(
      AccessControlCategory fieldValue) {
    List<AccessControlCommandTypeContainer> _values = new ArrayList<>();
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCategory() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static AccessControlCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
