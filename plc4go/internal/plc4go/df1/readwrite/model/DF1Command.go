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
type DF1Command struct {
	Status             uint8
	TransactionCounter uint16
	Child              IDF1CommandChild
}

// The corresponding interface
type IDF1Command interface {
	// GetCommandCode returns CommandCode (discriminator field)
	GetCommandCode() uint8
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetTransactionCounter returns TransactionCounter (property field)
	GetTransactionCounter() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IDF1CommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IDF1Command, serializeChildFunction func() error) error
	GetTypeName() string
}

type IDF1CommandChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *DF1Command, status uint8, transactionCounter uint16)
	GetTypeName() string
	IDF1Command
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *DF1Command) GetStatus() uint8 {
	return m.Status
}

func (m *DF1Command) GetTransactionCounter() uint16 {
	return m.TransactionCounter
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewDF1Command factory function for DF1Command
func NewDF1Command(status uint8, transactionCounter uint16) *DF1Command {
	return &DF1Command{Status: status, TransactionCounter: transactionCounter}
}

func CastDF1Command(structType interface{}) *DF1Command {
	if casted, ok := structType.(DF1Command); ok {
		return &casted
	}
	if casted, ok := structType.(*DF1Command); ok {
		return casted
	}
	return nil
}

func (m *DF1Command) GetTypeName() string {
	return "DF1Command"
}

func (m *DF1Command) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DF1Command) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *DF1Command) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *DF1Command) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1CommandParse(readBuffer utils.ReadBuffer) (*DF1Command, error) {
	if pullErr := readBuffer.PullContext("DF1Command"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode, _commandCodeErr := readBuffer.ReadUint8("commandCode", 8)
	if _commandCodeErr != nil {
		return nil, errors.Wrap(_commandCodeErr, "Error parsing 'commandCode' field")
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status

	// Simple Field (transactionCounter)
	_transactionCounter, _transactionCounterErr := readBuffer.ReadUint16("transactionCounter", 16)
	if _transactionCounterErr != nil {
		return nil, errors.Wrap(_transactionCounterErr, "Error parsing 'transactionCounter' field")
	}
	transactionCounter := _transactionCounter

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *DF1Command
	var typeSwitchError error
	switch {
	case commandCode == 0x01: // DF1UnprotectedReadRequest
		_parent, typeSwitchError = DF1UnprotectedReadRequestParse(readBuffer)
	case commandCode == 0x41: // DF1UnprotectedReadResponse
		_parent, typeSwitchError = DF1UnprotectedReadResponseParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("DF1Command"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, status, transactionCounter)
	return _parent, nil
}

func (m *DF1Command) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *DF1Command) SerializeParent(writeBuffer utils.WriteBuffer, child IDF1Command, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("DF1Command"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(child.GetCommandCode())
	_commandCodeErr := writeBuffer.WriteUint8("commandCode", 8, (commandCode))

	if _commandCodeErr != nil {
		return errors.Wrap(_commandCodeErr, "Error serializing 'commandCode' field")
	}

	// Simple Field (status)
	status := uint8(m.Status)
	_statusErr := writeBuffer.WriteUint8("status", 8, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Simple Field (transactionCounter)
	transactionCounter := uint16(m.TransactionCounter)
	_transactionCounterErr := writeBuffer.WriteUint16("transactionCounter", 16, (transactionCounter))
	if _transactionCounterErr != nil {
		return errors.Wrap(_transactionCounterErr, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1Command"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DF1Command) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
