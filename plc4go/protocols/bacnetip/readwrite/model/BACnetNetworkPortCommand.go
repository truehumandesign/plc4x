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

// BACnetNetworkPortCommand is an enum
type BACnetNetworkPortCommand uint8

type IBACnetNetworkPortCommand interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetNetworkPortCommand_IDLE                     BACnetNetworkPortCommand = 0
	BACnetNetworkPortCommand_DISCARD_CHANGES          BACnetNetworkPortCommand = 1
	BACnetNetworkPortCommand_RENEW_FD_REGISTRATION    BACnetNetworkPortCommand = 2
	BACnetNetworkPortCommand_RESTART_SLAVE_DISCOVERY  BACnetNetworkPortCommand = 3
	BACnetNetworkPortCommand_RENEW_DHCP               BACnetNetworkPortCommand = 4
	BACnetNetworkPortCommand_RESTART_AUTONEGOTIATION  BACnetNetworkPortCommand = 5
	BACnetNetworkPortCommand_DISCONNECT               BACnetNetworkPortCommand = 6
	BACnetNetworkPortCommand_RESTART_PORT             BACnetNetworkPortCommand = 7
	BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE BACnetNetworkPortCommand = 0xFF
)

var BACnetNetworkPortCommandValues []BACnetNetworkPortCommand

func init() {
	_ = errors.New
	BACnetNetworkPortCommandValues = []BACnetNetworkPortCommand{
		BACnetNetworkPortCommand_IDLE,
		BACnetNetworkPortCommand_DISCARD_CHANGES,
		BACnetNetworkPortCommand_RENEW_FD_REGISTRATION,
		BACnetNetworkPortCommand_RESTART_SLAVE_DISCOVERY,
		BACnetNetworkPortCommand_RENEW_DHCP,
		BACnetNetworkPortCommand_RESTART_AUTONEGOTIATION,
		BACnetNetworkPortCommand_DISCONNECT,
		BACnetNetworkPortCommand_RESTART_PORT,
		BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetNetworkPortCommandByValue(value uint8) (enum BACnetNetworkPortCommand, ok bool) {
	switch value {
	case 0:
		return BACnetNetworkPortCommand_IDLE, true
	case 0xFF:
		return BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetNetworkPortCommand_DISCARD_CHANGES, true
	case 2:
		return BACnetNetworkPortCommand_RENEW_FD_REGISTRATION, true
	case 3:
		return BACnetNetworkPortCommand_RESTART_SLAVE_DISCOVERY, true
	case 4:
		return BACnetNetworkPortCommand_RENEW_DHCP, true
	case 5:
		return BACnetNetworkPortCommand_RESTART_AUTONEGOTIATION, true
	case 6:
		return BACnetNetworkPortCommand_DISCONNECT, true
	case 7:
		return BACnetNetworkPortCommand_RESTART_PORT, true
	}
	return 0, false
}

func BACnetNetworkPortCommandByName(value string) (enum BACnetNetworkPortCommand, ok bool) {
	switch value {
	case "IDLE":
		return BACnetNetworkPortCommand_IDLE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE, true
	case "DISCARD_CHANGES":
		return BACnetNetworkPortCommand_DISCARD_CHANGES, true
	case "RENEW_FD_REGISTRATION":
		return BACnetNetworkPortCommand_RENEW_FD_REGISTRATION, true
	case "RESTART_SLAVE_DISCOVERY":
		return BACnetNetworkPortCommand_RESTART_SLAVE_DISCOVERY, true
	case "RENEW_DHCP":
		return BACnetNetworkPortCommand_RENEW_DHCP, true
	case "RESTART_AUTONEGOTIATION":
		return BACnetNetworkPortCommand_RESTART_AUTONEGOTIATION, true
	case "DISCONNECT":
		return BACnetNetworkPortCommand_DISCONNECT, true
	case "RESTART_PORT":
		return BACnetNetworkPortCommand_RESTART_PORT, true
	}
	return 0, false
}

func BACnetNetworkPortCommandKnows(value uint8) bool {
	for _, typeValue := range BACnetNetworkPortCommandValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetNetworkPortCommand(structType interface{}) BACnetNetworkPortCommand {
	castFunc := func(typ interface{}) BACnetNetworkPortCommand {
		if sBACnetNetworkPortCommand, ok := typ.(BACnetNetworkPortCommand); ok {
			return sBACnetNetworkPortCommand
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetNetworkPortCommand) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetNetworkPortCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNetworkPortCommandParse(readBuffer utils.ReadBuffer) (BACnetNetworkPortCommand, error) {
	val, err := readBuffer.ReadUint8("BACnetNetworkPortCommand", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetNetworkPortCommand")
	}
	if enum, ok := BACnetNetworkPortCommandByValue(val); !ok {
		return 0, utils.ParseAssertError{Message: fmt.Sprintf("no value %v found for BACnetNetworkPortCommand", val)}
	} else {
		return enum, nil
	}
}

func (e BACnetNetworkPortCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetNetworkPortCommand", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetNetworkPortCommand) PLC4XEnumName() string {
	switch e {
	case BACnetNetworkPortCommand_IDLE:
		return "IDLE"
	case BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetNetworkPortCommand_DISCARD_CHANGES:
		return "DISCARD_CHANGES"
	case BACnetNetworkPortCommand_RENEW_FD_REGISTRATION:
		return "RENEW_FD_REGISTRATION"
	case BACnetNetworkPortCommand_RESTART_SLAVE_DISCOVERY:
		return "RESTART_SLAVE_DISCOVERY"
	case BACnetNetworkPortCommand_RENEW_DHCP:
		return "RENEW_DHCP"
	case BACnetNetworkPortCommand_RESTART_AUTONEGOTIATION:
		return "RESTART_AUTONEGOTIATION"
	case BACnetNetworkPortCommand_DISCONNECT:
		return "DISCONNECT"
	case BACnetNetworkPortCommand_RESTART_PORT:
		return "RESTART_PORT"
	}
	return ""
}

func (e BACnetNetworkPortCommand) String() string {
	return e.PLC4XEnumName()
}
