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

// ObjectMeta ObjectMeta defines the set of fields that objects returned from the API have
// swagger:model ObjectMeta
type ObjectMeta struct {

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	// Format: date-time
	DeletionTimestamp strfmt.DateTime `json:"deletionTimestamp,omitempty"`

	// ID unique value that identifies the resource generated by the server
	ID string `json:"id,omitempty"`

	// Name represents human readable name for the resource
	Name string `json:"name,omitempty"`
}

// Validate validates this object meta
func (m *ObjectMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectMeta) validateCreationTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTimestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectMeta) validateDeletionTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("deletionTimestamp", "body", "date-time", m.DeletionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectMeta) UnmarshalBinary(b []byte) error {
	var res ObjectMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
