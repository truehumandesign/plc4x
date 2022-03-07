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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataMemoryRead struct {
	*ApduData
	NumBytes uint8
	Address  uint16

	// Arguments.
	DataLength uint8
}

// The corresponding interface
type IApduDataMemoryRead interface {
	IApduData
	// GetNumBytes returns NumBytes (property field)
	GetNumBytes() uint8
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryRead) ApciType() uint8 {
	return 0x8
}

func (m *ApduDataMemoryRead) GetApciType() uint8 {
	return 0x8
}

func (m *ApduDataMemoryRead) InitializeParent(parent *ApduData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryRead) GetNumBytes() uint8 {
	return m.NumBytes
}

func (m *ApduDataMemoryRead) GetAddress() uint16 {
	return m.Address
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataMemoryRead factory function for ApduDataMemoryRead
func NewApduDataMemoryRead(numBytes uint8, address uint16, dataLength uint8) *ApduData {
	child := &ApduDataMemoryRead{
		NumBytes: numBytes,
		Address:  address,
		ApduData: NewApduData(dataLength),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataMemoryRead(structType interface{}) *ApduDataMemoryRead {
	if casted, ok := structType.(ApduDataMemoryRead); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataMemoryRead); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataMemoryRead(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataMemoryRead(casted.Child)
	}
	return nil
}

func (m *ApduDataMemoryRead) GetTypeName() string {
	return "ApduDataMemoryRead"
}

func (m *ApduDataMemoryRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataMemoryRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numBytes)
	lengthInBits += 6

	// Simple field (address)
	lengthInBits += 16

	return lengthInBits
}

func (m *ApduDataMemoryRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataMemoryReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataMemoryRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (numBytes)
	_numBytes, _numBytesErr := readBuffer.ReadUint8("numBytes", 6)
	if _numBytesErr != nil {
		return nil, errors.Wrap(_numBytesErr, "Error parsing 'numBytes' field")
	}
	numBytes := _numBytes

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	if closeErr := readBuffer.CloseContext("ApduDataMemoryRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataMemoryRead{
		NumBytes: numBytes,
		Address:  address,
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataMemoryRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryRead"); pushErr != nil {
			return pushErr
		}

		// Simple Field (numBytes)
		numBytes := uint8(m.NumBytes)
		_numBytesErr := writeBuffer.WriteUint8("numBytes", 6, (numBytes))
		if _numBytesErr != nil {
			return errors.Wrap(_numBytesErr, "Error serializing 'numBytes' field")
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataMemoryRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
