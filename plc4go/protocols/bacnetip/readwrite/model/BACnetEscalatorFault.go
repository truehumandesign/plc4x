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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEscalatorFault is an enum
type BACnetEscalatorFault uint16

type IBACnetEscalatorFault interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetEscalatorFault_CONTROLLER_FAULT           BACnetEscalatorFault = 0
	BACnetEscalatorFault_DRIVE_AND_MOTOR_FAULT      BACnetEscalatorFault = 1
	BACnetEscalatorFault_MECHANICAL_COMPONENT_FAULT BACnetEscalatorFault = 2
	BACnetEscalatorFault_OVERSPEED_FAULT            BACnetEscalatorFault = 3
	BACnetEscalatorFault_POWER_SUPPLY_FAULT         BACnetEscalatorFault = 4
	BACnetEscalatorFault_SAFETY_DEVICE_FAULT        BACnetEscalatorFault = 5
	BACnetEscalatorFault_CONTROLLER_SUPPLY_FAULT    BACnetEscalatorFault = 6
	BACnetEscalatorFault_DRIVE_TEMPERATURE_EXCEEDED BACnetEscalatorFault = 7
	BACnetEscalatorFault_COMB_PLATE_FAULT           BACnetEscalatorFault = 8
	BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE   BACnetEscalatorFault = 0xFFFF
)

var BACnetEscalatorFaultValues []BACnetEscalatorFault

func init() {
	_ = errors.New
	BACnetEscalatorFaultValues = []BACnetEscalatorFault{
		BACnetEscalatorFault_CONTROLLER_FAULT,
		BACnetEscalatorFault_DRIVE_AND_MOTOR_FAULT,
		BACnetEscalatorFault_MECHANICAL_COMPONENT_FAULT,
		BACnetEscalatorFault_OVERSPEED_FAULT,
		BACnetEscalatorFault_POWER_SUPPLY_FAULT,
		BACnetEscalatorFault_SAFETY_DEVICE_FAULT,
		BACnetEscalatorFault_CONTROLLER_SUPPLY_FAULT,
		BACnetEscalatorFault_DRIVE_TEMPERATURE_EXCEEDED,
		BACnetEscalatorFault_COMB_PLATE_FAULT,
		BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEscalatorFaultByValue(value uint16) (enum BACnetEscalatorFault, ok bool) {
	switch value {
	case 0:
		return BACnetEscalatorFault_CONTROLLER_FAULT, true
	case 0xFFFF:
		return BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetEscalatorFault_DRIVE_AND_MOTOR_FAULT, true
	case 2:
		return BACnetEscalatorFault_MECHANICAL_COMPONENT_FAULT, true
	case 3:
		return BACnetEscalatorFault_OVERSPEED_FAULT, true
	case 4:
		return BACnetEscalatorFault_POWER_SUPPLY_FAULT, true
	case 5:
		return BACnetEscalatorFault_SAFETY_DEVICE_FAULT, true
	case 6:
		return BACnetEscalatorFault_CONTROLLER_SUPPLY_FAULT, true
	case 7:
		return BACnetEscalatorFault_DRIVE_TEMPERATURE_EXCEEDED, true
	case 8:
		return BACnetEscalatorFault_COMB_PLATE_FAULT, true
	}
	return 0, false
}

func BACnetEscalatorFaultByName(value string) (enum BACnetEscalatorFault, ok bool) {
	switch value {
	case "CONTROLLER_FAULT":
		return BACnetEscalatorFault_CONTROLLER_FAULT, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE, true
	case "DRIVE_AND_MOTOR_FAULT":
		return BACnetEscalatorFault_DRIVE_AND_MOTOR_FAULT, true
	case "MECHANICAL_COMPONENT_FAULT":
		return BACnetEscalatorFault_MECHANICAL_COMPONENT_FAULT, true
	case "OVERSPEED_FAULT":
		return BACnetEscalatorFault_OVERSPEED_FAULT, true
	case "POWER_SUPPLY_FAULT":
		return BACnetEscalatorFault_POWER_SUPPLY_FAULT, true
	case "SAFETY_DEVICE_FAULT":
		return BACnetEscalatorFault_SAFETY_DEVICE_FAULT, true
	case "CONTROLLER_SUPPLY_FAULT":
		return BACnetEscalatorFault_CONTROLLER_SUPPLY_FAULT, true
	case "DRIVE_TEMPERATURE_EXCEEDED":
		return BACnetEscalatorFault_DRIVE_TEMPERATURE_EXCEEDED, true
	case "COMB_PLATE_FAULT":
		return BACnetEscalatorFault_COMB_PLATE_FAULT, true
	}
	return 0, false
}

func BACnetEscalatorFaultKnows(value uint16) bool {
	for _, typeValue := range BACnetEscalatorFaultValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEscalatorFault(structType interface{}) BACnetEscalatorFault {
	castFunc := func(typ interface{}) BACnetEscalatorFault {
		if sBACnetEscalatorFault, ok := typ.(BACnetEscalatorFault); ok {
			return sBACnetEscalatorFault
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEscalatorFault) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetEscalatorFault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEscalatorFaultParse(readBuffer utils.ReadBuffer) (BACnetEscalatorFault, error) {
	val, err := readBuffer.ReadUint16("BACnetEscalatorFault", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetEscalatorFault")
	}
	if enum, ok := BACnetEscalatorFaultByValue(val); !ok {
		return 0, utils.ParseAssertError{Message: fmt.Sprintf("no value %v found for BACnetEscalatorFault", val)}
	} else {
		return enum, nil
	}
}

func (e BACnetEscalatorFault) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetEscalatorFault", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetEscalatorFault) PLC4XEnumName() string {
	switch e {
	case BACnetEscalatorFault_CONTROLLER_FAULT:
		return "CONTROLLER_FAULT"
	case BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEscalatorFault_DRIVE_AND_MOTOR_FAULT:
		return "DRIVE_AND_MOTOR_FAULT"
	case BACnetEscalatorFault_MECHANICAL_COMPONENT_FAULT:
		return "MECHANICAL_COMPONENT_FAULT"
	case BACnetEscalatorFault_OVERSPEED_FAULT:
		return "OVERSPEED_FAULT"
	case BACnetEscalatorFault_POWER_SUPPLY_FAULT:
		return "POWER_SUPPLY_FAULT"
	case BACnetEscalatorFault_SAFETY_DEVICE_FAULT:
		return "SAFETY_DEVICE_FAULT"
	case BACnetEscalatorFault_CONTROLLER_SUPPLY_FAULT:
		return "CONTROLLER_SUPPLY_FAULT"
	case BACnetEscalatorFault_DRIVE_TEMPERATURE_EXCEEDED:
		return "DRIVE_TEMPERATURE_EXCEEDED"
	case BACnetEscalatorFault_COMB_PLATE_FAULT:
		return "COMB_PLATE_FAULT"
	}
	return ""
}

func (e BACnetEscalatorFault) String() string {
	return e.PLC4XEnumName()
}
