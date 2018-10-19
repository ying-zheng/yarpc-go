// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sink

import (
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Hello_Sink_Args represents the arguments for the Hello.sink function.
//
// The arguments for sink are sent and received over the wire as this struct.
type Hello_Sink_Args struct {
	Snk *SinkRequest `json:"snk,omitempty"`
}

// ToWire translates a Hello_Sink_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Hello_Sink_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Snk != nil {
		w, err = v.Snk.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _SinkRequest_Read(w wire.Value) (*SinkRequest, error) {
	var v SinkRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Hello_Sink_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Hello_Sink_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Hello_Sink_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Hello_Sink_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Snk, err = _SinkRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Hello_Sink_Args
// struct.
func (v *Hello_Sink_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Snk != nil {
		fields[i] = fmt.Sprintf("Snk: %v", v.Snk)
		i++
	}

	return fmt.Sprintf("Hello_Sink_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Hello_Sink_Args match the
// provided Hello_Sink_Args.
//
// This function performs a deep comparison.
func (v *Hello_Sink_Args) Equals(rhs *Hello_Sink_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Snk == nil && rhs.Snk == nil) || (v.Snk != nil && rhs.Snk != nil && v.Snk.Equals(rhs.Snk))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Hello_Sink_Args.
func (v *Hello_Sink_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Snk != nil {
		err = multierr.Append(err, enc.AddObject("snk", v.Snk))
	}
	return err
}

// GetSnk returns the value of Snk if it is set or its
// zero value if it is unset.
func (v *Hello_Sink_Args) GetSnk() (o *SinkRequest) {
	if v.Snk != nil {
		return v.Snk
	}

	return
}

// IsSetSnk returns true if Snk is not nil.
func (v *Hello_Sink_Args) IsSetSnk() bool {
	return v.Snk != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "sink" for this struct.
func (v *Hello_Sink_Args) MethodName() string {
	return "sink"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be OneWay for this struct.
func (v *Hello_Sink_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

// Hello_Sink_Helper provides functions that aid in handling the
// parameters and return values of the Hello.sink
// function.
var Hello_Sink_Helper = struct {
	// Args accepts the parameters of sink in-order and returns
	// the arguments struct for the function.
	Args func(
		snk *SinkRequest,
	) *Hello_Sink_Args
}{}

func init() {
	Hello_Sink_Helper.Args = func(
		snk *SinkRequest,
	) *Hello_Sink_Args {
		return &Hello_Sink_Args{
			Snk: snk,
		}
	}

}
