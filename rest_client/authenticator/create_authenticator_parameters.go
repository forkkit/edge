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

package authenticator

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

	"github.com/openziti/edge/rest_model"
)

// NewCreateAuthenticatorParams creates a new CreateAuthenticatorParams object
// with the default values initialized.
func NewCreateAuthenticatorParams() *CreateAuthenticatorParams {
	var ()
	return &CreateAuthenticatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthenticatorParamsWithTimeout creates a new CreateAuthenticatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAuthenticatorParamsWithTimeout(timeout time.Duration) *CreateAuthenticatorParams {
	var ()
	return &CreateAuthenticatorParams{

		timeout: timeout,
	}
}

// NewCreateAuthenticatorParamsWithContext creates a new CreateAuthenticatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAuthenticatorParamsWithContext(ctx context.Context) *CreateAuthenticatorParams {
	var ()
	return &CreateAuthenticatorParams{

		Context: ctx,
	}
}

// NewCreateAuthenticatorParamsWithHTTPClient creates a new CreateAuthenticatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAuthenticatorParamsWithHTTPClient(client *http.Client) *CreateAuthenticatorParams {
	var ()
	return &CreateAuthenticatorParams{
		HTTPClient: client,
	}
}

/*CreateAuthenticatorParams contains all the parameters to send to the API endpoint
for the create authenticator operation typically these are written to a http.Request
*/
type CreateAuthenticatorParams struct {

	/*Body
	  A Authenticators create object

	*/
	Body *rest_model.AuthenticatorCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create authenticator params
func (o *CreateAuthenticatorParams) WithTimeout(timeout time.Duration) *CreateAuthenticatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create authenticator params
func (o *CreateAuthenticatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create authenticator params
func (o *CreateAuthenticatorParams) WithContext(ctx context.Context) *CreateAuthenticatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create authenticator params
func (o *CreateAuthenticatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create authenticator params
func (o *CreateAuthenticatorParams) WithHTTPClient(client *http.Client) *CreateAuthenticatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create authenticator params
func (o *CreateAuthenticatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create authenticator params
func (o *CreateAuthenticatorParams) WithBody(body *rest_model.AuthenticatorCreate) *CreateAuthenticatorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create authenticator params
func (o *CreateAuthenticatorParams) SetBody(body *rest_model.AuthenticatorCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthenticatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
