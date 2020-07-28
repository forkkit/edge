package persistence

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/storage/boltz"
)

func (m *Migrations) denormalizePolicies(step *boltz.MigrationStep) {
	log := pfxlog.Logger()

	err := m.stores.EdgeRouterPolicy.CheckIntegrity(step.Ctx.Tx(), true, func(err error) {
		log.WithError(err).Debug("updating edge router policies")
	})

	if step.SetError(err) {
		return
	}

	err = m.stores.ServiceEdgeRouterPolicy.CheckIntegrity(step.Ctx.Tx(), true, func(err error) {
		log.WithError(err).Debug("updating service edge router policies")
	})

	if step.SetError(err) {
		return
	}

	err = m.stores.ServicePolicy.CheckIntegrity(step.Ctx.Tx(), true, func(err error) {
		log.WithError(err).Debug("updating service policies")
	})

	if step.SetError(err) {
		return
	}
}