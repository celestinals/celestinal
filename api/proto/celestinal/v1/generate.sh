#!/bin/bash

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

CURRENT_DIR=$(pwd)
GOPATH_DIR=${GOPATH}

if [ -z "$GOPATH_DIR" ]; then
  echo "GOPATH is not set. Please set it before running this script."
  exit 1
fi

if [[ "$CURRENT_DIR" != *"$GOPATH_DIR"* ]]; then
  echo "Current directory ($CURRENT_DIR) does not contain GOPATH ($GOPATH_DIR)."
  exit 1
fi

CELESTINAL_PATH=$GOPATH/src/github.com/celestinals/celestinal
CELESTINAL_GEN_OUT=$GOPATH/src

protoc \
  -I"$CELESTINAL_PATH"/api/proto \
  -I"$CELESTINAL_PATH"/_submodules/googleapis \
  -I"$CELESTINAL_PATH"/_submodules/grpc-gateway \
  -I"$CELESTINAL_PATH"/_submodules/protovalidate/proto/protovalidate \
  --go_out="$CELESTINAL_GEN_OUT" \
  --go-grpc_out="$CELESTINAL_GEN_OUT" \
  --validate_out="lang=go,paths=:$CELESTINAL_GEN_OUT" \
  "$(pwd)"/*.proto
