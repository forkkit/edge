// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package service_edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDetailServiceEdgeRouterPolicyParams creates a new DetailServiceEdgeRouterPolicyParams object
// with the default values initialized.
func NewDetailServiceEdgeRouterPolicyParams() *DetailServiceEdgeRouterPolicyParams {
	var ()
	return &DetailServiceEdgeRouterPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetailServiceEdgeRouterPolicyParamsWithTimeout creates a new DetailServiceEdgeRouterPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetailServiceEdgeRouterPolicyParamsWithTimeout(timeout time.Duration) *DetailServiceEdgeRouterPolicyParams {
	var ()
	return &DetailServiceEdgeRouterPolicyParams{

		timeout: timeout,
	}
}

// NewDetailServiceEdgeRouterPolicyParamsWithContext creates a new DetailServiceEdgeRouterPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetailServiceEdgeRouterPolicyParamsWithContext(ctx context.Context) *DetailServiceEdgeRouterPolicyParams {
	var ()
	return &DetailServiceEdgeRouterPolicyParams{

		Context: ctx,
	}
}

// NewDetailServiceEdgeRouterPolicyParamsWithHTTPClient creates a new DetailServiceEdgeRouterPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDetailServiceEdgeRouterPolicyParamsWithHTTPClient(client *http.Client) *DetailServiceEdgeRouterPolicyParams {
	var ()
	return &DetailServiceEdgeRouterPolicyParams{
		HTTPClient: client,
	}
}

/*DetailServiceEdgeRouterPolicyParams contains all the parameters to send to the API endpoint
for the detail service edge router policy operation typically these are written to a http.Request
*/
type DetailServiceEdgeRouterPolicyParams struct {

	/*ID
	  The id of the requested resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) WithTimeout(timeout time.Duration) *DetailServiceEdgeRouterPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) WithContext(ctx context.Context) *DetailServiceEdgeRouterPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) WithHTTPClient(client *http.Client) *DetailServiceEdgeRouterPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) WithID(id string) *DetailServiceEdgeRouterPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the detail service edge router policy params
func (o *DetailServiceEdgeRouterPolicyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DetailServiceEdgeRouterPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}