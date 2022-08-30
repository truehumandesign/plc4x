/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CipUnconnectedRequest_ROUTE uint16 = 0x0001

// CipUnconnectedRequest is the corresponding interface of CipUnconnectedRequest
type CipUnconnectedRequest interface {
	utils.LengthAware
	utils.Serializable
	CipService
	// GetUnconnectedService returns UnconnectedService (property field)
	GetUnconnectedService() CipService
	// GetBackPlane returns BackPlane (property field)
	GetBackPlane() int8
	// GetSlot returns Slot (property field)
	GetSlot() int8
}

// CipUnconnectedRequestExactly can be used when we want exactly this type and not a type which fulfills CipUnconnectedRequest.
// This is useful for switch cases.
type CipUnconnectedRequestExactly interface {
	CipUnconnectedRequest
	isCipUnconnectedRequest() bool
}

// _CipUnconnectedRequest is the data-structure of this message
type _CipUnconnectedRequest struct {
	*_CipService
	UnconnectedService CipService
	BackPlane          int8
	Slot               int8
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
	reservedField2 *uint8
	reservedField3 *uint8
	reservedField4 *uint8
	reservedField5 *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipUnconnectedRequest) GetService() uint8 {
	return 0x52
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipUnconnectedRequest) InitializeParent(parent CipService) {}

func (m *_CipUnconnectedRequest) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipUnconnectedRequest) GetUnconnectedService() CipService {
	return m.UnconnectedService
}

func (m *_CipUnconnectedRequest) GetBackPlane() int8 {
	return m.BackPlane
}

