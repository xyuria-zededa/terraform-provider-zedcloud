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

// AppStatusFromTPController app status from t p controller
//
// swagger:model AppStatusFromTPController
type AppStatusFromTPController struct {

	// app Id
	AppID string `json:"appId,omitempty"`

	// app name
	AppName string `json:"appName,omitempty"`

	// azure status
	AzureStatus *AzureStatus `json:"azureStatus,omitempty"`

	// enterprise Id
	EnterpriseID string `json:"enterpriseId,omitempty"`

	// eve device Id
	EveDeviceID string `json:"eveDeviceId,omitempty"`

	// type
	Type *ControllerType `json:"type,omitempty"`

	// vce status
	VceStatus *VCEStatus `json:"vceStatus,omitempty"`
}

// Validate validates this app status from t p controller
func (m *AppStatusFromTPController) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVceStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppStatusFromTPController) validateAzureStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureStatus) { // not required
		return nil
	}

	if m.AzureStatus != nil {
		if err := m.AzureStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureStatus")
			}
			return err
		}
	}

	return nil
}

func (m *AppStatusFromTPController) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *AppStatusFromTPController) validateVceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.VceStatus) { // not required
		return nil
	}

	if m.VceStatus != nil {
		if err := m.VceStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vceStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vceStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app status from t p controller based on the context it is used
func (m *AppStatusFromTPController) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzureStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVceStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppStatusFromTPController) contextValidateAzureStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureStatus != nil {
		if err := m.AzureStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureStatus")
			}
			return err
		}
	}

	return nil
}

func (m *AppStatusFromTPController) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *AppStatusFromTPController) contextValidateVceStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.VceStatus != nil {
		if err := m.VceStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vceStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vceStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppStatusFromTPController) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppStatusFromTPController) UnmarshalBinary(b []byte) error {
	var res AppStatusFromTPController
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
