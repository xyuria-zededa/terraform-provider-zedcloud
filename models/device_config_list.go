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

// DeviceConfigList Device configuration payload detail
//
// # Device configuration request paylod
//
// swagger:model DeviceConfigList
type DeviceConfigList struct {

	// device config list
	// Required: true
	List []*DeviceConfigSummary `json:"list"`

	// filter next
	// Required: true
	Next *Cursor `json:"next"`

	// Summary by state
	// Required: true
	SummaryByState *Summary `json:"summaryByState"`

	// Summary by project distribution
	SummaryByTagDistribution *Summary `json:"summaryByTagDistribution,omitempty"`
}

// Validate validates this device config list
func (m *DeviceConfigList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByTagDistribution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceConfigList) validateList(formats strfmt.Registry) error {

	if err := validate.Required("list", "body", m.List); err != nil {
		return err
	}

	for i := 0; i < len(m.List); i++ {
		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {
			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceConfigList) validateNext(formats strfmt.Registry) error {

	if err := validate.Required("next", "body", m.Next); err != nil {
		return err
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceConfigList) validateSummaryByState(formats strfmt.Registry) error {

	if err := validate.Required("summaryByState", "body", m.SummaryByState); err != nil {
		return err
	}

	if m.SummaryByState != nil {
		if err := m.SummaryByState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceConfigList) validateSummaryByTagDistribution(formats strfmt.Registry) error {
	if swag.IsZero(m.SummaryByTagDistribution) { // not required
		return nil
	}

	if m.SummaryByTagDistribution != nil {
		if err := m.SummaryByTagDistribution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByTagDistribution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByTagDistribution")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device config list based on the context it is used
func (m *DeviceConfigList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByTagDistribution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceConfigList) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.List); i++ {

		if m.List[i] != nil {
			if err := m.List[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceConfigList) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceConfigList) contextValidateSummaryByState(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByState != nil {
		if err := m.SummaryByState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceConfigList) contextValidateSummaryByTagDistribution(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByTagDistribution != nil {
		if err := m.SummaryByTagDistribution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByTagDistribution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByTagDistribution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceConfigList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceConfigList) UnmarshalBinary(b []byte) error {
	var res DeviceConfigList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}