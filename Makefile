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

.PHONY: default

default: default.print \
	build.edge \
	build.x.greeter

default.print:
	@echo "[BUILD] TICKEX: build tickex and x services"

test.cover:
	@go test ./... -cover

#####################################################################
# Go linting tool                                              
#####################################################################
lint: lint.go lint.proto

lint.go:
	@golangci-lint run

lint.proto:
	@cd ./api && buf lint

#####################################################################
# lint external modules
#####################################################################
lint.x.greeter.v1: 
	cd ./x/greeter/v1 && golangci-lint run

updaterule:
	@echo "Updating the core ruleset"


#####################################################################
#####################################################################

build.edge: TKX_OUT ?= tickex-edge
build.edge:
	@go build -ldflags="-s -w" -o ./bin/$(TKX_OUT) ./cmd/edge
	@echo "[DONE]  TICKEX: edge ... ok"


build.x.greeter: TKX_OUT ?= tickex-greeter
build.x.greeter:
	@cd ./x/greeter/v1 && \
		go build -ldflags="-s -w" -o ../../../bin/$(TKX_OUT) ./cmd
	@echo "[DONE]  TICKEX: greeter.v1 ... ok"

#####################################################################

run.edge: TKX_OUT ?= tickex-edge
run.edge:
	@go build -ldflags="-s -w" -o ./bin/$(TKX_OUT) ./cmd/edge && \
 	./bin/$(TKX_OUT)

run.x.greeter: TKX_OUT ?= tickex-greeter
run.x.greeter:
	@cd ./x/greeter/v1 && \
		go build -ldflags="-s -w" -o ../../../bin/$(TKX_OUT) ./cmd
	@./bin/$(TKX_OUT)

#####################################################################
# Docker build commands
#####################################################################
mesh:
	@docker compose -f ./deploy/docker/mesh/docker-compose.yaml \
    -f ./deploy/docker/resource/docker-compose.resources.yaml up -d

build.image.edge: TAG ?= tickexvn/edge
build.image.edge:
	docker buildx build -f ./cmd/edge/Dockerfile -t $(TAG):latest .

build.image.x.greeter: TAG ?= tickexvn/tickex.x.greeter
build.image.x.greeter:
	docker buildx build -f ./x/greeter/v1/Dockerfile -t $(TAG):latest .