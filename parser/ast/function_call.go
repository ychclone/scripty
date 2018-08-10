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

	"github.com/sirupsen/logrus"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

type FunctionCall struct {
	FunctionName string
	Params       []string
}

func (fc *FunctionCall) GenCode(sc *ScopeContext, builder llvm.Builder) llvm.Value {
	args := make([]llvm.Value, len(fc.Params))
	for idx, param := range fc.Params {
		f, err := strconv.ParseFloat(param, 64)
		if err != nil {
			logrus.Errorf("Can't convert string '%s' to double", param)
		}
		args[idx] = llvm.ConstFloat(sc.LlvmCtx().DoubleType(), f)
	}

	var llvmFunc llvm.Value
	val := sc.Get(fc.FunctionName)
	if val == nil {
		llvmVal := sc.Module().NamedFunction(fc.FunctionName)

		if llvmVal.IsNil() {
			logrus.Errorf("Can't find function to call '%s'", fc.FunctionName)
			return llvm.Value{}
		}

		llvmFunc = llvmVal
	} else {
		llvmFunc = val.(llvm.Value)
	}

	return builder.CreateCall(llvmFunc, args, fc.FunctionName+strconv.Itoa(len(fc.Params)))
}
