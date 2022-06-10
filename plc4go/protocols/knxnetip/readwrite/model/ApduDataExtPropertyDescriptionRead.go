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

// ApduDataExtPropertyDescriptionRead is the data-structure of this message
type ApduDataExtPropertyDescriptionRead struct {
	*ApduDataExt
	ObjectIndex uint8
	PropertyId  uint8
	Index       uint8

	// Arguments.
	Length uint8
}

// IApduDataExtPropertyDescriptionRead is the corresponding interface of ApduDataExtPropertyDescriptionRead
type IApduDataExtPropertyDescriptionRead interface {
	IApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint8
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

func (m *ApduDataExtPropertyDescriptionRead) GetExtApciType() uint8 {
	return 0x18
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtPropertyDescriptionRead) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtPropertyDescriptionRead) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ApduDataExtPropertyDescriptionRead) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *ApduDataExtPropertyDescriptionRead) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *ApduDataExtPropertyDescriptionRead) GetIndex() uint8 {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyDescriptionRead factory function for ApduDataExtPropertyDescriptionRead
func NewApduDataExtPropertyDescriptionRead(objectIndex uint8, propertyId uint8, index uint8, length uint8) *ApduDataExtPropertyDescriptionRead {
	_result := &ApduDataExtPropertyDescriptionRead{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Index:       index,
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtPropertyDescriptionRead(structType interface{}) *ApduDataExtPropertyDescriptionRead {
	if casted, ok := structType.(ApduDataExtPropertyDescriptionRead); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyDescriptionRead); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtPropertyDescriptionRead(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtPropertyDescriptionRead(casted.Child)
	}
	return nil
}

func (m *ApduDataExtPropertyDescriptionRead) GetTypeName() string {
	return "ApduDataExtPropertyDescriptionRead"
}

func (m *ApduDataExtPropertyDescriptionRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtPropertyDescriptionRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (index)
	lengthInBits += 8

	return lengthInBits
}

func (m *ApduDataExtPropertyDescriptionRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtPropertyDescriptionReadParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtPropertyDescriptionRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyDescriptionRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyDescriptionRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIndex)
	_objectIndex, _objectIndexErr := readBuffer.ReadUint8("objectIndex", 8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field")
	}
	objectIndex := _objectIndex

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}
	propertyId := _propertyId

	// Simple Field (index)
	_index, _indexErr := readBuffer.ReadUint8("index", 8)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}
	index := _index

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyDescriptionRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyDescriptionRead")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtPropertyDescriptionRead{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Index:       index,
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtPropertyDescriptionRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyDescriptionRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyDescriptionRead")
		}

		// Simple Field (objectIndex)
		objectIndex := uint8(m.ObjectIndex)
		_objectIndexErr := writeBuffer.WriteUint8("objectIndex", 8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (index)
		index := uint8(m.Index)
		_indexErr := writeBuffer.WriteUint8("index", 8, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyDescriptionRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyDescriptionRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtPropertyDescriptionRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
