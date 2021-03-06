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
)

// GetRunLogsFileReader is a Reader for the GetRunLogsFile structure.
type GetRunLogsFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunLogsFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunLogsFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunLogsFileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunLogsFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunLogsFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunLogsFileOK creates a GetRunLogsFileOK with default headers values
func NewGetRunLogsFileOK() *GetRunLogsFileOK {
	return &GetRunLogsFileOK{}
}

/*GetRunLogsFileOK handles this case with default header values.

A successful response.
*/
type GetRunLogsFileOK struct {
	Payload string
}

func (o *GetRunLogsFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/logs/file][%d] getRunLogsFileOK  %+v", 200, o.Payload)
}

func (o *GetRunLogsFileOK) GetPayload() string {
	return o.Payload
}

func (o *GetRunLogsFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunLogsFileNoContent creates a GetRunLogsFileNoContent with default headers values
func NewGetRunLogsFileNoContent() *GetRunLogsFileNoContent {
	return &GetRunLogsFileNoContent{}
}

/*GetRunLogsFileNoContent handles this case with default header values.

No content.
*/
type GetRunLogsFileNoContent struct {
	Payload interface{}
}

func (o *GetRunLogsFileNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/logs/file][%d] getRunLogsFileNoContent  %+v", 204, o.Payload)
}

func (o *GetRunLogsFileNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunLogsFileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunLogsFileForbidden creates a GetRunLogsFileForbidden with default headers values
func NewGetRunLogsFileForbidden() *GetRunLogsFileForbidden {
	return &GetRunLogsFileForbidden{}
}

/*GetRunLogsFileForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetRunLogsFileForbidden struct {
	Payload interface{}
}

func (o *GetRunLogsFileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/logs/file][%d] getRunLogsFileForbidden  %+v", 403, o.Payload)
}

func (o *GetRunLogsFileForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunLogsFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunLogsFileNotFound creates a GetRunLogsFileNotFound with default headers values
func NewGetRunLogsFileNotFound() *GetRunLogsFileNotFound {
	return &GetRunLogsFileNotFound{}
}

/*GetRunLogsFileNotFound handles this case with default header values.

Resource does not exist.
*/
type GetRunLogsFileNotFound struct {
	Payload interface{}
}

func (o *GetRunLogsFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/logs/file][%d] getRunLogsFileNotFound  %+v", 404, o.Payload)
}

func (o *GetRunLogsFileNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunLogsFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
