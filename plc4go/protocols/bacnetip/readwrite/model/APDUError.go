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
)

// Code generated by code-generation. DO NOT EDIT.

// APDUError is the data-structure of this message
type APDUError struct {
	*APDU
	OriginalInvokeId uint8
	ErrorChoice      BACnetConfirmedServiceChoice
	Error            *BACnetError

	// Arguments.
	ApduLength uint16
}

// IAPDUError is the corresponding interface of APDUError
type IAPDUError interface {
	IAPDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetErrorChoice returns ErrorChoice (property field)
	GetErrorChoice() BACnetConfirmedServiceChoice
	// GetError returns Error (property field)
	GetError() *BACnetError
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

func (m *APDUError) GetApduType() ApduType {
	return ApduType_ERROR_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *APDUError) InitializeParent(parent *APDU) {}

func (m *APDUError) GetParent() *APDU {
	return m.APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *APDUError) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *APDUError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return m.ErrorChoice
}

func (m *APDUError) GetError() *BACnetError {
	return m.Error
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUError factory function for APDUError
func NewAPDUError(originalInvokeId uint8, errorChoice BACnetConfirmedServiceChoice, error *BACnetError, apduLength uint16) *APDUError {
	_result := &APDUError{
		OriginalInvokeId: originalInvokeId,
		ErrorChoice:      errorChoice,
		Error:            error,
		APDU:             NewAPDU(apduLength),
	}
	_result.Child = _result
	return _result
}

func CastAPDUError(structType interface{}) *APDUError {
	if casted, ok := structType.(APDUError); ok {
		return &casted
	}
	if casted, ok := structType.(*APDUError); ok {
		return casted
	}
	if casted, ok := structType.(APDU); ok {
		return CastAPDUError(casted.Child)
	}
	if casted, ok := structType.(*APDU); ok {
		return CastAPDUError(casted.Child)
	}
	return nil
}

func (m *APDUError) GetTypeName() string {
	return "APDUError"
}

func (m *APDUError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDUError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (errorChoice)
	lengthInBits += 8

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits()

	return lengthInBits
}

func (m *APDUError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUErrorParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDUError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (errorChoice)
	if pullErr := readBuffer.PullContext("errorChoice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorChoice")
	}
	_errorChoice, _errorChoiceErr := BACnetConfirmedServiceChoiceParse(readBuffer)
	if _errorChoiceErr != nil {
		return nil, errors.Wrap(_errorChoiceErr, "Error parsing 'errorChoice' field")
	}
	errorChoice := _errorChoice
	if closeErr := readBuffer.CloseContext("errorChoice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorChoice")
	}

	// Simple Field (error)
	if pullErr := readBuffer.PullContext("error"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for error")
	}
	_error, _errorErr := BACnetErrorParse(readBuffer, BACnetConfirmedServiceChoice(errorChoice))
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field")
	}
	error := CastBACnetError(_error)
	if closeErr := readBuffer.CloseContext("error"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for error")
	}

	if closeErr := readBuffer.CloseContext("APDUError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUError")
	}

	// Create a partially initialized instance
	_child := &APDUError{
		OriginalInvokeId: originalInvokeId,
		ErrorChoice:      errorChoice,
		Error:            CastBACnetError(error),
		APDU:             &APDU{},
	}
	_child.APDU.Child = _child
	return _child, nil
}

func (m *APDUError) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUError")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 4, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (errorChoice)
		if pushErr := writeBuffer.PushContext("errorChoice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorChoice")
		}
		_errorChoiceErr := m.ErrorChoice.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("errorChoice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorChoice")
		}
		if _errorChoiceErr != nil {
			return errors.Wrap(_errorChoiceErr, "Error serializing 'errorChoice' field")
		}

		// Simple Field (error)
		if pushErr := writeBuffer.PushContext("error"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for error")
		}
		_errorErr := m.Error.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("error"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for error")
		}
		if _errorErr != nil {
			return errors.Wrap(_errorErr, "Error serializing 'error' field")
		}

		if popErr := writeBuffer.PopContext("APDUError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUError")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
