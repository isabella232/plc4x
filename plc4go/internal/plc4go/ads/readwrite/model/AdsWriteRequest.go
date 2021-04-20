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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsWriteRequest struct {
	IndexGroup  uint32
	IndexOffset uint32
	Data        []int8
	Parent      *AdsData
}

// The corresponding interface
type IAdsWriteRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsWriteRequest) CommandId() CommandId {
	return CommandId_ADS_WRITE
}

func (m *AdsWriteRequest) Response() bool {
	return false
}

func (m *AdsWriteRequest) InitializeParent(parent *AdsData) {
}

func NewAdsWriteRequest(indexGroup uint32, indexOffset uint32, data []int8) *AdsData {
	child := &AdsWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Data:        data,
		Parent:      NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsWriteRequest(structType interface{}) *AdsWriteRequest {
	castFunc := func(typ interface{}) *AdsWriteRequest {
		if casted, ok := typ.(AdsWriteRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsWriteRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsWriteRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsWriteRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsWriteRequest) GetTypeName() string {
	return "AdsWriteRequest"
}

func (m *AdsWriteRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsWriteRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsWriteRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsWriteRequestParse(io utils.ReadBuffer) (*AdsData, error) {
	io.PullContext("AdsWriteRequest")

	// Simple Field (indexGroup)
	indexGroup, _indexGroupErr := io.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field")
	}

	// Simple Field (indexOffset)
	indexOffset, _indexOffsetErr := io.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field")
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := io.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Array field (data)
	io.PullContext("data", utils.WithRenderAsList(true))
	// Count array
	data := make([]int8, length)
	for curItem := uint16(0); curItem < uint16(length); curItem++ {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	io.CloseContext("data", utils.WithRenderAsList(true))

	io.CloseContext("AdsWriteRequest")

	// Create a partially initialized instance
	_child := &AdsWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Data:        data,
		Parent:      &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsWriteRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("AdsWriteRequest")

		// Simple Field (indexGroup)
		indexGroup := uint32(m.IndexGroup)
		_indexGroupErr := io.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.IndexOffset)
		_indexOffsetErr := io.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.Data)))
		_lengthErr := io.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		if m.Data != nil {
			io.PushContext("data", utils.WithRenderAsList(true))
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			io.PopContext("data", utils.WithRenderAsList(true))
		}

		io.PopContext("AdsWriteRequest")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsWriteRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "indexGroup":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IndexGroup = data
			case "indexOffset":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IndexOffset = data
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *AdsWriteRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.IndexGroup, xml.StartElement{Name: xml.Name{Local: "indexGroup"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.IndexOffset, xml.StartElement{Name: xml.Name{Local: "indexOffset"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m AdsWriteRequest) String() string {
	return string(m.Box("", 120))
}

func (m AdsWriteRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "AdsWriteRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("IndexGroup", m.IndexGroup, -1))
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("IndexOffset", m.IndexOffset, -1))
		// Implicit Field (length)
		length := uint32(uint32(len(m.Data)))
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Length", length, -1))
		// Array Field (data)
		if m.Data != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Data {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
