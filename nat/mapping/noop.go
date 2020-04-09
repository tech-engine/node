/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package mapping

import (
	"errors"

	"github.com/mysteriumnetwork/node/eventbus"
	"github.com/mysteriumnetwork/node/nat/event"
	"github.com/rs/zerolog/log"
)

// NewNoopPortMapper returns noop port mapper instance.
func NewNoopPortMapper(publisher eventbus.Publisher) PortMapper {
	return &noopPortMapper{
		publisher: publisher,
	}
}

type noopPortMapper struct {
	publisher eventbus.Publisher
}

func (p *noopPortMapper) Map(protocol string, port int, name string) (release func(), ok bool) {
	p.publisher.Publish(event.AppTopicTraversal, event.BuildFailureEvent(StageName, errors.New("noop mapping")))
	log.Debug().Msgf("Noop port mapping requested: %d", port)

	return func() {
		log.Debug().Msgf("Noop port mapping released: %d", port)
	}, false
}
