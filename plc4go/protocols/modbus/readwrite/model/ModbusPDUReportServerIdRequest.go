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

// ModbusPDUReportServerIdRequest is the data-structure of this message
type ModbusPDUReportServerIdRequest struct {
	*ModbusPDU
}

// IModbusPDUReportServerIdRequest is the corresponding interface of ModbusPDUReportServerIdRequest
type IModbusPDUReportServerIdRequest interface {
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

func (m *ModbusPDUReportServerIdRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReportServerIdRequest) GetFunctionFlag() uint8 {
	return 0x11
}

func (m *ModbusPDUReportServerIdRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReportServerIdRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReportServerIdRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

// NewModbusPDUReportServerIdRequest factory function for ModbusPDUReportServerIdRequest
func NewModbusPDUReportServerIdRequest() *ModbusPDUReportServerIdRequest {
	_result := &ModbusPDUReportServerIdRequest{
		ModbusPDU: NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUReportServerIdRequest(structType interface{}) *ModbusPDUReportServerIdRequest {
	if casted, ok := structType.(ModbusPDUReportServerIdRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReportServerIdRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReportServerIdRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReportServerIdRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReportServerIdRequest) GetTypeName() string {
	return "ModbusPDUReportServerIdRequest"
}

func (m *ModbusPDUReportServerIdRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReportServerIdRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ModbusPDUReportServerIdRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReportServerIdRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUReportServerIdRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReportServerIdRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReportServerIdRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUReportServerIdRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReportServerIdRequest")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReportServerIdRequest{
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUReportServerIdRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReportServerIdRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReportServerIdRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReportServerIdRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReportServerIdRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReportServerIdRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
