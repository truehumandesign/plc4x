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
type ApduDataExtAuthorizeResponse struct {
	*ApduDataExt
	Level uint8

	// Arguments.
	Length uint8
}

// The corresponding interface
type IApduDataExtAuthorizeResponse interface {
	IApduDataExt
	// GetLevel returns Level (property field)
	GetLevel() uint8
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
func (m *ApduDataExtAuthorizeResponse) ExtApciType() uint8 {
	return 0x12
}

func (m *ApduDataExtAuthorizeResponse) GetExtApciType() uint8 {
	return 0x12
}

func (m *ApduDataExtAuthorizeResponse) InitializeParent(parent *ApduDataExt) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ApduDataExtAuthorizeResponse) GetLevel() uint8 {
	return m.Level
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataExtAuthorizeResponse factory function for ApduDataExtAuthorizeResponse
func NewApduDataExtAuthorizeResponse(level uint8, length uint8) *ApduDataExt {
	child := &ApduDataExtAuthorizeResponse{
		Level:       level,
		ApduDataExt: NewApduDataExt(length),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtAuthorizeResponse(structType interface{}) *ApduDataExtAuthorizeResponse {
	if casted, ok := structType.(ApduDataExtAuthorizeResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtAuthorizeResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtAuthorizeResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtAuthorizeResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataExtAuthorizeResponse) GetTypeName() string {
	return "ApduDataExtAuthorizeResponse"
}

func (m *ApduDataExtAuthorizeResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtAuthorizeResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (level)
	lengthInBits += 8

	return lengthInBits
}

func (m *ApduDataExtAuthorizeResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtAuthorizeResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtAuthorizeResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (level)
	_level, _levelErr := readBuffer.ReadUint8("level", 8)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field")
	}
	level := _level

	if closeErr := readBuffer.CloseContext("ApduDataExtAuthorizeResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtAuthorizeResponse{
		Level:       level,
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtAuthorizeResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtAuthorizeResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (level)
		level := uint8(m.Level)
		_levelErr := writeBuffer.WriteUint8("level", 8, (level))
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtAuthorizeResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtAuthorizeResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
