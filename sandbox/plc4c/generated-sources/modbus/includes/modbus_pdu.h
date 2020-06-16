/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/
#ifndef PLC4C_MODBUS_READ_WRITE_MODBUS_PDU_H_
#define PLC4C_MODBUS_READ_WRITE_MODBUS_PDU_H_
#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/utils/list.h>
#include "modbus_pdu_write_file_record_request_item.h"
#include "modbus_pdu_read_file_record_request_item.h"
#include "modbus_pdu_read_file_record_response_item.h"
#include "modbus_pdu_write_file_record_response_item.h"
#include "modbus_pdu.h"

// Structure used to contain the discriminator values for discriminated types using this as a parent
struct plc4c_modbus_read_write_modbus_pdu_discriminator {
  bool error;
  unsigned int function;
  bool response;
};
typedef struct plc4c_modbus_read_write_modbus_pdu_discriminator plc4c_modbus_read_write_modbus_pdu_discriminator;

// Enum assigning each sub-type an individual id.
enum plc4c_modbus_read_write_modbus_pdu_type {
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_error = 0,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_discrete_inputs_request = 1,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_discrete_inputs_response = 2,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_coils_request = 3,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_coils_response = 4,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_single_coil_request = 5,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_single_coil_response = 6,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_multiple_coils_request = 7,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_multiple_coils_response = 8,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_input_registers_request = 9,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_input_registers_response = 10,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_holding_registers_request = 11,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_holding_registers_response = 12,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_single_register_request = 13,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_single_register_response = 14,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_multiple_holding_registers_request = 15,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_multiple_holding_registers_response = 16,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_write_multiple_holding_registers_request = 17,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_write_multiple_holding_registers_response = 18,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_mask_write_holding_register_request = 19,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_mask_write_holding_register_response = 20,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_fifo_queue_request = 21,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_fifo_queue_response = 22,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_file_record_request = 23,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_file_record_response = 24,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_file_record_request = 25,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_write_file_record_response = 26,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_exception_status_request = 27,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_exception_status_response = 28,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_diagnostic_request = 29,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_get_com_event_log_request = 30,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_get_com_event_log_response = 31,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_report_server_id_request = 32,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_report_server_id_response = 33,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_device_identification_request = 34,
  plc4c_modbus_read_write_modbus_pdu_type_modbus_read_write_modbus_pdu_read_device_identification_response = 35};
typedef enum plc4c_modbus_read_write_modbus_pdu_type plc4c_modbus_read_write_modbus_pdu_type;

// Function to get the discriminator values for a given type.
plc4c_modbus_read_write_modbus_pdu_discriminator plc4c_modbus_read_write_modbus_pdu_get_discriminator(plc4c_modbus_read_write_modbus_pdu_type type);

