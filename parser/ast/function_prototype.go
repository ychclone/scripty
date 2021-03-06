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

package ast

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

type FunctionPrototype struct {
	Name   string
	Params []antlr.TerminalNode
}

func (fp *FunctionPrototype) SignatureHash() string {
	return fp.Name + "_" + strconv.Itoa(len(fp.Params))
}

func (fp *FunctionPrototype) GenCode(sc *ScopeContext) llvm.Value {
	params := make([]llvm.Type, len(fp.Params))
	for idx := range params {
		params[idx] = sc.LlvmCtx().DoubleType()
	}

	f := llvm.AddFunction(
		sc.Module(),
		fp.Name,
		llvm.FunctionType(
			sc.LlvmCtx().DoubleType(),
			params,
			false,
		),
	)

	f.SetFunctionCallConv(llvm.CCallConv)

	for idx, param := range f.Params() {
		paramName := fp.Params[idx].GetText()
		param.SetName(paramName)
		sc.Set(paramName, param)
	}

	sc.Set(fp.Name, f)
	return f
}
