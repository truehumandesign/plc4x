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
type DIBSuppSvcFamilies struct {
	DescriptionType uint8
	ServiceIds      []*ServiceId
}

// The corresponding interface
type IDIBSuppSvcFamilies interface {
	// GetDescriptionType returns DescriptionType (property field)
	GetDescriptionType() uint8
	// GetServiceIds returns ServiceIds (property field)
	GetServiceIds() []*ServiceId
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *DIBSuppSvcFamilies) GetDescriptionType() uint8 {
	return m.DescriptionType
}

func (m *DIBSuppSvcFamilies) GetServiceIds() []*ServiceId {
	return m.ServiceIds
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewDIBSuppSvcFamilies factory function for DIBSuppSvcFamilies
func NewDIBSuppSvcFamilies(descriptionType uint8, serviceIds []*ServiceId) *DIBSuppSvcFamilies {
	return &DIBSuppSvcFamilies{DescriptionType: descriptionType, ServiceIds: serviceIds}
}

func CastDIBSuppSvcFamilies(structType interface{}) *DIBSuppSvcFamilies {
	if casted, ok := structType.(DIBSuppSvcFamilies); ok {
		return &casted
	}
	if casted, ok := structType.(*DIBSuppSvcFamilies); ok {
		return casted
	}
	return nil
}

func (m *DIBSuppSvcFamilies) GetTypeName() string {
	return "DIBSuppSvcFamilies"
}

func (m *DIBSuppSvcFamilies) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DIBSuppSvcFamilies) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (descriptionType)
	lengthInBits += 8

	// Array field
	if len(m.ServiceIds) > 0 {
		for _, element := range m.ServiceIds {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *DIBSuppSvcFamilies) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DIBSuppSvcFamiliesParse(readBuffer utils.ReadBuffer) (*DIBSuppSvcFamilies, error) {
	if pullErr := readBuffer.PullContext("DIBSuppSvcFamilies"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (descriptionType)
	_descriptionType, _descriptionTypeErr := readBuffer.ReadUint8("descriptionType", 8)
	if _descriptionTypeErr != nil {
		return nil, errors.Wrap(_descriptionTypeErr, "Error parsing 'descriptionType' field")
	}
	descriptionType := _descriptionType

	// Array field (serviceIds)
	if pullErr := readBuffer.PullContext("serviceIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	serviceIds := make([]*ServiceId, 0)
	{
		_serviceIdsLength := uint16(structureLength) - uint16(uint16(2))
		_serviceIdsEndPos := readBuffer.GetPos() + uint16(_serviceIdsLength)
		for readBuffer.GetPos() < _serviceIdsEndPos {
			_item, _err := ServiceIdParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'serviceIds' field")
			}
			serviceIds = append(serviceIds, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("serviceIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DIBSuppSvcFamilies"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDIBSuppSvcFamilies(descriptionType, serviceIds), nil
}

func (m *DIBSuppSvcFamilies) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("DIBSuppSvcFamilies"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (descriptionType)
	descriptionType := uint8(m.DescriptionType)
	_descriptionTypeErr := writeBuffer.WriteUint8("descriptionType", 8, (descriptionType))
	if _descriptionTypeErr != nil {
		return errors.Wrap(_descriptionTypeErr, "Error serializing 'descriptionType' field")
	}

	// Array Field (serviceIds)
	if m.ServiceIds != nil {
		if pushErr := writeBuffer.PushContext("serviceIds", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.ServiceIds {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'serviceIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("serviceIds", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("DIBSuppSvcFamilies"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DIBSuppSvcFamilies) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
