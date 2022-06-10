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

// BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer is the data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer struct {
	*BACnetUnconfirmedServiceRequest
	VendorId          *BACnetVendorIdTagged
	ServiceNumber     *BACnetContextTagUnsignedInteger
	ServiceParameters *BACnetConstructedData

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
type IBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer interface {
	IBACnetUnconfirmedServiceRequest
	// GetVendorId returns VendorId (property field)
	GetVendorId() *BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() *BACnetContextTagUnsignedInteger
	// GetServiceParameters returns ServiceParameters (property field)
	GetServiceParameters() *BACnetConstructedData
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

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetVendorId() *BACnetVendorIdTagged {
	return m.VendorId
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceNumber() *BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceParameters() *BACnetConstructedData {
	return m.ServiceParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer factory function for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
func NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(vendorId *BACnetVendorIdTagged, serviceNumber *BACnetContextTagUnsignedInteger, serviceParameters *BACnetConstructedData, serviceRequestLength uint16) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	_result := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:                        vendorId,
		ServiceNumber:                   serviceNumber,
		ServiceParameters:               serviceParameters,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits()

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits()

	// Optional Field (serviceParameters)
	if m.ServiceParameters != nil {
		lengthInBits += (*m.ServiceParameters).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}
	vendorId := CastBACnetVendorIdTagged(_vendorId)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}

	// Simple Field (serviceNumber)
	if pullErr := readBuffer.PullContext("serviceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceNumber")
	}
	_serviceNumber, _serviceNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field")
	}
	serviceNumber := CastBACnetContextTagUnsignedInteger(_serviceNumber)
	if closeErr := readBuffer.CloseContext("serviceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceNumber")
	}

	// Optional Field (serviceParameters) (Can be skipped, if a given expression evaluates to false)
	var serviceParameters *BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("serviceParameters"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for serviceParameters")
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(2), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'serviceParameters' field")
		default:
			serviceParameters = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("serviceParameters"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for serviceParameters")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:                        CastBACnetVendorIdTagged(vendorId),
		ServiceNumber:                   CastBACnetContextTagUnsignedInteger(serviceNumber),
		ServiceParameters:               CastBACnetConstructedData(serviceParameters),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := m.VendorId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (serviceNumber)
		if pushErr := writeBuffer.PushContext("serviceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serviceNumber")
		}
		_serviceNumberErr := m.ServiceNumber.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("serviceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serviceNumber")
		}
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Optional Field (serviceParameters) (Can be skipped, if the value is null)
		var serviceParameters *BACnetConstructedData = nil
		if m.ServiceParameters != nil {
			if pushErr := writeBuffer.PushContext("serviceParameters"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for serviceParameters")
			}
			serviceParameters = m.ServiceParameters
			_serviceParametersErr := serviceParameters.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("serviceParameters"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for serviceParameters")
			}
			if _serviceParametersErr != nil {
				return errors.Wrap(_serviceParametersErr, "Error serializing 'serviceParameters' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
