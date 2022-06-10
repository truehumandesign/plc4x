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

// BACnetConstructedDataCredentialDisable is the data-structure of this message
type BACnetConstructedDataCredentialDisable struct {
	*BACnetConstructedData
	CredentialDisable *BACnetAccessCredentialDisableTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCredentialDisable is the corresponding interface of BACnetConstructedDataCredentialDisable
type IBACnetConstructedDataCredentialDisable interface {
	IBACnetConstructedData
	// GetCredentialDisable returns CredentialDisable (property field)
	GetCredentialDisable() *BACnetAccessCredentialDisableTagged
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

func (m *BACnetConstructedDataCredentialDisable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCredentialDisable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIAL_DISABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCredentialDisable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCredentialDisable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCredentialDisable) GetCredentialDisable() *BACnetAccessCredentialDisableTagged {
	return m.CredentialDisable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialDisable factory function for BACnetConstructedDataCredentialDisable
func NewBACnetConstructedDataCredentialDisable(credentialDisable *BACnetAccessCredentialDisableTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCredentialDisable {
	_result := &BACnetConstructedDataCredentialDisable{
		CredentialDisable:     credentialDisable,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCredentialDisable(structType interface{}) *BACnetConstructedDataCredentialDisable {
	if casted, ok := structType.(BACnetConstructedDataCredentialDisable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialDisable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCredentialDisable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCredentialDisable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCredentialDisable) GetTypeName() string {
	return "BACnetConstructedDataCredentialDisable"
}

func (m *BACnetConstructedDataCredentialDisable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCredentialDisable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (credentialDisable)
	lengthInBits += m.CredentialDisable.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataCredentialDisable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCredentialDisableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCredentialDisable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialDisable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (credentialDisable)
	if pullErr := readBuffer.PullContext("credentialDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for credentialDisable")
	}
	_credentialDisable, _credentialDisableErr := BACnetAccessCredentialDisableTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _credentialDisableErr != nil {
		return nil, errors.Wrap(_credentialDisableErr, "Error parsing 'credentialDisable' field")
	}
	credentialDisable := CastBACnetAccessCredentialDisableTagged(_credentialDisable)
	if closeErr := readBuffer.CloseContext("credentialDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for credentialDisable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialDisable")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCredentialDisable{
		CredentialDisable:     CastBACnetAccessCredentialDisableTagged(credentialDisable),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCredentialDisable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialDisable")
		}

		// Simple Field (credentialDisable)
		if pushErr := writeBuffer.PushContext("credentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for credentialDisable")
		}
		_credentialDisableErr := m.CredentialDisable.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("credentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for credentialDisable")
		}
		if _credentialDisableErr != nil {
			return errors.Wrap(_credentialDisableErr, "Error serializing 'credentialDisable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialDisable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCredentialDisable) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
