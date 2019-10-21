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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EnableProjectCIReader is a Reader for the EnableProjectCI structure.
type EnableProjectCIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableProjectCIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableProjectCIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewEnableProjectCIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableProjectCINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEnableProjectCIOK creates a EnableProjectCIOK with default headers values
func NewEnableProjectCIOK() *EnableProjectCIOK {
	return &EnableProjectCIOK{}
}

/*EnableProjectCIOK handles this case with default header values.

A successful response.
*/
type EnableProjectCIOK struct {
	Payload interface{}
}

func (o *EnableProjectCIOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/ci][%d] enableProjectCIOK  %+v", 200, o.Payload)
}

func (o *EnableProjectCIOK) GetPayload() interface{} {
	return o.Payload
}

func (o *EnableProjectCIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableProjectCIForbidden creates a EnableProjectCIForbidden with default headers values
func NewEnableProjectCIForbidden() *EnableProjectCIForbidden {
	return &EnableProjectCIForbidden{}
}

/*EnableProjectCIForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type EnableProjectCIForbidden struct {
	Payload interface{}
}

func (o *EnableProjectCIForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/ci][%d] enableProjectCIForbidden  %+v", 403, o.Payload)
}

func (o *EnableProjectCIForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *EnableProjectCIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableProjectCINotFound creates a EnableProjectCINotFound with default headers values
func NewEnableProjectCINotFound() *EnableProjectCINotFound {
	return &EnableProjectCINotFound{}
}

/*EnableProjectCINotFound handles this case with default header values.

Resource does not exist.
*/
type EnableProjectCINotFound struct {
	Payload string
}

func (o *EnableProjectCINotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/ci][%d] enableProjectCINotFound  %+v", 404, o.Payload)
}

func (o *EnableProjectCINotFound) GetPayload() string {
	return o.Payload
}

func (o *EnableProjectCINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