struct plc4c_modbus_read_write_modbus_pdu {
  /* This is an abstract type so this property saves the type of this typed union */
  plc4c_modbus_read_write_modbus_pdu_type _type;
  /* Properties */
  union {
    struct { /* ModbusPDUError */
      uint8_t modbus_pdu_error_exception_code;
    };
    struct { /* ModbusPDUReadDiscreteInputsRequest */
      uint16_t modbus_pdu_read_discrete_inputs_request_starting_address;
      uint16_t modbus_pdu_read_discrete_inputs_request_quantity;
    };
    struct { /* ModbusPDUReadDiscreteInputsResponse */
      plc4c_list modbus_pdu_read_discrete_inputs_response_value;
    };
    struct { /* ModbusPDUReadCoilsRequest */
      uint16_t modbus_pdu_read_coils_request_starting_address;
      uint16_t modbus_pdu_read_coils_request_quantity;
    };
    struct { /* ModbusPDUReadCoilsResponse */
      plc4c_list modbus_pdu_read_coils_response_value;
    };
    struct { /* ModbusPDUWriteSingleCoilRequest */
      uint16_t modbus_pdu_write_single_coil_request_address;
      uint16_t modbus_pdu_write_single_coil_request_value;
    };
    struct { /* ModbusPDUWriteSingleCoilResponse */
      uint16_t modbus_pdu_write_single_coil_response_address;
      uint16_t modbus_pdu_write_single_coil_response_value;
    };
    struct { /* ModbusPDUWriteMultipleCoilsRequest */
      uint16_t modbus_pdu_write_multiple_coils_request_starting_address;
      uint16_t modbus_pdu_write_multiple_coils_request_quantity;
      plc4c_list modbus_pdu_write_multiple_coils_request_value;
    };
    struct { /* ModbusPDUWriteMultipleCoilsResponse */
      uint16_t modbus_pdu_write_multiple_coils_response_starting_address;
      uint16_t modbus_pdu_write_multiple_coils_response_quantity;
    };
    struct { /* ModbusPDUReadInputRegistersRequest */
      uint16_t modbus_pdu_read_input_registers_request_starting_address;
      uint16_t modbus_pdu_read_input_registers_request_quantity;
    };
    struct { /* ModbusPDUReadInputRegistersResponse */
      plc4c_list modbus_pdu_read_input_registers_response_value;
    };
    struct { /* ModbusPDUReadHoldingRegistersRequest */
      uint16_t modbus_pdu_read_holding_registers_request_starting_address;
      uint16_t modbus_pdu_read_holding_registers_request_quantity;
    };
    struct { /* ModbusPDUReadHoldingRegistersResponse */
      plc4c_list modbus_pdu_read_holding_registers_response_value;
    };
    struct { /* ModbusPDUWriteSingleRegisterRequest */
      uint16_t modbus_pdu_write_single_register_request_address;
      uint16_t modbus_pdu_write_single_register_request_value;
    };
    struct { /* ModbusPDUWriteSingleRegisterResponse */
      uint16_t modbus_pdu_write_single_register_response_address;
      uint16_t modbus_pdu_write_single_register_response_value;
    };
    struct { /* ModbusPDUWriteMultipleHoldingRegistersRequest */
      uint16_t modbus_pdu_write_multiple_holding_registers_request_starting_address;
      uint16_t modbus_pdu_write_multiple_holding_registers_request_quantity;
      plc4c_list modbus_pdu_write_multiple_holding_registers_request_value;
    };
    struct { /* ModbusPDUWriteMultipleHoldingRegistersResponse */
      uint16_t modbus_pdu_write_multiple_holding_registers_response_starting_address;
      uint16_t modbus_pdu_write_multiple_holding_registers_response_quantity;
    };
    struct { /* ModbusPDUReadWriteMultipleHoldingRegistersRequest */
      uint16_t modbus_pdu_read_write_multiple_holding_registers_request_read_starting_address;
      uint16_t modbus_pdu_read_write_multiple_holding_registers_request_read_quantity;
      uint16_t modbus_pdu_read_write_multiple_holding_registers_request_write_starting_address;
      uint16_t modbus_pdu_read_write_multiple_holding_registers_request_write_quantity;
      plc4c_list modbus_pdu_read_write_multiple_holding_registers_request_value;
    };
    struct { /* ModbusPDUReadWriteMultipleHoldingRegistersResponse */
      plc4c_list modbus_pdu_read_write_multiple_holding_registers_response_value;
    };
    struct { /* ModbusPDUMaskWriteHoldingRegisterRequest */
      uint16_t modbus_pdu_mask_write_holding_register_request_reference_address;
      uint16_t modbus_pdu_mask_write_holding_register_request_and_mask;
      uint16_t modbus_pdu_mask_write_holding_register_request_or_mask;
    };
    struct { /* ModbusPDUMaskWriteHoldingRegisterResponse */
      uint16_t modbus_pdu_mask_write_holding_register_response_reference_address;
      uint16_t modbus_pdu_mask_write_holding_register_response_and_mask;
      uint16_t modbus_pdu_mask_write_holding_register_response_or_mask;
    };
    struct { /* ModbusPDUReadFifoQueueRequest */
      uint16_t modbus_pdu_read_fifo_queue_request_fifo_pointer_address;
    };
    struct { /* ModbusPDUReadFifoQueueResponse */
      plc4c_list modbus_pdu_read_fifo_queue_response_fifo_value;
    };
    struct { /* ModbusPDUReadFileRecordRequest */
      plc4c_list modbus_pdu_read_file_record_request_items;
    };
    struct { /* ModbusPDUReadFileRecordResponse */
      plc4c_list modbus_pdu_read_file_record_response_items;
    };
    struct { /* ModbusPDUWriteFileRecordRequest */
      plc4c_list modbus_pdu_write_file_record_request_items;
    };
    struct { /* ModbusPDUWriteFileRecordResponse */
      plc4c_list modbus_pdu_write_file_record_response_items;
    };
    struct { /* ModbusPDUReadExceptionStatusRequest */
    };
    struct { /* ModbusPDUReadExceptionStatusResponse */
      uint8_t modbus_pdu_read_exception_status_response_value;
    };
    struct { /* ModbusPDUDiagnosticRequest */
      uint16_t modbus_pdu_diagnostic_request_status;
      uint16_t modbus_pdu_diagnostic_request_event_count;
    };
    struct { /* ModbusPDUGetComEventLogRequest */
    };
    struct { /* ModbusPDUGetComEventLogResponse */
      uint16_t modbus_pdu_get_com_event_log_response_status;
      uint16_t modbus_pdu_get_com_event_log_response_event_count;
      uint16_t modbus_pdu_get_com_event_log_response_message_count;
      plc4c_list modbus_pdu_get_com_event_log_response_events;
    };
    struct { /* ModbusPDUReportServerIdRequest */
    };
    struct { /* ModbusPDUReportServerIdResponse */
      plc4c_list modbus_pdu_report_server_id_response_value;
    };
    struct { /* ModbusPDUReadDeviceIdentificationRequest */
    };
    struct { /* ModbusPDUReadDeviceIdentificationResponse */
    };
  };
};
typedef struct plc4c_modbus_read_write_modbus_pdu plc4c_modbus_read_write_modbus_pdu;

plc4c_return_code plc4c_modbus_read_write_modbus_pdu_parse(plc4c_spi_read_buffer* buf, bool response, plc4c_modbus_read_write_modbus_pdu** message);

plc4c_return_code plc4c_modbus_read_write_modbus_pdu_serialize(plc4c_spi_write_buffer* buf, plc4c_modbus_read_write_modbus_pdu* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_MODBUS_READ_WRITE_MODBUS_PDU_H_
