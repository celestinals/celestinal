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

# Use Alpine edge version as the base image
FROM alpine:edge

# Add OCI metadata
LABEL org.opencontainers.image.title="Tickex Edge base on Alpine Edge base Image" \
      org.opencontainers.image.description="Edge version of Alpine Linux, lightweight and fast for containerized applications." \
      org.opencontainers.image.version="edge" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.url="https://tickex.io.vn" \
      org.opencontainers.image.source="https://github.com/tickexvn/tickex" \
      org.opencontainers.image.authors="Tickex Labs"

RUN apk --no-cache add ca-certificates

# Default command
CMD ["sh"]
