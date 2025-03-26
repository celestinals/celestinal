#!/bin/bash

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

TKXPATH=$GOPATH/src/github.com/tickexvn/tickex
TKXOUT=$GOPATH/src

protoc \
  -I"$TKXPATH"/api/proto \
  -I"$TKXPATH"/_submodules/googleapis \
  -I"$TKXPATH"/_submodules/grpc-gateway \
  -I"$TKXPATH"/_submodules/protovalidate/proto/protovalidate \
  --openapiv2_out="$TKXPATH"/public/swagger/api/v1 \
  --grpc-gateway_out="$TKXOUT" \
  --go_out="$TKXOUT" \
  --go-grpc_out="$TKXOUT" \
  --validate_out="lang=go,paths=:$TKXOUT" \
  "$(pwd)"/*.proto || exit 1

OLDPWD=$(pwd)

cd "$TKXPATH/public/swagger/api/v1" || exit 1
find . -mindepth 2 -type f -name "*.json" -exec mv {} ./ \; || exit 1
rm -rf ./tickex || exit 1
cd "$OLDPWD" || exit 1