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
	"github.com/mhelmich/scripty/parser/ast"
	"github.com/mhelmich/scripty/parser/parsergen"
	"github.com/sirupsen/logrus"
)

func NewScriptyListener() *scriptyListener {
	return &scriptyListener{
		functions: make(map[string]*ast.FunctionAst),
	}
}

type scriptyListener struct {
	*parsergen.BasescriptyListener

	functions map[string]*ast.FunctionAst
}

func (sl *scriptyListener) EnterProgram(ctx *parsergen.ProgramContext) {}

func (sl *scriptyListener) ExitProgram(ctx *parsergen.ProgramContext) {}

func (sl *scriptyListener) EnterExpression(ctx *parsergen.ExpressionContext) {
	logrus.Info("EnterExpression: " + ctx.ToStringTree(nil, nil))
}

func (sl *scriptyListener) EnterFunction_def(ctx *parsergen.Function_defContext) {
	fnCtx := ctx.Function_name().(*parsergen.Function_nameContext)
	funcName := fnCtx.NAME().GetText()
	paramsCtx := ctx.Parameter_defs().(*parsergen.Parameter_defsContext)
	params := paramsCtx.AllNAME()
	funcAst := &ast.FunctionAst{
		Name:   funcName,
		Params: params,
	}

	sl.functions[funcAst.SignatureHash()] = funcAst
}

func (sl *scriptyListener) GetFunctions() map[string]*ast.FunctionAst {
	return sl.functions
}
