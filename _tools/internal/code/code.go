/*
 * Copyright 2025 The Tickex Authors.
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

package code

import (
	"fmt"
	"strings"

	"github.com/tickexvn/tickex/api/gen/go/stdx/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// GenerateFile generates a _tickex.pb.go file containing gRPC service definitions.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_tickex.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-tickex. DO NOT EDIT.")
	g.P("// Copyright 2025 Duc-Hung Ho")
	g.P()

	g.P("package ", file.GoPackageName)
	g.P("\n")
	g.P("import (")
	g.P("\t\"fmt\"\n")
	g.P()
	g.P("\t\"github.com/tickexvn/tickex/api/gen/go/stdx/v1\"\n")
	g.P(")")
	g.P("\n")
	g.P("var (")
	g.P("\t_ stdx.Empty")
	g.P(")")
	g.P("\n")

	ascii(g, file)

	for _, service := range file.Services {
		for _, method := range service.Methods {
			tickexMethodOpt(g, service, method)
		}
	}

	return g
}

func ascii(g *protogen.GeneratedFile, file *protogen.File) {
	const asciiArt = `
 _______     __          
/_  __(_)___/ /_______ __	%s
 / / / / __/  '_/ -_) \ /	--------------
/_/ /_/\__/_/\_\\__/_\_\	%s
`
	n := "TICKEX // " + strings.ToUpper(string(file.GoPackageName))
	s := fmt.Sprintf(asciiArt, n, *file.Proto.Package)
	g.P("const ascii = `", s, "\n`")
	g.P("\n")

	g.P("// PrintASCII the ASCII art to the console.")
	g.P("func PrintASCII() {")
	g.P("\tfmt.Print(ascii)")
	g.P("}")
	g.P("\n")
}

func tickexMethodOpt(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	optVal := proto.GetExtension(method.Desc.Options(), stdx.E_Options)
	if optVal == nil {
		return
	}

	methodOpt, ok := optVal.(*stdx.TickexMethodOptions)
	if !ok || methodOpt.GetIgnore() {
		return
	}

	g.P("// HasRoleAt", service.GoName, "_", method.GoName, " checks if the role has access to the method")
	g.P("func HasRoleAt", service.GoName, "_", method.GoName, "(role stdx.Role) bool {")

	if len(methodOpt.GetRequire()) == 0 {
		g.P("\treturn true")
		g.P("}")
		g.P("\n")
		return
	}

	g.P("\troleMap := make(map[stdx.Role]bool, ", len(methodOpt.GetRequire()), ")")
	requires := methodOpt.GetRequire()
	for _, require := range requires {
		g.P("\troleMap[stdx.Role_", require.GetRole().String(), "] = true")
	}

	g.P()
	g.P("\thasRole, ok := roleMap[role]")
	g.P("\tif !ok {")
	g.P("\t\treturn false")
	g.P("\t}")
	g.P()
	g.P("\treturn ", "hasRole")
	g.P("}")
	g.P("\n")

	g.P("// HasPermissionAt", service.GoName, "_", method.GoName, " checks if the permission has access to the method")
	g.P("func HasPermissionAt", service.GoName, "_", method.GoName, "(permission stdx.Permission) bool {")

	if len(methodOpt.GetRequire()) == 0 {
		g.P("\treturn true")
		g.P("}")
		g.P("\n")
		return
	}

	g.P("\tpermissionMap := make(map[stdx.Permission]bool, ", len(methodOpt.GetRequire()), ")")
	for _, require := range methodOpt.GetRequire() {
		g.P("\tpermissionMap[stdx.Permission_", require.GetPermission().String(), "] = true")
	}

	g.P()
	g.P("\thasPermission, ok := permissionMap[permission]")
	g.P("\tif !ok {")
	g.P("\t\treturn false")
	g.P("\t}")
	g.P()
	g.P("\treturn ", "hasPermission")
	g.P("}")
	g.P("\n")
}
