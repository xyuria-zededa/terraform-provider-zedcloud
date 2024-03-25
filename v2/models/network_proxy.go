// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Proxy Configurations
//
// swagger:model Proxy
type Proxy struct {

	// Proxy exceptions
	Exceptions string `json:"exceptions,omitempty"`

	// Network proxy
	NetworkProxy bool `json:"networkProxy,omitempty"`

	// Proxy Certificates
	//
	// Network Proxy Certificates
	NetworkProxyCerts []string `json:"networkProxyCerts"`

	// Direct URL for wpad.dat download
	//
	// Network Proxy URL
	NetworkProxyURL string `json:"networkProxyURL,omitempty"`

	// proxy configuration in a pacfile
	//
	// proxy configuration in a pacfile
	Pacfile string `json:"pacfile,omitempty"`

	// protocol level proxies
	//
	// Net Proxy: protocol level proxies
	Proxies []*Server `json:"proxies"`
}

// Validate validates this net proxy config
func (m *Proxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProxies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Proxy) validateProxies(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxies) { // not required
		return nil
	}

	for i := 0; i < len(m.Proxies); i++ {
		if swag.IsZero(m.Proxies[i]) { // not required
			continue
		}

		if m.Proxies[i] != nil {
			if err := m.Proxies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this net proxy config based on the context it is used
func (m *Proxy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProxies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Proxy) contextValidateProxies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Proxies); i++ {

		if m.Proxies[i] != nil {
			if err := m.Proxies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Proxy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Proxy) UnmarshalBinary(b []byte) error {
	var res Proxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
