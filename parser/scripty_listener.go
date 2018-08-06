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

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mhelmich/scripty/parser/ast"
	"github.com/mhelmich/scripty/parser/parsergen"
	"github.com/sirupsen/logrus"
)

func newScriptyListener() *scriptyListener {
	return &scriptyListener{
		functions:   make(map[string]*ast.Function),
		expressions: make(map[antlr.ParserRuleContext]ast.CodeGenerator),
	}
}

type scriptyListener struct {
	*parsergen.BasescriptyListener

	functions   map[string]*ast.Function
	expressions map[antlr.ParserRuleContext]ast.CodeGenerator
}

////////////////////////////////////////////////////////
/////////////////////////////////////////
//              ENTER PASS

func (sl *scriptyListener) EnterProgram(ctx *parsergen.ProgramContext) {}

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

////////////////////////////////////////////////////////
/////////////////////////////////////////
//              EXIT PASS

func (sl *scriptyListener) ExitVar_or_literal(ctx *parsergen.Var_or_literalContext) {
	var key antlr.ParserRuleContext
	if ctx.Literal() != nil {
		key = ctx.Literal()
	} else if ctx.Variable() != nil {
		key = ctx.Variable()
	}

	gen, ok := sl.expressions[key]
	if !ok {
		logrus.Errorf("Can't find variable: %s", key.GetText())
	}
	sl.expressions[ctx] = &ast.VarOrLiteral{
		Child: gen,
	}
}
