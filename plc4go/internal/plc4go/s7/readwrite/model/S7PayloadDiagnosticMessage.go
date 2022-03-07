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
type S7PayloadDiagnosticMessage struct {
	*S7PayloadUserDataItem
	EventId       uint16
	PriorityClass uint8
	ObNumber      uint8
	DatId         uint16
	Info1         uint16
	Info2         uint32
	TimeStamp     *DateAndTime
}

// The corresponding interface
type IS7PayloadDiagnosticMessage interface {
	IS7PayloadUserDataItem
	// GetEventId returns EventId (property field)
	GetEventId() uint16
	// GetPriorityClass returns PriorityClass (property field)
	GetPriorityClass() uint8
	// GetObNumber returns ObNumber (property field)
	GetObNumber() uint8
	// GetDatId returns DatId (property field)
	GetDatId() uint16
	// GetInfo1 returns Info1 (property field)
	GetInfo1() uint16
	// GetInfo2 returns Info2 (property field)
	GetInfo2() uint32
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() *DateAndTime
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
func (m *S7PayloadDiagnosticMessage) CpuFunctionType() uint8 {
	return 0x00
}

func (m *S7PayloadDiagnosticMessage) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *S7PayloadDiagnosticMessage) CpuSubfunction() uint8 {
	return 0x03
}

func (m *S7PayloadDiagnosticMessage) GetCpuSubfunction() uint8 {
	return 0x03
}

func (m *S7PayloadDiagnosticMessage) DataLength() uint16 {
	return 0
}

func (m *S7PayloadDiagnosticMessage) GetDataLength() uint16 {
	return 0
}

func (m *S7PayloadDiagnosticMessage) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *S7PayloadDiagnosticMessage) GetEventId() uint16 {
	return m.EventId
}

func (m *S7PayloadDiagnosticMessage) GetPriorityClass() uint8 {
	return m.PriorityClass
}

func (m *S7PayloadDiagnosticMessage) GetObNumber() uint8 {
	return m.ObNumber
}

func (m *S7PayloadDiagnosticMessage) GetDatId() uint16 {
	return m.DatId
}

func (m *S7PayloadDiagnosticMessage) GetInfo1() uint16 {
	return m.Info1
}

func (m *S7PayloadDiagnosticMessage) GetInfo2() uint32 {
	return m.Info2
}

