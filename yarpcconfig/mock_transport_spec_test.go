// Copyright (c) 2022 Uber Technologies, Inc.
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

package yarpcconfig

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
	"go.uber.org/yarpc/api/transport"
)

// Builds mockTransportSpec objects.
//
//	mockSpec := mockTransportSpecBuilder{
//		Name: "...",
//		TransportConfig: reflect.TypeOf(myConfig{}),
//	}.Build()
//
//	mockSpec.EXPECT().BuildTransport(myConfig{...}).Return(...)
//	mockSpec.Spec()
type mockTransportSpecBuilder struct {
	Name            string
	TransportConfig reflect.Type

	// Any of the following may be nil to indicate that the transport does not
	// support that functionality.

	InboundConfig        reflect.Type
	UnaryOutboundConfig  reflect.Type
	OnewayOutboundConfig reflect.Type
	StreamOutboundConfig reflect.Type
}

// build the mockTransportSpec.
func (b mockTransportSpecBuilder) Build(ctrl *gomock.Controller) *mockTransportSpec {
	s := TransportSpec{Name: b.Name}
	m := mockTransportSpec{ctrl: ctrl, spec: &s}

	// Build a Spec where the Build* functions point to a dummy function
	// (generated by builderFunc) which calls into the mock controller to get
	// the value to return.

	s.BuildTransport = builderFunc(ctrl, &m, "BuildTransport", []reflect.Type{b.TransportConfig, _typeOfKit}, _typeOfTransport)

	if b.InboundConfig != nil {
		s.BuildInbound = builderFunc(ctrl, &m, "BuildInbound", []reflect.Type{b.InboundConfig, _typeOfTransport, _typeOfKit}, _typeOfInbound)
	}
	if b.UnaryOutboundConfig != nil {
		s.BuildUnaryOutbound = builderFunc(ctrl, &m, "BuildUnaryOutbound", []reflect.Type{b.UnaryOutboundConfig, _typeOfTransport, _typeOfKit}, _typeOfUnaryOutbound)
	}
	if b.OnewayOutboundConfig != nil {
		s.BuildOnewayOutbound = builderFunc(ctrl, &m, "BuildOnewayOutbound", []reflect.Type{b.OnewayOutboundConfig, _typeOfTransport, _typeOfKit}, _typeOfOnewayOutbound)
	}
	if b.StreamOutboundConfig != nil {
		s.BuildStreamOutbound = builderFunc(ctrl, &m, "BuildStreamOutbound", []reflect.Type{b.StreamOutboundConfig, _typeOfTransport, _typeOfKit}, _typeOfStreamOutbound)
	}

	return &m
}

// mockTransportSpec sets up a fake TransportSpec. The underlying Spec can be
// obtained using the Spec function.
type mockTransportSpec struct {
	spec *TransportSpec
	ctrl *gomock.Controller
}

// The following Build* functions are never called directly. They're just
// there to let gomock verify the signatures and return types.

func (m *mockTransportSpec) BuildTransport(interface{}, gomock.Matcher) (transport.Transport, error) {
	panic("This function should never be called")
}

func (m *mockTransportSpec) BuildInbound(interface{}, transport.Transport, gomock.Matcher) (transport.Inbound, error) {
	panic("This function should never be called")
}

func (m *mockTransportSpec) BuildUnaryOutbound(interface{}, transport.Transport, gomock.Matcher) (transport.UnaryOutbound, error) {
	panic("This function should never be called")
}

func (m *mockTransportSpec) BuildOnewayOutbound(interface{}, transport.Transport, gomock.Matcher) (transport.OnewayOutbound, error) {
	panic("This function should never be called")
}

func (m *mockTransportSpec) BuildStreamOutbound(interface{}, transport.Transport, gomock.Matcher) (transport.StreamOutbound, error) {
	panic("This function should never be called")
}

// EXPECT may be used to define expectations on the TransportSpec.
func (m *mockTransportSpec) EXPECT() *_transportSpecRecorder {
	return &_transportSpecRecorder{m: m, ctrl: m.ctrl}
}

// Spec returns a TransportSpec based on the expectations set on this
// mockTransportSpec.
func (m *mockTransportSpec) Spec() TransportSpec {
	return *m.spec
}

// Provides functions to record TransportSpec expectations.
type _transportSpecRecorder struct {
	m    *mockTransportSpec
	ctrl *gomock.Controller
}

func (r *_transportSpecRecorder) BuildTransport(cfg interface{}, kit gomock.Matcher) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildTransport", cfg, kit)
}

func (r *_transportSpecRecorder) BuildInbound(cfg interface{}, t transport.Transport, kit gomock.Matcher) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildInbound", cfg, t, kit)
}

func (r *_transportSpecRecorder) BuildUnaryOutbound(cfg interface{}, t transport.Transport, kit gomock.Matcher) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildUnaryOutbound", cfg, t, kit)
}

func (r *_transportSpecRecorder) BuildOnewayOutbound(cfg interface{}, t transport.Transport, kit gomock.Matcher) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildOnewayOutbound", cfg, t, kit)
}

func (r *_transportSpecRecorder) BuildStreamOutbound(cfg interface{}, t transport.Transport, kit gomock.Matcher) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildStreamOutbound", cfg, t, kit)
}

// kitMatcher matches attributes of a kit
type kitMatcher struct {
	ServiceName         string
	OutboundServiceName string
}

func (m kitMatcher) Matches(x interface{}) bool {
	k, ok := x.(*Kit)
	if !ok {
		return false
	}

	return k.ServiceName() == m.ServiceName && k.OutboundServiceName() == m.OutboundServiceName
}

func (m kitMatcher) String() string {
	return fmt.Sprintf("kit{name: %q, outboundName: %q}", m.ServiceName, m.OutboundServiceName)
}
