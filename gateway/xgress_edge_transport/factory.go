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

package xgress_edge_transport

import (
	"errors"
	"github.com/openziti/fabric/router/xgress"
	"github.com/openziti/foundation/identity/identity"
)

const BindingName = "edge_transport"

type factory struct {
	id      *identity.TokenId
	ctrl    xgress.CtrlChannel
	options *xgress.Options
}

// NewFactory returns a new Transport Xgress factory
func NewFactory(id *identity.TokenId, ctrl xgress.CtrlChannel) xgress.Factory {
	return &factory{id: id, ctrl: ctrl}
}

func (factory *factory) CreateListener(optionsData xgress.OptionsData) (xgress.Listener, error) {
	return nil, errors.New("listening not supported")
}

func (factory *factory) CreateDialer(optionsData xgress.OptionsData) (xgress.Dialer, error) {
	options := xgress.LoadOptions(optionsData)
	return newDialer(factory.id, factory.ctrl, options)
}
