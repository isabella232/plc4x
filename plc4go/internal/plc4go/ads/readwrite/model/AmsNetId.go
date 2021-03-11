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
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type AmsNetId struct {
    Octet1 uint8
    Octet2 uint8
    Octet3 uint8
    Octet4 uint8
    Octet5 uint8
    Octet6 uint8
    IAmsNetId
}

// The corresponding interface
type IAmsNetId interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewAmsNetId(octet1 uint8, octet2 uint8, octet3 uint8, octet4 uint8, octet5 uint8, octet6 uint8) *AmsNetId {
    return &AmsNetId{Octet1: octet1, Octet2: octet2, Octet3: octet3, Octet4: octet4, Octet5: octet5, Octet6: octet6}
}

func CastAmsNetId(structType interface{}) *AmsNetId {
    castFunc := func(typ interface{}) *AmsNetId {
        if casted, ok := typ.(AmsNetId); ok {
            return &casted
        }
        if casted, ok := typ.(*AmsNetId); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *AmsNetId) GetTypeName() string {
    return "AmsNetId"
}

func (m *AmsNetId) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (octet1)
    lengthInBits += 8

    // Simple field (octet2)
    lengthInBits += 8

    // Simple field (octet3)
    lengthInBits += 8

    // Simple field (octet4)
    lengthInBits += 8

    // Simple field (octet5)
    lengthInBits += 8

    // Simple field (octet6)
    lengthInBits += 8

    return lengthInBits
}

func (m *AmsNetId) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func AmsNetIdParse(io *utils.ReadBuffer) (*AmsNetId, error) {

    // Simple Field (octet1)
    octet1, _octet1Err := io.ReadUint8(8)
    if _octet1Err != nil {
        return nil, errors.New("Error parsing 'octet1' field " + _octet1Err.Error())
    }

    // Simple Field (octet2)
    octet2, _octet2Err := io.ReadUint8(8)
    if _octet2Err != nil {
        return nil, errors.New("Error parsing 'octet2' field " + _octet2Err.Error())
    }

    // Simple Field (octet3)
    octet3, _octet3Err := io.ReadUint8(8)
    if _octet3Err != nil {
        return nil, errors.New("Error parsing 'octet3' field " + _octet3Err.Error())
    }

    // Simple Field (octet4)
    octet4, _octet4Err := io.ReadUint8(8)
    if _octet4Err != nil {
        return nil, errors.New("Error parsing 'octet4' field " + _octet4Err.Error())
    }

    // Simple Field (octet5)
    octet5, _octet5Err := io.ReadUint8(8)
    if _octet5Err != nil {
        return nil, errors.New("Error parsing 'octet5' field " + _octet5Err.Error())
    }

    // Simple Field (octet6)
    octet6, _octet6Err := io.ReadUint8(8)
    if _octet6Err != nil {
        return nil, errors.New("Error parsing 'octet6' field " + _octet6Err.Error())
    }

    // Create the instance
    return NewAmsNetId(octet1, octet2, octet3, octet4, octet5, octet6), nil
}

func (m *AmsNetId) Serialize(io utils.WriteBuffer) error {

    // Simple Field (octet1)
    octet1 := uint8(m.Octet1)
    _octet1Err := io.WriteUint8(8, (octet1))
    if _octet1Err != nil {
        return errors.New("Error serializing 'octet1' field " + _octet1Err.Error())
    }

    // Simple Field (octet2)
    octet2 := uint8(m.Octet2)
    _octet2Err := io.WriteUint8(8, (octet2))
    if _octet2Err != nil {
        return errors.New("Error serializing 'octet2' field " + _octet2Err.Error())
    }

    // Simple Field (octet3)
    octet3 := uint8(m.Octet3)
    _octet3Err := io.WriteUint8(8, (octet3))
    if _octet3Err != nil {
        return errors.New("Error serializing 'octet3' field " + _octet3Err.Error())
    }

    // Simple Field (octet4)
    octet4 := uint8(m.Octet4)
    _octet4Err := io.WriteUint8(8, (octet4))
    if _octet4Err != nil {
        return errors.New("Error serializing 'octet4' field " + _octet4Err.Error())
    }

    // Simple Field (octet5)
    octet5 := uint8(m.Octet5)
    _octet5Err := io.WriteUint8(8, (octet5))
    if _octet5Err != nil {
        return errors.New("Error serializing 'octet5' field " + _octet5Err.Error())
    }

    // Simple Field (octet6)
    octet6 := uint8(m.Octet6)
    _octet6Err := io.WriteUint8(8, (octet6))
    if _octet6Err != nil {
        return errors.New("Error serializing 'octet6' field " + _octet6Err.Error())
    }

    return nil
}

func (m *AmsNetId) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    for {
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "octet1":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Octet1 = data
            case "octet2":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Octet2 = data
            case "octet3":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Octet3 = data
            case "octet4":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Octet4 = data
            case "octet5":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Octet5 = data
            case "octet6":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Octet6 = data
            }
        }
    }
}

func (m *AmsNetId) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.ads.readwrite.AmsNetId"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Octet1, xml.StartElement{Name: xml.Name{Local: "octet1"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Octet2, xml.StartElement{Name: xml.Name{Local: "octet2"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Octet3, xml.StartElement{Name: xml.Name{Local: "octet3"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Octet4, xml.StartElement{Name: xml.Name{Local: "octet4"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Octet5, xml.StartElement{Name: xml.Name{Local: "octet5"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Octet6, xml.StartElement{Name: xml.Name{Local: "octet6"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
