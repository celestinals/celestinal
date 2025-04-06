// Copyright 2025 The Celestinal Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package copier provides functions to copy objects.
package copier

import (
	"encoding/json"

	google "google.golang.org/protobuf/proto"

	"github.com/celestinals/celestinal/pkg/protobuf/proto"
)

// CopyProtoMessage copies the src message to the dst message.
func CopyProtoMessage(src, dst google.Message) error {
	bytes, err := proto.Marshal(src)
	if err != nil {
		return err
	}

	return proto.Unmarshal(bytes, dst)
}

// CopyJSON copies the src object to the dst object.
func CopyJSON(src, dst any) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dst)
}
