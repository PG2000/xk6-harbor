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

// Schedule schedule
//
// swagger:model Schedule
type Schedule struct {

	// the creation time of the schedule.
	// Read Only: true
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty" js:"creationTime"`

	// The id of the schedule.
	// Read Only: true
	ID int64 `json:"id,omitempty" js:"id"`

	// The parameters of schedule job
	Parameters map[string]interface{} `json:"parameters,omitempty" js:"parameters"`

	// schedule
	Schedule *ScheduleObj `json:"schedule,omitempty" js:"schedule"`

	// The status of the schedule.
	// Read Only: true
	Status string `json:"status,omitempty" js:"status"`

	// the update time of the schedule.
	// Read Only: true
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty" js:"updateTime"`
}

// Validate validates this schedule
func (m *Schedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schedule) validateCreationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *Schedule) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schedule based on the context it is used
func (m *Schedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreationTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schedule) contextValidateCreationTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "creation_time", "body", strfmt.DateTime(m.CreationTime)); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *Schedule) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "update_time", "body", strfmt.DateTime(m.UpdateTime)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schedule) UnmarshalBinary(b []byte) error {
	var res Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
