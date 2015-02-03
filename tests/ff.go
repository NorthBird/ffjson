/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package tff

import (
	"math"
	"time"
)

type FFFoo struct {
	Blah int
}

type FFRecord struct {
	Timestamp int64 `json:"id,omitempty"`
	OriginId  uint32
	Bar       FFFoo
	Method    string `json:"meth"`
	ReqId     string
	ServerIp  string
	RemoteIp  string
	BytesSent uint64
}

type mystring string

// fjson: skip
type TsortName struct {
	C string
	B int `json:"A"`
}
type XsortName struct {
	C string
	B int `json:"A"`
}

// fjson: skip
type Tobj struct {
	X Tint
}
type Xobj struct {
	X Xint
}

// fjson: skip
type Tduration struct {
	X time.Duration
}
type Xduration struct {
	X time.Duration
}

// fjson: skip
type TtimePtr struct {
	X *time.Time
}
type XtimePtr struct {
	X *time.Time
}

// fjson: skip
type Tarray struct {
	X []int
}
type Xarray struct {
	X []int
}

// fjson: skip
type TarrayPtr struct {
	X []*int
}
type XarrayPtr struct {
	X []*int
}

// fjson: skip
type Tstring struct {
	X string
}
type Xstring struct {
	X string
}

// fjson: skip
type Tmystring struct {
	X mystring
}
type Xmystring struct {
	X mystring
}

// fjson: skip
type TmystringPtr struct {
	X *mystring
}
type XmystringPtr struct {
	X *mystring
}

// fjson: skip
type TstringTagged struct {
	X string `json:",string"`
}
type XstringTagged struct {
	X string `json:",string"`
}

// fjson: skip
type TintTagged struct {
	X int `json:",string"`
}
type XintTagged struct {
	X int `json:",string"`
}

// fjson: skip
type TboolTagged struct {
	X int `json:",string"`
}
type XboolTagged struct {
	X int `json:",string"`
}

// fjson: skip
type Tbool struct {
	X bool
}
type Xbool struct {
	Tbool
}

// fjson: skip
type Tint struct {
	X int
}
type Xint struct {
	Tint
}

// fjson: skip
type Tint8 struct {
	X int8
}
type Xint8 struct {
	Tint8
}

// fjson: skip
type Tint16 struct {
	X int16
}
type Xint16 struct {
	Tint16
}

// fjson: skip
type Tint32 struct {
	X int32
}
type Xint32 struct {
	Tint32
}

// fjson: skip
type Tint64 struct {
	X int64
}
type Xint64 struct {
	Tint64
}

// fjson: skip
type Tuint struct {
	X uint
}
type Xuint struct {
	Tuint
}

// fjson: skip
type Tuint8 struct {
	X uint8
}
type Xuint8 struct {
	Tuint8
}

// fjson: skip
type Tuint16 struct {
	X uint16
}
type Xuint16 struct {
	Tuint16
}

// fjson: skip
type Tuint32 struct {
	X uint32
}
type Xuint32 struct {
	Tuint32
}

// fjson: skip
type Tuint64 struct {
	X uint64
}
type Xuint64 struct {
	Tuint64
}

// fjson: skip
type Tuintptr struct {
	X uintptr
}
type Xuintptr struct {
	Tuintptr
}

// fjson: skip
type Tfloat32 struct {
	X float32
}
type Xfloat32 struct {
	Tfloat32
}

// fjson: skip
type Tfloat64 struct {
	X float64
}
type Xfloat64 struct {
	Tfloat64
}

// Tests from golang test suite
type Optionals struct {
	Sr string `json:"sr"`
	So string `json:"so,omitempty"`
	Sw string `json:"-"`

	Ir int `json:"omitempty"` // actually named omitempty, not an option
	Io int `json:"io,omitempty"`

	Slr []string `json:"slr,random"`
	Slo []string `json:"slo,omitempty"`

	Mr map[string]interface{} `json:"mr"`
	Mo map[string]interface{} `json:",omitempty"`

	Fr float64 `json:"fr"`
	Fo float64 `json:"fo,omitempty"`

	Br bool `json:"br"`
	Bo bool `json:"bo,omitempty"`

	Ur uint `json:"ur"`
	Uo uint `json:"uo,omitempty"`

	Str struct{} `json:"str"`
	Sto struct{} `json:"sto,omitempty"`
}

