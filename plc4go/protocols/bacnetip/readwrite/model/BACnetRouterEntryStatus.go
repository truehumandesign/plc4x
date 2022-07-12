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

// BACnetRouterEntryStatus is an enum
type BACnetRouterEntryStatus uint8

type IBACnetRouterEntryStatus interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetRouterEntryStatus_AVAILABLE    BACnetRouterEntryStatus = 0
	BACnetRouterEntryStatus_BUSY         BACnetRouterEntryStatus = 1
	BACnetRouterEntryStatus_DISCONNECTED BACnetRouterEntryStatus = 2
)

var BACnetRouterEntryStatusValues []BACnetRouterEntryStatus

func init() {
	_ = errors.New
	BACnetRouterEntryStatusValues = []BACnetRouterEntryStatus{
		BACnetRouterEntryStatus_AVAILABLE,
		BACnetRouterEntryStatus_BUSY,
		BACnetRouterEntryStatus_DISCONNECTED,
	}
}

func BACnetRouterEntryStatusByValue(value uint8) (enum BACnetRouterEntryStatus, ok bool) {
	switch value {
	case 0:
		return BACnetRouterEntryStatus_AVAILABLE, true
	case 1:
		return BACnetRouterEntryStatus_BUSY, true
	case 2:
		return BACnetRouterEntryStatus_DISCONNECTED, true
	}
	return 0, false
}

func BACnetRouterEntryStatusByName(value string) (enum BACnetRouterEntryStatus, ok bool) {
	switch value {
	case "AVAILABLE":
		return BACnetRouterEntryStatus_AVAILABLE, true
	case "BUSY":
		return BACnetRouterEntryStatus_BUSY, true
	case "DISCONNECTED":
		return BACnetRouterEntryStatus_DISCONNECTED, true
	}
	return 0, false
}

func BACnetRouterEntryStatusKnows(value uint8) bool {
	for _, typeValue := range BACnetRouterEntryStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetRouterEntryStatus(structType interface{}) BACnetRouterEntryStatus {
	castFunc := func(typ interface{}) BACnetRouterEntryStatus {
		if sBACnetRouterEntryStatus, ok := typ.(BACnetRouterEntryStatus); ok {
			return sBACnetRouterEntryStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetRouterEntryStatus) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetRouterEntryStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRouterEntryStatusParse(readBuffer utils.ReadBuffer) (BACnetRouterEntryStatus, error) {
	val, err := readBuffer.ReadUint8("BACnetRouterEntryStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetRouterEntryStatus")
	}
	if enum, ok := BACnetRouterEntryStatusByValue(val); !ok {
		return 0, utils.ParseAssertError{Message: fmt.Sprintf("no value %v found for BACnetRouterEntryStatus", val)}
	} else {
		return enum, nil
	}
}

func (e BACnetRouterEntryStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetRouterEntryStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetRouterEntryStatus) PLC4XEnumName() string {
	switch e {
	case BACnetRouterEntryStatus_AVAILABLE:
		return "AVAILABLE"
	case BACnetRouterEntryStatus_BUSY:
		return "BUSY"
	case BACnetRouterEntryStatus_DISCONNECTED:
		return "DISCONNECTED"
	}
	return ""
}

func (e BACnetRouterEntryStatus) String() string {
	return e.PLC4XEnumName()
}
