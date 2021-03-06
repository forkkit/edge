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

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CheckDataIntegrityOKCode is the HTTP code returned for type CheckDataIntegrityOK
const CheckDataIntegrityOKCode int = 200

/*CheckDataIntegrityOK A list of data integrity issues found

swagger:response checkDataIntegrityOK
*/
type CheckDataIntegrityOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DataIntegrityCheckResultEnvelope `json:"body,omitempty"`
}

// NewCheckDataIntegrityOK creates CheckDataIntegrityOK with default headers values
func NewCheckDataIntegrityOK() *CheckDataIntegrityOK {

	return &CheckDataIntegrityOK{}
}

// WithPayload adds the payload to the check data integrity o k response
func (o *CheckDataIntegrityOK) WithPayload(payload *rest_model.DataIntegrityCheckResultEnvelope) *CheckDataIntegrityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check data integrity o k response
func (o *CheckDataIntegrityOK) SetPayload(payload *rest_model.DataIntegrityCheckResultEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDataIntegrityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDataIntegrityUnauthorizedCode is the HTTP code returned for type CheckDataIntegrityUnauthorized
const CheckDataIntegrityUnauthorizedCode int = 401

/*CheckDataIntegrityUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response checkDataIntegrityUnauthorized
*/
type CheckDataIntegrityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCheckDataIntegrityUnauthorized creates CheckDataIntegrityUnauthorized with default headers values
func NewCheckDataIntegrityUnauthorized() *CheckDataIntegrityUnauthorized {

	return &CheckDataIntegrityUnauthorized{}
}

// WithPayload adds the payload to the check data integrity unauthorized response
func (o *CheckDataIntegrityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CheckDataIntegrityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check data integrity unauthorized response
func (o *CheckDataIntegrityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDataIntegrityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
