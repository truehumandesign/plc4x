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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry interface {
	utils.LengthAware
	utils.Serializable
	// GetMonitoredProperty returns MonitoredProperty (property field)
	GetMonitoredProperty() BACnetPropertyReferenceEnclosed
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetContextTagReal
	// GetTimestamped returns Timestamped (property field)
	GetTimestamped() BACnetContextTagBoolean
}

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry.
// This is useful for switch cases.
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntryExactly interface {
	BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry
	isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry() bool
}

// _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry is the data-structure of this message
type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry struct {
	MonitoredProperty BACnetPropertyReferenceEnclosed
	CovIncrement      BACnetContextTagReal
	Timestamped       BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) GetMonitoredProperty() BACnetPropertyReferenceEnclosed {
	return m.MonitoredProperty
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) GetCovIncrement() BACnetContextTagReal {
	return m.CovIncrement
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) GetTimestamped() BACnetContextTagBoolean {
	return m.Timestamped
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry factory function for _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry(monitoredProperty BACnetPropertyReferenceEnclosed, covIncrement BACnetContextTagReal, timestamped BACnetContextTagBoolean) *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry {
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry{MonitoredProperty: monitoredProperty, CovIncrement: covIncrement, Timestamped: timestamped}
}

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry(structType interface{}) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry"
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredProperty)
	lengthInBits += m.MonitoredProperty.GetLengthInBits()

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += m.CovIncrement.GetLengthInBits()
	}

	// Simple field (timestamped)
	lengthInBits += m.Timestamped.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntryParse(readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (monitoredProperty)
	if pullErr := readBuffer.PullContext("monitoredProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredProperty")
	}
	_monitoredProperty, _monitoredPropertyErr := BACnetPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(0)))
	if _monitoredPropertyErr != nil {
		return nil, errors.Wrap(_monitoredPropertyErr, "Error parsing 'monitoredProperty' field of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry")
	}
	monitoredProperty := _monitoredProperty.(BACnetPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("monitoredProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredProperty")
	}

	// Optional Field (covIncrement) (Can be skipped, if a given expression evaluates to false)
	var covIncrement BACnetContextTagReal = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for covIncrement")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_REAL)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'covIncrement' field of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry")
		default:
			covIncrement = _val.(BACnetContextTagReal)
			if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
			}
		}
	}

	// Simple Field (timestamped)
	if pullErr := readBuffer.PullContext("timestamped"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamped")
	}
	_timestamped, _timestampedErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _timestampedErr != nil {
		return nil, errors.Wrap(_timestampedErr, "Error parsing 'timestamped' field of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry")
	}
	timestamped := _timestamped.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("timestamped"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamped")
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry")
	}

	// Create the instance
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry{
		MonitoredProperty: monitoredProperty,
		CovIncrement:      covIncrement,
		Timestamped:       timestamped,
	}, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry")
	}

	// Simple Field (monitoredProperty)
	if pushErr := writeBuffer.PushContext("monitoredProperty"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredProperty")
	}
	_monitoredPropertyErr := writeBuffer.WriteSerializable(m.GetMonitoredProperty())
	if popErr := writeBuffer.PopContext("monitoredProperty"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredProperty")
	}
	if _monitoredPropertyErr != nil {
		return errors.Wrap(_monitoredPropertyErr, "Error serializing 'monitoredProperty' field")
	}

	// Optional Field (covIncrement) (Can be skipped, if the value is null)
	var covIncrement BACnetContextTagReal = nil
	if m.GetCovIncrement() != nil {
		if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for covIncrement")
		}
		covIncrement = m.GetCovIncrement()
		_covIncrementErr := writeBuffer.WriteSerializable(covIncrement)
		if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for covIncrement")
		}
		if _covIncrementErr != nil {
			return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
		}
	}

	// Simple Field (timestamped)
	if pushErr := writeBuffer.PushContext("timestamped"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamped")
	}
	_timestampedErr := writeBuffer.WriteSerializable(m.GetTimestamped())
	if popErr := writeBuffer.PopContext("timestamped"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamped")
	}
	if _timestampedErr != nil {
		return errors.Wrap(_timestampedErr, "Error serializing 'timestamped' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry")
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry() bool {
	return true
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