func (m *_CipUnconnectedRequest) GetSlot() int8 {
	return m.Slot
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CipUnconnectedRequest) GetRoute() uint16 {
	return CipUnconnectedRequest_ROUTE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipUnconnectedRequest factory function for _CipUnconnectedRequest
func NewCipUnconnectedRequest(unconnectedService CipService, backPlane int8, slot int8, serviceLen uint16) *_CipUnconnectedRequest {
	_result := &_CipUnconnectedRequest{
		UnconnectedService: unconnectedService,
		BackPlane:          backPlane,
		Slot:               slot,
		_CipService:        NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipUnconnectedRequest(structType interface{}) CipUnconnectedRequest {
	if casted, ok := structType.(CipUnconnectedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipUnconnectedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipUnconnectedRequest) GetTypeName() string {
	return "CipUnconnectedRequest"
}

func (m *_CipUnconnectedRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CipUnconnectedRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (messageSize)
	lengthInBits += 16

	// Simple field (unconnectedService)
	lengthInBits += m.UnconnectedService.GetLengthInBits()

	// Const Field (route)
	lengthInBits += 16

	// Simple field (backPlane)
	lengthInBits += 8

	// Simple field (slot)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipUnconnectedRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipUnconnectedRequestParse(readBuffer utils.ReadBuffer, serviceLen uint16) (CipUnconnectedRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipUnconnectedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipUnconnectedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipUnconnectedRequest")
		}
		if reserved != uint8(0x02) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x02),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	var reservedField1 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipUnconnectedRequest")
		}
		if reserved != uint8(0x20) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x20),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	var reservedField2 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipUnconnectedRequest")
		}
		if reserved != uint8(0x06) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x06),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField2 = &reserved
		}
	}

	var reservedField3 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipUnconnectedRequest")
		}
		if reserved != uint8(0x24) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x24),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField3 = &reserved
		}
	}

	var reservedField4 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipUnconnectedRequest")
		}
		if reserved != uint8(0x01) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x01),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField4 = &reserved
		}
	}

	var reservedField5 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipUnconnectedRequest")
		}
		if reserved != uint16(0x9D05) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x9D05),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField5 = &reserved
		}
	}

	// Implicit Field (messageSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	messageSize, _messageSizeErr := readBuffer.ReadUint16("messageSize", 16)
	_ = messageSize
	if _messageSizeErr != nil {
		return nil, errors.Wrap(_messageSizeErr, "Error parsing 'messageSize' field of CipUnconnectedRequest")
	}

	// Simple Field (unconnectedService)
	if pullErr := readBuffer.PullContext("unconnectedService"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unconnectedService")
	}
	_unconnectedService, _unconnectedServiceErr := CipServiceParse(readBuffer, uint16(messageSize))
	if _unconnectedServiceErr != nil {
		return nil, errors.Wrap(_unconnectedServiceErr, "Error parsing 'unconnectedService' field of CipUnconnectedRequest")
	}
	unconnectedService := _unconnectedService.(CipService)
	if closeErr := readBuffer.CloseContext("unconnectedService"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unconnectedService")
	}

	// Const Field (route)
	route, _routeErr := readBuffer.ReadUint16("route", 16)
	if _routeErr != nil {
		return nil, errors.Wrap(_routeErr, "Error parsing 'route' field of CipUnconnectedRequest")
	}
	if route != CipUnconnectedRequest_ROUTE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipUnconnectedRequest_ROUTE) + " but got " + fmt.Sprintf("%d", route))
	}

	// Simple Field (backPlane)
	_backPlane, _backPlaneErr := readBuffer.ReadInt8("backPlane", 8)
	if _backPlaneErr != nil {
		return nil, errors.Wrap(_backPlaneErr, "Error parsing 'backPlane' field of CipUnconnectedRequest")
	}
	backPlane := _backPlane

	// Simple Field (slot)
	_slot, _slotErr := readBuffer.ReadInt8("slot", 8)
	if _slotErr != nil {
		return nil, errors.Wrap(_slotErr, "Error parsing 'slot' field of CipUnconnectedRequest")
	}
	slot := _slot

	if closeErr := readBuffer.CloseContext("CipUnconnectedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipUnconnectedRequest")
	}

	// Create a partially initialized instance
	_child := &_CipUnconnectedRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		UnconnectedService: unconnectedService,
		BackPlane:          backPlane,
		Slot:               slot,
		reservedField0:     reservedField0,
		reservedField1:     reservedField1,
		reservedField2:     reservedField2,
		reservedField3:     reservedField3,
		reservedField4:     reservedField4,
		reservedField5:     reservedField5,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipUnconnectedRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipUnconnectedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipUnconnectedRequest")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x02)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x02),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x20)
			if m.reservedField1 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x20),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x06)
			if m.reservedField2 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x06),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField2
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x24)
			if m.reservedField3 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x24),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField3
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x01)
			if m.reservedField4 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x01),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField4
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved uint16 = uint16(0x9D05)
			if m.reservedField5 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint16(0x9D05),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField5
			}
			_err := writeBuffer.WriteUint16("reserved", 16, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Implicit Field (messageSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		messageSize := uint16(uint16(uint16(uint16(m.GetLengthInBytes()))-uint16(uint16(10))) - uint16(uint16(4)))
		_messageSizeErr := writeBuffer.WriteUint16("messageSize", 16, (messageSize))
		if _messageSizeErr != nil {
			return errors.Wrap(_messageSizeErr, "Error serializing 'messageSize' field")
		}

		// Simple Field (unconnectedService)
		if pushErr := writeBuffer.PushContext("unconnectedService"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unconnectedService")
		}
		_unconnectedServiceErr := writeBuffer.WriteSerializable(m.GetUnconnectedService())
		if popErr := writeBuffer.PopContext("unconnectedService"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unconnectedService")
		}
		if _unconnectedServiceErr != nil {
			return errors.Wrap(_unconnectedServiceErr, "Error serializing 'unconnectedService' field")
		}

		// Const Field (route)
		_routeErr := writeBuffer.WriteUint16("route", 16, 0x0001)
		if _routeErr != nil {
			return errors.Wrap(_routeErr, "Error serializing 'route' field")
		}

		// Simple Field (backPlane)
		backPlane := int8(m.GetBackPlane())
		_backPlaneErr := writeBuffer.WriteInt8("backPlane", 8, (backPlane))
		if _backPlaneErr != nil {
			return errors.Wrap(_backPlaneErr, "Error serializing 'backPlane' field")
		}

		// Simple Field (slot)
		slot := int8(m.GetSlot())
		_slotErr := writeBuffer.WriteInt8("slot", 8, (slot))
		if _slotErr != nil {
			return errors.Wrap(_slotErr, "Error serializing 'slot' field")
		}

		if popErr := writeBuffer.PopContext("CipUnconnectedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipUnconnectedRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CipUnconnectedRequest) isCipUnconnectedRequest() bool {
	return true
}

func (m *_CipUnconnectedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
