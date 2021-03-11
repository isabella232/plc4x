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
type AdsStampHeader struct {
    Timestamp uint64
    Samples uint32
    AdsNotificationSamples []*AdsNotificationSample
    IAdsStampHeader
}

// The corresponding interface
type IAdsStampHeader interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewAdsStampHeader(timestamp uint64, samples uint32, adsNotificationSamples []*AdsNotificationSample) *AdsStampHeader {
    return &AdsStampHeader{Timestamp: timestamp, Samples: samples, AdsNotificationSamples: adsNotificationSamples}
}

func CastAdsStampHeader(structType interface{}) *AdsStampHeader {
    castFunc := func(typ interface{}) *AdsStampHeader {
        if casted, ok := typ.(AdsStampHeader); ok {
            return &casted
        }
        if casted, ok := typ.(*AdsStampHeader); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *AdsStampHeader) GetTypeName() string {
    return "AdsStampHeader"
}

func (m *AdsStampHeader) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (timestamp)
    lengthInBits += 64

    // Simple field (samples)
    lengthInBits += 32

    // Array field
    if len(m.AdsNotificationSamples) > 0 {
        for _, element := range m.AdsNotificationSamples {
            lengthInBits += element.LengthInBits()
        }
    }

    return lengthInBits
}

func (m *AdsStampHeader) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func AdsStampHeaderParse(io *utils.ReadBuffer) (*AdsStampHeader, error) {

    // Simple Field (timestamp)
    timestamp, _timestampErr := io.ReadUint64(64)
    if _timestampErr != nil {
        return nil, errors.New("Error parsing 'timestamp' field " + _timestampErr.Error())
    }

    // Simple Field (samples)
    samples, _samplesErr := io.ReadUint32(32)
    if _samplesErr != nil {
        return nil, errors.New("Error parsing 'samples' field " + _samplesErr.Error())
    }

    // Array field (adsNotificationSamples)
    // Count array
    adsNotificationSamples := make([]*AdsNotificationSample, samples)
    for curItem := uint16(0); curItem < uint16(samples); curItem++ {
        _item, _err := AdsNotificationSampleParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'adsNotificationSamples' field " + _err.Error())
        }
        adsNotificationSamples[curItem] = _item
    }

    // Create the instance
    return NewAdsStampHeader(timestamp, samples, adsNotificationSamples), nil
}

func (m *AdsStampHeader) Serialize(io utils.WriteBuffer) error {

    // Simple Field (timestamp)
    timestamp := uint64(m.Timestamp)
    _timestampErr := io.WriteUint64(64, (timestamp))
    if _timestampErr != nil {
        return errors.New("Error serializing 'timestamp' field " + _timestampErr.Error())
    }

    // Simple Field (samples)
    samples := uint32(m.Samples)
    _samplesErr := io.WriteUint32(32, (samples))
    if _samplesErr != nil {
        return errors.New("Error serializing 'samples' field " + _samplesErr.Error())
    }

    // Array Field (adsNotificationSamples)
    if m.AdsNotificationSamples != nil {
        for _, _element := range m.AdsNotificationSamples {
            _elementErr := _element.Serialize(io)
            if _elementErr != nil {
                return errors.New("Error serializing 'adsNotificationSamples' field " + _elementErr.Error())
            }
        }
    }

    return nil
}

func (m *AdsStampHeader) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "timestamp":
                var data uint64
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Timestamp = data
            case "samples":
                var data uint32
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Samples = data
            case "adsNotificationSamples":
                var data []*AdsNotificationSample
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.AdsNotificationSamples = data
            }
        }
    }
}

func (m *AdsStampHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.ads.readwrite.AdsStampHeader"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Timestamp, xml.StartElement{Name: xml.Name{Local: "timestamp"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Samples, xml.StartElement{Name: xml.Name{Local: "samples"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "adsNotificationSamples"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.AdsNotificationSamples, xml.StartElement{Name: xml.Name{Local: "adsNotificationSamples"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "adsNotificationSamples"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
