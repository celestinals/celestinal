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

// Package namespace provide name of all service at tickex
package namespace

// Namespace is a type that represents the namespace of a service.
type Namespace string

func (ns Namespace) String() string {
	return string(ns)
}

const (
	// Edge namespace info
	Edge Namespace = "tickex.edge"

	// GreeterV1 namespace info
	GreeterV1 Namespace = "tickex.x.greeter.v1"

	// TicketV1 namespace info
	TicketV1 Namespace = "tickex.x.ticket.v1"
)
