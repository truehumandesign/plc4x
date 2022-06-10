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

// BACnetConfirmedServiceRequestReadRangeRangeByTime is the data-structure of this message
type BACnetConfirmedServiceRequestReadRangeRangeByTime struct {
	*BACnetConfirmedServiceRequestReadRangeRange
	ReferenceTime *BACnetDateTime
	Count         *BACnetApplicationTagSignedInteger
}

// IBACnetConfirmedServiceRequestReadRangeRangeByTime is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeByTime
type IBACnetConfirmedServiceRequestReadRangeRangeByTime interface {
	IBACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceTime returns ReferenceTime (property field)
	GetReferenceTime() *BACnetDateTime
	// GetCount returns Count (property field)
	GetCount() *BACnetApplicationTagSignedInteger
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

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) InitializeParent(parent *BACnetConfirmedServiceRequestReadRangeRange, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConfirmedServiceRequestReadRangeRange.PeekedTagHeader = peekedTagHeader
	m.BACnetConfirmedServiceRequestReadRangeRange.OpeningTag = openingTag
	m.BACnetConfirmedServiceRequestReadRangeRange.ClosingTag = closingTag
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) GetParent() *BACnetConfirmedServiceRequestReadRangeRange {
	return m.BACnetConfirmedServiceRequestReadRangeRange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) GetReferenceTime() *BACnetDateTime {
	return m.ReferenceTime
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) GetCount() *BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRangeRangeByTime factory function for BACnetConfirmedServiceRequestReadRangeRangeByTime
func NewBACnetConfirmedServiceRequestReadRangeRangeByTime(referenceTime *BACnetDateTime, count *BACnetApplicationTagSignedInteger, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) *BACnetConfirmedServiceRequestReadRangeRangeByTime {
	_result := &BACnetConfirmedServiceRequestReadRangeRangeByTime{
		ReferenceTime: referenceTime,
		Count:         count,
		BACnetConfirmedServiceRequestReadRangeRange: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestReadRangeRangeByTime(structType interface{}) *BACnetConfirmedServiceRequestReadRangeRangeByTime {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeByTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeByTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRange); ok {
		return CastBACnetConfirmedServiceRequestReadRangeRangeByTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRange); ok {
		return CastBACnetConfirmedServiceRequestReadRangeRangeByTime(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeByTime"
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (referenceTime)
	lengthInBits += m.ReferenceTime.GetLengthInBits()

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadRangeRangeByTimeParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceRequestReadRangeRangeByTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRangeByTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceTime)
	if pullErr := readBuffer.PullContext("referenceTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceTime")
	}
	_referenceTime, _referenceTimeErr := BACnetDateTimeParse(readBuffer)
	if _referenceTimeErr != nil {
		return nil, errors.Wrap(_referenceTimeErr, "Error parsing 'referenceTime' field")
	}
	referenceTime := CastBACnetDateTime(_referenceTime)
	if closeErr := readBuffer.CloseContext("referenceTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceTime")
	}

	// Simple Field (count)
	if pullErr := readBuffer.PullContext("count"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for count")
	}
	_count, _countErr := BACnetApplicationTagParse(readBuffer)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := CastBACnetApplicationTagSignedInteger(_count)
	if closeErr := readBuffer.CloseContext("count"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for count")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRangeByTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestReadRangeRangeByTime{
		ReferenceTime: CastBACnetDateTime(referenceTime),
		Count:         CastBACnetApplicationTagSignedInteger(count),
		BACnetConfirmedServiceRequestReadRangeRange: &BACnetConfirmedServiceRequestReadRangeRange{},
	}
	_child.BACnetConfirmedServiceRequestReadRangeRange.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRangeByTime")
		}

		// Simple Field (referenceTime)
		if pushErr := writeBuffer.PushContext("referenceTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceTime")
		}
		_referenceTimeErr := m.ReferenceTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("referenceTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceTime")
		}
		if _referenceTimeErr != nil {
			return errors.Wrap(_referenceTimeErr, "Error serializing 'referenceTime' field")
		}

		// Simple Field (count)
		if pushErr := writeBuffer.PushContext("count"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for count")
		}
		_countErr := m.Count.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("count"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for count")
		}
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRangeByTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeByTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
