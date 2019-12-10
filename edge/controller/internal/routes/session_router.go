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

package routes

import (
	"github.com/netfoundry/ziti-edge/edge/controller/env"
	"github.com/netfoundry/ziti-edge/edge/controller/internal/permissions"
	"github.com/netfoundry/ziti-edge/edge/controller/model"
	"github.com/netfoundry/ziti-edge/edge/controller/response"
	"github.com/netfoundry/ziti-edge/edge/controller/util"

	"github.com/Jeffail/gabs"
	"github.com/michaelquigley/pfxlog"
)

func init() {
	r := NewSessionRouter()
	env.AddRouter(r)
}

type SessionRouter struct {
	BasePath string
	IdType   response.IdType
}

func NewSessionRouter() *SessionRouter {
	return &SessionRouter{
		BasePath: "/" + EntityNameNetworkSession,
		IdType:   response.IdTypeUuid,
	}
}

func (ir *SessionRouter) Register(ae *env.AppEnv) {
	registerCreateReadDeleteRouter(ae, ae.RootRouter, ir.BasePath, ir, &crudResolvers{
		Create:  permissions.IsAuthenticated(),
		Read:    permissions.IsAuthenticated(),
		Delete:  permissions.IsAuthenticated(),
		Default: permissions.IsAdmin(),
	})
}

func (ir *SessionRouter) List(ae *env.AppEnv, rc *response.RequestContext) {
	// ListWithHandler won't do search limiting by logged in user
	List(rc, func(rc *response.RequestContext, queryOptions *model.QueryOptions) (*QueryResult, error) {
		result, err := ae.Handlers.Session.HandleListForIdentity(rc.Identity, queryOptions)
		if err != nil {
			return nil, err
		}
		sessions, err := MapSessionsToApiEntities(ae, rc, result.Sessions)
		if err != nil {
			return nil, err
		}
		return NewQueryResult(sessions, &result.QueryMetaData), nil
	})
}

func (ir *SessionRouter) Detail(ae *env.AppEnv, rc *response.RequestContext) {
	// DetailWithHandler won't do search limiting by logged in user
	Detail(rc, ir.IdType, func(rc *response.RequestContext, id string) (BaseApiEntity, error) {
		service, err := ae.Handlers.Session.HandleReadForIdentity(id, rc.Session.IdentityId)
		if err != nil {
			return nil, err
		}
		return MapSessionToApiEntity(ae, rc, service)
	})
}

func (ir *SessionRouter) Delete(ae *env.AppEnv, rc *response.RequestContext) {
	Delete(rc, ir.IdType, func(rc *response.RequestContext, id string) error {
		return ae.Handlers.Session.HandleDeleteForIdentity(id, rc.Session.IdentityId)
	})
}

func (ir *SessionRouter) Create(ae *env.AppEnv, rc *response.RequestContext) {
	//todo re-enable this check w/ a new auth table or allow any auth'ed session to have a short term NS cert
	//if rc.Identity.AuthenticatorCert == nil {
	//	rc.RequestResponder.RespondWithApiError(&response.ApiError{
	//		Code:           response.NetworkSessionsRequireCertificateAuthCode,
	//		Message:        response.NetworkSessionsRequireCertificateAuthMessage,
	//		HttpStatusCode: http.StatusBadRequest,
	//	})
	//	return
	//}

	sessionCreate := &SessionApiPost{}
	responder := &SessionRequestResponder{ae: ae, RequestResponder: rc.RequestResponder}
	Create(rc, responder, ae.Schemes.NetworkSession.Post, sessionCreate, (&SessionApiList{}).BuildSelfLink, func() (string, error) {
		return ae.Handlers.Session.HandleCreate(sessionCreate.ToModel(rc))
	})
}

type SessionRequestResponder struct {
	response.RequestResponder
	ae *env.AppEnv
}

type SessionEdgeRouter struct {
	Hostname *string           `json:"hostname"`
	Name     *string           `json:"name"`
	Urls     map[string]string `json:"urls"`
}

func getSessionEdgeRouters(ae *env.AppEnv, ns *model.Session) ([]*SessionEdgeRouter, error) {
	var edgeRouters []*SessionEdgeRouter

	service, err := ae.Handlers.Service.HandleRead(ns.ServiceId)
	if err != nil {
		return nil, err
	}
	for _, c := range service.Clusters {
		gs := ae.Broker.GetOnlineEdgeRoutersByCluster(c)

		for _, g := range gs {
			c := &SessionEdgeRouter{
				Hostname: g.Hostname,
				Name:     &g.Name,
				Urls:     map[string]string{},
			}

			for p, url := range g.EdgeRouterProtocols {
				c.Urls[p] = url
			}

			edgeRouters = append(edgeRouters, c)
		}
	}

	return edgeRouters, nil
}

func (nsr *SessionRequestResponder) RespondWithCreatedId(id string, link *response.Link) {
	ns, _ := nsr.ae.GetHandlers().Session.HandleRead(id)

	gws, err := getSessionEdgeRouters(nsr.ae, ns)
	if err != nil {
		if util.IsErrNotFoundErr(err) {
			nsr.RespondWithNotFound()
		} else {
			nsr.RespondWithError(err)
		}
		return
	}

	json := gabs.New()

	if _, err := json.SetP(id, "id"); err != nil {
		pfxlog.Logger().WithField("cause", err).Error("could not set value by path")
	}

	if _, err := json.SetP(gws, "gateways"); err != nil {
		pfxlog.Logger().WithField("cause", err).Error("could not set value by path")
	}

	if _, err := json.SetP(map[string]*response.Link{"self": link}, "_links"); err != nil {
		pfxlog.Logger().WithField("cause", err).Error("could not set value by path")
	}

	if _, err := json.SetP(ns.Token, "token"); err != nil {
		pfxlog.Logger().WithField("cause", err).Error("could not set value by path")
	}

	if _, err := json.SetP(ns.IsHosting, "hosting"); err != nil {
		pfxlog.Logger().WithField("cause", err).Error("could not set value by path")
	}

	nsr.RequestResponder.RespondWithCreated(json.Data(), nil, link)
}