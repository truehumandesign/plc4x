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

// BACnetConstructedDataMusterPoint is the data-structure of this message
type BACnetConstructedDataMusterPoint struct {
	*BACnetConstructedData
	MusterPoint *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMusterPoint is the corresponding interface of BACnetConstructedDataMusterPoint
type IBACnetConstructedDataMusterPoint interface {
	IBACnetConstructedData
	// GetMusterPoint returns MusterPoint (property field)
	GetMusterPoint() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataMusterPoint) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMusterPoint) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MUSTER_POINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMusterPoint) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMusterPoint) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMusterPoint) GetMusterPoint() *BACnetApplicationTagBoolean {
	return m.MusterPoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMusterPoint factory function for BACnetConstructedDataMusterPoint
func NewBACnetConstructedDataMusterPoint(musterPoint *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMusterPoint {
	_result := &BACnetConstructedDataMusterPoint{
		MusterPoint:           musterPoint,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMusterPoint(structType interface{}) *BACnetConstructedDataMusterPoint {
	if casted, ok := structType.(BACnetConstructedDataMusterPoint); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMusterPoint); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMusterPoint(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMusterPoint(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMusterPoint) GetTypeName() string {
	return "BACnetConstructedDataMusterPoint"
}

func (m *BACnetConstructedDataMusterPoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMusterPoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (musterPoint)
	lengthInBits += m.MusterPoint.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMusterPoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMusterPointParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMusterPoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMusterPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMusterPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (musterPoint)
	if pullErr := readBuffer.PullContext("musterPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for musterPoint")
	}
	_musterPoint, _musterPointErr := BACnetApplicationTagParse(readBuffer)
	if _musterPointErr != nil {
		return nil, errors.Wrap(_musterPointErr, "Error parsing 'musterPoint' field")
	}
	musterPoint := CastBACnetApplicationTagBoolean(_musterPoint)
	if closeErr := readBuffer.CloseContext("musterPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for musterPoint")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMusterPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMusterPoint")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMusterPoint{
		MusterPoint:           CastBACnetApplicationTagBoolean(musterPoint),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMusterPoint) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMusterPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMusterPoint")
		}

		// Simple Field (musterPoint)
		if pushErr := writeBuffer.PushContext("musterPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for musterPoint")
		}
		_musterPointErr := m.MusterPoint.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("musterPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for musterPoint")
		}
		if _musterPointErr != nil {
			return errors.Wrap(_musterPointErr, "Error serializing 'musterPoint' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMusterPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMusterPoint")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMusterPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
