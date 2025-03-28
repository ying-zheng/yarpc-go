// Code generated by thriftrw v1.32.0. DO NOT EDIT.
// @generated

// Copyright (c) 2025 Uber Technologies, Inc.
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

package kv

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type ResourceDoesNotExist struct {
	Key     string  `json:"key,required"`
	Message *string `json:"message,omitempty"`
}

// ToWire translates a ResourceDoesNotExist struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *ResourceDoesNotExist) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Key), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Message != nil {
		w, err = wire.NewValueString(*(v.Message)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ResourceDoesNotExist struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ResourceDoesNotExist struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v ResourceDoesNotExist
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *ResourceDoesNotExist) FromWire(w wire.Value) error {
	var err error

	keyIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Key, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				keyIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Message = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !keyIsSet {
		return errors.New("field Key of ResourceDoesNotExist is required")
	}

	return nil
}

// Encode serializes a ResourceDoesNotExist struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a ResourceDoesNotExist struct could not be encoded.
func (v *ResourceDoesNotExist) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Key); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if v.Message != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Message)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a ResourceDoesNotExist struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a ResourceDoesNotExist struct could not be generated from the wire
// representation.
func (v *ResourceDoesNotExist) Decode(sr stream.Reader) error {

	keyIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Key, err = sr.ReadString()
			if err != nil {
				return err
			}
			keyIsSet = true
		case fh.ID == 2 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Message = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !keyIsSet {
		return errors.New("field Key of ResourceDoesNotExist is required")
	}

	return nil
}

// String returns a readable string representation of a ResourceDoesNotExist
// struct.
func (v *ResourceDoesNotExist) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Key: %v", v.Key)
	i++
	if v.Message != nil {
		fields[i] = fmt.Sprintf("Message: %v", *(v.Message))
		i++
	}

	return fmt.Sprintf("ResourceDoesNotExist{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*ResourceDoesNotExist) ErrorName() string {
	return "ResourceDoesNotExist"
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this ResourceDoesNotExist match the
// provided ResourceDoesNotExist.
//
// This function performs a deep comparison.
func (v *ResourceDoesNotExist) Equals(rhs *ResourceDoesNotExist) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Key == rhs.Key) {
		return false
	}
	if !_String_EqualsPtr(v.Message, rhs.Message) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ResourceDoesNotExist.
func (v *ResourceDoesNotExist) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("key", v.Key)
	if v.Message != nil {
		enc.AddString("message", *v.Message)
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *ResourceDoesNotExist) GetKey() (o string) {
	if v != nil {
		o = v.Key
	}
	return
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *ResourceDoesNotExist) GetMessage() (o string) {
	if v != nil && v.Message != nil {
		return *v.Message
	}

	return
}

// IsSetMessage returns true if Message is not nil.
func (v *ResourceDoesNotExist) IsSetMessage() bool {
	return v != nil && v.Message != nil
}

func (v *ResourceDoesNotExist) Error() string {
	return v.String()
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "kv",
	Package:  "go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/kv",
	FilePath: "kv.thrift",
	SHA1:     "9e8c1c30d0b6bd7d83426a92962269aaef706295",
	Raw:      rawIDL,
}

const rawIDL = "exception ResourceDoesNotExist {\n    1: required string key\n    2: optional string message\n}\n\nservice KeyValue {\n    string getValue(1: string key)\n        throws (1: ResourceDoesNotExist doesNotExist)\n    void setValue(1: string key, 2: string value)\n}\n"

// KeyValue_GetValue_Args represents the arguments for the KeyValue.getValue function.
//
// The arguments for getValue are sent and received over the wire as this struct.
type KeyValue_GetValue_Args struct {
	Key *string `json:"key,omitempty"`
}

