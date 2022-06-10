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

// BACnetAssignedAccessRights is the data-structure of this message
type BACnetAssignedAccessRights struct {
	AssignedAccessRights *BACnetDeviceObjectReferenceEnclosed
	Enable               *BACnetContextTagBoolean
}

// IBACnetAssignedAccessRights is the corresponding interface of BACnetAssignedAccessRights
type IBACnetAssignedAccessRights interface {
	// GetAssignedAccessRights returns AssignedAccessRights (property field)
	GetAssignedAccessRights() *BACnetDeviceObjectReferenceEnclosed
	// GetEnable returns Enable (property field)
	GetEnable() *BACnetContextTagBoolean
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

func (m *BACnetAssignedAccessRights) GetAssignedAccessRights() *BACnetDeviceObjectReferenceEnclosed {
	return m.AssignedAccessRights
}

func (m *BACnetAssignedAccessRights) GetEnable() *BACnetContextTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAssignedAccessRights factory function for BACnetAssignedAccessRights
func NewBACnetAssignedAccessRights(assignedAccessRights *BACnetDeviceObjectReferenceEnclosed, enable *BACnetContextTagBoolean) *BACnetAssignedAccessRights {
	return &BACnetAssignedAccessRights{AssignedAccessRights: assignedAccessRights, Enable: enable}
}

func CastBACnetAssignedAccessRights(structType interface{}) *BACnetAssignedAccessRights {
	if casted, ok := structType.(BACnetAssignedAccessRights); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetAssignedAccessRights); ok {
		return casted
	}
	return nil
}

func (m *BACnetAssignedAccessRights) GetTypeName() string {
	return "BACnetAssignedAccessRights"
}

func (m *BACnetAssignedAccessRights) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetAssignedAccessRights) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (assignedAccessRights)
	lengthInBits += m.AssignedAccessRights.GetLengthInBits()

	// Simple field (enable)
	lengthInBits += m.Enable.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetAssignedAccessRights) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAssignedAccessRightsParse(readBuffer utils.ReadBuffer) (*BACnetAssignedAccessRights, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedAccessRights"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedAccessRights")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (assignedAccessRights)
	if pullErr := readBuffer.PullContext("assignedAccessRights"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for assignedAccessRights")
	}
	_assignedAccessRights, _assignedAccessRightsErr := BACnetDeviceObjectReferenceEnclosedParse(readBuffer, uint8(uint8(0)))
	if _assignedAccessRightsErr != nil {
		return nil, errors.Wrap(_assignedAccessRightsErr, "Error parsing 'assignedAccessRights' field")
	}
	assignedAccessRights := CastBACnetDeviceObjectReferenceEnclosed(_assignedAccessRights)
	if closeErr := readBuffer.CloseContext("assignedAccessRights"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for assignedAccessRights")
	}

	// Simple Field (enable)
	if pullErr := readBuffer.PullContext("enable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enable")
	}
	_enable, _enableErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _enableErr != nil {
		return nil, errors.Wrap(_enableErr, "Error parsing 'enable' field")
	}
	enable := CastBACnetContextTagBoolean(_enable)
	if closeErr := readBuffer.CloseContext("enable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enable")
	}

	if closeErr := readBuffer.CloseContext("BACnetAssignedAccessRights"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedAccessRights")
	}

	// Create the instance
	return NewBACnetAssignedAccessRights(assignedAccessRights, enable), nil
}

func (m *BACnetAssignedAccessRights) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAssignedAccessRights"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedAccessRights")
	}

	// Simple Field (assignedAccessRights)
	if pushErr := writeBuffer.PushContext("assignedAccessRights"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for assignedAccessRights")
	}
	_assignedAccessRightsErr := m.AssignedAccessRights.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("assignedAccessRights"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for assignedAccessRights")
	}
	if _assignedAccessRightsErr != nil {
		return errors.Wrap(_assignedAccessRightsErr, "Error serializing 'assignedAccessRights' field")
	}

	// Simple Field (enable)
	if pushErr := writeBuffer.PushContext("enable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for enable")
	}
	_enableErr := m.Enable.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("enable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for enable")
	}
	if _enableErr != nil {
		return errors.Wrap(_enableErr, "Error serializing 'enable' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedAccessRights"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedAccessRights")
	}
	return nil
}

func (m *BACnetAssignedAccessRights) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
