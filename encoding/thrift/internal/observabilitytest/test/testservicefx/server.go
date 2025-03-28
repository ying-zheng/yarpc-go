// Code generated by thriftrw-plugin-yarpc
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

package testservicefx

import (
	fx "go.uber.org/fx"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	testserviceserver "go.uber.org/yarpc/encoding/thrift/internal/observabilitytest/test/testserviceserver"
)

// ServerParams defines the dependencies for the TestService server.
type ServerParams struct {
	fx.In

	Handler testserviceserver.Interface
}

// ServerResult defines the output of TestService server module. It provides the
// procedures of a TestService handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures []transport.Procedure `group:"yarpcfx"`
}

// Server provides procedures for TestService to an Fx application. It expects a
// testservicefx.Interface to be present in the container.
//
//	fx.Provide(
//		func(h *MyTestServiceHandler) testserviceserver.Interface {
//			return h
//		},
//		testservicefx.Server(),
//	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := testserviceserver.New(p.Handler, opts...)
		return ServerResult{Procedures: procedures}
	}
}
