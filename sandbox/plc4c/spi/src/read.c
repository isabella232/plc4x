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

#include <stdlib.h>
#include <plc4c/read.h>
#include <plc4c/spi/types_private.h>


return_code plc4c_read_request_execute(plc4c_read_request *read_request,
                                       plc4c_read_request_execution **read_request_execution) {
    plc4c_system_task *system_task;
    read_request->connection->driver->read_function(read_request, &system_task);
    // Add the new task to the task-list.
    plc4c_utils_list_insert_tail_value(read_request->connection->system->task_list, system_task);

    plc4c_read_request_execution *new_read_request_execution = malloc(sizeof(plc4c_read_request_execution));
    new_read_request_execution->read_request = read_request;
    new_read_request_execution->read_response = NULL;
    new_read_request_execution->system_task = system_task;

    *read_request_execution = new_read_request_execution;
    return OK;
}

bool plc4c_read_request_finished_successfully(plc4c_read_request_execution *read_request_execution) {
    return true;
}

bool plc4c_read_request_has_error(plc4c_read_request_execution *read_request_execution) {
    return false;
}

plc4c_read_response *plc4c_read_request_get_response(plc4c_read_request_execution *read_request_execution) {
    return NULL;
}

void plc4c_read_request_destroy(plc4c_read_request *read_request) {
    free(read_request);
}

void plc4c_read_request_execution_destroy(plc4c_read_request_execution *read_request_execution) {
    free(read_request_execution);
}

