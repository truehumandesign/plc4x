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

// BACnetConstructedDataNetworkNumberQuality is the data-structure of this message
type BACnetConstructedDataNetworkNumberQuality struct {
	*BACnetConstructedData
	NetworkNumberQuality *BACnetNetworkNumberQualityTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataNetworkNumberQuality is the corresponding interface of BACnetConstructedDataNetworkNumberQuality
type IBACnetConstructedDataNetworkNumberQuality interface {
	IBACnetConstructedData
	// GetNetworkNumberQuality returns NetworkNumberQuality (property field)
	GetNetworkNumberQuality() *BACnetNetworkNumberQualityTagged
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

func (m *BACnetConstructedDataNetworkNumberQuality) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataNetworkNumberQuality) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_NUMBER_QUALITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataNetworkNumberQuality) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataNetworkNumberQuality) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataNetworkNumberQuality) GetNetworkNumberQuality() *BACnetNetworkNumberQualityTagged {
	return m.NetworkNumberQuality
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkNumberQuality factory function for BACnetConstructedDataNetworkNumberQuality
func NewBACnetConstructedDataNetworkNumberQuality(networkNumberQuality *BACnetNetworkNumberQualityTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataNetworkNumberQuality {
	_result := &BACnetConstructedDataNetworkNumberQuality{
		NetworkNumberQuality:  networkNumberQuality,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataNetworkNumberQuality(structType interface{}) *BACnetConstructedDataNetworkNumberQuality {
	if casted, ok := structType.(BACnetConstructedDataNetworkNumberQuality); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkNumberQuality); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkNumberQuality(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkNumberQuality(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataNetworkNumberQuality) GetTypeName() string {
	return "BACnetConstructedDataNetworkNumberQuality"
}

func (m *BACnetConstructedDataNetworkNumberQuality) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataNetworkNumberQuality) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkNumberQuality)
	lengthInBits += m.NetworkNumberQuality.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataNetworkNumberQuality) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNetworkNumberQualityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataNetworkNumberQuality, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkNumberQuality"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkNumberQuality")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkNumberQuality)
	if pullErr := readBuffer.PullContext("networkNumberQuality"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkNumberQuality")
	}
	_networkNumberQuality, _networkNumberQualityErr := BACnetNetworkNumberQualityTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _networkNumberQualityErr != nil {
		return nil, errors.Wrap(_networkNumberQualityErr, "Error parsing 'networkNumberQuality' field")
	}
	networkNumberQuality := CastBACnetNetworkNumberQualityTagged(_networkNumberQuality)
	if closeErr := readBuffer.CloseContext("networkNumberQuality"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkNumberQuality")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkNumberQuality"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkNumberQuality")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataNetworkNumberQuality{
		NetworkNumberQuality:  CastBACnetNetworkNumberQualityTagged(networkNumberQuality),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataNetworkNumberQuality) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkNumberQuality"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkNumberQuality")
		}

		// Simple Field (networkNumberQuality)
		if pushErr := writeBuffer.PushContext("networkNumberQuality"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkNumberQuality")
		}
		_networkNumberQualityErr := m.NetworkNumberQuality.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("networkNumberQuality"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkNumberQuality")
		}
		if _networkNumberQualityErr != nil {
			return errors.Wrap(_networkNumberQualityErr, "Error serializing 'networkNumberQuality' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkNumberQuality"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkNumberQuality")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataNetworkNumberQuality) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
