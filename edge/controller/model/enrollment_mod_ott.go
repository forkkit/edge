/*
	Copyright 2019 Netfoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package model

import (
	"github.com/netfoundry/ziti-edge/edge/controller/apierror"
	"github.com/netfoundry/ziti-edge/edge/controller/persistence"
	"github.com/netfoundry/ziti-edge/edge/internal/cert"
	"encoding/pem"
	"github.com/google/uuid"
)

type EnrollModuleOtt struct {
	env                  Env
	method               string
	fingerprintGenerator cert.FingerprintGenerator
}

func NewEnrollModuleOtt(env Env) *EnrollModuleOtt {
	handler := &EnrollModuleOtt{
		env:                  env,
		method:               persistence.MethodEnrollOtt,
		fingerprintGenerator: cert.NewFingerprintGenerator(),
	}

	return handler
}

func (module *EnrollModuleOtt) CanHandle(method string) bool {
	return method == module.method
}

func (module *EnrollModuleOtt) Process(ctx EnrollmentContext) (*EnrollmentResult, error) {
	enrollment, err := module.env.GetHandlers().Enrollment.HandleReadByToken(ctx.GetToken())
	if err != nil {
		return nil, err
	}

	if enrollment == nil {
		return nil, apierror.NewInvalidEnrollmentToken()
	}

	identity, err := module.env.GetHandlers().Identity.HandleRead(enrollment.IdentityId)

	if err != nil {
		return nil, err
	}

	if identity == nil {
		return nil, apierror.NewInvalidEnrollmentToken()
	}

	certRaw, err := module.env.GetApiClientCsrSigner().Sign(ctx.GetDataAsByteArray(), &cert.SigningOpts{})

	if err != nil {
		apiErr := apierror.NewCouldNotProcessCsr()
		apiErr.Cause = err
		apiErr.AppendCause = true
		return nil, apiErr
	}

	fp := module.fingerprintGenerator.FromRaw(certRaw)

	newAuthenticator := &Authenticator{
		BaseModelEntityImpl: BaseModelEntityImpl{
			Id: uuid.New().String(),
		},
		Method:     persistence.MethodAuthenticatorCert,
		IdentityId: enrollment.IdentityId,
		SubType: &AuthenticatorCert{
			Fingerprint: fp,
		},
	}

	b := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certRaw,
	}
	certPem := pem.EncodeToMemory(b)

	err = module.env.GetHandlers().Enrollment.HandleReplaceWithAuthenticator(enrollment.Id, newAuthenticator)

	if err != nil {
		return nil, err
	}

	return &EnrollmentResult{
		Identity:      identity,
		Authenticator: newAuthenticator,
		Content:       certPem,
		ContentType:   "application/x-pem-file",
		Status:        200,
	}, nil

}