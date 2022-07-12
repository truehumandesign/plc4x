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

// BACnetNodeType is an enum
type BACnetNodeType uint8

type IBACnetNodeType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetNodeType_UNKNOWN        BACnetNodeType = 0x00
	BACnetNodeType_SYSTEM         BACnetNodeType = 0x01
	BACnetNodeType_NETWORK        BACnetNodeType = 0x02
	BACnetNodeType_DEVICE         BACnetNodeType = 0x03
	BACnetNodeType_ORGANIZATIONAL BACnetNodeType = 0x04
	BACnetNodeType_AREA           BACnetNodeType = 0x05
	BACnetNodeType_EQUIPMENT      BACnetNodeType = 0x06
	BACnetNodeType_POINT          BACnetNodeType = 0x07
	BACnetNodeType_COLLECTION     BACnetNodeType = 0x08
	BACnetNodeType_PROPERTY       BACnetNodeType = 0x09
	BACnetNodeType_FUNCTIONAL     BACnetNodeType = 0x0A
	BACnetNodeType_OTHER          BACnetNodeType = 0x0B
	BACnetNodeType_SUBSYSTEM      BACnetNodeType = 0x0C
	BACnetNodeType_BUILDING       BACnetNodeType = 0x0D
	BACnetNodeType_FLOOR          BACnetNodeType = 0x0E
	BACnetNodeType_SECTION        BACnetNodeType = 0x0F
	BACnetNodeType_MODULE         BACnetNodeType = 0x10
	BACnetNodeType_TREE           BACnetNodeType = 0x11
	BACnetNodeType_MEMBER         BACnetNodeType = 0x12
	BACnetNodeType_PROTOCOL       BACnetNodeType = 0x13
	BACnetNodeType_ROOM           BACnetNodeType = 0x14
	BACnetNodeType_ZONE           BACnetNodeType = 0x15
)

var BACnetNodeTypeValues []BACnetNodeType

func init() {
	_ = errors.New
	BACnetNodeTypeValues = []BACnetNodeType{
		BACnetNodeType_UNKNOWN,
		BACnetNodeType_SYSTEM,
		BACnetNodeType_NETWORK,
		BACnetNodeType_DEVICE,
		BACnetNodeType_ORGANIZATIONAL,
		BACnetNodeType_AREA,
		BACnetNodeType_EQUIPMENT,
		BACnetNodeType_POINT,
		BACnetNodeType_COLLECTION,
		BACnetNodeType_PROPERTY,
		BACnetNodeType_FUNCTIONAL,
		BACnetNodeType_OTHER,
		BACnetNodeType_SUBSYSTEM,
		BACnetNodeType_BUILDING,
		BACnetNodeType_FLOOR,
		BACnetNodeType_SECTION,
		BACnetNodeType_MODULE,
		BACnetNodeType_TREE,
		BACnetNodeType_MEMBER,
		BACnetNodeType_PROTOCOL,
		BACnetNodeType_ROOM,
		BACnetNodeType_ZONE,
	}
}

func BACnetNodeTypeByValue(value uint8) (enum BACnetNodeType, ok bool) {
	switch value {
	case 0x00:
		return BACnetNodeType_UNKNOWN, true
	case 0x01:
		return BACnetNodeType_SYSTEM, true
	case 0x02:
		return BACnetNodeType_NETWORK, true
	case 0x03:
		return BACnetNodeType_DEVICE, true
	case 0x04:
		return BACnetNodeType_ORGANIZATIONAL, true
	case 0x05:
		return BACnetNodeType_AREA, true
	case 0x06:
		return BACnetNodeType_EQUIPMENT, true
	case 0x07:
		return BACnetNodeType_POINT, true
	case 0x08:
		return BACnetNodeType_COLLECTION, true
	case 0x09:
		return BACnetNodeType_PROPERTY, true
	case 0x0A:
		return BACnetNodeType_FUNCTIONAL, true
	case 0x0B:
		return BACnetNodeType_OTHER, true
	case 0x0C:
		return BACnetNodeType_SUBSYSTEM, true
	case 0x0D:
		return BACnetNodeType_BUILDING, true
	case 0x0E:
		return BACnetNodeType_FLOOR, true
	case 0x0F:
		return BACnetNodeType_SECTION, true
	case 0x10:
		return BACnetNodeType_MODULE, true
	case 0x11:
		return BACnetNodeType_TREE, true
	case 0x12:
		return BACnetNodeType_MEMBER, true
	case 0x13:
		return BACnetNodeType_PROTOCOL, true
	case 0x14:
		return BACnetNodeType_ROOM, true
	case 0x15:
		return BACnetNodeType_ZONE, true
	}
	return 0, false
}

