// Code generated by go-swagger; DO NOT EDIT.

package resource_group_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource group status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resource group status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ResourceGroupStatusGetResourceGroupStatusByID(params *ResourceGroupStatusGetResourceGroupStatusByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusGetResourceGroupStatusByIDOK, error)

	ResourceGroupStatusGetResourceGroupStatusByName(params *ResourceGroupStatusGetResourceGroupStatusByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusGetResourceGroupStatusByNameOK, error)

	ResourceGroupStatusQueryResourceGroupStatus(params *ResourceGroupStatusQueryResourceGroupStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusQueryResourceGroupStatusOK, error)

	ResourceGroupStatusQueryResourceGroupStatusConfig(params *ResourceGroupStatusQueryResourceGroupStatusConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusQueryResourceGroupStatusConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ResourceGroupStatusGetResourceGroupStatusByID gets resource group status

Get the status (without security details) of a resource group record.
*/
func (a *Client) ResourceGroupStatusGetResourceGroupStatusByID(params *ResourceGroupStatusGetResourceGroupStatusByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusGetResourceGroupStatusByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourceGroupStatusGetResourceGroupStatusByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroupStatus_GetResourceGroupStatusById",
		Method:             "GET",
		PathPattern:        "/v1/projects/id/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourceGroupStatusGetResourceGroupStatusByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResourceGroupStatusGetResourceGroupStatusByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResourceGroupStatusGetResourceGroupStatusByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResourceGroupStatusGetResourceGroupStatusByName gets resource group status

Get the status (without security details) of a resource group record.
*/
func (a *Client) ResourceGroupStatusGetResourceGroupStatusByName(params *ResourceGroupStatusGetResourceGroupStatusByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusGetResourceGroupStatusByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourceGroupStatusGetResourceGroupStatusByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroupStatus_GetResourceGroupStatusByName",
		Method:             "GET",
		PathPattern:        "/v1/projects/name/{name}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourceGroupStatusGetResourceGroupStatusByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResourceGroupStatusGetResourceGroupStatusByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResourceGroupStatusGetResourceGroupStatusByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResourceGroupStatusQueryResourceGroupStatus queries resource groups status

Query the resource group status records.
*/
func (a *Client) ResourceGroupStatusQueryResourceGroupStatus(params *ResourceGroupStatusQueryResourceGroupStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusQueryResourceGroupStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourceGroupStatusQueryResourceGroupStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroupStatus_QueryResourceGroupStatus",
		Method:             "GET",
		PathPattern:        "/v1/projects/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourceGroupStatusQueryResourceGroupStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResourceGroupStatusQueryResourceGroupStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResourceGroupStatusQueryResourceGroupStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResourceGroupStatusQueryResourceGroupStatusConfig queries resource groups status and config

Query the resource group status and config records.
*/
func (a *Client) ResourceGroupStatusQueryResourceGroupStatusConfig(params *ResourceGroupStatusQueryResourceGroupStatusConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceGroupStatusQueryResourceGroupStatusConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourceGroupStatusQueryResourceGroupStatusConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroupStatus_QueryResourceGroupStatusConfig",
		Method:             "GET",
		PathPattern:        "/v1/projects/status-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourceGroupStatusQueryResourceGroupStatusConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResourceGroupStatusQueryResourceGroupStatusConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResourceGroupStatusQueryResourceGroupStatusConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
