// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DNSInfo DNSInfo consists of DNS Configurations.
//
// DNSInfo is used to store Domain Name Server configuration.
//
// swagger:model DNSInfo
type DNSInfo struct {

	// domain name
	// Required: true
	Domain *string `json:"domain"`

	// Array of search strings
	// Required: true
	Search []string `json:"search"`

	// Array of dns server
	// Required: true
	Servers []string `json:"servers"`
}

// Validate validates this DNS info
func (m *DNSInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSInfo) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *DNSInfo) validateSearch(formats strfmt.Registry) error {

	if err := validate.Required("search", "body", m.Search); err != nil {
		return err
	}

	return nil
}

func (m *DNSInfo) validateServers(formats strfmt.Registry) error {

	if err := validate.Required("servers", "body", m.Servers); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this DNS info based on context it is used
func (m *DNSInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSInfo) UnmarshalBinary(b []byte) error {
	var res DNSInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
