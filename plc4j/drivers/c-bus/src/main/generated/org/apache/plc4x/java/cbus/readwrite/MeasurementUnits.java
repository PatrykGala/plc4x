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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum MeasurementUnits {
  CELSIUS((short) 0x00),
  AMPS((short) 0x01),
  ANGLE_DEGREES((short) 0x02),
  COULOMB((short) 0x03),
  BOOLEANLOGIC((short) 0x04),
  FARADS((short) 0x05),
  HENRYS((short) 0x06),
  HERTZ((short) 0x07),
  JOULES((short) 0x08),
  KATAL((short) 0x09),
  KG_PER_M3((short) 0x0A),
  KILOGRAMS((short) 0x0B),
  LITRES((short) 0x0C),
  LITRES_PER_HOUR((short) 0x0D),
  LITRES_PER_MINUTE((short) 0x0E),
  LITRES_PER_SECOND((short) 0x0F),
  LUX((short) 0x10),
  METRES((short) 0x11),
  METRES_PER_MINUTE((short) 0x12),
  METRES_PER_SECOND((short) 0x13),
  METRES_PER_S_SQUARED((short) 0x14),
  MOLE((short) 0x15),
  NEWTON_METRE((short) 0x16),
  NEWTONS((short) 0x17),
  OHMS((short) 0x18),
  PASCAL((short) 0x19),
  PERCENT((short) 0x1A),
  DECIBELS((short) 0x1B),
  PPM((short) 0x1C),
  RPM((short) 0x1D),
  SECOND((short) 0x1E),
  MINUTES((short) 0x1F),
  HOURS((short) 0x20),
  SIEVERTS((short) 0x21),
  STERADIAN((short) 0x22),
  TESLA((short) 0x23),
  VOLTS((short) 0x24),
  WATT_HOURS((short) 0x25),
  WATTS((short) 0x26),
  WEBERS((short) 0x27),
  NO_UNITS((short) 0xFE),
  CUSTOM((short) 0xFF);
  private static final Map<Short, MeasurementUnits> map;

  static {
    map = new HashMap<>();
    for (MeasurementUnits value : MeasurementUnits.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private short value;

  MeasurementUnits(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static MeasurementUnits enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}