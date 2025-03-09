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

FROM golang:1.24-alpine

LABEL org.opencontainers.image.title="Tickex Go compiler base on Go 1.24 Alpine Base Image" \
      org.opencontainers.image.description="Go 1.24 with a lightweight and fast Alpine Linux environment." \
      org.opencontainers.image.version="1.24" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.url="https://tickex.io.vn" \
      org.opencontainers.image.source="https://github.com/tickexvn/tickex" \
      org.opencontainers.image.authors="Tickex Labs"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git ca-certificates

WORKDIR /home

CMD ["go", "version"]