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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetPropertyValue struct {
	PropertyIdentifier *BACnetContextTagPropertyIdentifier
	PropertyArrayIndex *BACnetContextTagUnsignedInteger
	PropertyValue      *BACnetConstructedDataElement
	Priority           *BACnetContextTagUnsignedInteger

	// Arguments.
	ObjectType BACnetObjectType
}

// The corresponding interface
type IBACnetPropertyValue interface {
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() *BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() *BACnetConstructedDataElement
	// GetPriority returns Priority (property field)
	GetPriority() *BACnetContextTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetPropertyValue) GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier {
	return m.PropertyIdentifier
}

func (m *BACnetPropertyValue) GetPropertyArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *BACnetPropertyValue) GetPropertyValue() *BACnetConstructedDataElement {
	return m.PropertyValue
}

func (m *BACnetPropertyValue) GetPriority() *BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetPropertyValue factory function for BACnetPropertyValue
func NewBACnetPropertyValue(propertyIdentifier *BACnetContextTagPropertyIdentifier, propertyArrayIndex *BACnetContextTagUnsignedInteger, propertyValue *BACnetConstructedDataElement, priority *BACnetContextTagUnsignedInteger, objectType BACnetObjectType) *BACnetPropertyValue {
	return &BACnetPropertyValue{PropertyIdentifier: propertyIdentifier, PropertyArrayIndex: propertyArrayIndex, PropertyValue: propertyValue, Priority: priority, ObjectType: objectType}
}

func CastBACnetPropertyValue(structType interface{}) *BACnetPropertyValue {
	if casted, ok := structType.(BACnetPropertyValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyValue); ok {
		return casted
	}
	return nil
}

func (m *BACnetPropertyValue) GetTypeName() string {
	return "BACnetPropertyValue"
}

func (m *BACnetPropertyValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += (*m.PropertyArrayIndex).GetLengthInBits()
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += (*m.PropertyValue).GetLengthInBits()
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += (*m.Priority).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetPropertyValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyValueParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType) (*BACnetPropertyValue, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_PROPERTY_IDENTIFIER))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetContextTagPropertyIdentifier(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if a given expression evaluates to false)
	var propertyArrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("propertyArrayIndex"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyArrayIndex' field")
		default:
			propertyArrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("propertyArrayIndex"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue *BACnetConstructedDataElement = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataElementParse(readBuffer, objectType, propertyIdentifier)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field")
		default:
			propertyValue = CastBACnetConstructedDataElement(_val)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field")
		default:
			priority = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetPropertyValue(propertyIdentifier, propertyArrayIndex, propertyValue, priority, objectType), nil
}

func (m *BACnetPropertyValue) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetPropertyValue"); pushErr != nil {
		return pushErr
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return pushErr
	}
	_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return popErr
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if the value is null)
	var propertyArrayIndex *BACnetContextTagUnsignedInteger = nil
	if m.PropertyArrayIndex != nil {
		if pushErr := writeBuffer.PushContext("propertyArrayIndex"); pushErr != nil {
			return pushErr
		}
		propertyArrayIndex = m.PropertyArrayIndex
		_propertyArrayIndexErr := propertyArrayIndex.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyArrayIndex"); popErr != nil {
			return popErr
		}
		if _propertyArrayIndexErr != nil {
			return errors.Wrap(_propertyArrayIndexErr, "Error serializing 'propertyArrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue *BACnetConstructedDataElement = nil
	if m.PropertyValue != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return pushErr
		}
		propertyValue = m.PropertyValue
		_propertyValueErr := propertyValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return popErr
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority *BACnetContextTagUnsignedInteger = nil
	if m.Priority != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return pushErr
		}
		priority = m.Priority
		_priorityErr := priority.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return popErr
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyValue"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetPropertyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
