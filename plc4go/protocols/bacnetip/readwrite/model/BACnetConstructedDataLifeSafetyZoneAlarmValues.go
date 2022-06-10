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

// BACnetConstructedDataLifeSafetyZoneAlarmValues is the data-structure of this message
type BACnetConstructedDataLifeSafetyZoneAlarmValues struct {
	*BACnetConstructedData
	AlarmValues []*BACnetLifeSafetyStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLifeSafetyZoneAlarmValues is the corresponding interface of BACnetConstructedDataLifeSafetyZoneAlarmValues
type IBACnetConstructedDataLifeSafetyZoneAlarmValues interface {
	IBACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []*BACnetLifeSafetyStateTagged
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

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetAlarmValues() []*BACnetLifeSafetyStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyZoneAlarmValues factory function for BACnetConstructedDataLifeSafetyZoneAlarmValues
func NewBACnetConstructedDataLifeSafetyZoneAlarmValues(alarmValues []*BACnetLifeSafetyStateTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLifeSafetyZoneAlarmValues {
	_result := &BACnetConstructedDataLifeSafetyZoneAlarmValues{
		AlarmValues:           alarmValues,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLifeSafetyZoneAlarmValues(structType interface{}) *BACnetConstructedDataLifeSafetyZoneAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZoneAlarmValues); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZoneAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyZoneAlarmValues(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyZoneAlarmValues(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZoneAlarmValues"
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyZoneAlarmValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLifeSafetyZoneAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZoneAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (alarmValues)
	if pullErr := readBuffer.PullContext("alarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmValues")
	}
	// Terminated array
	alarmValues := make([]*BACnetLifeSafetyStateTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'alarmValues' field")
			}
			alarmValues = append(alarmValues, CastBACnetLifeSafetyStateTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("alarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZoneAlarmValues")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLifeSafetyZoneAlarmValues{
		AlarmValues:           alarmValues,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZoneAlarmValues")
		}

		// Array Field (alarmValues)
		if m.AlarmValues != nil {
			if pushErr := writeBuffer.PushContext("alarmValues", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for alarmValues")
			}
			for _, _element := range m.AlarmValues {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'alarmValues' field")
				}
			}
			if popErr := writeBuffer.PopContext("alarmValues", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for alarmValues")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZoneAlarmValues")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLifeSafetyZoneAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
