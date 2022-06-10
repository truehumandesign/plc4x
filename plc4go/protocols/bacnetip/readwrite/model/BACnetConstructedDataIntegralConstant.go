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

// BACnetConstructedDataIntegralConstant is the data-structure of this message
type BACnetConstructedDataIntegralConstant struct {
	*BACnetConstructedData
	IntegralConstant *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataIntegralConstant is the corresponding interface of BACnetConstructedDataIntegralConstant
type IBACnetConstructedDataIntegralConstant interface {
	IBACnetConstructedData
	// GetIntegralConstant returns IntegralConstant (property field)
	GetIntegralConstant() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataIntegralConstant) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIntegralConstant) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTEGRAL_CONSTANT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIntegralConstant) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIntegralConstant) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIntegralConstant) GetIntegralConstant() *BACnetApplicationTagReal {
	return m.IntegralConstant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegralConstant factory function for BACnetConstructedDataIntegralConstant
func NewBACnetConstructedDataIntegralConstant(integralConstant *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataIntegralConstant {
	_result := &BACnetConstructedDataIntegralConstant{
		IntegralConstant:      integralConstant,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIntegralConstant(structType interface{}) *BACnetConstructedDataIntegralConstant {
	if casted, ok := structType.(BACnetConstructedDataIntegralConstant); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegralConstant); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIntegralConstant(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIntegralConstant(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIntegralConstant) GetTypeName() string {
	return "BACnetConstructedDataIntegralConstant"
}

func (m *BACnetConstructedDataIntegralConstant) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIntegralConstant) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integralConstant)
	lengthInBits += m.IntegralConstant.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIntegralConstant) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntegralConstantParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataIntegralConstant, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegralConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegralConstant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integralConstant)
	if pullErr := readBuffer.PullContext("integralConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integralConstant")
	}
	_integralConstant, _integralConstantErr := BACnetApplicationTagParse(readBuffer)
	if _integralConstantErr != nil {
		return nil, errors.Wrap(_integralConstantErr, "Error parsing 'integralConstant' field")
	}
	integralConstant := CastBACnetApplicationTagReal(_integralConstant)
	if closeErr := readBuffer.CloseContext("integralConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integralConstant")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegralConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegralConstant")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIntegralConstant{
		IntegralConstant:      CastBACnetApplicationTagReal(integralConstant),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIntegralConstant) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegralConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegralConstant")
		}

		// Simple Field (integralConstant)
		if pushErr := writeBuffer.PushContext("integralConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integralConstant")
		}
		_integralConstantErr := m.IntegralConstant.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("integralConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integralConstant")
		}
		if _integralConstantErr != nil {
			return errors.Wrap(_integralConstantErr, "Error serializing 'integralConstant' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegralConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegralConstant")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIntegralConstant) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
