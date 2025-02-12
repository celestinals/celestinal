/*
 * Copyright 2024 The Tickex Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package api contains the proto files used in the Tickex ecosystem,
// including proto files used in the runtime, services, and domain layers.
// All functions are separated and run completely independently of each other.
//
// # Architecture Layer
//
// Runtime: defines common system services such as service registry, proxy, etc...
//
// Services: defines tasks by combining business logic through the domain layer
//
// Domain: defines business logic and data types
//
// #
//
// Package api chứa các tệp proto được sử dụng trong hệ sinh thái Tickex,
// bao gồm các tệp proto được sử dụng trong runtime, services và các lớp domain.
// Tất cả các chức năng đều được tách biệt và chạy hoàn toàn độc lập với nhau.
//
// # Lớp Kiến trúc
//
// Runtime: định nghĩa các dịch vụ hệ thống chung như đăng ký dịch vụ, proxy, v.v...
//
// Service: định nghĩa các nhiệm vụ bằng cách kết hợp logic kinh doanh thông qua lớp miền
//
// Domain: định nghĩa logic kinh doanh và các kiểu dữ liệu
package api
