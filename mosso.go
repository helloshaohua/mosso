// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package mosso

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

const (
	DefaultWriteFileName = "mosso.txt"
)

// DebugShowContentWithJSON Output debugging data information,
// which can be customized by option type method.
func DebugShowContentWithJSON(data interface{}, ops ...Option) {
	fmt.Print(DebugContentWithJSON(data, ops...))
}

// DebugContentWithJSON Get debugging data information,
// which can be customized by option type method.
func DebugContentWithJSON(data interface{}, ops ...Option) string {
	options := options{
		format:          true,
		visual:          true,
		typ:             true,
		fileInfo:        true,
		writeFile:       false,
		bottomSpaceLine: 1,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	var show, storage = bytes.Buffer{},bytes.Buffer{}
	defer show.Reset()
	defer storage.Reset()

	if options.fileInfo {
		_, file, line, _ := runtime.Caller(3)
		show.WriteString(fmt.Sprintf("%s:%d\n", file, line))
	}

	// Main content.
	content := DebugConvertInterfaceToJSONString(data, options.format)

	if options.visual {
		if options.typ {
			show.WriteString(fmt.Sprintf("--- CurrentType(%T) | Start ------------------\n", data))
			show.WriteString(fmt.Sprintf("%s\n", content))
			show.WriteString(fmt.Sprintf("--- CurrentType(%T) | Finish ------------------\n", data))

		} else {
			show.WriteString(content)
		}
	} else {
		show.WriteString(content)
	}

	// Space line.
	show.WriteString(fmt.Sprintf("%s", getBottomLinesSpaceStr(options.bottomSpaceLine)))

	// Write content to file.
	writerFileContent(options, storage, content)

	return show.String()
}

func writerFileContent(options options, storage bytes.Buffer, content string) {
	if options.writeFile {
		storage.WriteString(fmt.Sprintf("%s\n", content))
		ioutil.WriteFile(DefaultWriteFileName, storage.Bytes(), os.ModePerm)
	}
}
