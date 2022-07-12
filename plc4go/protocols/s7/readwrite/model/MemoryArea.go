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

// MemoryArea is an enum
type MemoryArea uint8

type IMemoryArea interface {
	ShortName() string
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	MemoryArea_COUNTERS                 MemoryArea = 0x1C
	MemoryArea_TIMERS                   MemoryArea = 0x1D
	MemoryArea_DIRECT_PERIPHERAL_ACCESS MemoryArea = 0x80
	MemoryArea_INPUTS                   MemoryArea = 0x81
	MemoryArea_OUTPUTS                  MemoryArea = 0x82
	MemoryArea_FLAGS_MARKERS            MemoryArea = 0x83
	MemoryArea_DATA_BLOCKS              MemoryArea = 0x84
	MemoryArea_INSTANCE_DATA_BLOCKS     MemoryArea = 0x85
	MemoryArea_LOCAL_DATA               MemoryArea = 0x86
)

var MemoryAreaValues []MemoryArea

func init() {
	_ = errors.New
	MemoryAreaValues = []MemoryArea{
		MemoryArea_COUNTERS,
		MemoryArea_TIMERS,
		MemoryArea_DIRECT_PERIPHERAL_ACCESS,
		MemoryArea_INPUTS,
		MemoryArea_OUTPUTS,
		MemoryArea_FLAGS_MARKERS,
		MemoryArea_DATA_BLOCKS,
		MemoryArea_INSTANCE_DATA_BLOCKS,
		MemoryArea_LOCAL_DATA,
	}
}

func (e MemoryArea) ShortName() string {
	switch e {
	case 0x1C:
		{ /* '0x1C' */
			return "C"
		}
	case 0x1D:
		{ /* '0x1D' */
			return "T"
		}
	case 0x80:
		{ /* '0x80' */
			return "D"
		}
	case 0x81:
		{ /* '0x81' */
			return "I"
		}
	case 0x82:
		{ /* '0x82' */
			return "Q"
		}
	case 0x83:
		{ /* '0x83' */
			return "M"
		}
	case 0x84:
		{ /* '0x84' */
			return "DB"
		}
	case 0x85:
		{ /* '0x85' */
			return "DBI"
		}
	case 0x86:
		{ /* '0x86' */
			return "LD"
		}
	default:
		{
			return ""
		}
	}
}

func MemoryAreaFirstEnumForFieldShortName(value string) (MemoryArea, error) {
	for _, sizeValue := range MemoryAreaValues {
		if sizeValue.ShortName() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing ShortName not found", value)
}
func MemoryAreaByValue(value uint8) (enum MemoryArea, ok bool) {
	switch value {
	case 0x1C:
		return MemoryArea_COUNTERS, true
	case 0x1D:
		return MemoryArea_TIMERS, true
	case 0x80:
		return MemoryArea_DIRECT_PERIPHERAL_ACCESS, true
	case 0x81:
		return MemoryArea_INPUTS, true
	case 0x82:
		return MemoryArea_OUTPUTS, true
	case 0x83:
		return MemoryArea_FLAGS_MARKERS, true
	case 0x84:
		return MemoryArea_DATA_BLOCKS, true
	case 0x85:
		return MemoryArea_INSTANCE_DATA_BLOCKS, true
	case 0x86:
		return MemoryArea_LOCAL_DATA, true
	}
	return 0, false
}

func MemoryAreaByName(value string) (enum MemoryArea, ok bool) {
	switch value {
	case "COUNTERS":
		return MemoryArea_COUNTERS, true
	case "TIMERS":
		return MemoryArea_TIMERS, true
	case "DIRECT_PERIPHERAL_ACCESS":
		return MemoryArea_DIRECT_PERIPHERAL_ACCESS, true
	case "INPUTS":
		return MemoryArea_INPUTS, true
	case "OUTPUTS":
		return MemoryArea_OUTPUTS, true
	case "FLAGS_MARKERS":
		return MemoryArea_FLAGS_MARKERS, true
	case "DATA_BLOCKS":
		return MemoryArea_DATA_BLOCKS, true
	case "INSTANCE_DATA_BLOCKS":
		return MemoryArea_INSTANCE_DATA_BLOCKS, true
	case "LOCAL_DATA":
		return MemoryArea_LOCAL_DATA, true
	}
	return 0, false
}

func MemoryAreaKnows(value uint8) bool {
	for _, typeValue := range MemoryAreaValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMemoryArea(structType interface{}) MemoryArea {
	castFunc := func(typ interface{}) MemoryArea {
		if sMemoryArea, ok := typ.(MemoryArea); ok {
			return sMemoryArea
		}
		return 0
	}
	return castFunc(structType)
}

func (m MemoryArea) GetLengthInBits() uint16 {
	return 8
}

func (m MemoryArea) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MemoryAreaParse(readBuffer utils.ReadBuffer) (MemoryArea, error) {
	val, err := readBuffer.ReadUint8("MemoryArea", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MemoryArea")
	}
	if enum, ok := MemoryAreaByValue(val); !ok {
		return 0, utils.ParseAssertError{Message: fmt.Sprintf("no value %v found for MemoryArea", val)}
	} else {
		return enum, nil
	}
}

func (e MemoryArea) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MemoryArea", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MemoryArea) PLC4XEnumName() string {
	switch e {
	case MemoryArea_COUNTERS:
		return "COUNTERS"
	case MemoryArea_TIMERS:
		return "TIMERS"
	case MemoryArea_DIRECT_PERIPHERAL_ACCESS:
		return "DIRECT_PERIPHERAL_ACCESS"
	case MemoryArea_INPUTS:
		return "INPUTS"
	case MemoryArea_OUTPUTS:
		return "OUTPUTS"
	case MemoryArea_FLAGS_MARKERS:
		return "FLAGS_MARKERS"
	case MemoryArea_DATA_BLOCKS:
		return "DATA_BLOCKS"
	case MemoryArea_INSTANCE_DATA_BLOCKS:
		return "INSTANCE_DATA_BLOCKS"
	case MemoryArea_LOCAL_DATA:
		return "LOCAL_DATA"
	}
	return ""
}

func (e MemoryArea) String() string {
	return e.PLC4XEnumName()
}
