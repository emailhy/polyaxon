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

package queues_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateQueueReader is a Reader for the CreateQueue structure.
type CreateQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateQueueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateQueueOK creates a CreateQueueOK with default headers values
func NewCreateQueueOK() *CreateQueueOK {
	return &CreateQueueOK{}
}

/*CreateQueueOK handles this case with default header values.

A successful response.
*/
type CreateQueueOK struct {
	Payload *service_model.V1Agent
}

func (o *CreateQueueOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] createQueueOK  %+v", 200, o.Payload)
}

func (o *CreateQueueOK) GetPayload() *service_model.V1Agent {
	return o.Payload
}

func (o *CreateQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueueNoContent creates a CreateQueueNoContent with default headers values
func NewCreateQueueNoContent() *CreateQueueNoContent {
	return &CreateQueueNoContent{}
}

/*CreateQueueNoContent handles this case with default header values.

No content.
*/
type CreateQueueNoContent struct {
	Payload interface{}
}

func (o *CreateQueueNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] createQueueNoContent  %+v", 204, o.Payload)
}

func (o *CreateQueueNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateQueueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueueForbidden creates a CreateQueueForbidden with default headers values
func NewCreateQueueForbidden() *CreateQueueForbidden {
	return &CreateQueueForbidden{}
}

/*CreateQueueForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateQueueForbidden struct {
	Payload interface{}
}

func (o *CreateQueueForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] createQueueForbidden  %+v", 403, o.Payload)
}

func (o *CreateQueueForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueueNotFound creates a CreateQueueNotFound with default headers values
func NewCreateQueueNotFound() *CreateQueueNotFound {
	return &CreateQueueNotFound{}
}

/*CreateQueueNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateQueueNotFound struct {
	Payload interface{}
}

func (o *CreateQueueNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents/{agent}/queues][%d] createQueueNotFound  %+v", 404, o.Payload)
}

func (o *CreateQueueNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
