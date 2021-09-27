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

type Option interface {
	apply(*options)
}

type options struct {
	format          bool
	visual          bool
	typ             bool
	fileInfo        bool
	writeFile       bool
	bottomSpaceLine int
}

type optionsFunc func(*options)

func (o optionsFunc) apply(ops *options) {
	o(ops)
}

func WithSpecificOfFormat(format bool) Option {
	return optionsFunc(func(ops *options) {
		ops.format = format
	})
}

func WithSpecificOfVisual(visual bool) Option {
	return optionsFunc(func(ops *options) {
		ops.visual = visual
	})
}

func WithSpecificOfShowType(typ bool) Option {
	return optionsFunc(func(ops *options) {
		ops.typ = typ
	})
}

func WithSpecificOfLineNumber(lineNum bool) Option {
	return optionsFunc(func(ops *options) {
		ops.fileInfo = lineNum
	})
}

func WithSpecificOfBottomSpaceLine(number int) Option {
	return optionsFunc(func(ops *options) {
		ops.bottomSpaceLine = number
	})
}

func WithSpecificOfWriteFile(writeFile bool) Option {
	return optionsFunc(func(ops *options) {
		ops.writeFile = writeFile
	})
}
