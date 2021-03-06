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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetRunSettingsReader is a Reader for the GetRunSettings structure.
type GetRunSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunSettingsOK creates a GetRunSettingsOK with default headers values
func NewGetRunSettingsOK() *GetRunSettingsOK {
	return &GetRunSettingsOK{}
}

/*GetRunSettingsOK handles this case with default header values.

A successful response.
*/
type GetRunSettingsOK struct {
	Payload *service_model.V1RunSettings
}

func (o *GetRunSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/settings][%d] getRunSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRunSettingsOK) GetPayload() *service_model.V1RunSettings {
	return o.Payload
}

func (o *GetRunSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1RunSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunSettingsNoContent creates a GetRunSettingsNoContent with default headers values
func NewGetRunSettingsNoContent() *GetRunSettingsNoContent {
	return &GetRunSettingsNoContent{}
}

/*GetRunSettingsNoContent handles this case with default header values.

No content.
*/
type GetRunSettingsNoContent struct {
	Payload interface{}
}

func (o *GetRunSettingsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/settings][%d] getRunSettingsNoContent  %+v", 204, o.Payload)
}

func (o *GetRunSettingsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunSettingsForbidden creates a GetRunSettingsForbidden with default headers values
func NewGetRunSettingsForbidden() *GetRunSettingsForbidden {
	return &GetRunSettingsForbidden{}
}

/*GetRunSettingsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetRunSettingsForbidden struct {
	Payload interface{}
}

func (o *GetRunSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/settings][%d] getRunSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRunSettingsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunSettingsNotFound creates a GetRunSettingsNotFound with default headers values
func NewGetRunSettingsNotFound() *GetRunSettingsNotFound {
	return &GetRunSettingsNotFound{}
}

/*GetRunSettingsNotFound handles this case with default header values.

Resource does not exist.
*/
type GetRunSettingsNotFound struct {
	Payload interface{}
}

func (o *GetRunSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/settings][%d] getRunSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunSettingsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
