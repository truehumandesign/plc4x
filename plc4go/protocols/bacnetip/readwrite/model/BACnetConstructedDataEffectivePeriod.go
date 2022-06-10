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

// BACnetConstructedDataEffectivePeriod is the data-structure of this message
type BACnetConstructedDataEffectivePeriod struct {
	*BACnetConstructedData
	DateRange *BACnetDateRange

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEffectivePeriod is the corresponding interface of BACnetConstructedDataEffectivePeriod
type IBACnetConstructedDataEffectivePeriod interface {
	IBACnetConstructedData
	// GetDateRange returns DateRange (property field)
	GetDateRange() *BACnetDateRange
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

func (m *BACnetConstructedDataEffectivePeriod) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEffectivePeriod) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EFFECTIVE_PERIOD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEffectivePeriod) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEffectivePeriod) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEffectivePeriod) GetDateRange() *BACnetDateRange {
	return m.DateRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEffectivePeriod factory function for BACnetConstructedDataEffectivePeriod
func NewBACnetConstructedDataEffectivePeriod(dateRange *BACnetDateRange, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEffectivePeriod {
	_result := &BACnetConstructedDataEffectivePeriod{
		DateRange:             dateRange,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEffectivePeriod(structType interface{}) *BACnetConstructedDataEffectivePeriod {
	if casted, ok := structType.(BACnetConstructedDataEffectivePeriod); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEffectivePeriod); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEffectivePeriod(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEffectivePeriod(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEffectivePeriod) GetTypeName() string {
	return "BACnetConstructedDataEffectivePeriod"
}

func (m *BACnetConstructedDataEffectivePeriod) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEffectivePeriod) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataEffectivePeriod) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEffectivePeriodParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEffectivePeriod, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEffectivePeriod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEffectivePeriod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateRange)
	if pullErr := readBuffer.PullContext("dateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateRange")
	}
	_dateRange, _dateRangeErr := BACnetDateRangeParse(readBuffer)
	if _dateRangeErr != nil {
		return nil, errors.Wrap(_dateRangeErr, "Error parsing 'dateRange' field")
	}
	dateRange := CastBACnetDateRange(_dateRange)
	if closeErr := readBuffer.CloseContext("dateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateRange")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEffectivePeriod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEffectivePeriod")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEffectivePeriod{
		DateRange:             CastBACnetDateRange(dateRange),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEffectivePeriod) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEffectivePeriod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEffectivePeriod")
		}

		// Simple Field (dateRange)
		if pushErr := writeBuffer.PushContext("dateRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateRange")
		}
		_dateRangeErr := m.DateRange.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dateRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateRange")
		}
		if _dateRangeErr != nil {
			return errors.Wrap(_dateRangeErr, "Error serializing 'dateRange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEffectivePeriod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEffectivePeriod")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEffectivePeriod) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
