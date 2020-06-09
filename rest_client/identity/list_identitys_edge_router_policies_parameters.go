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

package identity

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

// NewListIdentitysEdgeRouterPoliciesParams creates a new ListIdentitysEdgeRouterPoliciesParams object
// with the default values initialized.
func NewListIdentitysEdgeRouterPoliciesParams() *ListIdentitysEdgeRouterPoliciesParams {
	var ()
	return &ListIdentitysEdgeRouterPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListIdentitysEdgeRouterPoliciesParamsWithTimeout creates a new ListIdentitysEdgeRouterPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIdentitysEdgeRouterPoliciesParamsWithTimeout(timeout time.Duration) *ListIdentitysEdgeRouterPoliciesParams {
	var ()
	return &ListIdentitysEdgeRouterPoliciesParams{

		timeout: timeout,
	}
}

// NewListIdentitysEdgeRouterPoliciesParamsWithContext creates a new ListIdentitysEdgeRouterPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIdentitysEdgeRouterPoliciesParamsWithContext(ctx context.Context) *ListIdentitysEdgeRouterPoliciesParams {
	var ()
	return &ListIdentitysEdgeRouterPoliciesParams{

		Context: ctx,
	}
}

// NewListIdentitysEdgeRouterPoliciesParamsWithHTTPClient creates a new ListIdentitysEdgeRouterPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIdentitysEdgeRouterPoliciesParamsWithHTTPClient(client *http.Client) *ListIdentitysEdgeRouterPoliciesParams {
	var ()
	return &ListIdentitysEdgeRouterPoliciesParams{
		HTTPClient: client,
	}
}

/*ListIdentitysEdgeRouterPoliciesParams contains all the parameters to send to the API endpoint
for the list identitys edge router policies operation typically these are written to a http.Request
*/
type ListIdentitysEdgeRouterPoliciesParams struct {

	/*ID
	  The id of the requested resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) WithTimeout(timeout time.Duration) *ListIdentitysEdgeRouterPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) WithContext(ctx context.Context) *ListIdentitysEdgeRouterPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) WithHTTPClient(client *http.Client) *ListIdentitysEdgeRouterPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) WithID(id string) *ListIdentitysEdgeRouterPoliciesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list identitys edge router policies params
func (o *ListIdentitysEdgeRouterPoliciesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListIdentitysEdgeRouterPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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