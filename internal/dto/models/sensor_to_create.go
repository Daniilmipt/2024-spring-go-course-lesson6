// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SensorToCreate SensorToCreate
//
// Датчик умного дома, который надо создать
// Example: {"description":"Датчик температуры","is_active":true,"serial_number":"1234567890","type":"cc"}
//
// swagger:model SensorToCreate
type SensorToCreate struct {

	// Описание
	// Required: true
	Description *string `json:"description"`

	// Флаг активности датчика
	// Required: true
	IsActive *bool `json:"is_active"`

	// Серийный номер
	// Required: true
	// Pattern: ^\d{10}$
	SerialNumber *string `json:"serial_number"`

	// Тип
	// Required: true
	// Enum: ["cc","adc"]
	Type *string `json:"type"`
}

// Validate validates this sensor to create
func (m *SensorToCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorToCreate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *SensorToCreate) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("is_active", "body", m.IsActive); err != nil {
		return err
	}

	return nil
}

func (m *SensorToCreate) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.Required("serial_number", "body", m.SerialNumber); err != nil {
		return err
	}

	if err := validate.Pattern("serial_number", "body", *m.SerialNumber, `^\d{10}$`); err != nil {
		return err
	}

	return nil
}

var sensorToCreateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cc","adc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorToCreateTypeTypePropEnum = append(sensorToCreateTypeTypePropEnum, v)
	}
}

const (

	// SensorToCreateTypeCc captures enum value "cc"
	SensorToCreateTypeCc string = "cc"

	// SensorToCreateTypeAdc captures enum value "adc"
	SensorToCreateTypeAdc string = "adc"
)

// prop value enum
func (m *SensorToCreate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorToCreateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SensorToCreate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sensor to create based on context it is used
func (m *SensorToCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SensorToCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorToCreate) UnmarshalBinary(b []byte) error {
	var res SensorToCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
