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

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// UpdateRunReader is a Reader for the UpdateRun structure.
type UpdateRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRunOK creates a UpdateRunOK with default headers values
func NewUpdateRunOK() *UpdateRunOK {
	return &UpdateRunOK{}
}

/*UpdateRunOK handles this case with default header values.

A successful response.
*/
type UpdateRunOK struct {
	Payload *service_model.V1Run
}

func (o *UpdateRunOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunOK  %+v", 200, o.Payload)
}

func (o *UpdateRunOK) GetPayload() *service_model.V1Run {
	return o.Payload
}

func (o *UpdateRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunForbidden creates a UpdateRunForbidden with default headers values
func NewUpdateRunForbidden() *UpdateRunForbidden {
	return &UpdateRunForbidden{}
}

/*UpdateRunForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateRunForbidden struct {
	Payload interface{}
}

func (o *UpdateRunForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunNotFound creates a UpdateRunNotFound with default headers values
func NewUpdateRunNotFound() *UpdateRunNotFound {
	return &UpdateRunNotFound{}
}

/*UpdateRunNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateRunNotFound struct {
	Payload string
}

func (o *UpdateRunNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRunNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
