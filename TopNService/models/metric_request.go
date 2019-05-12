// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MetricRequest metric request
// swagger:model metricRequest
type MetricRequest struct {

	// metrickey
	Metrickey string `json:"metrickey,omitempty"`

	// metricvalue
	Metricvalue string `json:"metricvalue,omitempty"`

	// pageno
	Pageno int64 `json:"pageno,omitempty"`

	// pagesize
	Pagesize int64 `json:"pagesize,omitempty"`

	// topvaluecount
	Topvaluecount int64 `json:"topvaluecount,omitempty"`
}

// Validate validates this metric request
func (m *MetricRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricRequest) UnmarshalBinary(b []byte) error {
	var res MetricRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}