func BACnetNodeTypeByName(value string) (enum BACnetNodeType, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetNodeType_UNKNOWN, true
	case "SYSTEM":
		return BACnetNodeType_SYSTEM, true
	case "NETWORK":
		return BACnetNodeType_NETWORK, true
	case "DEVICE":
		return BACnetNodeType_DEVICE, true
	case "ORGANIZATIONAL":
		return BACnetNodeType_ORGANIZATIONAL, true
	case "AREA":
		return BACnetNodeType_AREA, true
	case "EQUIPMENT":
		return BACnetNodeType_EQUIPMENT, true
	case "POINT":
		return BACnetNodeType_POINT, true
	case "COLLECTION":
		return BACnetNodeType_COLLECTION, true
	case "PROPERTY":
		return BACnetNodeType_PROPERTY, true
	case "FUNCTIONAL":
		return BACnetNodeType_FUNCTIONAL, true
	case "OTHER":
		return BACnetNodeType_OTHER, true
	case "SUBSYSTEM":
		return BACnetNodeType_SUBSYSTEM, true
	case "BUILDING":
		return BACnetNodeType_BUILDING, true
	case "FLOOR":
		return BACnetNodeType_FLOOR, true
	case "SECTION":
		return BACnetNodeType_SECTION, true
	case "MODULE":
		return BACnetNodeType_MODULE, true
	case "TREE":
		return BACnetNodeType_TREE, true
	case "MEMBER":
		return BACnetNodeType_MEMBER, true
	case "PROTOCOL":
		return BACnetNodeType_PROTOCOL, true
	case "ROOM":
		return BACnetNodeType_ROOM, true
	case "ZONE":
		return BACnetNodeType_ZONE, true
	}
	return 0, false
}

func BACnetNodeTypeKnows(value uint8) bool {
	for _, typeValue := range BACnetNodeTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetNodeType(structType interface{}) BACnetNodeType {
	castFunc := func(typ interface{}) BACnetNodeType {
		if sBACnetNodeType, ok := typ.(BACnetNodeType); ok {
			return sBACnetNodeType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetNodeType) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetNodeType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNodeTypeParse(readBuffer utils.ReadBuffer) (BACnetNodeType, error) {
	val, err := readBuffer.ReadUint8("BACnetNodeType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetNodeType")
	}
	if enum, ok := BACnetNodeTypeByValue(val); !ok {
		return 0, utils.ParseAssertError{Message: fmt.Sprintf("no value %v found for BACnetNodeType", val)}
	} else {
		return enum, nil
	}
}

func (e BACnetNodeType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetNodeType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetNodeType) PLC4XEnumName() string {
	switch e {
	case BACnetNodeType_UNKNOWN:
		return "UNKNOWN"
	case BACnetNodeType_SYSTEM:
		return "SYSTEM"
	case BACnetNodeType_NETWORK:
		return "NETWORK"
	case BACnetNodeType_DEVICE:
		return "DEVICE"
	case BACnetNodeType_ORGANIZATIONAL:
		return "ORGANIZATIONAL"
	case BACnetNodeType_AREA:
		return "AREA"
	case BACnetNodeType_EQUIPMENT:
		return "EQUIPMENT"
	case BACnetNodeType_POINT:
		return "POINT"
	case BACnetNodeType_COLLECTION:
		return "COLLECTION"
	case BACnetNodeType_PROPERTY:
		return "PROPERTY"
	case BACnetNodeType_FUNCTIONAL:
		return "FUNCTIONAL"
	case BACnetNodeType_OTHER:
		return "OTHER"
	case BACnetNodeType_SUBSYSTEM:
		return "SUBSYSTEM"
	case BACnetNodeType_BUILDING:
		return "BUILDING"
	case BACnetNodeType_FLOOR:
		return "FLOOR"
	case BACnetNodeType_SECTION:
		return "SECTION"
	case BACnetNodeType_MODULE:
		return "MODULE"
	case BACnetNodeType_TREE:
		return "TREE"
	case BACnetNodeType_MEMBER:
		return "MEMBER"
	case BACnetNodeType_PROTOCOL:
		return "PROTOCOL"
	case BACnetNodeType_ROOM:
		return "ROOM"
	case BACnetNodeType_ZONE:
		return "ZONE"
	}
	return ""
}

func (e BACnetNodeType) String() string {
	return e.PLC4XEnumName()
}