var unsupportedValues = []interface{}{
	math.NaN(),
	math.Inf(-1),
	math.Inf(1),
}

var optionalsExpected = `{
 "sr": "",
 "omitempty": 0,
 "slr": null,
 "mr": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "sto": {}
}`

type StringTag struct {
	BoolStr bool    `json:",string"`
	IntStr  int64   `json:",string"`
	FltStr  float64 `json:",string"`
	StrStr  string  `json:",string"`
}

var stringTagExpected = `{
 "BoolStr": "true",
 "IntStr": "42",
 "FltStr": "0",
 "StrStr": "\"xzbit\""
}`

type OmitAll struct {
	Ostr    string                 `json:",omitempty"`
	Oint    int                    `json:",omitempty"`
	Obool   bool                   `json:",omitempty"`
	Odouble float64                `json:",omitempty"`
	Ointer  interface{}            `json:",omitempty"`
	Omap    map[string]interface{} `json:",omitempty"`
	OstrP   *string                `json:",omitempty"`
	OintP   *int                   `json:",omitempty"`
	// TODO: Re-enable when issue #55 is fixed.
	//	OboolP *bool                   `json:",omitempty"`
	OmapP   *map[string]interface{} `json:",omitempty"`
	Astr    []string                `json:",omitempty"`
	Aint    []int                   `json:",omitempty"`
	Abool   []bool                  `json:",omitempty"`
	Adouble []float64               `json:",omitempty"`
}

var omitAllExpected = `{}`

type NoExported struct {
	field1 string
	field2 string
	field3 string
}

var noExportedExpected = `{}`

type OmitFirst struct {
	Ostr string `json:",omitempty"`
	Str  string
}

var omitFirstExpected = `{
 "Str": ""
}`

type OmitLast struct {
	Xstr string `json:",omitempty"`
	Str  string
}

var omitLastExpected = `{
 "Str": ""
}`

// byte slices are special even if they're renamed types.
type renamedByte byte
type renamedByteSlice []byte
type renamedRenamedByteSlice []renamedByte

// Ref has Marshaler and Unmarshaler methods with pointer receiver.
type Ref int

func (*Ref) MarshalJSON() ([]byte, error) {
	return []byte(`"ref"`), nil
}

func (r *Ref) UnmarshalJSON([]byte) error {
	*r = 12
	return nil
}

// Val has Marshaler methods with value receiver.
type Val int

func (Val) MarshalJSON() ([]byte, error) {
	return []byte(`"val"`), nil
}

// RefText has Marshaler and Unmarshaler methods with pointer receiver.
type RefText int

func (*RefText) MarshalText() ([]byte, error) {
	return []byte(`"ref"`), nil
}

func (r *RefText) UnmarshalText([]byte) error {
	*r = 13
	return nil
}

// ValText has Marshaler methods with value receiver.
type ValText int

func (ValText) MarshalText() ([]byte, error) {
	return []byte(`"val"`), nil
}

// C implements Marshaler and returns unescaped JSON.
type C int

func (C) MarshalJSON() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

// CText implements Marshaler and returns unescaped text.
type CText int

func (CText) MarshalText() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

type IntType int

type MyStruct struct {
	IntType
}

type BugA struct {
	S string
}

type BugB struct {
	BugA
	S string
}

type BugC struct {
	S string
}

// Legal Go: We never use the repeated embedded field (S).
type BugX struct {
	A int
	BugA
	BugB
}

type BugD struct { // Same as BugA after tagging.
	XXX string `json:"S"`
}

// BugD's tagged S field should dominate BugA's.
type BugY struct {
	BugA
	BugD
}

// There are no tags here, so S should not appear.
type BugZ struct {
	BugA
	BugC
	BugY // Contains a tagged S field through BugD; should not dominate.
}
