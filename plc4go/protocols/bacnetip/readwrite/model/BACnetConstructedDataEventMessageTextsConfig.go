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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEventMessageTextsConfig is the data-structure of this message
type BACnetConstructedDataEventMessageTextsConfig struct {
	*BACnetConstructedData
	NumberOfDataElements    *BACnetApplicationTagUnsignedInteger
	EventMessageTextsConfig []*BACnetOptionalCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEventMessageTextsConfig is the corresponding interface of BACnetConstructedDataEventMessageTextsConfig
type IBACnetConstructedDataEventMessageTextsConfig interface {
	IBACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger
	// GetEventMessageTextsConfig returns EventMessageTextsConfig (property field)
	GetEventMessageTextsConfig() []*BACnetOptionalCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetToOffnormalTextConfig returns ToOffnormalTextConfig (virtual field)
	GetToOffnormalTextConfig() *BACnetOptionalCharacterString
	// GetToFaultTextConfig returns ToFaultTextConfig (virtual field)
	GetToFaultTextConfig() *BACnetOptionalCharacterString
	// GetToNormalTextConfig returns ToNormalTextConfig (virtual field)
	GetToNormalTextConfig() *BACnetOptionalCharacterString
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

func (m *BACnetConstructedDataEventMessageTextsConfig) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_MESSAGE_TEXTS_CONFIG
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEventMessageTextsConfig) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEventMessageTextsConfig) GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetEventMessageTextsConfig() []*BACnetOptionalCharacterString {
	return m.EventMessageTextsConfig
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataEventMessageTextsConfig) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetToOffnormalTextConfig() *BACnetOptionalCharacterString {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTextsConfig())) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(m.GetEventMessageTextsConfig()[0]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetToFaultTextConfig() *BACnetOptionalCharacterString {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTextsConfig())) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(m.GetEventMessageTextsConfig()[1]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetToNormalTextConfig() *BACnetOptionalCharacterString {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTextsConfig())) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(m.GetEventMessageTextsConfig()[2]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventMessageTextsConfig factory function for BACnetConstructedDataEventMessageTextsConfig
func NewBACnetConstructedDataEventMessageTextsConfig(numberOfDataElements *BACnetApplicationTagUnsignedInteger, eventMessageTextsConfig []*BACnetOptionalCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEventMessageTextsConfig {
	_result := &BACnetConstructedDataEventMessageTextsConfig{
		NumberOfDataElements:    numberOfDataElements,
		EventMessageTextsConfig: eventMessageTextsConfig,
		BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEventMessageTextsConfig(structType interface{}) *BACnetConstructedDataEventMessageTextsConfig {
	if casted, ok := structType.(BACnetConstructedDataEventMessageTextsConfig); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventMessageTextsConfig); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventMessageTextsConfig(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventMessageTextsConfig(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetTypeName() string {
	return "BACnetConstructedDataEventMessageTextsConfig"
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += (*m.NumberOfDataElements).GetLengthInBits()
	}

	// Array field
	if len(m.EventMessageTextsConfig) > 0 {
		for _, element := range m.EventMessageTextsConfig {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataEventMessageTextsConfig) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventMessageTextsConfigParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEventMessageTextsConfig, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventMessageTextsConfig"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventMessageTextsConfig")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field")
		default:
			numberOfDataElements = CastBACnetApplicationTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (eventMessageTextsConfig)
	if pullErr := readBuffer.PullContext("eventMessageTextsConfig", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventMessageTextsConfig")
	}
	// Terminated array
	eventMessageTextsConfig := make([]*BACnetOptionalCharacterString, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetOptionalCharacterStringParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'eventMessageTextsConfig' field")
			}
			eventMessageTextsConfig = append(eventMessageTextsConfig, CastBACnetOptionalCharacterString(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("eventMessageTextsConfig", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventMessageTextsConfig")
	}

	// Virtual field
	_toOffnormalTextConfig := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTextsConfig)) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(eventMessageTextsConfig[0]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) }))
	toOffnormalTextConfig := CastBACnetOptionalCharacterString(_toOffnormalTextConfig)
	_ = toOffnormalTextConfig

	// Virtual field
	_toFaultTextConfig := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTextsConfig)) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(eventMessageTextsConfig[1]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) }))
	toFaultTextConfig := CastBACnetOptionalCharacterString(_toFaultTextConfig)
	_ = toFaultTextConfig

	// Virtual field
	_toNormalTextConfig := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTextsConfig)) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(eventMessageTextsConfig[2]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) }))
	toNormalTextConfig := CastBACnetOptionalCharacterString(_toNormalTextConfig)
	_ = toNormalTextConfig

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(eventMessageTextsConfig)) == (3)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"eventMessageTextsConfig should have exactly 3 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventMessageTextsConfig"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventMessageTextsConfig")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventMessageTextsConfig{
		NumberOfDataElements:    CastBACnetApplicationTagUnsignedInteger(numberOfDataElements),
		EventMessageTextsConfig: eventMessageTextsConfig,
		BACnetConstructedData:   &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEventMessageTextsConfig) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventMessageTextsConfig"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventMessageTextsConfig")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
		if m.NumberOfDataElements != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.NumberOfDataElements
			_numberOfDataElementsErr := numberOfDataElements.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (eventMessageTextsConfig)
		if m.EventMessageTextsConfig != nil {
			if pushErr := writeBuffer.PushContext("eventMessageTextsConfig", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for eventMessageTextsConfig")
			}
			for _, _element := range m.EventMessageTextsConfig {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'eventMessageTextsConfig' field")
				}
			}
			if popErr := writeBuffer.PopContext("eventMessageTextsConfig", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for eventMessageTextsConfig")
			}
		}
		// Virtual field
		if _toOffnormalTextConfigErr := writeBuffer.WriteVirtual("toOffnormalTextConfig", m.GetToOffnormalTextConfig()); _toOffnormalTextConfigErr != nil {
			return errors.Wrap(_toOffnormalTextConfigErr, "Error serializing 'toOffnormalTextConfig' field")
		}
		// Virtual field
		if _toFaultTextConfigErr := writeBuffer.WriteVirtual("toFaultTextConfig", m.GetToFaultTextConfig()); _toFaultTextConfigErr != nil {
			return errors.Wrap(_toFaultTextConfigErr, "Error serializing 'toFaultTextConfig' field")
		}
		// Virtual field
		if _toNormalTextConfigErr := writeBuffer.WriteVirtual("toNormalTextConfig", m.GetToNormalTextConfig()); _toNormalTextConfigErr != nil {
			return errors.Wrap(_toNormalTextConfigErr, "Error serializing 'toNormalTextConfig' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventMessageTextsConfig"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventMessageTextsConfig")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventMessageTextsConfig) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
