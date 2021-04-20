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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7PayloadWriteVarRequest struct {
	Items  []*S7VarPayloadDataItem
	Parent *S7Payload
}

// The corresponding interface
type IS7PayloadWriteVarRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadWriteVarRequest) ParameterParameterType() uint8 {
	return 0x05
}

func (m *S7PayloadWriteVarRequest) MessageType() uint8 {
	return 0x01
}

func (m *S7PayloadWriteVarRequest) InitializeParent(parent *S7Payload) {
}

func NewS7PayloadWriteVarRequest(items []*S7VarPayloadDataItem) *S7Payload {
	child := &S7PayloadWriteVarRequest{
		Items:  items,
		Parent: NewS7Payload(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7PayloadWriteVarRequest(structType interface{}) *S7PayloadWriteVarRequest {
	castFunc := func(typ interface{}) *S7PayloadWriteVarRequest {
		if casted, ok := typ.(S7PayloadWriteVarRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadWriteVarRequest); ok {
			return casted
		}
		if casted, ok := typ.(S7Payload); ok {
			return CastS7PayloadWriteVarRequest(casted.Child)
		}
		if casted, ok := typ.(*S7Payload); ok {
			return CastS7PayloadWriteVarRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadWriteVarRequest) GetTypeName() string {
	return "S7PayloadWriteVarRequest"
}

func (m *S7PayloadWriteVarRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadWriteVarRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadWriteVarRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadWriteVarRequestParse(io utils.ReadBuffer, parameter *S7Parameter) (*S7Payload, error) {
	io.PullContext("S7PayloadWriteVarRequest")

	// Array field (items)
	io.PullContext("items", utils.WithRenderAsList(true))
	// Count array
	items := make([]*S7VarPayloadDataItem, uint16(len(CastS7ParameterWriteVarRequest(parameter).Items)))
	for curItem := uint16(0); curItem < uint16(uint16(len(CastS7ParameterWriteVarRequest(parameter).Items))); curItem++ {
		lastItem := curItem == uint16((len(CastS7ParameterWriteVarRequest(parameter).Items))-1)
		_item, _err := S7VarPayloadDataItemParse(io, lastItem)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}
	io.CloseContext("items", utils.WithRenderAsList(true))

	io.CloseContext("S7PayloadWriteVarRequest")

	// Create a partially initialized instance
	_child := &S7PayloadWriteVarRequest{
		Items:  items,
		Parent: &S7Payload{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7PayloadWriteVarRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("S7PayloadWriteVarRequest")

		// Array Field (items)
		if m.Items != nil {
			io.PushContext("items", utils.WithRenderAsList(true))
			itemCount := uint16(len(m.Items))
			var curItem uint16 = 0
			for _, _element := range m.Items {
				var lastItem bool = curItem == (itemCount - 1)
				_elementErr := _element.Serialize(io, lastItem)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
				curItem++
			}
			io.PopContext("items", utils.WithRenderAsList(true))
		}

		io.PopContext("S7PayloadWriteVarRequest")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7PayloadWriteVarRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "items":
				var data []*S7VarPayloadDataItem
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Items = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *S7PayloadWriteVarRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.Items {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	return nil
}

func (m S7PayloadWriteVarRequest) String() string {
	return string(m.Box("", 120))
}

func (m S7PayloadWriteVarRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "S7PayloadWriteVarRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Array Field (items)
		if m.Items != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Items {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Items", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
