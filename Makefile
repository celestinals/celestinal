# Copyright 2025 The Tickex Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

test:
	@go test ./... -cover

# Go linting tool
lint: lint.go lint.proto
lint.go:
	@golangci-lint run
lint.proto:
	@cd ./api && buf lint

# lint external modules
lint.x.greeter.v1: 
	cd ./x/greeter/v1 && golangci-lint run

updaterule:
	@echo "Updating the core ruleset"

build.tickex: TKX_OUT ?= tickex-edge
build.tickex:
	@go build -ldflags="-s -w" -o ./_build/$(TKX_OUT) ./cmd/tickex

run.tickex: TKX_OUT ?= tickex-edge
run.tickex:
	@go build -ldflags="-s -w" -o ./_build/$(TKX_OUT) ./cmd/tickex && \
 	./_build/$(TKX_OUT)

run.x.greeter: TKX_OUT ?= tickex-x-greeter-v1
run.x.greeter:
	@cd ./x/greeter/v1 && go build -ldflags="-s -w" -o ../../../_build/$(TKX_OUT) ./cmd
	@./_build/$(TKX_OUT)


mesh:
	@docker compose -f ./deploy/docker/mesh/docker-compose.yaml \
    -f ./deploy/docker/resource/docker-compose.resources.yaml up -d

# Docker build commands
build.d.tickex: TAG ?= tickexvn/tickex
build.d.tickex:
	docker buildx build -f ./cmd/tickex/Dockerfile -t $(TAG):latest .

build.d.x.greeter: TAG ?= tickexvn/tickex.x.greeter
build.d.x.greeter:
	docker buildx build -f ./x/greeter/v1/Dockerfile -t $(TAG):latest .