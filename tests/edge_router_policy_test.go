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

func Test_EdgeRouterPolicy(t *testing.T) {
	ctx := NewTestContext(t)
	defer ctx.Teardown()
	ctx.StartServer()
	ctx.RequireAdminLogin()

	edgeRouterRole1 := eid.New()
	edgeRouterRole2 := eid.New()
	identityRole1 := eid.New()
	identityRole2 := eid.New()

	edgeRouter1 := ctx.AdminSession.requireNewEdgeRouter(edgeRouterRole1)
	edgeRouter2 := ctx.AdminSession.requireNewEdgeRouter(edgeRouterRole1, edgeRouterRole2)
	edgeRouter3 := ctx.AdminSession.requireNewEdgeRouter()

	identity1 := ctx.AdminSession.requireNewIdentity(false, identityRole1)
	identity2 := ctx.AdminSession.requireNewIdentity(false, identityRole1, identityRole2)
	identity3 := ctx.AdminSession.requireNewIdentity(false, identityRole2)

	policy1 := ctx.AdminSession.requireNewEdgeRouterPolicy(s("#"+edgeRouterRole1), s("#"+identityRole1))
	policy2 := ctx.AdminSession.requireNewEdgeRouterPolicy(s("#"+edgeRouterRole1, "@"+edgeRouter3.id), s("#"+identityRole1, "@"+identity3.Id))
	policy3 := ctx.AdminSession.requireNewEdgeRouterPolicy(s("@"+edgeRouter2.id, "@"+edgeRouter3.id), s("@"+identity2.Id, "@"+identity3.Id))
	policy4 := ctx.AdminSession.requireNewEdgeRouterPolicy(s("#all"), s("#all"))
	policy5 := ctx.AdminSession.requireNewEdgeRouterPolicyWithSemantic("AllOf", s("#"+edgeRouterRole1, "#"+edgeRouterRole2), s("#"+identityRole1, "#"+identityRole2))
	policy6 := ctx.AdminSession.requireNewEdgeRouterPolicyWithSemantic("AnyOf", s("#"+edgeRouterRole1, "#"+edgeRouterRole2), s("#"+identityRole1, "#"+identityRole2))

	t.Run("policy 1 created", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateEntityWithQuery(policy1)
	})

	t.Run("policy 2 created", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateEntityWithLookup(policy2)
	})

	t.Run("policy 3 created", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateEntityWithLookup(policy3)
	})

	t.Run("policy 4 created", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateEntityWithLookup(policy4)
	})

	t.Run("policy 5 created", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateEntityWithLookup(policy5)
	})
	t.Run("policy 6 created", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateEntityWithLookup(policy6)
	})

	t.Run("policy 1 has edge routers 1, 2", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy1, "edge-routers", edgeRouter1, edgeRouter2)
	})

	t.Run("policy 2 has edge routers 1, 2, 3", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy2, "edge-routers", edgeRouter1, edgeRouter2, edgeRouter3)
	})

	t.Run("policy 3 has edge routers 2, 3", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy3, "edge-routers", edgeRouter2, edgeRouter3)
	})

	t.Run("policy 4 has edge routers 1, 2, 3", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociationContains(policy4, "edge-routers", edgeRouter1, edgeRouter2, edgeRouter3)
	})

	t.Run("policy 5 has edge router 2", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy5, "edge-routers", edgeRouter2)
	})

	t.Run("policy 6 has edge routers 1, 2", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy6, "edge-routers", edgeRouter1, edgeRouter2)
	})

	t.Run("policy 1 has identities 1, 2", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy1, "identities", identity1, identity2)
	})

	t.Run("policy 2 has identities 1, 2, 3", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy2, "identities", identity1, identity2, identity3)
	})

	t.Run("policy 3 has identities 2, 3", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy3, "identities", identity2, identity3)
	})

	t.Run("policy 4 has identities 1, 2, 3", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociationContains(policy4, "identities", identity1, identity2, identity3)
	})

	t.Run("policy 5 has identity 2", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy5, "identities", identity2)
	})

	t.Run("policy 6 has identities 1, 2, 3", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(policy6, "identities", identity1, identity2, identity3)
	})

	t.Run("edge router 1 has edge router policies 1, 2, 3, 4, 6", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(edgeRouter1, "edge-router-policies", policy1, policy2, policy4, policy6)
	})

	t.Run("edge router 2 has edge router policies 1, 2, 3, 4, 5, 6", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(edgeRouter2, "edge-router-policies", policy1, policy2, policy3, policy4, policy5, policy6)
	})

	t.Run("edge router 3 has edge router policies 2, 3, 4", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(edgeRouter3, "edge-router-policies", policy2, policy3, policy4)
	})

	t.Run("identity 1 has edge router policies 1, 2, 4, 6", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(identity1, "edge-router-policies", policy1, policy2, policy4, policy6)
	})

	t.Run("identity 2 has policies 1, 2, 3, 4 ,5, 6", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(identity2, "edge-router-policies", policy1, policy2, policy3, policy4, policy5, policy6)
	})

	t.Run("identity 3 has edge router policy 2, 3, 4, 6", func(t *testing.T) {
		ctx.testContextChanged(t)
		ctx.AdminSession.validateAssociations(identity3, "edge-router-policies", policy2, policy3, policy4, policy6)
	})

	t.Run("updated policy 1 ", func(t *testing.T) {
		ctx.testContextChanged(t)
		policy1.edgeRouterRoles = append(policy1.edgeRouterRoles, "#"+edgeRouterRole2)
		policy1.identityRoles = s("#" + identityRole2)
		ctx.AdminSession.requireUpdateEntity(policy1)

		t.Run("has  edge router 2 and not 1", func(t *testing.T) {
			ctx.testContextChanged(t)

			ctx.AdminSession.validateAssociations(policy1, "edge-routers", edgeRouter2)
		})

		t.Run("policy 1 has identities 2, 3 and not 1", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(policy1, "identities", identity2, identity3)
		})

		t.Run("edge router 1 has edge router policies 2, 4, 6 and not 1", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(edgeRouter1, "edge-router-policies", policy2, policy4, policy6)
		})

		t.Run("edge router 2 retained edge router policies 1, 2 ,3 ,4 ,5, 6", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(edgeRouter2, "edge-router-policies", policy1, policy2, policy3, policy4, policy5, policy6)
		})

		t.Run("edge router 3 retained edge router policies: 2, 3, 4", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(edgeRouter3, "edge-router-policies", policy2, policy3, policy4)
		})

		t.Run("identity 1 retained edge router policies 2, 4, 6 but not 1", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(identity1, "edge-router-policies", policy2, policy4, policy6)
		})

		t.Run("identity 2 retained edge router policies 1, 2, 3, 4, 5, 6", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(identity2, "edge-router-policies", policy1, policy2, policy3, policy4, policy5, policy6)
		})

		t.Run("identity 3 retained edge router policies 2, 3, 4, 6 and gained 1", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(identity3, "edge-router-policies", policy1, policy2, policy3, policy4, policy6)
		})
	})

	t.Run("delete policy 2", func(t *testing.T) {
		ctx.AdminSession.requireDeleteEntity(policy2)

		t.Run("edge router 1 has edge router policies 4, 6 and not 2", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(edgeRouter1, "edge-router-policies", policy4, policy6)
		})

		t.Run("edge router 2 has edge router policies 1, 3, 4, 5, 6 and not 2", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(edgeRouter2, "edge-router-policies", policy1, policy3, policy4, policy5, policy6)
		})

		t.Run("edge router 3 has edge router policies 3, 4 and not 2", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(edgeRouter3, "edge-router-policies", policy3, policy4)
		})

		t.Run("identity 1 retained edge router policies 4, 6 and not 2", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(identity1, "edge-router-policies", policy4, policy6)
		})

		t.Run("identity 2 retained edge router policies 1, 3, 4, 5, 6 and not 2", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(identity2, "edge-router-policies", policy1, policy3, policy4, policy5, policy6)
		})

		t.Run("identity 3 retained edge router policies 1, 3 , 4, 6 and not 2", func(t *testing.T) {
			ctx.testContextChanged(t)
			ctx.AdminSession.validateAssociations(identity3, "edge-router-policies", policy1, policy3, policy4, policy6)
		})

	})

}
