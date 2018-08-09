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

type Program struct {
	Functions           map[string]*Function
	TopLevelExpressions []*Expression
}

func (p *Program) GenCode(sc *ScopeContext) llvm.Value {
	if len(p.TopLevelExpressions) > 0 {
		f := llvm.AddFunction(
			sc.Module(),
			"top-level-function",
			llvm.FunctionType(
				llvm.DoubleType(),
				[]llvm.Type{llvm.DoubleType(), llvm.DoubleType()},
				false,
			),
		)

		topLevelBB := sc.LlvmCtx().AddBasicBlock(f, "entry-point")
		builder := sc.LlvmCtx().NewBuilder()
		defer builder.Dispose()
		builder.SetInsertPointAtEnd(topLevelBB)
		for _, e := range p.TopLevelExpressions {
			builder.CreateRet(e.GenCode(sc, builder))
		}

		return f
	} else if len(p.Functions) > 0 {
		for _, fn := range p.Functions {
			logrus.Infof("GenCode for function: %s", fn.Proto.Name)
			return fn.GenCode(sc)
		}
	}

	return llvm.Value{}
}
