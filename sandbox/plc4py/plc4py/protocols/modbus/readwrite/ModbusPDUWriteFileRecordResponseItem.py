#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License") you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from abc import ABC, abstractmethod
from dataclasses import dataclass


from ctypes import c_byte
from ctypes import c_uint16
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
import math

    
@dataclass
class ModbusPDUWriteFileRecordResponseItem(PlcMessage):
    referenceType: c_uint8
    fileNumber: c_uint16
    recordNumber: c_uint16
    recordData: []c_byte



    def __post_init__(self):
        super().__init__( )



    def getReferenceType(self) -> c_uint8:
        return referenceType

    def getFileNumber(self) -> c_uint16:
        return fileNumber

    def getRecordNumber(self) -> c_uint16:
        return recordNumber

    def getRecordData(self) -> []c_byte:
        return recordData


    def serialize(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusPDUWriteFileRecordResponseItem")

        # Simple Field (referenceType)
        writeSimpleField("referenceType", referenceType, writeUnsignedShort(writeBuffer, 8))

        # Simple Field (fileNumber)
        writeSimpleField("fileNumber", fileNumber, writeUnsignedInt(writeBuffer, 16))

        # Simple Field (recordNumber)
        writeSimpleField("recordNumber", recordNumber, writeUnsignedInt(writeBuffer, 16))

        # Implicit Field (recordLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        c_uint16 recordLength = (c_uint16) ((COUNT(getRecordData())) / (2))
        writeImplicitField("recordLength", recordLength, writeUnsignedInt(writeBuffer, 16))

        # Array Field (recordData)
        writeByteArrayField("recordData", recordData, writeByteArray(writeBuffer, 8))

        writeBuffer.popContext("ModbusPDUWriteFileRecordResponseItem")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = 0
        _value: ModbusPDUWriteFileRecordResponseItem = self

        # Simple field (referenceType)
        lengthInBits += 8

        # Simple field (fileNumber)
        lengthInBits += 16

        # Simple field (recordNumber)
        lengthInBits += 16

        # Implicit Field (recordLength)
        lengthInBits += 16

        # Array field
        if recordData is not None):
            lengthInBits += 8 * recordData.length


        return lengthInBits


    def staticParse(readBuffer: ReadBuffer , args) -> ModbusPDUWriteFileRecordResponseItem:
        positionAware: PositionAware = readBuffer
        return staticParse(readBuffer)


    @staticmethod
    def staticParseContext(readBuffer: ReadBuffer) -> ModbusPDUWriteFileRecordResponseItem:
        readBuffer.pullContext("ModbusPDUWriteFileRecordResponseItem")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        referenceType: c_uint8 = readSimpleField("referenceType", readUnsignedShort(readBuffer, 8))

        fileNumber: c_uint16 = readSimpleField("fileNumber", readUnsignedInt(readBuffer, 16))

        recordNumber: c_uint16 = readSimpleField("recordNumber", readUnsignedInt(readBuffer, 16))

        recordLength: c_uint16 = readImplicitField("recordLength", readUnsignedInt(readBuffer, 16))

        recordData: byte[] = readBuffer.readByteArray("recordData", Math.toIntExact(recordLength))

        readBuffer.closeContext("ModbusPDUWriteFileRecordResponseItem")
        # Create the instance
        _modbusPDUWriteFileRecordResponseItem: ModbusPDUWriteFileRecordResponseItem = ModbusPDUWriteFileRecordResponseItem(referenceType, fileNumber, recordNumber, recordData )
        return _modbusPDUWriteFileRecordResponseItem


    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUWriteFileRecordResponseItem):
            return False

        that: ModbusPDUWriteFileRecordResponseItem = ModbusPDUWriteFileRecordResponseItem(o)
        return (getReferenceType() == that.getReferenceType()) && (getFileNumber() == that.getFileNumber()) && (getRecordNumber() == that.getRecordNumber()) && (getRecordData() == that.getRecordData()) && True

    def hashCode(self) -> int:
        return hash(getReferenceType(), getFileNumber(), getRecordNumber(), getRecordData() )

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"




