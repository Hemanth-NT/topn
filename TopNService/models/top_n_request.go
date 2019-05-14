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

// TopNRequest Model for TopN request API
// swagger:model topNRequest
type TopNRequest struct {

	// filter criteria
	// Required: true
	FilterCriteria *Criteria `json:"filterCriteria"`

	// Json field on which
	// Required: true
	Metrickey *string `json:"metrickey"`

	// page number of response
	// Required: true
	Pageno *int64 `json:"pageno"`

	// page size of response
	// Required: true
	Pagesize *int64 `json:"pagesize"`

	// Top N count needed
	// Required: true
	Topvaluecount *int64 `json:"topvaluecount"`
}

// Validate validates this top n request
func (m *TopNRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrickey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageno(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagesize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopvaluecount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopNRequest) validateFilterCriteria(formats strfmt.Registry) error {

	if err := validate.Required("filterCriteria", "body", m.FilterCriteria); err != nil {
		return err
	}

	if m.FilterCriteria != nil {
		if err := m.FilterCriteria.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterCriteria")
			}
			return err
		}
	}

	return nil
}

func (m *TopNRequest) validateMetrickey(formats strfmt.Registry) error {

	if err := validate.Required("metrickey", "body", m.Metrickey); err != nil {
		return err
	}

	return nil
}

func (m *TopNRequest) validatePageno(formats strfmt.Registry) error {

	if err := validate.Required("pageno", "body", m.Pageno); err != nil {
		return err
	}

	return nil
}

func (m *TopNRequest) validatePagesize(formats strfmt.Registry) error {

	if err := validate.Required("pagesize", "body", m.Pagesize); err != nil {
		return err
	}

	return nil
}

func (m *TopNRequest) validateTopvaluecount(formats strfmt.Registry) error {

	if err := validate.Required("topvaluecount", "body", m.Topvaluecount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopNRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopNRequest) UnmarshalBinary(b []byte) error {
	var res TopNRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}