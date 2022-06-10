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

// BACnetPropertyStatesIntegerValue is the data-structure of this message
type BACnetPropertyStatesIntegerValue struct {
	*BACnetPropertyStates
	IntegerValue *BACnetContextTagSignedInteger
}

// IBACnetPropertyStatesIntegerValue is the corresponding interface of BACnetPropertyStatesIntegerValue
type IBACnetPropertyStatesIntegerValue interface {
	IBACnetPropertyStates
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() *BACnetContextTagSignedInteger
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

func (m *BACnetPropertyStatesIntegerValue) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesIntegerValue) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesIntegerValue) GetIntegerValue() *BACnetContextTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesIntegerValue factory function for BACnetPropertyStatesIntegerValue
func NewBACnetPropertyStatesIntegerValue(integerValue *BACnetContextTagSignedInteger, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesIntegerValue {
	_result := &BACnetPropertyStatesIntegerValue{
		IntegerValue:         integerValue,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesIntegerValue(structType interface{}) *BACnetPropertyStatesIntegerValue {
	if casted, ok := structType.(BACnetPropertyStatesIntegerValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesIntegerValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesIntegerValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesIntegerValue(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesIntegerValue) GetTypeName() string {
	return "BACnetPropertyStatesIntegerValue"
}

func (m *BACnetPropertyStatesIntegerValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesIntegerValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesIntegerValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesIntegerValueParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesIntegerValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesIntegerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesIntegerValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
	_integerValue, _integerValueErr := BACnetContextTagParse(readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field")
	}
	integerValue := CastBACnetContextTagSignedInteger(_integerValue)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesIntegerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesIntegerValue")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesIntegerValue{
		IntegerValue:         CastBACnetContextTagSignedInteger(integerValue),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesIntegerValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesIntegerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesIntegerValue")
		}

		// Simple Field (integerValue)
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerValue")
		}
		_integerValueErr := m.IntegerValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerValue")
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesIntegerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesIntegerValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesIntegerValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
