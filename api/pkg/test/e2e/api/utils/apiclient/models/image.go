// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Image Image represents an Image returned by the Compute API.
// swagger:model Image
type Image struct {

	// Created is the date when the image was created.
	Created string `json:"Created,omitempty"`

	// ID is the unique ID of an image.
	ID string `json:"ID,omitempty"`

	// Metadata provides free-form key/value pairs that further describe the
	// image.
	Metadata map[string]interface{} `json:"Metadata,omitempty"`

	// MinDisk is the minimum amount of disk a flavor must have to be able
	// to create a server based on the image, measured in GB.
	MinDisk int64 `json:"MinDisk,omitempty"`

	// MinRAM is the minimum amount of RAM a flavor must have to be able
	// to create a server based on the image, measured in MB.
	MinRAM int64 `json:"MinRAM,omitempty"`

	// Name provides a human-readable moniker for the OS image.
	Name string `json:"Name,omitempty"`

	// The Progress and Status fields indicate image-creation status.
	Progress int64 `json:"Progress,omitempty"`

	// Status is the current status of the image.
	Status string `json:"Status,omitempty"`

	// Update is the date when the image was updated.
	Updated string `json:"Updated,omitempty"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
