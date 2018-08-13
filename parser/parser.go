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
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mhelmich/scripty/parser/ast"
	"github.com/mhelmich/scripty/parser/parsergen"
	"github.com/sirupsen/logrus"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

var context llvm.Context
var module llvm.Module

func init() {
	llvm.LinkInMCJIT()
	err := llvm.InitializeNativeTarget()
	if err != nil {
		logrus.Errorf("Can't init native target: %s", err.Error())
	}

	err = llvm.InitializeNativeAsmPrinter()
	if err != nil {
		logrus.Errorf("Can't init native asm printer: %s", err.Error())
	}

	context = llvm.NewContext()
	module = context.NewModule("root-module")
}

func PrintIR() {
	fmt.Printf("generated IR:\n%s", module.String())
}

func Parse(input string) {
	is := antlr.NewInputStream(input)
	lexer := parsergen.NewscriptyLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parsergen.NewscriptyParser(ts)
	rootCtx := parser.Program()

	listener := newScriptyListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, rootCtx)
	sc := ast.NewScopeContext(context, module)
	listener.theProgram.GenCode(sc)

	err := llvm.VerifyModule(module, llvm.PrintMessageAction)
	if err != nil {
		logrus.Errorf("Verification failed: %s", err.Error())
	}

	logrus.Debugf("generated IR:\n%s", module.String())
	PrintIR()

	if len(listener.theProgram.TopLevelExpressions) > 0 {
		options := llvm.NewMCJITCompilerOptions()
		options.SetMCJITOptimizationLevel(2)
		options.SetMCJITEnableFastISel(true)
		options.SetMCJITNoFramePointerElim(true)
		options.SetMCJITCodeModel(llvm.CodeModelJITDefault)
		engine, err := llvm.NewMCJITCompiler(module, options)
		if err != nil {
			logrus.Errorf("Error creating JIT: %s", err)
			return
		}
		defer engine.Dispose()

		pass := llvm.NewPassManager()
		defer pass.Dispose()
		pass.AddConstantPropagationPass()
		pass.AddInstructionCombiningPass()
		pass.AddReassociatePass()
		pass.AddPromoteMemoryToRegisterPass()
		pass.AddGVNPass()
		pass.AddCFGSimplificationPass()
		pass.Run(module)

		args := []llvm.GenericValue{}
		res := engine.RunFunction(module.NamedFunction("top-level-function"), args)
		defer res.Dispose()
		fmt.Printf("result: %f\n", res.Float(sc.LlvmCtx().DoubleType()))
	}
}
