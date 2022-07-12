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

// BACnetAction is an enum
type BACnetAction uint8

type IBACnetAction interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetAction_DIRECT  BACnetAction = 0
	BACnetAction_REVERSE BACnetAction = 1
)

var BACnetActionValues []BACnetAction

func init() {
	_ = errors.New
	BACnetActionValues = []BACnetAction{
		BACnetAction_DIRECT,
		BACnetAction_REVERSE,
	}
}

func BACnetActionByValue(value uint8) (enum BACnetAction, ok bool) {
	switch value {
	case 0:
		return BACnetAction_DIRECT, true
	case 1:
		return BACnetAction_REVERSE, true
	}
	return 0, false
}

func BACnetActionByName(value string) (enum BACnetAction, ok bool) {
	switch value {
	case "DIRECT":
		return BACnetAction_DIRECT, true
	case "REVERSE":
		return BACnetAction_REVERSE, true
	}
	return 0, false
}

func BACnetActionKnows(value uint8) bool {
	for _, typeValue := range BACnetActionValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAction(structType interface{}) BACnetAction {
	castFunc := func(typ interface{}) BACnetAction {
		if sBACnetAction, ok := typ.(BACnetAction); ok {
			return sBACnetAction
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAction) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetAction) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetActionParse(readBuffer utils.ReadBuffer) (BACnetAction, error) {
	val, err := readBuffer.ReadUint8("BACnetAction", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAction")
	}
	if enum, ok := BACnetActionByValue(val); !ok {
		return 0, utils.ParseAssertError{Message: fmt.Sprintf("no value %v found for BACnetAction", val)}
	} else {
		return enum, nil
	}
}

func (e BACnetAction) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetAction", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAction) PLC4XEnumName() string {
	switch e {
	case BACnetAction_DIRECT:
		return "DIRECT"
	case BACnetAction_REVERSE:
		return "REVERSE"
	}
	return ""
}

func (e BACnetAction) String() string {
	return e.PLC4XEnumName()
}
