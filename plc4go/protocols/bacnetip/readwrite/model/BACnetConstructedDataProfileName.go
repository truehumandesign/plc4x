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

// BACnetConstructedDataProfileName is the data-structure of this message
type BACnetConstructedDataProfileName struct {
	*BACnetConstructedData
	ProfileName *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataProfileName is the corresponding interface of BACnetConstructedDataProfileName
type IBACnetConstructedDataProfileName interface {
	IBACnetConstructedData
	// GetProfileName returns ProfileName (property field)
	GetProfileName() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataProfileName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataProfileName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROFILE_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataProfileName) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataProfileName) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataProfileName) GetProfileName() *BACnetApplicationTagCharacterString {
	return m.ProfileName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProfileName factory function for BACnetConstructedDataProfileName
func NewBACnetConstructedDataProfileName(profileName *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataProfileName {
	_result := &BACnetConstructedDataProfileName{
		ProfileName:           profileName,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataProfileName(structType interface{}) *BACnetConstructedDataProfileName {
	if casted, ok := structType.(BACnetConstructedDataProfileName); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProfileName); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataProfileName(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataProfileName(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataProfileName) GetTypeName() string {
	return "BACnetConstructedDataProfileName"
}

func (m *BACnetConstructedDataProfileName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataProfileName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (profileName)
	lengthInBits += m.ProfileName.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataProfileName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProfileNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataProfileName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProfileName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProfileName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (profileName)
	if pullErr := readBuffer.PullContext("profileName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for profileName")
	}
	_profileName, _profileNameErr := BACnetApplicationTagParse(readBuffer)
	if _profileNameErr != nil {
		return nil, errors.Wrap(_profileNameErr, "Error parsing 'profileName' field")
	}
	profileName := CastBACnetApplicationTagCharacterString(_profileName)
	if closeErr := readBuffer.CloseContext("profileName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for profileName")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProfileName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProfileName")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataProfileName{
		ProfileName:           CastBACnetApplicationTagCharacterString(profileName),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataProfileName) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProfileName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProfileName")
		}

		// Simple Field (profileName)
		if pushErr := writeBuffer.PushContext("profileName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for profileName")
		}
		_profileNameErr := m.ProfileName.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("profileName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for profileName")
		}
		if _profileNameErr != nil {
			return errors.Wrap(_profileNameErr, "Error serializing 'profileName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProfileName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProfileName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataProfileName) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
