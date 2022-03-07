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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type S7ParameterUserDataItem struct {
	Child IS7ParameterUserDataItemChild
}

// The corresponding interface
type IS7ParameterUserDataItem interface {
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7ParameterUserDataItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7ParameterUserDataItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7ParameterUserDataItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7ParameterUserDataItem)
	GetTypeName() string
	IS7ParameterUserDataItem
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewS7ParameterUserDataItem factory function for S7ParameterUserDataItem
func NewS7ParameterUserDataItem() *S7ParameterUserDataItem {
	return &S7ParameterUserDataItem{}
}

func CastS7ParameterUserDataItem(structType interface{}) *S7ParameterUserDataItem {
	if casted, ok := structType.(S7ParameterUserDataItem); ok {
		return &casted
	}
	if casted, ok := structType.(*S7ParameterUserDataItem); ok {
		return casted
	}
	return nil
}

func (m *S7ParameterUserDataItem) GetTypeName() string {
	return "S7ParameterUserDataItem"
}

func (m *S7ParameterUserDataItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7ParameterUserDataItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *S7ParameterUserDataItem) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7ParameterUserDataItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterUserDataItemParse(readBuffer utils.ReadBuffer) (*S7ParameterUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7ParameterUserDataItem"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType, _itemTypeErr := readBuffer.ReadUint8("itemType", 8)
	if _itemTypeErr != nil {
		return nil, errors.Wrap(_itemTypeErr, "Error parsing 'itemType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7ParameterUserDataItem
	var typeSwitchError error
	switch {
	case itemType == 0x12: // S7ParameterUserDataItemCPUFunctions
		_parent, typeSwitchError = S7ParameterUserDataItemCPUFunctionsParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7ParameterUserDataItem"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *S7ParameterUserDataItem) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7ParameterUserDataItem) SerializeParent(writeBuffer utils.WriteBuffer, child IS7ParameterUserDataItem, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7ParameterUserDataItem"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType := uint8(child.GetItemType())
	_itemTypeErr := writeBuffer.WriteUint8("itemType", 8, (itemType))

	if _itemTypeErr != nil {
		return errors.Wrap(_itemTypeErr, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7ParameterUserDataItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7ParameterUserDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
