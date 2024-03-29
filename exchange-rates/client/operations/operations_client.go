// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetDate(params *GetDateParams) (*GetDateOK, error)

	GetHistory(params *GetHistoryParams) (*GetHistoryOK, error)

	GetLatest(params *GetLatestParams) (*GetLatestOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDate returns the historical exchange for given date
*/
func (a *Client) GetDate(params *GetDateParams) (*GetDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDate",
		Method:             "GET",
		PathPattern:        "/{date}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetHistory returns the historical exchange rates
*/
func (a *Client) GetHistory(params *GetHistoryParams) (*GetHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHistory",
		Method:             "GET",
		PathPattern:        "/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLatest returns the latest foreign exchange reference rates
*/
func (a *Client) GetLatest(params *GetLatestParams) (*GetLatestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLatestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLatest",
		Method:             "GET",
		PathPattern:        "/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLatestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLatestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLatest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