func (m *S7PayloadDiagnosticMessage) GetTimeStamp() *DateAndTime {
	return m.TimeStamp
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewS7PayloadDiagnosticMessage factory function for S7PayloadDiagnosticMessage
func NewS7PayloadDiagnosticMessage(EventId uint16, PriorityClass uint8, ObNumber uint8, DatId uint16, Info1 uint16, Info2 uint32, TimeStamp *DateAndTime, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadDiagnosticMessage{
		EventId:               EventId,
		PriorityClass:         PriorityClass,
		ObNumber:              ObNumber,
		DatId:                 DatId,
		Info1:                 Info1,
		Info2:                 Info2,
		TimeStamp:             TimeStamp,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Child = child
	return child.S7PayloadUserDataItem
}

func CastS7PayloadDiagnosticMessage(structType interface{}) *S7PayloadDiagnosticMessage {
	if casted, ok := structType.(S7PayloadDiagnosticMessage); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadDiagnosticMessage); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadDiagnosticMessage(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadDiagnosticMessage(casted.Child)
	}
	return nil
}

func (m *S7PayloadDiagnosticMessage) GetTypeName() string {
	return "S7PayloadDiagnosticMessage"
}

func (m *S7PayloadDiagnosticMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadDiagnosticMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (EventId)
	lengthInBits += 16

	// Simple field (PriorityClass)
	lengthInBits += 8

	// Simple field (ObNumber)
	lengthInBits += 8

	// Simple field (DatId)
	lengthInBits += 16

	// Simple field (Info1)
	lengthInBits += 16

	// Simple field (Info2)
	lengthInBits += 32

	// Simple field (TimeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits()

	return lengthInBits
}

func (m *S7PayloadDiagnosticMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadDiagnosticMessageParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadDiagnosticMessage"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (EventId)
	_EventId, _EventIdErr := readBuffer.ReadUint16("EventId", 16)
	if _EventIdErr != nil {
		return nil, errors.Wrap(_EventIdErr, "Error parsing 'EventId' field")
	}
	EventId := _EventId

	// Simple Field (PriorityClass)
	_PriorityClass, _PriorityClassErr := readBuffer.ReadUint8("PriorityClass", 8)
	if _PriorityClassErr != nil {
		return nil, errors.Wrap(_PriorityClassErr, "Error parsing 'PriorityClass' field")
	}
	PriorityClass := _PriorityClass

	// Simple Field (ObNumber)
	_ObNumber, _ObNumberErr := readBuffer.ReadUint8("ObNumber", 8)
	if _ObNumberErr != nil {
		return nil, errors.Wrap(_ObNumberErr, "Error parsing 'ObNumber' field")
	}
	ObNumber := _ObNumber

	// Simple Field (DatId)
	_DatId, _DatIdErr := readBuffer.ReadUint16("DatId", 16)
	if _DatIdErr != nil {
		return nil, errors.Wrap(_DatIdErr, "Error parsing 'DatId' field")
	}
	DatId := _DatId

	// Simple Field (Info1)
	_Info1, _Info1Err := readBuffer.ReadUint16("Info1", 16)
	if _Info1Err != nil {
		return nil, errors.Wrap(_Info1Err, "Error parsing 'Info1' field")
	}
	Info1 := _Info1

	// Simple Field (Info2)
	_Info2, _Info2Err := readBuffer.ReadUint32("Info2", 32)
	if _Info2Err != nil {
		return nil, errors.Wrap(_Info2Err, "Error parsing 'Info2' field")
	}
	Info2 := _Info2

	// Simple Field (TimeStamp)
	if pullErr := readBuffer.PullContext("TimeStamp"); pullErr != nil {
		return nil, pullErr
	}
	_TimeStamp, _TimeStampErr := DateAndTimeParse(readBuffer)
	if _TimeStampErr != nil {
		return nil, errors.Wrap(_TimeStampErr, "Error parsing 'TimeStamp' field")
	}
	TimeStamp := CastDateAndTime(_TimeStamp)
	if closeErr := readBuffer.CloseContext("TimeStamp"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadDiagnosticMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadDiagnosticMessage{
		EventId:               EventId,
		PriorityClass:         PriorityClass,
		ObNumber:              ObNumber,
		DatId:                 DatId,
		Info1:                 Info1,
		Info2:                 Info2,
		TimeStamp:             CastDateAndTime(TimeStamp),
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child.S7PayloadUserDataItem, nil
}

func (m *S7PayloadDiagnosticMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadDiagnosticMessage"); pushErr != nil {
			return pushErr
		}

		// Simple Field (EventId)
		EventId := uint16(m.EventId)
		_EventIdErr := writeBuffer.WriteUint16("EventId", 16, (EventId))
		if _EventIdErr != nil {
			return errors.Wrap(_EventIdErr, "Error serializing 'EventId' field")
		}

		// Simple Field (PriorityClass)
		PriorityClass := uint8(m.PriorityClass)
		_PriorityClassErr := writeBuffer.WriteUint8("PriorityClass", 8, (PriorityClass))
		if _PriorityClassErr != nil {
			return errors.Wrap(_PriorityClassErr, "Error serializing 'PriorityClass' field")
		}

		// Simple Field (ObNumber)
		ObNumber := uint8(m.ObNumber)
		_ObNumberErr := writeBuffer.WriteUint8("ObNumber", 8, (ObNumber))
		if _ObNumberErr != nil {
			return errors.Wrap(_ObNumberErr, "Error serializing 'ObNumber' field")
		}

		// Simple Field (DatId)
		DatId := uint16(m.DatId)
		_DatIdErr := writeBuffer.WriteUint16("DatId", 16, (DatId))
		if _DatIdErr != nil {
			return errors.Wrap(_DatIdErr, "Error serializing 'DatId' field")
		}

		// Simple Field (Info1)
		Info1 := uint16(m.Info1)
		_Info1Err := writeBuffer.WriteUint16("Info1", 16, (Info1))
		if _Info1Err != nil {
			return errors.Wrap(_Info1Err, "Error serializing 'Info1' field")
		}

		// Simple Field (Info2)
		Info2 := uint32(m.Info2)
		_Info2Err := writeBuffer.WriteUint32("Info2", 32, (Info2))
		if _Info2Err != nil {
			return errors.Wrap(_Info2Err, "Error serializing 'Info2' field")
		}

		// Simple Field (TimeStamp)
		if pushErr := writeBuffer.PushContext("TimeStamp"); pushErr != nil {
			return pushErr
		}
		_TimeStampErr := m.TimeStamp.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("TimeStamp"); popErr != nil {
			return popErr
		}
		if _TimeStampErr != nil {
			return errors.Wrap(_TimeStampErr, "Error serializing 'TimeStamp' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadDiagnosticMessage"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadDiagnosticMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
