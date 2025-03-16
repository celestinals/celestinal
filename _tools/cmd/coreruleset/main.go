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

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var rulepath string
var gengopath string

func scanfromconfig() *strings.Builder {
	file, err := os.Open(rulepath)
	if err != nil {
		fmt.Println("open file error:", err)
		return nil
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.ReplaceAll(line, "`", "` + \"`\" + `") // Escape backtick (`)
		line = strings.TrimSpace(line)

		sb.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("read file error:", err)
		return nil
	}

	return &sb
}

func writetogo(sb *strings.Builder) {
	if err := os.MkdirAll("_tools/cmd/coreruleset/gen", os.ModePerm); err != nil {
		fmt.Println("create folder error:", err)
		return
	}

	output, err := os.Create(gengopath)
	if err != nil {
		fmt.Println("create output file :", err)
		return
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	_, _ = fmt.Fprintln(writer, "package main")
	_, _ = fmt.Fprintln(writer, "var rules = `")
	_, _ = writer.WriteString(sb.String())
	_, _ = fmt.Fprintln(writer, "`")
	_ = writer.Flush()

	fmt.Println("create coreruleset .go successfully")
}

func main() {
	flag.StringVar(&rulepath, "rule",
		"_submodules/coreruleset/rules/REQUEST-932-APPLICATION-ATTACK-RCE.conf",
		"path to a rule file")

	flag.StringVar(&gengopath, "out",
		"_tools/cmd/coreruleset/gen/REQUEST-932-APPLICATION-ATTACK-RCE.go",
		"path to a generate go directory")

	flag.Parse()

	writetogo(scanfromconfig())
}
