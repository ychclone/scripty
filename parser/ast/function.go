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
	"github.com/sirupsen/logrus"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

type Function struct {
	Proto *FunctionPrototype
	Body  *FunctionBody
}

func (f *Function) SignatureHash() string {
	return f.Proto.SignatureHash()
}

func (f *Function) GenCode(sc *ScopeContext) llvm.Value {
	fsc := sc.CreateChildScope()
	fction := f.Proto.GenCode(fsc)
	bb := fsc.LlvmCtx().AddBasicBlock(fction, "entry-point-"+f.Proto.Name)
	builder := fsc.LlvmCtx().NewBuilder()
	defer builder.Dispose()
	builder.SetInsertPointAtEnd(bb)
	builder.CreateRet(f.Body.GenCode(fsc, builder))

	err := llvm.VerifyFunction(fction, llvm.PrintMessageAction)
	if err != nil {
		logrus.Errorf("Verification failed: %s", err.Error())
	}

	return fction
}
