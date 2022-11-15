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


from ctypes import c_bool
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
import math

    
@dataclass
class ModbusPDUReadExceptionStatusRequest(PlcMessage,ModbusPDU):

    # Accessors for discriminator values.
    def c_bool getErrorFlag() {
        return (c_bool) false
    def c_uint8 getFunctionFlag() {
        return (c_uint8) 0x07
    def c_bool getResponse() {
        return (c_bool) false


    def __post_init__(self):
        super().__init__( )




    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusPDUReadExceptionStatusRequest")

        writeBuffer.popContext("ModbusPDUReadExceptionStatusRequest")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUReadExceptionStatusRequest = self

        return lengthInBits


    @staticmethod
    def staticParseBuilder(readBuffer: ReadBuffer, response: c_bool) -> ModbusPDUReadExceptionStatusRequestBuilder:
        readBuffer.pullContext("ModbusPDUReadExceptionStatusRequest")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        readBuffer.closeContext("ModbusPDUReadExceptionStatusRequest")
        # Create the instance
        return ModbusPDUReadExceptionStatusRequestBuilder()


    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadExceptionStatusRequest):
            return False

        that: ModbusPDUReadExceptionStatusRequest = ModbusPDUReadExceptionStatusRequest(o)
        return super().equals(that) && True

    def hashCode(self) -> int:
        return hash(super().hashCode() )

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"


class ModbusPDUReadExceptionStatusRequestBuilder(ModbusPDUModbusPDUBuilder:def ModbusPDUReadExceptionStatusRequestBuilder( ):

        def build(
        ) -> ModbusPDUReadExceptionStatusRequest:
        modbusPDUReadExceptionStatusRequest: ModbusPDUReadExceptionStatusRequest = ModbusPDUReadExceptionStatusRequest(
)
        return modbusPDUReadExceptionStatusRequest



