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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry is the data-structure of this message
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry struct {
	MonitoredObjectIdentifier *BACnetContextTagObjectIdentifier
	ListOfCovReferences       *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
}

// IBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
type IBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry interface {
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetListOfCovReferences returns ListOfCovReferences (property field)
	GetListOfCovReferences() *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
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

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetListOfCovReferences() *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences {
	return m.ListOfCovReferences
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry factory function for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(monitoredObjectIdentifier *BACnetContextTagObjectIdentifier, listOfCovReferences *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	return &BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry{MonitoredObjectIdentifier: monitoredObjectIdentifier, ListOfCovReferences: listOfCovReferences}
}

func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(structType interface{}) *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry); ok {
		return casted
	}
	return nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits()

	// Simple field (listOfCovReferences)
	lengthInBits += m.ListOfCovReferences.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParse(readBuffer utils.ReadBuffer) (*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field")
	}
	monitoredObjectIdentifier := CastBACnetContextTagObjectIdentifier(_monitoredObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredObjectIdentifier")
	}

	// Simple Field (listOfCovReferences)
	if pullErr := readBuffer.PullContext("listOfCovReferences"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfCovReferences")
	}
	_listOfCovReferences, _listOfCovReferencesErr := BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesParse(readBuffer, uint8(uint8(1)))
	if _listOfCovReferencesErr != nil {
		return nil, errors.Wrap(_listOfCovReferencesErr, "Error parsing 'listOfCovReferences' field")
	}
	listOfCovReferences := CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences(_listOfCovReferences)
	if closeErr := readBuffer.CloseContext("listOfCovReferences"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfCovReferences")
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}

	// Create the instance
	return NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(monitoredObjectIdentifier, listOfCovReferences), nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifierErr := m.MonitoredObjectIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredObjectIdentifier")
	}
	if _monitoredObjectIdentifierErr != nil {
		return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
	}

	// Simple Field (listOfCovReferences)
	if pushErr := writeBuffer.PushContext("listOfCovReferences"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfCovReferences")
	}
	_listOfCovReferencesErr := m.ListOfCovReferences.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("listOfCovReferences"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfCovReferences")
	}
	if _listOfCovReferencesErr != nil {
		return errors.Wrap(_listOfCovReferencesErr, "Error serializing 'listOfCovReferences' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}
	return nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
