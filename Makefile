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
	apiserver.build \
	greeter.build

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
#####################################################################

apiserver.build: CELESTINAL_OUT ?= apiserver
apiserver.build:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/apiserver
	@echo "[DONE]  CELESTINAL: apiserver ... ok"

apiserver.run: CELESTINAL_OUT ?= apiserver
apiserver.run:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/apiserver && \
 	./bin/$(CELESTINAL_OUT)

apiserver.build.image: TAG ?= celestinals/celestinal
apiserver.build.image:
	docker buildx build -f ./cmd/apiserver/Dockerfile -t $(TAG):latest .


greeter.build: CELESTINAL_OUT ?= greeter
greeter.build:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/greeter/v1
	@echo "[DONE]  CELESTINAL: greeter.v1 ... ok"

greeter.run: CELESTINAL_OUT ?= greeter
greeter.run:
	@go build -ldflags="-s -w" -o ./bin/$(CELESTINAL_OUT) ./cmd/greeter/v1 && \
	./bin/$(CELESTINAL_OUT)

greeter.build.image: TAG ?= celestinals/celestinal.greeter
greeter.build.image:
	docker buildx build -f ./cmd/greeter/v1/Dockerfile -t $(TAG):latest .