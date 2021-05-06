//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type LBusmonInd struct {
	AdditionalInformationLength uint8
	AdditionalInformation       []*CEMIAdditionalInformation
	DataFrame                   *LDataFrame
	Crc                         *uint8
	Parent                      *CEMI
}

// The corresponding interface
type ILBusmonInd interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LBusmonInd) MessageCode() uint8 {
	return 0x2B
}

func (m *LBusmonInd) InitializeParent(parent *CEMI) {
}

func NewLBusmonInd(additionalInformationLength uint8, additionalInformation []*CEMIAdditionalInformation, dataFrame *LDataFrame, crc *uint8) *CEMI {
	child := &LBusmonInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Crc:                         crc,
		Parent:                      NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLBusmonInd(structType interface{}) *LBusmonInd {
	castFunc := func(typ interface{}) *LBusmonInd {
		if casted, ok := typ.(LBusmonInd); ok {
			return &casted
		}
		if casted, ok := typ.(*LBusmonInd); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastLBusmonInd(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastLBusmonInd(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LBusmonInd) GetTypeName() string {
	return "LBusmonInd"
}

func (m *LBusmonInd) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LBusmonInd) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.LengthInBits()
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.LengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *LBusmonInd) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LBusmonIndParse(readBuffer utils.ReadBuffer) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("LBusmonInd"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (additionalInformationLength)
	additionalInformationLength, _additionalInformationLengthErr := readBuffer.ReadUint8("additionalInformationLength", 8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field")
	}

	// Array field (additionalInformation)
	if pullErr := readBuffer.PullContext("additionalInformation", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	additionalInformation := make([]*CEMIAdditionalInformation, 0)
	_additionalInformationLength := additionalInformationLength
	_additionalInformationEndPos := readBuffer.GetPos() + uint16(_additionalInformationLength)
	for readBuffer.GetPos() < _additionalInformationEndPos {
		_item, _err := CEMIAdditionalInformationParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field")
		}
		additionalInformation = append(additionalInformation, _item)
	}
	if closeErr := readBuffer.CloseContext("additionalInformation", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("dataFrame"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (dataFrame)
	dataFrame, _dataFrameErr := LDataFrameParse(readBuffer)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field")
	}
	if closeErr := readBuffer.CloseContext("dataFrame"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc *uint8 = nil
	if CastLDataFrame(dataFrame).Child.NotAckFrame() {
		_val, _err := readBuffer.ReadUint8("crc", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'crc' field")
		}
		crc = &_val
	}

	if closeErr := readBuffer.CloseContext("LBusmonInd"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LBusmonInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Crc:                         crc,
		Parent:                      &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LBusmonInd) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LBusmonInd"); pushErr != nil {
			return pushErr
		}

		// Simple Field (additionalInformationLength)
		additionalInformationLength := uint8(m.AdditionalInformationLength)
		_additionalInformationLengthErr := writeBuffer.WriteUint8("additionalInformationLength", 8, (additionalInformationLength))
		if _additionalInformationLengthErr != nil {
			return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
		}

		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			if pushErr := writeBuffer.PushContext("additionalInformation", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.AdditionalInformation {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
				}
			}
			if popErr := writeBuffer.PopContext("additionalInformation", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Simple Field (dataFrame)
		if pushErr := writeBuffer.PushContext("dataFrame"); pushErr != nil {
			return pushErr
		}
		_dataFrameErr := m.DataFrame.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dataFrame"); popErr != nil {
			return popErr
		}
		if _dataFrameErr != nil {
			return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
		}

		// Optional Field (crc) (Can be skipped, if the value is null)
		var crc *uint8 = nil
		if m.Crc != nil {
			crc = m.Crc
			_crcErr := writeBuffer.WriteUint8("crc", 8, *(crc))
			if _crcErr != nil {
				return errors.Wrap(_crcErr, "Error serializing 'crc' field")
			}
		}

		if popErr := writeBuffer.PopContext("LBusmonInd"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
