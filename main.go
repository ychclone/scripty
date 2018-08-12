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

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mhelmich/scripty/parser"
	"github.com/sirupsen/logrus"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

var shouldRun bool = true

func main222() {
	testFactorial()
}

func main() {
	logrus.Info("Scripty interactive shell!!")
	logrus.Infof("llvm version: %s", llvm.Version)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	interpreterLoop()
}

func cleanup() {
	shouldRun = false
	parser.PrintIR()
}

func interpreterLoop() {
	for shouldRun {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		parser.Parse(text)
	}
}

func testFactorial() {
	llvm.LinkInMCJIT()
	llvm.InitializeNativeTarget()
	llvm.InitializeNativeAsmPrinter()

	mod := llvm.NewModule("fac_module")

	fac_args := []llvm.Type{llvm.Int32Type()}
	fac_type := llvm.FunctionType(llvm.Int32Type(), fac_args, false)
	fac := llvm.AddFunction(mod, "fac", fac_type)
	fac.SetFunctionCallConv(llvm.CCallConv)
	n := fac.Param(0)

	entry := llvm.AddBasicBlock(fac, "entry")
	iftrue := llvm.AddBasicBlock(fac, "iftrue")
	iffalse := llvm.AddBasicBlock(fac, "iffalse")
	end := llvm.AddBasicBlock(fac, "end")

	builder := llvm.NewBuilder()
	defer builder.Dispose()

	builder.SetInsertPointAtEnd(entry)
	If := builder.CreateICmp(llvm.IntEQ, n, llvm.ConstInt(llvm.Int32Type(), 0, false), "cmptmp")
	builder.CreateCondBr(If, iftrue, iffalse)

	builder.SetInsertPointAtEnd(iftrue)
	res_iftrue := llvm.ConstInt(llvm.Int32Type(), 1, false)
	builder.CreateBr(end)

	builder.SetInsertPointAtEnd(iffalse)
	n_minus := builder.CreateSub(n, llvm.ConstInt(llvm.Int32Type(), 1, false), "subtmp")
	call_fac_args := []llvm.Value{n_minus}
	call_fac := builder.CreateCall(fac, call_fac_args, "calltmp")
	res_iffalse := builder.CreateMul(n, call_fac, "multmp")
	builder.CreateBr(end)

	builder.SetInsertPointAtEnd(end)
	res := builder.CreatePHI(llvm.Int32Type(), "result")
	phi_vals := []llvm.Value{res_iftrue, res_iffalse}
	phi_blocks := []llvm.BasicBlock{iftrue, iffalse}
	res.AddIncoming(phi_vals, phi_blocks)
	builder.CreateRet(res)

	err := llvm.VerifyModule(mod, llvm.ReturnStatusAction)
	if err != nil {
		logrus.Errorf("Error verifying module: %s", err)
		return
	}

	options := llvm.NewMCJITCompilerOptions()
	options.SetMCJITOptimizationLevel(2)
	options.SetMCJITEnableFastISel(true)
	options.SetMCJITNoFramePointerElim(true)
	options.SetMCJITCodeModel(llvm.CodeModelJITDefault)
	engine, err := llvm.NewMCJITCompiler(mod, options)
	if err != nil {
		logrus.Errorf("Error creating JIT: %s", err)
		return
	}
	defer engine.Dispose()

	pass := llvm.NewPassManager()
	defer pass.Dispose()

	pass.AddConstantPropagationPass()
	pass.AddInstructionCombiningPass()
	pass.AddPromoteMemoryToRegisterPass()
	pass.AddGVNPass()
	pass.AddCFGSimplificationPass()
	pass.Run(mod)

	exec_args := []llvm.GenericValue{llvm.NewGenericValueFromInt(llvm.Int32Type(), 10, false)}
	exec_res := engine.RunFunction(fac, exec_args)
	logrus.Infof("computed %d", exec_res.Int(false))
	var fac10 uint64 = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	if exec_res.Int(false) != fac10 {
		logrus.Errorf("Expected %d, got %d", fac10, exec_res.Int(false))
	}
}
