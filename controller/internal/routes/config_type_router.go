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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/controller/internal/permissions"
	"github.com/openziti/edge/controller/response"
	"github.com/openziti/edge/rest_server/operations/config"
)

func init() {
	r := NewConfigTypeRouter()
	env.AddRouter(r)
}

type ConfigTypeRouter struct {
	BasePath string
}

func NewConfigTypeRouter() *ConfigTypeRouter {
	return &ConfigTypeRouter{
		BasePath: "/" + EntityNameConfigType,
	}
}

func (r *ConfigTypeRouter) Register(ae *env.AppEnv) {
	ae.Api.ConfigDeleteConfigTypeHandler = config.DeleteConfigTypeHandlerFunc(func(params config.DeleteConfigTypeParams, _ interface{}) middleware.Responder {
		return ae.IsAllowed(r.Delete, params.HTTPRequest, params.ID, "", permissions.IsAdmin())
	})

	ae.Api.ConfigDetailConfigTypeHandler = config.DetailConfigTypeHandlerFunc(func(params config.DetailConfigTypeParams, _ interface{}) middleware.Responder {
		return ae.IsAllowed(r.Detail, params.HTTPRequest, params.ID, "", permissions.IsAdmin())
	})

	ae.Api.ConfigListConfigTypesHandler = config.ListConfigTypesHandlerFunc(func(params config.ListConfigTypesParams, _ interface{}) middleware.Responder {
		return ae.IsAllowed(r.List, params.HTTPRequest, "", "", permissions.IsAdmin())
	})

	ae.Api.ConfigUpdateConfigTypeHandler = config.UpdateConfigTypeHandlerFunc(func(params config.UpdateConfigTypeParams, _ interface{}) middleware.Responder {
		return ae.IsAllowed(func(ae *env.AppEnv, rc *response.RequestContext) { r.Update(ae, rc, params) }, params.HTTPRequest, params.ID, "", permissions.IsAdmin())
	})

	ae.Api.ConfigCreateConfigTypeHandler = config.CreateConfigTypeHandlerFunc(func(params config.CreateConfigTypeParams, _ interface{}) middleware.Responder {
		return ae.IsAllowed(func(ae *env.AppEnv, rc *response.RequestContext) { r.Create(ae, rc, params) }, params.HTTPRequest, "", "", permissions.IsAdmin())
	})

	ae.Api.ConfigPatchConfigTypeHandler = config.PatchConfigTypeHandlerFunc(func(params config.PatchConfigTypeParams, _ interface{}) middleware.Responder {
		return ae.IsAllowed(func(ae *env.AppEnv, rc *response.RequestContext) { r.Patch(ae, rc, params) }, params.HTTPRequest, params.ID, "", permissions.IsAdmin())
	})

	ae.Api.ConfigListConfigsForConfigTypeHandler = config.ListConfigsForConfigTypeHandlerFunc(func(params config.ListConfigsForConfigTypeParams, _ interface{}) middleware.Responder {
		return ae.IsAllowed(func(ae *env.AppEnv, rc *response.RequestContext) { r.ListConfigs(ae, rc, params) }, params.HTTPRequest, "", "", permissions.IsAdmin())
	})
}

func (r *ConfigTypeRouter) List(ae *env.AppEnv, rc *response.RequestContext) {
	ListWithHandler(ae, rc, ae.Handlers.ConfigType, MapConfigTypeToRestEntity)
}

func (r *ConfigTypeRouter) Detail(ae *env.AppEnv, rc *response.RequestContext) {
	DetailWithHandler(ae, rc, ae.Handlers.ConfigType, MapConfigTypeToRestEntity)
}

func (r *ConfigTypeRouter) Create(ae *env.AppEnv, rc *response.RequestContext, params config.CreateConfigTypeParams) {
	if params.Body.Schema != nil {
		if _, ok := params.Body.Schema.(map[string]interface{}); !ok {
			ctx := middleware.MatchedRouteFrom(rc.Request)
			ae.Api.ServeErrorFor(ctx.Operation.ID)(rc.ResponseWriter, rc.Request, errors.InvalidType("schema", "body", "object", params.Body.Schema))
			return
		}
	}

	Create(rc, rc, ConfigTypeLinkFactory, func() (string, error) {
		return ae.Handlers.ConfigType.Create(MapCreateConfigTypeToModel(params.Body))
	})
}

func (r *ConfigTypeRouter) Delete(ae *env.AppEnv, rc *response.RequestContext) {
	DeleteWithHandler(rc, ae.Handlers.ConfigType)
}

func (r *ConfigTypeRouter) Update(ae *env.AppEnv, rc *response.RequestContext, params config.UpdateConfigTypeParams) {
	if params.Body.Schema != nil {
		if _, ok := params.Body.Schema.(map[string]interface{}); !ok {
			ctx := middleware.MatchedRouteFrom(rc.Request)
			ae.Api.ServeErrorFor(ctx.Operation.ID)(rc.ResponseWriter, rc.Request, errors.InvalidType("schema", "body", "object", params.Body.Schema))
			return
		}
	}

	Update(rc, func(id string) error {
		return ae.Handlers.ConfigType.Update(MapUpdateConfigTypeToModel(params.ID, params.Body))
	})
}

func (r *ConfigTypeRouter) Patch(ae *env.AppEnv, rc *response.RequestContext, params config.PatchConfigTypeParams) {

	if _, ok := params.Body.Schema.(map[string]interface{}); !ok {
		ctx := middleware.MatchedRouteFrom(rc.Request)
		ae.Api.ServeErrorFor(ctx.Operation.ID)(rc.ResponseWriter, rc.Request, errors.InvalidType("schema", "body", "object", params.Body.Schema))
		return
	}
	if params.Body.Schema == nil {
		ctx := middleware.MatchedRouteFrom(rc.Request)
		ae.Api.ServeErrorFor(ctx.Operation.ID)(rc.ResponseWriter, rc.Request, errors.Required("schema", "body", nil))
		return
	}

	Patch(rc, func(id string, fields JsonFields) error {
		return ae.Handlers.ConfigType.Patch(MapPatchConfigTypeToModel(params.ID, params.Body), fields.FilterMaps("tags", "schema"))
	})
}

func (r *ConfigTypeRouter) ListConfigs(ae *env.AppEnv, rc *response.RequestContext, params config.ListConfigsForConfigTypeParams) {
	ListAssociationWithHandler(ae, rc, ae.Handlers.ConfigType, ae.Handlers.Config, MapConfigToRestEntity)
}
