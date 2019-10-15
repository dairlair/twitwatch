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

// CreateTopic create topic
// swagger:model CreateTopic
type CreateTopic struct {

	// is active
	// Required: true
	IsActive *bool `json:"isActive"`

	// name
	// Required: true
	Name *string `json:"name"`

	// tracks
	// Required: true
	Tracks []string `json:"tracks"`
}

// Validate validates this create topic
func (m *CreateTopic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTracks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTopic) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("isActive", "body", m.IsActive); err != nil {
		return err
	}

	return nil
}

func (m *CreateTopic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateTopic) validateTracks(formats strfmt.Registry) error {

	if err := validate.Required("tracks", "body", m.Tracks); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTopic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTopic) UnmarshalBinary(b []byte) error {
	var res CreateTopic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}