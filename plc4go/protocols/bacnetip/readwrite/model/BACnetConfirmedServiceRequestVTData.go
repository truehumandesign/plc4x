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

// BACnetConfirmedServiceRequestVTData is the data-structure of this message
type BACnetConfirmedServiceRequestVTData struct {
	*BACnetConfirmedServiceRequest
	VtSessionIdentifier *BACnetApplicationTagUnsignedInteger
	VtNewData           *BACnetApplicationTagOctetString
	VtDataFlag          *BACnetApplicationTagUnsignedInteger

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestVTData is the corresponding interface of BACnetConfirmedServiceRequestVTData
type IBACnetConfirmedServiceRequestVTData interface {
	IBACnetConfirmedServiceRequest
	// GetVtSessionIdentifier returns VtSessionIdentifier (property field)
	GetVtSessionIdentifier() *BACnetApplicationTagUnsignedInteger
	// GetVtNewData returns VtNewData (property field)
	GetVtNewData() *BACnetApplicationTagOctetString
	// GetVtDataFlag returns VtDataFlag (property field)
	GetVtDataFlag() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConfirmedServiceRequestVTData) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_DATA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestVTData) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestVTData) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestVTData) GetVtSessionIdentifier() *BACnetApplicationTagUnsignedInteger {
	return m.VtSessionIdentifier
}

func (m *BACnetConfirmedServiceRequestVTData) GetVtNewData() *BACnetApplicationTagOctetString {
	return m.VtNewData
}

func (m *BACnetConfirmedServiceRequestVTData) GetVtDataFlag() *BACnetApplicationTagUnsignedInteger {
	return m.VtDataFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestVTData factory function for BACnetConfirmedServiceRequestVTData
func NewBACnetConfirmedServiceRequestVTData(vtSessionIdentifier *BACnetApplicationTagUnsignedInteger, vtNewData *BACnetApplicationTagOctetString, vtDataFlag *BACnetApplicationTagUnsignedInteger, serviceRequestLength uint16) *BACnetConfirmedServiceRequestVTData {
	_result := &BACnetConfirmedServiceRequestVTData{
		VtSessionIdentifier:           vtSessionIdentifier,
		VtNewData:                     vtNewData,
		VtDataFlag:                    vtDataFlag,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestVTData(structType interface{}) *BACnetConfirmedServiceRequestVTData {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTData); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTData); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestVTData(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestVTData(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestVTData) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTData"
}

func (m *BACnetConfirmedServiceRequestVTData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestVTData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vtSessionIdentifier)
	lengthInBits += m.VtSessionIdentifier.GetLengthInBits()

	// Simple field (vtNewData)
	lengthInBits += m.VtNewData.GetLengthInBits()

	// Simple field (vtDataFlag)
	lengthInBits += m.VtDataFlag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestVTData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestVTDataParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestVTData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vtSessionIdentifier)
	if pullErr := readBuffer.PullContext("vtSessionIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtSessionIdentifier")
	}
	_vtSessionIdentifier, _vtSessionIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _vtSessionIdentifierErr != nil {
		return nil, errors.Wrap(_vtSessionIdentifierErr, "Error parsing 'vtSessionIdentifier' field")
	}
	vtSessionIdentifier := CastBACnetApplicationTagUnsignedInteger(_vtSessionIdentifier)
	if closeErr := readBuffer.CloseContext("vtSessionIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtSessionIdentifier")
	}

	// Simple Field (vtNewData)
	if pullErr := readBuffer.PullContext("vtNewData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtNewData")
	}
	_vtNewData, _vtNewDataErr := BACnetApplicationTagParse(readBuffer)
	if _vtNewDataErr != nil {
		return nil, errors.Wrap(_vtNewDataErr, "Error parsing 'vtNewData' field")
	}
	vtNewData := CastBACnetApplicationTagOctetString(_vtNewData)
	if closeErr := readBuffer.CloseContext("vtNewData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtNewData")
	}

	// Simple Field (vtDataFlag)
	if pullErr := readBuffer.PullContext("vtDataFlag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtDataFlag")
	}
	_vtDataFlag, _vtDataFlagErr := BACnetApplicationTagParse(readBuffer)
	if _vtDataFlagErr != nil {
		return nil, errors.Wrap(_vtDataFlagErr, "Error parsing 'vtDataFlag' field")
	}
	vtDataFlag := CastBACnetApplicationTagUnsignedInteger(_vtDataFlag)
	if closeErr := readBuffer.CloseContext("vtDataFlag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtDataFlag")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTData")
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestVTData{
		VtSessionIdentifier:           CastBACnetApplicationTagUnsignedInteger(vtSessionIdentifier),
		VtNewData:                     CastBACnetApplicationTagOctetString(vtNewData),
		VtDataFlag:                    CastBACnetApplicationTagUnsignedInteger(vtDataFlag),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestVTData) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTData")
		}

		// Simple Field (vtSessionIdentifier)
		if pushErr := writeBuffer.PushContext("vtSessionIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtSessionIdentifier")
		}
		_vtSessionIdentifierErr := m.VtSessionIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vtSessionIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtSessionIdentifier")
		}
		if _vtSessionIdentifierErr != nil {
			return errors.Wrap(_vtSessionIdentifierErr, "Error serializing 'vtSessionIdentifier' field")
		}

		// Simple Field (vtNewData)
		if pushErr := writeBuffer.PushContext("vtNewData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtNewData")
		}
		_vtNewDataErr := m.VtNewData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vtNewData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtNewData")
		}
		if _vtNewDataErr != nil {
			return errors.Wrap(_vtNewDataErr, "Error serializing 'vtNewData' field")
		}

		// Simple Field (vtDataFlag)
		if pushErr := writeBuffer.PushContext("vtDataFlag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtDataFlag")
		}
		_vtDataFlagErr := m.VtDataFlag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vtDataFlag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtDataFlag")
		}
		if _vtDataFlagErr != nil {
			return errors.Wrap(_vtDataFlagErr, "Error serializing 'vtDataFlag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTData")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestVTData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
