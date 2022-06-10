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

// BACnetConstructedDataModificationDate is the data-structure of this message
type BACnetConstructedDataModificationDate struct {
	*BACnetConstructedData
	ModificationDate *BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataModificationDate is the corresponding interface of BACnetConstructedDataModificationDate
type IBACnetConstructedDataModificationDate interface {
	IBACnetConstructedData
	// GetModificationDate returns ModificationDate (property field)
	GetModificationDate() *BACnetDateTime
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

func (m *BACnetConstructedDataModificationDate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataModificationDate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MODIFICATION_DATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataModificationDate) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataModificationDate) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataModificationDate) GetModificationDate() *BACnetDateTime {
	return m.ModificationDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataModificationDate factory function for BACnetConstructedDataModificationDate
func NewBACnetConstructedDataModificationDate(modificationDate *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataModificationDate {
	_result := &BACnetConstructedDataModificationDate{
		ModificationDate:      modificationDate,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataModificationDate(structType interface{}) *BACnetConstructedDataModificationDate {
	if casted, ok := structType.(BACnetConstructedDataModificationDate); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataModificationDate); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataModificationDate(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataModificationDate(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataModificationDate) GetTypeName() string {
	return "BACnetConstructedDataModificationDate"
}

func (m *BACnetConstructedDataModificationDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataModificationDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (modificationDate)
	lengthInBits += m.ModificationDate.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataModificationDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataModificationDateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataModificationDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataModificationDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataModificationDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (modificationDate)
	if pullErr := readBuffer.PullContext("modificationDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for modificationDate")
	}
	_modificationDate, _modificationDateErr := BACnetDateTimeParse(readBuffer)
	if _modificationDateErr != nil {
		return nil, errors.Wrap(_modificationDateErr, "Error parsing 'modificationDate' field")
	}
	modificationDate := CastBACnetDateTime(_modificationDate)
	if closeErr := readBuffer.CloseContext("modificationDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for modificationDate")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataModificationDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataModificationDate")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataModificationDate{
		ModificationDate:      CastBACnetDateTime(modificationDate),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataModificationDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataModificationDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataModificationDate")
		}

		// Simple Field (modificationDate)
		if pushErr := writeBuffer.PushContext("modificationDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for modificationDate")
		}
		_modificationDateErr := m.ModificationDate.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("modificationDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for modificationDate")
		}
		if _modificationDateErr != nil {
			return errors.Wrap(_modificationDateErr, "Error serializing 'modificationDate' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataModificationDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataModificationDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataModificationDate) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
