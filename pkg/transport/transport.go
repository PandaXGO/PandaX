package transport

import (
	"context"
)

type Type string

const (
	TypeHTTP Type = "HTTP"
	TypeGRPC Type = "GRPC"
)

// Server is transport server.
type Server interface {
	Type() Type
	Start(context.Context) error
	Stop(context.Context) error
}
