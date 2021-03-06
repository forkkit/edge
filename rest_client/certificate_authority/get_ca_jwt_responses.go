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

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// GetCaJwtReader is a Reader for the GetCaJwt structure.
type GetCaJwtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCaJwtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCaJwtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCaJwtUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCaJwtNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCaJwtOK creates a GetCaJwtOK with default headers values
func NewGetCaJwtOK() *GetCaJwtOK {
	return &GetCaJwtOK{}
}

/*GetCaJwtOK handles this case with default header values.

The result is the JWT text to validate the CA
*/
type GetCaJwtOK struct {
	Payload string
}

func (o *GetCaJwtOK) Error() string {
	return fmt.Sprintf("[GET /cas/{id}/jwt][%d] getCaJwtOK  %+v", 200, o.Payload)
}

func (o *GetCaJwtOK) GetPayload() string {
	return o.Payload
}

func (o *GetCaJwtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCaJwtUnauthorized creates a GetCaJwtUnauthorized with default headers values
func NewGetCaJwtUnauthorized() *GetCaJwtUnauthorized {
	return &GetCaJwtUnauthorized{}
}

/*GetCaJwtUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type GetCaJwtUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *GetCaJwtUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cas/{id}/jwt][%d] getCaJwtUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCaJwtUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetCaJwtUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCaJwtNotFound creates a GetCaJwtNotFound with default headers values
func NewGetCaJwtNotFound() *GetCaJwtNotFound {
	return &GetCaJwtNotFound{}
}

/*GetCaJwtNotFound handles this case with default header values.

The requested resource does not exist
*/
type GetCaJwtNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *GetCaJwtNotFound) Error() string {
	return fmt.Sprintf("[GET /cas/{id}/jwt][%d] getCaJwtNotFound  %+v", 404, o.Payload)
}

func (o *GetCaJwtNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetCaJwtNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
