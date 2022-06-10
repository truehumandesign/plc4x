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

// BACnetLogRecordLogDatum is the data-structure of this message
type BACnetLogRecordLogDatum struct {
	OpeningTag      *BACnetOpeningTag
	PeekedTagHeader *BACnetTagHeader
	ClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber uint8
	Child     IBACnetLogRecordLogDatumChild
}

// IBACnetLogRecordLogDatum is the corresponding interface of BACnetLogRecordLogDatum
type IBACnetLogRecordLogDatum interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetLogRecordLogDatumParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetLogRecordLogDatum, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetLogRecordLogDatumChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetLogRecordLogDatum, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag)
	GetParent() *BACnetLogRecordLogDatum

	GetTypeName() string
	IBACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLogRecordLogDatum) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetLogRecordLogDatum) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetLogRecordLogDatum) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetLogRecordLogDatum) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatum factory function for BACnetLogRecordLogDatum
func NewBACnetLogRecordLogDatum(openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetLogRecordLogDatum {
	return &BACnetLogRecordLogDatum{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetLogRecordLogDatum(structType interface{}) *BACnetLogRecordLogDatum {
	if casted, ok := structType.(BACnetLogRecordLogDatum); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatum); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetLogRecordLogDatumChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetLogRecordLogDatum) GetTypeName() string {
	return "BACnetLogRecordLogDatum"
}

func (m *BACnetLogRecordLogDatum) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogRecordLogDatum) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetLogRecordLogDatum) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLogRecordLogDatum) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetLogRecordLogDatum, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetLogRecordLogDatumChild interface {
		InitializeParent(*BACnetLogRecordLogDatum, *BACnetOpeningTag, *BACnetTagHeader, *BACnetClosingTag)
		GetParent() *BACnetLogRecordLogDatum
	}
	var _child BACnetLogRecordLogDatumChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetLogRecordLogDatumLogStatus
		_child, typeSwitchError = BACnetLogRecordLogDatumLogStatusParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(1): // BACnetLogRecordLogDatumBooleanValue
		_child, typeSwitchError = BACnetLogRecordLogDatumBooleanValueParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(2): // BACnetLogRecordLogDatumRealValue
		_child, typeSwitchError = BACnetLogRecordLogDatumRealValueParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(3): // BACnetLogRecordLogDatumEnumeratedValue
		_child, typeSwitchError = BACnetLogRecordLogDatumEnumeratedValueParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(4): // BACnetLogRecordLogDatumUnsignedValue
		_child, typeSwitchError = BACnetLogRecordLogDatumUnsignedValueParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(5): // BACnetLogRecordLogDatumIntegerValue
		_child, typeSwitchError = BACnetLogRecordLogDatumIntegerValueParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(6): // BACnetLogRecordLogDatumBitStringValue
		_child, typeSwitchError = BACnetLogRecordLogDatumBitStringValueParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(7): // BACnetLogRecordLogDatumNullValue
		_child, typeSwitchError = BACnetLogRecordLogDatumNullValueParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(8): // BACnetLogRecordLogDatumFailure
		_child, typeSwitchError = BACnetLogRecordLogDatumFailureParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(9): // BACnetLogRecordLogDatumTimeChange
		_child, typeSwitchError = BACnetLogRecordLogDatumTimeChangeParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(10): // BACnetLogRecordLogDatumAnyValue
		_child, typeSwitchError = BACnetLogRecordLogDatumAnyValueParse(readBuffer, tagNumber)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatum")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), openingTag, peekedTagHeader, closingTag)
	return _child.GetParent(), nil
}

func (m *BACnetLogRecordLogDatum) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetLogRecordLogDatum) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetLogRecordLogDatum, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatum")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatum")
	}
	return nil
}

func (m *BACnetLogRecordLogDatum) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
