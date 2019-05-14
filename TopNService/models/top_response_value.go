// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TopResponseValue top response value
// swagger:model topResponseValue
type TopResponseValue struct {

	// Values in sorted order
	// Required: true
	MetricValue *string `json:"metricValue"`

	// JSON field on which TOP N service filtered
	// Required: true
	Metrickey *string `json:"metrickey"`
}

// Validate validates this top response value
func (m *TopResponseValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetricValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrickey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopResponseValue) validateMetricValue(formats strfmt.Registry) error {

	if err := validate.Required("metricValue", "body", m.MetricValue); err != nil {
		return err
	}

	return nil
}

func (m *TopResponseValue) validateMetrickey(formats strfmt.Registry) error {

	if err := validate.Required("metrickey", "body", m.Metrickey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopResponseValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopResponseValue) UnmarshalBinary(b []byte) error {
	var res TopResponseValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}