// ToWire translates a KeyValue_GetValue_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *KeyValue_GetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Key != nil {
		w, err = wire.NewValueString(*(v.Key)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_GetValue_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_GetValue_Args struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v KeyValue_GetValue_Args
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *KeyValue_GetValue_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Key = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a KeyValue_GetValue_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a KeyValue_GetValue_Args struct could not be encoded.
func (v *KeyValue_GetValue_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Key != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Key)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a KeyValue_GetValue_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a KeyValue_GetValue_Args struct could not be generated from the wire
// representation.
func (v *KeyValue_GetValue_Args) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Key = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a KeyValue_GetValue_Args
// struct.
func (v *KeyValue_GetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}

	return fmt.Sprintf("KeyValue_GetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_GetValue_Args match the
// provided KeyValue_GetValue_Args.
//
// This function performs a deep comparison.
func (v *KeyValue_GetValue_Args) Equals(rhs *KeyValue_GetValue_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Key, rhs.Key) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_GetValue_Args.
func (v *KeyValue_GetValue_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Key != nil {
		enc.AddString("key", *v.Key)
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetValue_Args) GetKey() (o string) {
	if v != nil && v.Key != nil {
		return *v.Key
	}

	return
}

// IsSetKey returns true if Key is not nil.
func (v *KeyValue_GetValue_Args) IsSetKey() bool {
	return v != nil && v.Key != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "getValue" for this struct.
func (v *KeyValue_GetValue_Args) MethodName() string {
	return "getValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *KeyValue_GetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// KeyValue_GetValue_Helper provides functions that aid in handling the
// parameters and return values of the KeyValue.getValue
// function.
var KeyValue_GetValue_Helper = struct {
	// Args accepts the parameters of getValue in-order and returns
	// the arguments struct for the function.
	Args func(
		key *string,
	) *KeyValue_GetValue_Args

	// IsException returns true if the given error can be thrown
	// by getValue.
	//
	// An error can be thrown by getValue only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for getValue
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// getValue into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by getValue
	//
	//   value, err := getValue(args)
	//   result, err := KeyValue_GetValue_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from getValue: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*KeyValue_GetValue_Result, error)

	// UnwrapResponse takes the result struct for getValue
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if getValue threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := KeyValue_GetValue_Helper.UnwrapResponse(result)
	UnwrapResponse func(*KeyValue_GetValue_Result) (string, error)
}{}

func init() {
	KeyValue_GetValue_Helper.Args = func(
		key *string,
	) *KeyValue_GetValue_Args {
		return &KeyValue_GetValue_Args{
			Key: key,
		}
	}

	KeyValue_GetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *ResourceDoesNotExist:
			return true
		default:
			return false
		}
	}

	KeyValue_GetValue_Helper.WrapResponse = func(success string, err error) (*KeyValue_GetValue_Result, error) {
		if err == nil {
			return &KeyValue_GetValue_Result{Success: &success}, nil
		}

		switch e := err.(type) {
		case *ResourceDoesNotExist:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_GetValue_Result.DoesNotExist")
			}
			return &KeyValue_GetValue_Result{DoesNotExist: e}, nil
		}

		return nil, err
	}
	KeyValue_GetValue_Helper.UnwrapResponse = func(result *KeyValue_GetValue_Result) (success string, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// KeyValue_GetValue_Result represents the result of a KeyValue.getValue function call.
//
// The result of a getValue execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type KeyValue_GetValue_Result struct {
	// Value returned by getValue after a successful execution.
	Success      *string               `json:"success,omitempty"`
	DoesNotExist *ResourceDoesNotExist `json:"doesNotExist,omitempty"`
}

// ToWire translates a KeyValue_GetValue_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *KeyValue_GetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ResourceDoesNotExist_Read(w wire.Value) (*ResourceDoesNotExist, error) {
	var v ResourceDoesNotExist
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a KeyValue_GetValue_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_GetValue_Result struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v KeyValue_GetValue_Result
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *KeyValue_GetValue_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _ResourceDoesNotExist_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// Encode serializes a KeyValue_GetValue_Result struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a KeyValue_GetValue_Result struct could not be encoded.
func (v *KeyValue_GetValue_Result) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Success != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 0, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Success)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.DoesNotExist != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.DoesNotExist.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}

	if count != 1 {
		return fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", count)
	}

	return sw.WriteStructEnd()
}

func _ResourceDoesNotExist_Decode(sr stream.Reader) (*ResourceDoesNotExist, error) {
	var v ResourceDoesNotExist
	err := v.Decode(sr)
	return &v, err
}

