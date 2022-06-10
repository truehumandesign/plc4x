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

// BACnetTimeStampsEnclosed is the data-structure of this message
type BACnetTimeStampsEnclosed struct {
	OpeningTag *BACnetOpeningTag
	Timestamps []*BACnetTimeStamp
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetTimeStampsEnclosed is the corresponding interface of BACnetTimeStampsEnclosed
type IBACnetTimeStampsEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetTimestamps returns Timestamps (property field)
	GetTimestamps() []*BACnetTimeStamp
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetTimeStampsEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetTimeStampsEnclosed) GetTimestamps() []*BACnetTimeStamp {
	return m.Timestamps
}

func (m *BACnetTimeStampsEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampsEnclosed factory function for BACnetTimeStampsEnclosed
func NewBACnetTimeStampsEnclosed(openingTag *BACnetOpeningTag, timestamps []*BACnetTimeStamp, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetTimeStampsEnclosed {
	return &BACnetTimeStampsEnclosed{OpeningTag: openingTag, Timestamps: timestamps, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetTimeStampsEnclosed(structType interface{}) *BACnetTimeStampsEnclosed {
	if casted, ok := structType.(BACnetTimeStampsEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTimeStampsEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetTimeStampsEnclosed) GetTypeName() string {
	return "BACnetTimeStampsEnclosed"
}

func (m *BACnetTimeStampsEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimeStampsEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.Timestamps) > 0 {
		for _, element := range m.Timestamps {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetTimeStampsEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimeStampsEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetTimeStampsEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampsEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampsEnclosed")
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

	// Array field (timestamps)
	if pullErr := readBuffer.PullContext("timestamps", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamps")
	}
	// Terminated array
	timestamps := make([]*BACnetTimeStamp, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetTimeStampParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'timestamps' field")
			}
			timestamps = append(timestamps, CastBACnetTimeStamp(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("timestamps", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamps")
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

	if closeErr := readBuffer.CloseContext("BACnetTimeStampsEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampsEnclosed")
	}

	// Create the instance
	return NewBACnetTimeStampsEnclosed(openingTag, timestamps, closingTag, tagNumber), nil
}

func (m *BACnetTimeStampsEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTimeStampsEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampsEnclosed")
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

	// Array Field (timestamps)
	if m.Timestamps != nil {
		if pushErr := writeBuffer.PushContext("timestamps", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timestamps")
		}
		for _, _element := range m.Timestamps {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'timestamps' field")
			}
		}
		if popErr := writeBuffer.PopContext("timestamps", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timestamps")
		}
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

	if popErr := writeBuffer.PopContext("BACnetTimeStampsEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeStampsEnclosed")
	}
	return nil
}

func (m *BACnetTimeStampsEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
