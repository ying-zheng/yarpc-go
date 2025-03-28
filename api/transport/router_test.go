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

package transport

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestProcedureLogMarshaling(t *testing.T) {
	p := Procedure{
		Name:    "name",
		Service: "service",
		HandlerSpec: NewUnaryHandlerSpec(UnaryHandlerFunc(func(context.Context, *Request, ResponseWriter) error {
			return nil
		})),
		Encoding:  "raw",
		Signature: "signature",
	}
	enc := zapcore.NewMapObjectEncoder()
	assert.NoError(t, p.MarshalLogObject(enc), "Unexpected error marshaling procedure.")
	assert.Equal(t, map[string]interface{}{
		"name":      "name",
		"service":   "service",
		"handler":   map[string]interface{}{"rpcType": "Unary"},
		"encoding":  "raw",
		"signature": "signature",
	}, enc.Fields, "Unexpected output from marshaling procedure.")
}