// Decode deserializes a KeyValue_GetValue_Result struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a KeyValue_GetValue_Result struct could not be generated from the wire
// representation.
func (v *KeyValue_GetValue_Result) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 0 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Success = &x
			if err != nil {
				return err
			}

		case fh.ID == 1 && fh.Type == wire.TStruct:
			v.DoesNotExist, err = _ResourceDoesNotExist_Decode(sr)
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a KeyValue_GetValue_Result
// struct.
func (v *KeyValue_GetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}

	return fmt.Sprintf("KeyValue_GetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_GetValue_Result match the
// provided KeyValue_GetValue_Result.
//
// This function performs a deep comparison.
func (v *KeyValue_GetValue_Result) Equals(rhs *KeyValue_GetValue_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_GetValue_Result.
func (v *KeyValue_GetValue_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddString("success", *v.Success)
	}
	if v.DoesNotExist != nil {
		err = multierr.Append(err, enc.AddObject("doesNotExist", v.DoesNotExist))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetValue_Result) GetSuccess() (o string) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *KeyValue_GetValue_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetDoesNotExist returns the value of DoesNotExist if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetValue_Result) GetDoesNotExist() (o *ResourceDoesNotExist) {
	if v != nil && v.DoesNotExist != nil {
		return v.DoesNotExist
	}

	return
}

// IsSetDoesNotExist returns true if DoesNotExist is not nil.
func (v *KeyValue_GetValue_Result) IsSetDoesNotExist() bool {
	return v != nil && v.DoesNotExist != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "getValue" for this struct.
func (v *KeyValue_GetValue_Result) MethodName() string {
	return "getValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *KeyValue_GetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

// KeyValue_SetValue_Args represents the arguments for the KeyValue.setValue function.
//
// The arguments for setValue are sent and received over the wire as this struct.
type KeyValue_SetValue_Args struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// ToWire translates a KeyValue_SetValue_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *KeyValue_SetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Key != nil {
		w, err = wire.NewValueString(*(v.Key)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Value != nil {
		w, err = wire.NewValueString(*(v.Value)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_SetValue_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_SetValue_Args struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v KeyValue_SetValue_Args
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *KeyValue_SetValue_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Key = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Value = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a KeyValue_SetValue_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a KeyValue_SetValue_Args struct could not be encoded.
func (v *KeyValue_SetValue_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Key != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Key)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.Value != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Value)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a KeyValue_SetValue_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a KeyValue_SetValue_Args struct could not be generated from the wire
// representation.
func (v *KeyValue_SetValue_Args) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Key = &x
			if err != nil {
				return err
			}

		case fh.ID == 2 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Value = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a KeyValue_SetValue_Args
// struct.
func (v *KeyValue_SetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	if v.Value != nil {
		fields[i] = fmt.Sprintf("Value: %v", *(v.Value))
		i++
	}

	return fmt.Sprintf("KeyValue_SetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_SetValue_Args match the
// provided KeyValue_SetValue_Args.
//
// This function performs a deep comparison.
func (v *KeyValue_SetValue_Args) Equals(rhs *KeyValue_SetValue_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Key, rhs.Key) {
		return false
	}
	if !_String_EqualsPtr(v.Value, rhs.Value) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_SetValue_Args.
func (v *KeyValue_SetValue_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Key != nil {
		enc.AddString("key", *v.Key)
	}
	if v.Value != nil {
		enc.AddString("value", *v.Value)
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *KeyValue_SetValue_Args) GetKey() (o string) {
	if v != nil && v.Key != nil {
		return *v.Key
	}

	return
}

// IsSetKey returns true if Key is not nil.
func (v *KeyValue_SetValue_Args) IsSetKey() bool {
	return v != nil && v.Key != nil
}

// GetValue returns the value of Value if it is set or its
// zero value if it is unset.
func (v *KeyValue_SetValue_Args) GetValue() (o string) {
	if v != nil && v.Value != nil {
		return *v.Value
	}

	return
}

// IsSetValue returns true if Value is not nil.
func (v *KeyValue_SetValue_Args) IsSetValue() bool {
	return v != nil && v.Value != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "setValue" for this struct.
func (v *KeyValue_SetValue_Args) MethodName() string {
	return "setValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *KeyValue_SetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// KeyValue_SetValue_Helper provides functions that aid in handling the
// parameters and return values of the KeyValue.setValue
// function.
var KeyValue_SetValue_Helper = struct {
	// Args accepts the parameters of setValue in-order and returns
	// the arguments struct for the function.
	Args func(
		key *string,
		value *string,
	) *KeyValue_SetValue_Args

	// IsException returns true if the given error can be thrown
	// by setValue.
	//
	// An error can be thrown by setValue only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for setValue
	// given the error returned by it. The provided error may
	// be nil if setValue did not fail.
	//
	// This allows mapping errors returned by setValue into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// setValue
	//
	//   err := setValue(args)
	//   result, err := KeyValue_SetValue_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from setValue: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*KeyValue_SetValue_Result, error)

	// UnwrapResponse takes the result struct for setValue
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if setValue threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := KeyValue_SetValue_Helper.UnwrapResponse(result)
	UnwrapResponse func(*KeyValue_SetValue_Result) error
}{}

func init() {
	KeyValue_SetValue_Helper.Args = func(
		key *string,
		value *string,
	) *KeyValue_SetValue_Args {
		return &KeyValue_SetValue_Args{
			Key:   key,
			Value: value,
		}
	}

	KeyValue_SetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	KeyValue_SetValue_Helper.WrapResponse = func(err error) (*KeyValue_SetValue_Result, error) {
		if err == nil {
			return &KeyValue_SetValue_Result{}, nil
		}

		return nil, err
	}
	KeyValue_SetValue_Helper.UnwrapResponse = func(result *KeyValue_SetValue_Result) (err error) {
		return
	}

}

// KeyValue_SetValue_Result represents the result of a KeyValue.setValue function call.
//
// The result of a setValue execution is sent and received over the wire as this struct.
type KeyValue_SetValue_Result struct {
}

// ToWire translates a KeyValue_SetValue_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *KeyValue_SetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_SetValue_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_SetValue_Result struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v KeyValue_SetValue_Result
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *KeyValue_SetValue_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// Encode serializes a KeyValue_SetValue_Result struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a KeyValue_SetValue_Result struct could not be encoded.
func (v *KeyValue_SetValue_Result) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a KeyValue_SetValue_Result struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a KeyValue_SetValue_Result struct could not be generated from the wire
// representation.
func (v *KeyValue_SetValue_Result) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a KeyValue_SetValue_Result
// struct.
func (v *KeyValue_SetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("KeyValue_SetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_SetValue_Result match the
// provided KeyValue_SetValue_Result.
//
// This function performs a deep comparison.
func (v *KeyValue_SetValue_Result) Equals(rhs *KeyValue_SetValue_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_SetValue_Result.
func (v *KeyValue_SetValue_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "setValue" for this struct.
func (v *KeyValue_SetValue_Result) MethodName() string {
	return "setValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *KeyValue_SetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
