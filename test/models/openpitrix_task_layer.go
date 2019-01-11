// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixTaskLayer openpitrix task layer
// swagger:model openpitrixTaskLayer
type OpenpitrixTaskLayer struct {

	// child
	Child *OpenpitrixTaskLayer `json:"child,omitempty"`

	// tasks
	Tasks OpenpitrixTaskLayerTasks `json:"tasks"`
}

// Validate validates this openpitrix task layer
func (m *OpenpitrixTaskLayer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChild(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixTaskLayer) validateChild(formats strfmt.Registry) error {

	if swag.IsZero(m.Child) { // not required
		return nil
	}

	if m.Child != nil {

		if err := m.Child.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("child")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixTaskLayer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixTaskLayer) UnmarshalBinary(b []byte) error {
	var res OpenpitrixTaskLayer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
