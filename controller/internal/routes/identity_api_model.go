/*
	Copyright NetFoundry, Inc.

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

package routes

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/controller/model"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/edge/controller/response"
	"github.com/openziti/edge/rest_model"
	"github.com/openziti/fabric/controller/models"
	"github.com/openziti/foundation/util/stringz"
)

const (
	EntityNameIdentity              = "identities"
	EntityNameIdentityServiceConfig = "service-configs"
)

type PermissionsApi []string

var IdentityLinkFactory = NewIdentityLinkFactory(NewBasicLinkFactory(EntityNameIdentity))

func NewIdentityLinkFactory(selfFactory *BasicLinkFactory) *IdentityLinkFactoryImpl {
	return &IdentityLinkFactoryImpl{
		BasicLinkFactory: *selfFactory,
	}
}

type IdentityLinkFactoryImpl struct {
	BasicLinkFactory
}

func (factory *IdentityLinkFactoryImpl) Links(entity models.Entity) rest_model.Links {
	links := factory.BasicLinkFactory.Links(entity)
	links[EntityNameEdgeRouterPolicy] = factory.NewNestedLink(entity, EntityNameEdgeRouter)
	links[EntityNameEdgeRouterPolicy] = factory.NewNestedLink(entity, EntityNameEdgeRouter)

	return links
}

func MapCreateIdentityToModel(identity *rest_model.IdentityCreate, identityTypeId string) (*model.Identity, []*model.Enrollment) {
	var enrollments []*model.Enrollment

	ret := &model.Identity{
		BaseEntity: models.BaseEntity{
			Tags: identity.Tags,
		},
		Name:           stringz.OrEmpty(identity.Name),
		IdentityTypeId: identityTypeId,
		IsDefaultAdmin: false,
		IsAdmin:        *identity.IsAdmin,
		RoleAttributes: identity.RoleAttributes,
	}

	if identity.Enrollment != nil {
		if identity.Enrollment.Ott {
			enrollments = append(enrollments, &model.Enrollment{
				BaseEntity: models.BaseEntity{},
				Method:     persistence.MethodEnrollOtt,
				Token:      uuid.New().String(),
			})
		} else if identity.Enrollment.Ottca != "" {
			caId := identity.Enrollment.Ottca
			enrollments = append(enrollments, &model.Enrollment{
				BaseEntity: models.BaseEntity{},
				Method:     persistence.MethodEnrollOttCa,
				Token:      uuid.New().String(),
				CaId:       &caId,
			})
		} else if identity.Enrollment.Updb != "" {
			username := identity.Enrollment.Updb
			enrollments = append(enrollments, &model.Enrollment{
				BaseEntity: models.BaseEntity{},
				Method:     persistence.MethodEnrollUpdb,
				Token:      uuid.New().String(),
				Username:   &username,
			})
		}
	}

	return ret, enrollments
}

func MapUpdateIdentityToModel(id string, identity *rest_model.IdentityUpdate, identityTypeId string) *model.Identity {
	ret := &model.Identity{
		BaseEntity: models.BaseEntity{
			Tags: identity.Tags,
			Id:   id,
		},
		Name:           stringz.OrEmpty(identity.Name),
		IdentityTypeId: identityTypeId,
		IsAdmin:        *identity.IsAdmin,
		RoleAttributes: identity.RoleAttributes,
	}

	return ret
}

func MapPatchIdentityToModel(id string, identity *rest_model.IdentityPatch, identityTypeId string) *model.Identity {
	ret := &model.Identity{
		BaseEntity: models.BaseEntity{
			Tags: identity.Tags,
			Id:   id,
		},
		Name:           identity.Name,
		IdentityTypeId: identityTypeId,
		IsAdmin:        identity.IsAdmin,
		RoleAttributes: identity.RoleAttributes,
	}

	return ret
}

func MapIdentityToRestEntity(ae *env.AppEnv, _ *response.RequestContext, e models.Entity) (interface{}, error) {
	identity, ok := e.(*model.Identity)

	if !ok {
		err := fmt.Errorf("entity is not a Identity \"%s\"", e.GetId())
		log := pfxlog.Logger()
		log.Error(err)
		return nil, err
	}

	restModel, err := MapIdentityToRestModel(ae, identity)

	if err != nil {
		err := fmt.Errorf("could not convert to API entity \"%s\": %s", e.GetId(), err)
		log := pfxlog.Logger()
		log.Error(err)
		return nil, err
	}
	return restModel, nil
}

func MapIdentityToRestModel(ae *env.AppEnv, identity *model.Identity) (*rest_model.IdentityDetail, error) {

	identityType, err := ae.Handlers.IdentityType.ReadByIdOrName(identity.IdentityTypeId)

	if err != nil {
		return nil, err
	}

	ret := &rest_model.IdentityDetail{
		BaseEntity:     BaseEntityToRestModel(identity, IdentityLinkFactory),
		IsAdmin:        &identity.IsAdmin,
		IsDefaultAdmin: &identity.IsDefaultAdmin,
		Name:           &identity.Name,
		RoleAttributes: identity.RoleAttributes,
		Type:           ToEntityRef(identityType.Name, identityType, IdentityTypeLinkFactory),
		TypeID:         &identityType.Id,
	}
	fillInfo(ret, identity.EnvInfo, identity.SdkInfo)

	ret.Authenticators = &rest_model.IdentityAuthenticators{}
	if err = ae.GetHandlers().Identity.CollectAuthenticators(identity.Id, func(entity *model.Authenticator) error {
		if entity.Method == persistence.MethodAuthenticatorUpdb {
			ret.Authenticators.Updb = &rest_model.IdentityAuthenticatorsUpdb{
				Username: entity.ToUpdb().Username,
			}
		}

		if entity.Method == persistence.MethodAuthenticatorCert {
			ret.Authenticators.Cert = &rest_model.IdentityAuthenticatorsCert{
				Fingerprint: entity.ToCert().Fingerprint,
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	ret.Enrollment = &rest_model.IdentityEnrollments{}
	if err := ae.GetHandlers().Identity.CollectEnrollments(identity.Id, func(entity *model.Enrollment) error {
		var expiresAt strfmt.DateTime
		if entity.ExpiresAt != nil {
			expiresAt = strfmt.DateTime(*entity.ExpiresAt)
		}

		if entity.Method == persistence.MethodEnrollUpdb {

			ret.Enrollment.Updb = &rest_model.IdentityEnrollmentsUpdb{
				Jwt:       entity.Jwt,
				Token:     entity.Token,
				ExpiresAt: expiresAt,
			}
		}

		if entity.Method == persistence.MethodEnrollOtt {
			ret.Enrollment.Ott = &rest_model.IdentityEnrollmentsOtt{
				Jwt:   entity.Jwt,
				Token: entity.Token,
				ExpiresAt: expiresAt,
			}
		}

		if entity.Method == persistence.MethodEnrollOttCa {
			ca, err := ae.Handlers.Ca.Read(*entity.CaId)

			if err != nil {
				return err
			}

			ret.Enrollment.Ottca = &rest_model.IdentityEnrollmentsOttca{
				Ca:    ToEntityRef(ca.Name, ca, CaLinkFactory),
				CaID:  ca.Id,
				Jwt:   entity.Jwt,
				Token: entity.Token,
				ExpiresAt: expiresAt,
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func fillInfo(identity *rest_model.IdentityDetail, envInfo *model.EnvInfo, sdkInfo *model.SdkInfo) {
	if envInfo != nil {
		identity.EnvInfo = &rest_model.EnvInfo{
			Arch:      envInfo.Arch,
			Os:        envInfo.Os,
			OsRelease: envInfo.OsRelease,
			OsVersion: envInfo.OsVersion,
		}
	} else {
		identity.EnvInfo = &rest_model.EnvInfo{}
	}

	if sdkInfo != nil {
		identity.SdkInfo = &rest_model.SdkInfo{
			Branch:   sdkInfo.Branch,
			Revision: sdkInfo.Revision,
			Type:     sdkInfo.Type,
			Version:  sdkInfo.Version,
		}
	} else {
		identity.SdkInfo = &rest_model.SdkInfo{}
	}
}

func MapServiceConfigToModel(config rest_model.ServiceConfigAssign) model.ServiceConfig {
	return model.ServiceConfig{
		Service: stringz.OrEmpty(config.ServiceID),
		Config:  stringz.OrEmpty(config.ConfigID),
	}
}
func MapAdvisorServiceReachabilityToRestEntity(entity *model.AdvisorServiceReachability) *rest_model.PolicyAdvice {

	var commonRouters []*rest_model.RouterEntityRef

	for _, router := range entity.CommonRouters {
		commonRouters = append(commonRouters, &rest_model.RouterEntityRef{
			EntityRef: *ToEntityRef(router.Router.Name, router.Router, EdgeRouterLinkFactory),
			IsOnline:  &router.IsOnline,
		})
	}

	result := &rest_model.PolicyAdvice{
		IdentityID:          entity.Identity.Id,
		Identity:            ToEntityRef(entity.Identity.Name, entity.Identity, IdentityLinkFactory),
		ServiceID:           entity.Service.Id,
		Service:             ToEntityRef(entity.Service.Name, entity.Service, ServiceLinkFactory),
		IsBindAllowed:       entity.IsBindAllowed,
		IsDialAllowed:       entity.IsDialAllowed,
		IdentityRouterCount: int32(entity.IdentityRouterCount),
		ServiceRouterCount:  int32(entity.ServiceRouterCount),
		CommonRouters:       commonRouters,
	}

	return result
}
