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
	"strconv"

	"github.com/mhelmich/scripty/parser/ast"
	"github.com/mhelmich/scripty/parser/parsergen"
	"github.com/sirupsen/logrus"
)

func newScriptyListener() *scriptyListener {
	return &scriptyListener{
		functions: make(map[string]*ast.Function),
	}
}

type scriptyListener struct {
	*parsergen.BasescriptyListener

	functions   map[string]*ast.Function
	expressions map[*parsergen.LiteralContext]ast.CodeGenerator
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
	funcAst := &ast.Function{
		Name:   funcName,
		Params: params,
	}

	sl.functions[funcAst.SignatureHash()] = funcAst
}

func (sl *scriptyListener) EnterLiteral(ctx *parsergen.LiteralContext) {
	if ctx.NUMBER() != nil {
		numAsString := ctx.NUMBER().GetText()
		f, err := strconv.ParseFloat(numAsString, 64)
		if err != nil {
			logrus.Errorf("Can't convert string to number: %s", err.Error())
		}
		nle := &ast.NumberLiteralExpression{
			Num: f,
		}
		sl.expressions[ctx] = nle
	} else if ctx.STRING() != nil {
		str := ctx.STRING().GetText()
		sle := &ast.StringLiteralExpression{
			Str: str,
		}
		sl.expressions[ctx] = sle
	}
}

func (sl *scriptyListener) GetFunctions() map[string]*ast.Function {
	return sl.functions
}
