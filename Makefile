# Copyright 2025 The Celestinal Authors.
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
	build.greeter

default.print:
	@echo "[BUILD] CELESTINAL: build celestinal and services"

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

build.edge: CELESTINAL_OUT ?= edge
build.edge:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/edge
	@echo "[DONE]  CELESTINAL: edge ... ok"


build.greeter: CELESTINAL_OUT ?= greeter
build.greeter:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/greeter/v1
	@echo "[DONE]  CELESTINAL: greeter.v1 ... ok"

#####################################################################

run.edge: CELESTINAL_OUT ?= edge
run.edge:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/edge && \
 	./bin/$(CELESTINAL_OUT)

run.greeter: CELESTINAL_OUT ?= greeter
run.greeter:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/greeter/v1 && \
	./bin/$(CELESTINAL_OUT)

#####################################################################
# Docker build commands
#####################################################################
mesh:
	@docker compose -f ./deploy/docker/mesh/docker-compose.yaml \
    -f ./deploy/docker/resource/docker-compose.resources.yaml up -d

build.image.edge: TAG ?= celestinals/celestinal
build.image.edge:
	docker buildx build -f ./cmd/edge/Dockerfile -t $(TAG):latest .

build.image.x.greeter: TAG ?= celestinals/celestinal.greeter
build.image.x.greeter:
	docker buildx build -f ./x/greeter/v1/Dockerfile -t $(TAG):latest .