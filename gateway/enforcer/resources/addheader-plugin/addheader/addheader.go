package main

import (
	"context"

	"github.com/wso2/apk/gateway/enforcer/pkg/plugins"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	extprocv3 "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
)

// AddHeader is a plugin that adds headers to requests and responses.
type AddHeader struct{}

// Name returns the name of the plugin.
func (a *AddHeader) Name() string {
	return "AddHeader"
}

// ApplyRequestHeaders adds a header to the request.
func (a *AddHeader) ApplyRequestHeaders(ctx context.Context, req *extprocv3.ProcessingRequest) ([]*corev3.HeaderValueOption, error) {
	return []*corev3.HeaderValueOption{
		{Header: &corev3.HeaderValue{Key: "x-added-req", Value: "header-plugin"}},
	}, nil
}

// ApplyResponseHeaders adds a header to the response.
func (a *AddHeader) ApplyResponseHeaders(ctx context.Context, req *extprocv3.ProcessingRequest) ([]*corev3.HeaderValueOption, error) {
	return []*corev3.HeaderValueOption{
		{Header: &corev3.HeaderValue{Key: "x-added-res", Value: "header-plugin"}},
	}, nil
}

// ApplyRequestBody does not modify the request body.
func (a *AddHeader) ApplyRequestBody(ctx context.Context, req *extprocv3.ProcessingRequest) ([]byte, error) {
	return nil, nil
}

// ApplyResponseBody does not modify the response body.
func (a *AddHeader) ApplyResponseBody(ctx context.Context, req *extprocv3.ProcessingRequest) ([]byte, error) {
	return nil, nil
}

// Plugin is exported plugin.
var Plugin plugins.Policy = &AddHeader{}
