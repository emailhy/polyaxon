// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1Status Status specification
// swagger:model v1Status
type V1Status struct {

	// The current status
	Status string `json:"status,omitempty"`

	// The status conditions timeline
	StatusConditions []*V1StatusCondition `json:"status_conditions"`

	// The uuid of the run
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 status
func (m *V1Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatusConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Status) validateStatusConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusConditions); i++ {
		if swag.IsZero(m.StatusConditions[i]) { // not required
			continue
		}

		if m.StatusConditions[i] != nil {
			if err := m.StatusConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Status) UnmarshalBinary(b []byte) error {
	var res V1Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
