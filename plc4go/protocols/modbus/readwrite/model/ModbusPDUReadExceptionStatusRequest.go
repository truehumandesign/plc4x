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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadExceptionStatusRequest is the data-structure of this message
type ModbusPDUReadExceptionStatusRequest struct {
	*ModbusPDU
}

// IModbusPDUReadExceptionStatusRequest is the corresponding interface of ModbusPDUReadExceptionStatusRequest
type IModbusPDUReadExceptionStatusRequest interface {
	IModbusPDU
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *ModbusPDUReadExceptionStatusRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadExceptionStatusRequest) GetFunctionFlag() uint8 {
	return 0x07
}

func (m *ModbusPDUReadExceptionStatusRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReadExceptionStatusRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReadExceptionStatusRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

// NewModbusPDUReadExceptionStatusRequest factory function for ModbusPDUReadExceptionStatusRequest
func NewModbusPDUReadExceptionStatusRequest() *ModbusPDUReadExceptionStatusRequest {
	_result := &ModbusPDUReadExceptionStatusRequest{
		ModbusPDU: NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUReadExceptionStatusRequest(structType interface{}) *ModbusPDUReadExceptionStatusRequest {
	if casted, ok := structType.(ModbusPDUReadExceptionStatusRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadExceptionStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadExceptionStatusRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadExceptionStatusRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadExceptionStatusRequest) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusRequest"
}

func (m *ModbusPDUReadExceptionStatusRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadExceptionStatusRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ModbusPDUReadExceptionStatusRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadExceptionStatusRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUReadExceptionStatusRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadExceptionStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadExceptionStatusRequest")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadExceptionStatusRequest{
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUReadExceptionStatusRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadExceptionStatusRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadExceptionStatusRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadExceptionStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
