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
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1ConfigResourceKind config resource kinds
// swagger:model v1ConfigResourceKind
type V1ConfigResourceKind string

const (

	// V1ConfigResourceKindConfigMap captures enum value "config_map"
	V1ConfigResourceKindConfigMap V1ConfigResourceKind = "config_map"

	// V1ConfigResourceKindSecret captures enum value "secret"
	V1ConfigResourceKindSecret V1ConfigResourceKind = "secret"

	// V1ConfigResourceKindServiceAccount captures enum value "service_account"
	V1ConfigResourceKindServiceAccount V1ConfigResourceKind = "service_account"

	// V1ConfigResourceKindImagePullSecret captures enum value "image_pull_secret"
	V1ConfigResourceKindImagePullSecret V1ConfigResourceKind = "image_pull_secret"
)

// for schema
var v1ConfigResourceKindEnum []interface{}

func init() {
	var res []V1ConfigResourceKind
	if err := json.Unmarshal([]byte(`["config_map","secret","service_account","image_pull_secret"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ConfigResourceKindEnum = append(v1ConfigResourceKindEnum, v)
	}
}

func (m V1ConfigResourceKind) validateV1ConfigResourceKindEnum(path, location string, value V1ConfigResourceKind) error {
	if err := validate.Enum(path, location, value, v1ConfigResourceKindEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 config resource kind
func (m V1ConfigResourceKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ConfigResourceKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
