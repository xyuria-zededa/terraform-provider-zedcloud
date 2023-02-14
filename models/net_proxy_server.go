// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Server Net Proxy Server
//
// # Net Proxy Server
//
// swagger:model Server
type Server struct {

	// Net Proxy Port
	Port int64 `json:"port,omitempty"`

	// Net Proxy proto
	Proto *Proto `json:"proto,omitempty"`

	// Net Proxy Server
	Server string `json:"server,omitempty"`
}

// Validate validates this net proxy server
func (m *Server) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProto(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Server) validateProto(formats strfmt.Registry) error {
	if swag.IsZero(m.Proto) { // not required
		return nil
	}

	if m.Proto != nil {
		if err := m.Proto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proto")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this net proxy server based on the context it is used
func (m *Server) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Server) contextValidateProto(ctx context.Context, formats strfmt.Registry) error {

	if m.Proto != nil {
		if err := m.Proto.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proto")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Server) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Server) UnmarshalBinary(b []byte) error {
	var res Server
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}