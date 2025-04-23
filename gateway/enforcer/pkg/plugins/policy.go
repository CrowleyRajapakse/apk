package plugins

import (
	"context"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	extprocv3 "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
)

// Policy is an interface that defines the methods for a policy plugin.
type Policy interface {
	Name() string

	ApplyRequestHeaders(ctx context.Context, headers *extprocv3.ProcessingRequest) ([]*corev3.HeaderValueOption, error)
	ApplyRequestBody(ctx context.Context, body *extprocv3.ProcessingRequest) ([]byte, error)
	ApplyResponseHeaders(ctx context.Context, headers *extprocv3.ProcessingRequest) ([]*corev3.HeaderValueOption, error)
	ApplyResponseBody(ctx context.Context, body *extprocv3.ProcessingRequest) ([]byte, error)
}
