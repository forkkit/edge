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

package policy

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/runner"
	"time"
)

type SessionEnforcer struct {
	appEnv         *env.AppEnv
	sessionTimeout time.Duration
	*runner.BaseOperation
}

func NewSessionEnforcer(appEnv *env.AppEnv, frequency time.Duration, sessionTimeout time.Duration) *SessionEnforcer {
	if sessionTimeout < 60*time.Second {
		pfxlog.Logger().Panic("sessionTimeout can not be less than 60 seconds")
	}

	pfxlog.Logger().
		WithField("sessionTimeout", sessionTimeout.String()).
		WithField("frequency", frequency.String()).
		Info("session enforcer configured")

	return &SessionEnforcer{
		appEnv:         appEnv,
		sessionTimeout: sessionTimeout,
		BaseOperation:  runner.NewBaseOperation("SessionEnforcer", frequency),
	}
}

func (s *SessionEnforcer) Run() error {
	oldest := time.Now().Add(s.sessionTimeout * -1)
	query := fmt.Sprintf("updatedAt < datetime(%s)", oldest.UTC().Format(time.RFC3339))

	return s.appEnv.GetHandlers().ApiSession.StreamIds(query, func(id string, err error) error {
		_ = s.appEnv.GetHandlers().ApiSession.Delete(id)
		return nil
	})
}
