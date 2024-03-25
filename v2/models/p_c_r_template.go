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
	"github.com/go-openapi/validate"
)

// PCRTemplate PCR template
//
// # PCR template data for the specified eve version and firmware version
//
// swagger:model PCRTemplate
type PCRTemplate struct {

	// List of PCR values for the PCR template
	// Required: true
	PCRValues []*PCRValue `json:"PCRValues"`

	// EVE version related to the PCR template
	// Required: true
	EveVersion *string `json:"eveVersion"`

	// Firmware version related to the PCR template. If user doesn't set it, it will be set to '*'
	FirmwareVersion string `json:"firmwareVersion,omitempty"`

	// System defined universally unique Id of the PCR template. If not set in POST / PUT API calls, this will be treated as a new entry and a unique System Generated ID assigned.
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// Name of the PCR template. The name is Unique among PCR templates for that System Model. If it is not specified, a system-generated name will be assigned.
	// Max Length: 256
	Name string `json:"name,omitempty"`
}

// Validate validates this p c r template
func (m *PCRTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePCRValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEveVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PCRTemplate) validatePCRValues(formats strfmt.Registry) error {

	if err := validate.Required("PCRValues", "body", m.PCRValues); err != nil {
		return err
	}

	for i := 0; i < len(m.PCRValues); i++ {
		if swag.IsZero(m.PCRValues[i]) { // not required
			continue
		}

		if m.PCRValues[i] != nil {
			if err := m.PCRValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PCRValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PCRValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PCRTemplate) validateEveVersion(formats strfmt.Registry) error {

	if err := validate.Required("eveVersion", "body", m.EveVersion); err != nil {
		return err
	}

	return nil
}

func (m *PCRTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *PCRTemplate) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p c r template based on the context it is used
func (m *PCRTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePCRValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PCRTemplate) contextValidatePCRValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PCRValues); i++ {

		if m.PCRValues[i] != nil {
			if err := m.PCRValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PCRValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PCRValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PCRTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PCRTemplate) UnmarshalBinary(b []byte) error {
	var res PCRTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
