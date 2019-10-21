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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListBookmarkedRunsParams creates a new ListBookmarkedRunsParams object
// with the default values initialized.
func NewListBookmarkedRunsParams() *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBookmarkedRunsParamsWithTimeout creates a new ListBookmarkedRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBookmarkedRunsParamsWithTimeout(timeout time.Duration) *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{

		timeout: timeout,
	}
}

// NewListBookmarkedRunsParamsWithContext creates a new ListBookmarkedRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBookmarkedRunsParamsWithContext(ctx context.Context) *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{

		Context: ctx,
	}
}

// NewListBookmarkedRunsParamsWithHTTPClient creates a new ListBookmarkedRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBookmarkedRunsParamsWithHTTPClient(client *http.Client) *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{
		HTTPClient: client,
	}
}

/*ListBookmarkedRunsParams contains all the parameters to send to the API endpoint
for the list bookmarked runs operation typically these are written to a http.Request
*/
type ListBookmarkedRunsParams struct {

	/*Limit
	  Limit size.

	*/
	Limit *int32
	/*Offset
	  Pagination offset.

	*/
	Offset *int32
	/*Query
	  Query filter the search search.

	*/
	Query *string
	/*Sort
	  Sort to order the search.

	*/
	Sort *string
	/*User
	  User

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithTimeout(timeout time.Duration) *ListBookmarkedRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithContext(ctx context.Context) *ListBookmarkedRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithHTTPClient(client *http.Client) *ListBookmarkedRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithLimit(limit *int32) *ListBookmarkedRunsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithOffset(offset *int32) *ListBookmarkedRunsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithQuery adds the query to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithQuery(query *string) *ListBookmarkedRunsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithSort(sort *string) *ListBookmarkedRunsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithUser adds the user to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithUser(user string) *ListBookmarkedRunsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *ListBookmarkedRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
