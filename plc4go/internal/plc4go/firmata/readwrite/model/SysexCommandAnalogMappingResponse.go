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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type SysexCommandAnalogMappingResponse struct {
	Parent *SysexCommand
}

// The corresponding interface
type ISysexCommandAnalogMappingResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandAnalogMappingResponse) CommandType() uint8 {
	return 0x6A
}

func (m *SysexCommandAnalogMappingResponse) Response() bool {
	return false
}

func (m *SysexCommandAnalogMappingResponse) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandAnalogMappingResponse() *SysexCommand {
	child := &SysexCommandAnalogMappingResponse{
		Parent: NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandAnalogMappingResponse(structType interface{}) *SysexCommandAnalogMappingResponse {
	castFunc := func(typ interface{}) *SysexCommandAnalogMappingResponse {
		if casted, ok := typ.(SysexCommandAnalogMappingResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandAnalogMappingResponse); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandAnalogMappingResponse(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandAnalogMappingResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandAnalogMappingResponse) GetTypeName() string {
	return "SysexCommandAnalogMappingResponse"
}

func (m *SysexCommandAnalogMappingResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandAnalogMappingResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandAnalogMappingResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandAnalogMappingResponseParse(readBuffer utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingResponse"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandAnalogMappingResponse{
		Parent: &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandAnalogMappingResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandAnalogMappingResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}