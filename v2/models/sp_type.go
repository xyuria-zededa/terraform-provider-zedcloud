// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SpType clone of types defined in API repo.
//
//   - SPTYPE_MAPSERVER: mapping service for zededa overlay service
//   - SPTYPE_SUPPORTSERVER: if device has support feature enabled, this ia cloud service
//
// which device can be reached.
//
// swagger:model spType
type SpType string

func NewSpType(value SpType) *SpType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SpType.
func (m SpType) Pointer() *SpType {
	return &m
}

const (

	// SpTypeSPTYPEINVALIDSRV captures enum value "SPTYPE_INVALIDSRV"
	SpTypeSPTYPEINVALIDSRV SpType = "SPTYPE_INVALIDSRV"

	// SpTypeSPTYPEMAPSERVER captures enum value "SPTYPE_MAPSERVER"
	SpTypeSPTYPEMAPSERVER SpType = "SPTYPE_MAPSERVER"

	// SpTypeSPTYPESUPPORTSERVER captures enum value "SPTYPE_SUPPORTSERVER"
	SpTypeSPTYPESUPPORTSERVER SpType = "SPTYPE_SUPPORTSERVER"
)

// for schema
var spTypeEnum []interface{}

func init() {
	var res []SpType
	if err := json.Unmarshal([]byte(`["SPTYPE_INVALIDSRV","SPTYPE_MAPSERVER","SPTYPE_SUPPORTSERVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		spTypeEnum = append(spTypeEnum, v)
	}
}

func (m SpType) validateSpTypeEnum(path, location string, value SpType) error {
	if err := validate.EnumCase(path, location, value, spTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sp type
func (m SpType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSpTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sp type based on context it is used
func (m SpType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
