// +build apitests

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

package tests

import (
	"github.com/openziti/edge/eid"
	"testing"
)

func Test_ServicePolicy(t *testing.T) {
	ctx := NewTestContext(t)
	defer ctx.Teardown()
	ctx.StartServer()
	ctx.RequireAdminLogin()

	serviceRole1 := eid.New()
	serviceRole2 := eid.New()
	identityRole1 := eid.New()
	identityRole2 := eid.New()

	service1 := ctx.AdminSession.requireNewService(s(serviceRole1), nil)
	service2 := ctx.AdminSession.requireNewService(s(serviceRole1, serviceRole2), nil)
	service3 := ctx.AdminSession.requireNewService(nil, nil)

	identity1 := ctx.AdminSession.requireNewIdentity(false, identityRole1)
	identity2 := ctx.AdminSession.requireNewIdentity(false, identityRole1, identityRole2)
	identity3 := ctx.AdminSession.requireNewIdentity(false, identityRole2)

	policy1 := ctx.AdminSession.requireNewServicePolicy("Dial", s("#"+serviceRole1), s("#"+identityRole1))
	policy2 := ctx.AdminSession.requireNewServicePolicy("Dial", s("#"+serviceRole1, "@"+service3.Id), s("#"+identityRole1, "@"+identity3.Id))
	policy3 := ctx.AdminSession.requireNewServicePolicy("Dial", s("@"+service2.Id, "@"+service3.Id), s("@"+identity2.Id, "@"+identity3.Id))
	policy4 := ctx.AdminSession.requireNewServicePolicy("Dial", s("#all"), s("#all"))
	policy5 := ctx.AdminSession.requireNewServicePolicyWithSemantic("Dial", "AllOf", s("#"+serviceRole1, "#"+serviceRole2), s("#"+identityRole1, "#"+identityRole2))
	policy6 := ctx.AdminSession.requireNewServicePolicyWithSemantic("Dial", "AnyOf", s("#"+serviceRole1, "#"+serviceRole2), s("#"+identityRole1, "#"+identityRole2))

	ctx.AdminSession.validateEntityWithQuery(policy1)
	ctx.AdminSession.validateEntityWithLookup(policy2)
	ctx.AdminSession.validateEntityWithLookup(policy3)
	ctx.AdminSession.validateEntityWithLookup(policy4)
	ctx.AdminSession.validateEntityWithLookup(policy5)
	ctx.AdminSession.validateEntityWithLookup(policy6)

	ctx.AdminSession.validateAssociations(policy1, "services", service1, service2)
	ctx.AdminSession.validateAssociations(policy2, "services", service1, service2, service3)
	ctx.AdminSession.validateAssociations(policy3, "services", service2, service3)
	ctx.AdminSession.validateAssociationContains(policy4, "services", service1, service2, service3)
	ctx.AdminSession.validateAssociationContains(policy5, "services", service2)
	ctx.AdminSession.validateAssociationContains(policy6, "services", service1, service2)

	ctx.AdminSession.validateAssociations(policy1, "identities", identity1, identity2)
	ctx.AdminSession.validateAssociations(policy2, "identities", identity1, identity2, identity3)
	ctx.AdminSession.validateAssociations(policy3, "identities", identity2, identity3)
	ctx.AdminSession.validateAssociationContains(policy4, "identities", identity1, identity2, identity3)
	ctx.AdminSession.validateAssociations(policy5, "identities", identity2)
	ctx.AdminSession.validateAssociations(policy6, "identities", identity1, identity2, identity3)

	ctx.AdminSession.validateAssociations(service1, "service-policies", policy1, policy2, policy4, policy6)
	ctx.AdminSession.validateAssociations(service2, "service-policies", policy1, policy2, policy3, policy4, policy5, policy6)
	ctx.AdminSession.validateAssociations(service3, "service-policies", policy2, policy3, policy4)

	ctx.AdminSession.validateAssociations(identity1, "service-policies", policy1, policy2, policy4, policy6)
	ctx.AdminSession.validateAssociations(identity2, "service-policies", policy1, policy2, policy3, policy4, policy5, policy6)
	ctx.AdminSession.validateAssociations(identity3, "service-policies", policy2, policy3, policy4, policy6)

	policy1.serviceRoles = append(policy1.serviceRoles, "#"+serviceRole2)
	policy1.identityRoles = s("#" + identityRole2)
	ctx.AdminSession.requireUpdateEntity(policy1)

	ctx.AdminSession.validateAssociations(policy1, "services", service2)
	ctx.AdminSession.validateAssociations(policy1, "identities", identity2, identity3)

	ctx.AdminSession.validateAssociations(service1, "service-policies", policy2, policy4, policy6)
	ctx.AdminSession.validateAssociations(service2, "service-policies", policy1, policy2, policy3, policy4, policy5, policy6)
	ctx.AdminSession.validateAssociations(service3, "service-policies", policy2, policy3, policy4)

	ctx.AdminSession.validateAssociations(identity1, "service-policies", policy2, policy4, policy6)
	ctx.AdminSession.validateAssociations(identity2, "service-policies", policy1, policy2, policy3, policy4, policy5, policy6)
	ctx.AdminSession.validateAssociations(identity3, "service-policies", policy1, policy2, policy3, policy4, policy6)

	ctx.AdminSession.requireDeleteEntity(policy2)

	ctx.AdminSession.validateAssociations(service1, "service-policies", policy4, policy6)
	ctx.AdminSession.validateAssociations(service2, "service-policies", policy1, policy3, policy4, policy5, policy6)
	ctx.AdminSession.validateAssociations(service3, "service-policies", policy3, policy4)

	ctx.AdminSession.validateAssociations(identity1, "service-policies", policy4, policy6)
	ctx.AdminSession.validateAssociations(identity2, "service-policies", policy1, policy3, policy4, policy5, policy6)
	ctx.AdminSession.validateAssociations(identity3, "service-policies", policy1, policy3, policy4, policy6)
}
