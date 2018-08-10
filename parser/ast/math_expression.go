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

type MathExpression struct {
	LeftHandSide  CodeGenerator
	RightHandSide CodeGenerator
	Operand       string
}

func (me *MathExpression) GenCode(sc *ScopeContext, builder llvm.Builder) llvm.Value {
	lv := me.LeftHandSide.GenCode(sc, builder)
	rv := me.RightHandSide.GenCode(sc, builder)

	switch me.Operand {
	case "+":
		return builder.CreateFAdd(lv, rv, "add_tmp")
	case "-":
		return builder.CreateFSub(lv, rv, "sub_tmp")
	case "*":
		return builder.CreateFMul(lv, rv, "mul_tmp")
	case "/":
		return builder.CreateFDiv(lv, rv, "div_tmp")
	default:
		logrus.Errorf("Don't know operand: %s", me.Operand)
	}

	logrus.Panic("unreachable")
	return llvm.Value{}
}
