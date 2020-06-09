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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// DetailIdentityReader is a Reader for the DetailIdentity structure.
type DetailIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailIdentityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetailIdentityOK creates a DetailIdentityOK with default headers values
func NewDetailIdentityOK() *DetailIdentityOK {
	return &DetailIdentityOK{}
}

/*DetailIdentityOK handles this case with default header values.

A signle identity
*/
type DetailIdentityOK struct {
	Payload *rest_model.DetailIdentityEnvelope
}

func (o *DetailIdentityOK) Error() string {
	return fmt.Sprintf("[GET /identities/{id}][%d] detailIdentityOK  %+v", 200, o.Payload)
}

func (o *DetailIdentityOK) GetPayload() *rest_model.DetailIdentityEnvelope {
	return o.Payload
}

func (o *DetailIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailIdentityEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailIdentityUnauthorized creates a DetailIdentityUnauthorized with default headers values
func NewDetailIdentityUnauthorized() *DetailIdentityUnauthorized {
	return &DetailIdentityUnauthorized{}
}

/*DetailIdentityUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DetailIdentityUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailIdentityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /identities/{id}][%d] detailIdentityUnauthorized  %+v", 401, o.Payload)
}

func (o *DetailIdentityUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailIdentityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailIdentityNotFound creates a DetailIdentityNotFound with default headers values
func NewDetailIdentityNotFound() *DetailIdentityNotFound {
	return &DetailIdentityNotFound{}
}

/*DetailIdentityNotFound handles this case with default header values.

The requested resource does not exist
*/
type DetailIdentityNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailIdentityNotFound) Error() string {
	return fmt.Sprintf("[GET /identities/{id}][%d] detailIdentityNotFound  %+v", 404, o.Payload)
}

func (o *DetailIdentityNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}