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

// DeviceLoad device load
//
// swagger:model DeviceLoad
type DeviceLoad string

func NewDeviceLoad(value DeviceLoad) *DeviceLoad {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DeviceLoad.
func (m DeviceLoad) Pointer() *DeviceLoad {
	return &m
}

const (

	// DeviceLoadDEVICELOADUNSPECIFIED captures enum value "DEVICE_LOAD_UNSPECIFIED"
	DeviceLoadDEVICELOADUNSPECIFIED DeviceLoad = "DEVICE_LOAD_UNSPECIFIED"

	// DeviceLoadDEVICELOADFREE captures enum value "DEVICE_LOAD_FREE"
	DeviceLoadDEVICELOADFREE DeviceLoad = "DEVICE_LOAD_FREE"

	// DeviceLoadDEVICELOADMODERATE captures enum value "DEVICE_LOAD_MODERATE"
	DeviceLoadDEVICELOADMODERATE DeviceLoad = "DEVICE_LOAD_MODERATE"

	// DeviceLoadDEVICELOADHEAVY captures enum value "DEVICE_LOAD_HEAVY"
	DeviceLoadDEVICELOADHEAVY DeviceLoad = "DEVICE_LOAD_HEAVY"
)

// for schema
var deviceLoadEnum []interface{}

func init() {
	var res []DeviceLoad
	if err := json.Unmarshal([]byte(`["DEVICE_LOAD_UNSPECIFIED","DEVICE_LOAD_FREE","DEVICE_LOAD_MODERATE","DEVICE_LOAD_HEAVY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceLoadEnum = append(deviceLoadEnum, v)
	}
}

func (m DeviceLoad) validateDeviceLoadEnum(path, location string, value DeviceLoad) error {
	if err := validate.EnumCase(path, location, value, deviceLoadEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this device load
func (m DeviceLoad) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeviceLoadEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this device load based on context it is used
func (m DeviceLoad) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
