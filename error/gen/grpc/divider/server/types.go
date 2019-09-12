// Code generated by goa v3.0.6, DO NOT EDIT.
//
// divider gRPC server types
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package server

import (
	divider "goa.design/examples/error/gen/divider"
	dividerpb "goa.design/examples/error/gen/grpc/divider/pb"
)

// NewIntegerDividePayload builds the payload of the "integer_divide" endpoint
// of the "divider" service from the gRPC request type.
func NewIntegerDividePayload(message *dividerpb.IntegerDivideRequest) *divider.IntOperands {
	v := &divider.IntOperands{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewIntegerDivideResponse builds the gRPC response type from the result of
// the "integer_divide" endpoint of the "divider" service.
func NewIntegerDivideResponse(result int) *dividerpb.IntegerDivideResponse {
	message := &dividerpb.IntegerDivideResponse{}
	message.Field = int32(result)
	return message
}

// NewDividePayload builds the payload of the "divide" endpoint of the
// "divider" service from the gRPC request type.
func NewDividePayload(message *dividerpb.DivideRequest) *divider.FloatOperands {
	v := &divider.FloatOperands{
		A: message.A,
		B: message.B,
	}
	return v
}

// NewDivideResponse builds the gRPC response type from the result of the
// "divide" endpoint of the "divider" service.
func NewDivideResponse(result float64) *dividerpb.DivideResponse {
	message := &dividerpb.DivideResponse{}
	message.Field = result
	return message
}
