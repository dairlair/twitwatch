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

// Tweet tweet
// swagger:model Tweet
type Tweet struct {

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// full text
	// Required: true
	FullText *string `json:"fullText"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// twitter Id
	// Required: true
	TwitterID *int64 `json:"twitterId"`

	// twitter user Id
	// Required: true
	TwitterUserID *int64 `json:"twitterUserId"`

	// twitter username
	// Required: true
	TwitterUsername *string `json:"twitterUsername"`
}

// Validate validates this tweet
func (m *Tweet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwitterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwitterUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwitterUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tweet) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Tweet) validateFullText(formats strfmt.Registry) error {

	if err := validate.Required("fullText", "body", m.FullText); err != nil {
		return err
	}

	return nil
}

func (m *Tweet) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Tweet) validateTwitterID(formats strfmt.Registry) error {

	if err := validate.Required("twitterId", "body", m.TwitterID); err != nil {
		return err
	}

	return nil
}

func (m *Tweet) validateTwitterUserID(formats strfmt.Registry) error {

	if err := validate.Required("twitterUserId", "body", m.TwitterUserID); err != nil {
		return err
	}

	return nil
}

func (m *Tweet) validateTwitterUsername(formats strfmt.Registry) error {

	if err := validate.Required("twitterUsername", "body", m.TwitterUsername); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tweet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tweet) UnmarshalBinary(b []byte) error {
	var res Tweet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
