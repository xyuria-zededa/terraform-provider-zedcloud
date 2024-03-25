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

// ObjectTagFilter Object tag list filter criteria
//
// Filter criteria for the object tag list.
//
// swagger:model ObjectTagFilter
type ObjectTagFilter struct {

	// Object Id which tags are associated.
	ObjID string `json:"objId,omitempty"`

	// Object name which tags are associated.
	ObjName string `json:"objName,omitempty"`

	// Object type
	ObjType *ObjectType `json:"objType,omitempty"`
}

// Validate validates this object tag filter
func (m *ObjectTagFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectTagFilter) validateObjType(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjType) { // not required
		return nil
	}

	if m.ObjType != nil {
		if err := m.ObjType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object tag filter based on the context it is used
func (m *ObjectTagFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectTagFilter) contextValidateObjType(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjType != nil {
		if err := m.ObjType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectTagFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectTagFilter) UnmarshalBinary(b []byte) error {
	var res ObjectTagFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
