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

FROM node:20-alpine

LABEL org.opencontainers.image.title="Tickex Node Runtime base on Node.js 20 Base Image" \
      org.opencontainers.image.description="Node.js 20 LTS, a JavaScript runtime built on Chrome's V8 JavaScript engine." \
      org.opencontainers.image.version="20" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.url="https://tickex.io.vn" \
      org.opencontainers.image.source="https://github.com/tickexvn/tickex" \
      org.opencontainers.image.authors="Tickex Team"

RUN apk --no-cache add ca-certificates

WORKDIR /usr/src/app

CMD ["node", "--version"]