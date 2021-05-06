//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type COTPPacketData struct {
	Eot     bool
	TpduRef uint8
	Parent  *COTPPacket
}

// The corresponding interface
type ICOTPPacketData interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPPacketData) TpduCode() uint8 {
	return 0xF0
}

func (m *COTPPacketData) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.Parent.Parameters = parameters
	m.Parent.Payload = payload
}

func NewCOTPPacketData(eot bool, tpduRef uint8, parameters []*COTPParameter, payload *S7Message) *COTPPacket {
	child := &COTPPacketData{
		Eot:     eot,
		TpduRef: tpduRef,
		Parent:  NewCOTPPacket(parameters, payload),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCOTPPacketData(structType interface{}) *COTPPacketData {
	castFunc := func(typ interface{}) *COTPPacketData {
		if casted, ok := typ.(COTPPacketData); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPPacketData); ok {
			return casted
		}
		if casted, ok := typ.(COTPPacket); ok {
			return CastCOTPPacketData(casted.Child)
		}
		if casted, ok := typ.(*COTPPacket); ok {
			return CastCOTPPacketData(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPPacketData) GetTypeName() string {
	return "COTPPacketData"
}

func (m *COTPPacketData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *COTPPacketData) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (eot)
	lengthInBits += 1

	// Simple field (tpduRef)
	lengthInBits += 7

	return lengthInBits
}

func (m *COTPPacketData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPPacketDataParse(readBuffer utils.ReadBuffer) (*COTPPacket, error) {
	if pullErr := readBuffer.PullContext("COTPPacketData"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (eot)
	eot, _eotErr := readBuffer.ReadBit("eot")
	if _eotErr != nil {
		return nil, errors.Wrap(_eotErr, "Error parsing 'eot' field")
	}

	// Simple Field (tpduRef)
	tpduRef, _tpduRefErr := readBuffer.ReadUint8("tpduRef", 7)
	if _tpduRefErr != nil {
		return nil, errors.Wrap(_tpduRefErr, "Error parsing 'tpduRef' field")
	}

	if closeErr := readBuffer.CloseContext("COTPPacketData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPPacketData{
		Eot:     eot,
		TpduRef: tpduRef,
		Parent:  &COTPPacket{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *COTPPacketData) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketData"); pushErr != nil {
			return pushErr
		}

		// Simple Field (eot)
		eot := bool(m.Eot)
		_eotErr := writeBuffer.WriteBit("eot", (eot))
		if _eotErr != nil {
			return errors.Wrap(_eotErr, "Error serializing 'eot' field")
		}

		// Simple Field (tpduRef)
		tpduRef := uint8(m.TpduRef)
		_tpduRefErr := writeBuffer.WriteUint8("tpduRef", 7, (tpduRef))
		if _tpduRefErr != nil {
			return errors.Wrap(_tpduRefErr, "Error serializing 'tpduRef' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketData"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
