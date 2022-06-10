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

// BACnetLogDataLogDataTimeChange is the data-structure of this message
type BACnetLogDataLogDataTimeChange struct {
	*BACnetLogData
	TimeChange *BACnetContextTagReal

	// Arguments.
	TagNumber uint8
}

// IBACnetLogDataLogDataTimeChange is the corresponding interface of BACnetLogDataLogDataTimeChange
type IBACnetLogDataLogDataTimeChange interface {
	IBACnetLogData
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() *BACnetContextTagReal
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetLogDataLogDataTimeChange) InitializeParent(parent *BACnetLogData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetLogData.OpeningTag = openingTag
	m.BACnetLogData.PeekedTagHeader = peekedTagHeader
	m.BACnetLogData.ClosingTag = closingTag
}

func (m *BACnetLogDataLogDataTimeChange) GetParent() *BACnetLogData {
	return m.BACnetLogData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLogDataLogDataTimeChange) GetTimeChange() *BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataTimeChange factory function for BACnetLogDataLogDataTimeChange
func NewBACnetLogDataLogDataTimeChange(timeChange *BACnetContextTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetLogDataLogDataTimeChange {
	_result := &BACnetLogDataLogDataTimeChange{
		TimeChange:    timeChange,
		BACnetLogData: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetLogDataLogDataTimeChange(structType interface{}) *BACnetLogDataLogDataTimeChange {
	if casted, ok := structType.(BACnetLogDataLogDataTimeChange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetLogData); ok {
		return CastBACnetLogDataLogDataTimeChange(casted.Child)
	}
	if casted, ok := structType.(*BACnetLogData); ok {
		return CastBACnetLogDataLogDataTimeChange(casted.Child)
	}
	return nil
}

func (m *BACnetLogDataLogDataTimeChange) GetTypeName() string {
	return "BACnetLogDataLogDataTimeChange"
}

func (m *BACnetLogDataLogDataTimeChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogDataLogDataTimeChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLogDataLogDataTimeChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataTimeChangeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetLogDataLogDataTimeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeChange)
	if pullErr := readBuffer.PullContext("timeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeChange")
	}
	_timeChange, _timeChangeErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _timeChangeErr != nil {
		return nil, errors.Wrap(_timeChangeErr, "Error parsing 'timeChange' field")
	}
	timeChange := CastBACnetContextTagReal(_timeChange)
	if closeErr := readBuffer.CloseContext("timeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeChange")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataTimeChange")
	}

	// Create a partially initialized instance
	_child := &BACnetLogDataLogDataTimeChange{
		TimeChange:    CastBACnetContextTagReal(timeChange),
		BACnetLogData: &BACnetLogData{},
	}
	_child.BACnetLogData.Child = _child
	return _child, nil
}

func (m *BACnetLogDataLogDataTimeChange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataTimeChange")
		}

		// Simple Field (timeChange)
		if pushErr := writeBuffer.PushContext("timeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeChange")
		}
		_timeChangeErr := m.TimeChange.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeChange")
		}
		if _timeChangeErr != nil {
			return errors.Wrap(_timeChangeErr, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataTimeChange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetLogDataLogDataTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
