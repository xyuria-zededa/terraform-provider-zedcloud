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

// EdgeApplication Edge application detailed configuration
//
// Edge application gets installed/uninstalled to/from edge node(s) and perform specific edge computing tasks. Lifecycle of Edge application (upgrade/restart) on Edge node can be managed and monitored by ZEDEDA Cloud controller through this detailed configuration.
// Example: {"name":"sample-app","originType":"ORIGIN_LOCAL","title":"Sample Edge Applications"}
//
// swagger:model EdgeApplication
type EdgeApplication struct {

	// user defined cpus for bundle
	Cpus int64 `json:"cpus,omitempty"`

	// Detailed description of the edge application
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// user defined drives
	// Read Only: true
	Drives int64 `json:"drives,omitempty"`

	// System defined universally unique Id of the edge application
	// Read Only: true
	// Pattern: [0-9A-Za-z-]+
	ID string `json:"id,omitempty"`

	// Flag to represent where app bundle is already imported
	IsImported bool `json:"isImported,omitempty"`

	// user defined manifest in JSON format
	ManifestJSON *VMManifest `json:"manifestJSON,omitempty"`

	// user defined memory for bundle
	Memory int64 `json:"memory,omitempty"`

	// User defined name of the edge application, unique across the enterprise. Once object is created, name can’t be changed
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// user defined network options
	Networks int64 `json:"networks,omitempty"`

	// origin of object
	// Required: true
	OriginType *Origin `json:"originType"`

	// origin and parent related details
	ParentDetail *ObjectParentDetail `json:"parentDetail,omitempty"`

	// project access list of the app bundle
	ProjectAccessList []string `json:"projectAccessList"`

	// system defined info
	Revision *ObjectRevision `json:"revision,omitempty"`

	// user defined storage for bundle
	Storage int64 `json:"storage,omitempty"`

	// User defined title of the edge application. Title can be changed at any time
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$
	Title *string `json:"title"`

	// User defined version for the given edge-app
	UserDefinedVersion string `json:"userDefinedVersion,omitempty"`
}

// Validate validates this app
func (m *EdgeApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManifestJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeApplication) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *EdgeApplication) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *EdgeApplication) validateManifestJSON(formats strfmt.Registry) error {
	if swag.IsZero(m.ManifestJSON) { // not required
		return nil
	}

	if m.ManifestJSON != nil {
		if err := m.ManifestJSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manifestJSON")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manifestJSON")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeApplication) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *EdgeApplication) validateOriginType(formats strfmt.Registry) error {

	if err := validate.Required("originType", "body", m.OriginType); err != nil {
		return err
	}

	if err := validate.Required("originType", "body", m.OriginType); err != nil {
		return err
	}

	if m.OriginType != nil {
		if err := m.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeApplication) validateParentDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDetail) { // not required
		return nil
	}

	if m.ParentDetail != nil {
		if err := m.ParentDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentDetail")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeApplication) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeApplication) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app based on the context it is used
func (m *EdgeApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDrives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManifestJSON(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeApplication) contextValidateDrives(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "drives", "body", int64(m.Drives)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeApplication) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeApplication) contextValidateManifestJSON(ctx context.Context, formats strfmt.Registry) error {

	if m.ManifestJSON != nil {
		if err := m.ManifestJSON.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manifestJSON")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manifestJSON")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeApplication) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginType != nil {
		if err := m.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeApplication) contextValidateParentDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentDetail != nil {
		if err := m.ParentDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentDetail")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeApplication) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {
		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeApplication) UnmarshalBinary(b []byte) error {
	var res EdgeApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
