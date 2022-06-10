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

// BACnetConstructedDataSetpoint is the data-structure of this message
type BACnetConstructedDataSetpoint struct {
	*BACnetConstructedData
	Setpoint *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataSetpoint is the corresponding interface of BACnetConstructedDataSetpoint
type IBACnetConstructedDataSetpoint interface {
	IBACnetConstructedData
	// GetSetpoint returns Setpoint (property field)
	GetSetpoint() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataSetpoint) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSetpoint) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SETPOINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSetpoint) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSetpoint) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSetpoint) GetSetpoint() *BACnetApplicationTagReal {
	return m.Setpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSetpoint factory function for BACnetConstructedDataSetpoint
func NewBACnetConstructedDataSetpoint(setpoint *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataSetpoint {
	_result := &BACnetConstructedDataSetpoint{
		Setpoint:              setpoint,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSetpoint(structType interface{}) *BACnetConstructedDataSetpoint {
	if casted, ok := structType.(BACnetConstructedDataSetpoint); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSetpoint); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSetpoint(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSetpoint(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSetpoint) GetTypeName() string {
	return "BACnetConstructedDataSetpoint"
}

func (m *BACnetConstructedDataSetpoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSetpoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (setpoint)
	lengthInBits += m.Setpoint.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataSetpoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSetpointParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataSetpoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSetpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSetpoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (setpoint)
	if pullErr := readBuffer.PullContext("setpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for setpoint")
	}
	_setpoint, _setpointErr := BACnetApplicationTagParse(readBuffer)
	if _setpointErr != nil {
		return nil, errors.Wrap(_setpointErr, "Error parsing 'setpoint' field")
	}
	setpoint := CastBACnetApplicationTagReal(_setpoint)
	if closeErr := readBuffer.CloseContext("setpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for setpoint")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSetpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSetpoint")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSetpoint{
		Setpoint:              CastBACnetApplicationTagReal(setpoint),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSetpoint) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSetpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSetpoint")
		}

		// Simple Field (setpoint)
		if pushErr := writeBuffer.PushContext("setpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for setpoint")
		}
		_setpointErr := m.Setpoint.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("setpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for setpoint")
		}
		if _setpointErr != nil {
			return errors.Wrap(_setpointErr, "Error serializing 'setpoint' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSetpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSetpoint")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSetpoint) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
