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

// BACnetConstructedDataVerificationTime is the data-structure of this message
type BACnetConstructedDataVerificationTime struct {
	*BACnetConstructedData
	VerificationTime *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataVerificationTime is the corresponding interface of BACnetConstructedDataVerificationTime
type IBACnetConstructedDataVerificationTime interface {
	IBACnetConstructedData
	// GetVerificationTime returns VerificationTime (property field)
	GetVerificationTime() *BACnetApplicationTagSignedInteger
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

func (m *BACnetConstructedDataVerificationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataVerificationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VERIFICATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataVerificationTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataVerificationTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataVerificationTime) GetVerificationTime() *BACnetApplicationTagSignedInteger {
	return m.VerificationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVerificationTime factory function for BACnetConstructedDataVerificationTime
func NewBACnetConstructedDataVerificationTime(verificationTime *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataVerificationTime {
	_result := &BACnetConstructedDataVerificationTime{
		VerificationTime:      verificationTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataVerificationTime(structType interface{}) *BACnetConstructedDataVerificationTime {
	if casted, ok := structType.(BACnetConstructedDataVerificationTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVerificationTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataVerificationTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataVerificationTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataVerificationTime) GetTypeName() string {
	return "BACnetConstructedDataVerificationTime"
}

func (m *BACnetConstructedDataVerificationTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataVerificationTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (verificationTime)
	lengthInBits += m.VerificationTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataVerificationTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVerificationTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataVerificationTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVerificationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVerificationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (verificationTime)
	if pullErr := readBuffer.PullContext("verificationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for verificationTime")
	}
	_verificationTime, _verificationTimeErr := BACnetApplicationTagParse(readBuffer)
	if _verificationTimeErr != nil {
		return nil, errors.Wrap(_verificationTimeErr, "Error parsing 'verificationTime' field")
	}
	verificationTime := CastBACnetApplicationTagSignedInteger(_verificationTime)
	if closeErr := readBuffer.CloseContext("verificationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for verificationTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVerificationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVerificationTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataVerificationTime{
		VerificationTime:      CastBACnetApplicationTagSignedInteger(verificationTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataVerificationTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVerificationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVerificationTime")
		}

		// Simple Field (verificationTime)
		if pushErr := writeBuffer.PushContext("verificationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for verificationTime")
		}
		_verificationTimeErr := m.VerificationTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("verificationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for verificationTime")
		}
		if _verificationTimeErr != nil {
			return errors.Wrap(_verificationTimeErr, "Error serializing 'verificationTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVerificationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVerificationTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataVerificationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
