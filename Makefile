# Copyright 2024 The Tickex Authors.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

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


# Docker build commands
build.tickex:
	docker buildx build -f ./cmd/tickex/Dockerfile -t tx.tickex:latest .

build.x.greeter:
	docker buildx build -f ./x/greeter/v1/Dockerfile -t tx.x.greeter.v1:latest .