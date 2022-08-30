/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataControlGroups is the corresponding interface of BACnetConstructedDataControlGroups
type BACnetConstructedDataControlGroups interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetControlGroups returns ControlGroups (property field)
	GetControlGroups() []BACnetApplicationTagUnsignedInteger
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataControlGroupsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataControlGroups.
// This is useful for switch cases.
type BACnetConstructedDataControlGroupsExactly interface {
	BACnetConstructedDataControlGroups
	isBACnetConstructedDataControlGroups() bool
}

// _BACnetConstructedDataControlGroups is the data-structure of this message
type _BACnetConstructedDataControlGroups struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ControlGroups        []BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlGroups) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlGroups) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROL_GROUPS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlGroups) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataControlGroups) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlGroups) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataControlGroups) GetControlGroups() []BACnetApplicationTagUnsignedInteger {
	return m.ControlGroups
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlGroups) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataControlGroups factory function for _BACnetConstructedDataControlGroups
func NewBACnetConstructedDataControlGroups(numberOfDataElements BACnetApplicationTagUnsignedInteger, controlGroups []BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlGroups {
	_result := &_BACnetConstructedDataControlGroups{
		NumberOfDataElements:   numberOfDataElements,
		ControlGroups:          controlGroups,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataControlGroups(structType interface{}) BACnetConstructedDataControlGroups {
	if casted, ok := structType.(BACnetConstructedDataControlGroups); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlGroups); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlGroups) GetTypeName() string {
	return "BACnetConstructedDataControlGroups"
}

func (m *_BACnetConstructedDataControlGroups) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataControlGroups) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.ControlGroups) > 0 {
		for _, element := range m.ControlGroups {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataControlGroups) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataControlGroupsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataControlGroups, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlGroups"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlGroups")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataControlGroups")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (controlGroups)
	if pullErr := readBuffer.PullContext("controlGroups", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for controlGroups")
	}
	// Terminated array
	var controlGroups []BACnetApplicationTagUnsignedInteger
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'controlGroups' field of BACnetConstructedDataControlGroups")
			}
			controlGroups = append(controlGroups, _item.(BACnetApplicationTagUnsignedInteger))

		}
	}
	if closeErr := readBuffer.CloseContext("controlGroups", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for controlGroups")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlGroups"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlGroups")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataControlGroups{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		ControlGroups:        controlGroups,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataControlGroups) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlGroups"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlGroups")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (controlGroups)
		if pushErr := writeBuffer.PushContext("controlGroups", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for controlGroups")
		}
		for _, _element := range m.GetControlGroups() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'controlGroups' field")
			}
		}
		if popErr := writeBuffer.PopContext("controlGroups", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for controlGroups")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlGroups"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlGroups")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlGroups) isBACnetConstructedDataControlGroups() bool {
	return true
}

func (m *_BACnetConstructedDataControlGroups) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
