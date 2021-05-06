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
type HPAIDiscoveryEndpoint struct {
	HostProtocolCode HostProtocolCode
	IpAddress        *IPAddress
	IpPort           uint16
}

// The corresponding interface
type IHPAIDiscoveryEndpoint interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewHPAIDiscoveryEndpoint(hostProtocolCode HostProtocolCode, ipAddress *IPAddress, ipPort uint16) *HPAIDiscoveryEndpoint {
	return &HPAIDiscoveryEndpoint{HostProtocolCode: hostProtocolCode, IpAddress: ipAddress, IpPort: ipPort}
}

func CastHPAIDiscoveryEndpoint(structType interface{}) *HPAIDiscoveryEndpoint {
	castFunc := func(typ interface{}) *HPAIDiscoveryEndpoint {
		if casted, ok := typ.(HPAIDiscoveryEndpoint); ok {
			return &casted
		}
		if casted, ok := typ.(*HPAIDiscoveryEndpoint); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *HPAIDiscoveryEndpoint) GetTypeName() string {
	return "HPAIDiscoveryEndpoint"
}

func (m *HPAIDiscoveryEndpoint) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *HPAIDiscoveryEndpoint) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (hostProtocolCode)
	lengthInBits += 8

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.LengthInBits()

	// Simple field (ipPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *HPAIDiscoveryEndpoint) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func HPAIDiscoveryEndpointParse(readBuffer utils.ReadBuffer) (*HPAIDiscoveryEndpoint, error) {
	if pullErr := readBuffer.PullContext("HPAIDiscoveryEndpoint"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	if pullErr := readBuffer.PullContext("hostProtocolCode"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (hostProtocolCode)
	hostProtocolCode, _hostProtocolCodeErr := HostProtocolCodeParse(readBuffer)
	if _hostProtocolCodeErr != nil {
		return nil, errors.Wrap(_hostProtocolCodeErr, "Error parsing 'hostProtocolCode' field")
	}
	if closeErr := readBuffer.CloseContext("hostProtocolCode"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("ipAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (ipAddress)
	ipAddress, _ipAddressErr := IPAddressParse(readBuffer)
	if _ipAddressErr != nil {
		return nil, errors.Wrap(_ipAddressErr, "Error parsing 'ipAddress' field")
	}
	if closeErr := readBuffer.CloseContext("ipAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (ipPort)
	ipPort, _ipPortErr := readBuffer.ReadUint16("ipPort", 16)
	if _ipPortErr != nil {
		return nil, errors.Wrap(_ipPortErr, "Error parsing 'ipPort' field")
	}

	if closeErr := readBuffer.CloseContext("HPAIDiscoveryEndpoint"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewHPAIDiscoveryEndpoint(hostProtocolCode, ipAddress, ipPort), nil
}

func (m *HPAIDiscoveryEndpoint) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("HPAIDiscoveryEndpoint"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (hostProtocolCode)
	if pushErr := writeBuffer.PushContext("hostProtocolCode"); pushErr != nil {
		return pushErr
	}
	_hostProtocolCodeErr := m.HostProtocolCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("hostProtocolCode"); popErr != nil {
		return popErr
	}
	if _hostProtocolCodeErr != nil {
		return errors.Wrap(_hostProtocolCodeErr, "Error serializing 'hostProtocolCode' field")
	}

	// Simple Field (ipAddress)
	if pushErr := writeBuffer.PushContext("ipAddress"); pushErr != nil {
		return pushErr
	}
	_ipAddressErr := m.IpAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("ipAddress"); popErr != nil {
		return popErr
	}
	if _ipAddressErr != nil {
		return errors.Wrap(_ipAddressErr, "Error serializing 'ipAddress' field")
	}

	// Simple Field (ipPort)
	ipPort := uint16(m.IpPort)
	_ipPortErr := writeBuffer.WriteUint16("ipPort", 16, (ipPort))
	if _ipPortErr != nil {
		return errors.Wrap(_ipPortErr, "Error serializing 'ipPort' field")
	}

	if popErr := writeBuffer.PopContext("HPAIDiscoveryEndpoint"); popErr != nil {
		return popErr
	}
	return nil
}
