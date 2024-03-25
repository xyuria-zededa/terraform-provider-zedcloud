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

// AppInstanceLogsQueryResponseItem App instance query log response
//
// # AppInstanceLogsQueryResponseItem is used as the response for querying the logs of app instance
//
// swagger:model AppInstanceLogsQueryResponseItem
type AppInstanceLogsQueryResponseItem struct {

	// app instance logs
	Content string `json:"content,omitempty"`

	// message Id
	Msgid string `json:"msgid,omitempty"`

	// timestamp of query time
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this app instance logs query response item
func (m *AppInstanceLogsQueryResponseItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstanceLogsQueryResponseItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this app instance logs query response item based on context it is used
func (m *AppInstanceLogsQueryResponseItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppInstanceLogsQueryResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstanceLogsQueryResponseItem) UnmarshalBinary(b []byte) error {
	var res AppInstanceLogsQueryResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
