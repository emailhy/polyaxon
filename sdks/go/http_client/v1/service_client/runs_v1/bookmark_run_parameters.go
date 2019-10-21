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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBookmarkRunParams creates a new BookmarkRunParams object
// with the default values initialized.
func NewBookmarkRunParams() *BookmarkRunParams {
	var ()
	return &BookmarkRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBookmarkRunParamsWithTimeout creates a new BookmarkRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBookmarkRunParamsWithTimeout(timeout time.Duration) *BookmarkRunParams {
	var ()
	return &BookmarkRunParams{

		timeout: timeout,
	}
}

// NewBookmarkRunParamsWithContext creates a new BookmarkRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewBookmarkRunParamsWithContext(ctx context.Context) *BookmarkRunParams {
	var ()
	return &BookmarkRunParams{

		Context: ctx,
	}
}

// NewBookmarkRunParamsWithHTTPClient creates a new BookmarkRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBookmarkRunParamsWithHTTPClient(client *http.Client) *BookmarkRunParams {
	var ()
	return &BookmarkRunParams{
		HTTPClient: client,
	}
}

/*BookmarkRunParams contains all the parameters to send to the API endpoint
for the bookmark run operation typically these are written to a http.Request
*/
type BookmarkRunParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the experiement will be assigned

	*/
	Project string
	/*UUID
	  Unique integer identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bookmark run params
func (o *BookmarkRunParams) WithTimeout(timeout time.Duration) *BookmarkRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bookmark run params
func (o *BookmarkRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bookmark run params
func (o *BookmarkRunParams) WithContext(ctx context.Context) *BookmarkRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bookmark run params
func (o *BookmarkRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bookmark run params
func (o *BookmarkRunParams) WithHTTPClient(client *http.Client) *BookmarkRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bookmark run params
func (o *BookmarkRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the bookmark run params
func (o *BookmarkRunParams) WithOwner(owner string) *BookmarkRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the bookmark run params
func (o *BookmarkRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the bookmark run params
func (o *BookmarkRunParams) WithProject(project string) *BookmarkRunParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the bookmark run params
func (o *BookmarkRunParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the bookmark run params
func (o *BookmarkRunParams) WithUUID(uuid string) *BookmarkRunParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the bookmark run params
func (o *BookmarkRunParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *BookmarkRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
