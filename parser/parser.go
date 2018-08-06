/*
 * Copyright 2018 Marco Helmich
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

package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mhelmich/scripty/parser/parsergen"
	"github.com/sirupsen/logrus"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

var symbols map[string]llvm.Value
var llvmCtx llvm.Context
var builder llvm.Builder
var module llvm.Module

func init() {
	llvm.LinkInMCJIT()
	err := llvm.InitializeNativeTarget()
	if err != nil {
		logrus.Errorf("Can't init native target: %s", err.Error())
	}

	symbols = make(map[string]llvm.Value)
	llvmCtx = llvm.NewContext()
	builder = llvmCtx.NewBuilder()
	module = llvmCtx.NewModule("module")
}

func Parse(input string) {
	is := antlr.NewInputStream(input)
	lexer := parsergen.NewscriptyLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parsergen.NewscriptyParser(ts)
	rootCtx := parser.Program()

	listener := newScriptyListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, rootCtx)
	listener.theProgram.GenCode(llvmCtx, module)
	logrus.Infof("generated IR:\n%s", module.String())
}
