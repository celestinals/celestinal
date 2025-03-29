/*
 * Copyright 2025 The Tickex Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package watch is watching service registry when service info was changed
package watch

import (
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"github.com/tickexvn/tickex/internal/utils/eventq"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/txlog"
)

const timeout = time.Second * 2
const consulNamespace = "consul"

// Serve is watching function consul when service info was changed
func Serve(_ core.HTTPServer, config *tickex.Config) {
	client, err := newConsulClient(config)
	if err != nil {
		txlog.Errorf("[watch] failed to create consul client: %v", err)
		return
	}

	go service(client)
}

func service(client *api.Client) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	var lastIndex uint64

	for {
		services, meta, err := client.Catalog().Services(&api.QueryOptions{})
		if err != nil {
			txlog.Errorf("[watch] service event: %v", err)
			<-ticker.C

			continue
		}

		if meta.LastIndex != lastIndex {
			lastIndex = meta.LastIndex
			for serviceName := range services {
				// pass consul service
				if serviceName == consulNamespace {
					continue
				}

				entries, _, err := client.Health().Service(serviceName, "", true, nil)
				if err != nil {
					txlog.Debug("[watch] healthcheck: ", err)
					continue
				}

				if len(entries) != 0 {
					txlog.Infof("[watch] found: name=%s index=%d", serviceName, lastIndex)
					eventq.Publish(serviceName, fmt.Sprintf(
						"%s:%d", entries[0].Service.Address, entries[0].Service.Port))
				}
			}
		}

		<-ticker.C
	}
}

func newConsulClient(config *tickex.Config) (*api.Client, error) {
	conf := api.DefaultConfig()
	conf.Address = config.GetServiceRegistryAddr()
	return api.NewClient(conf)
}
