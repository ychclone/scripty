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

	theProgram  *ast.Program
	functions   map[string]*ast.Function
	expressions map[antlr.ParserRuleContext]ast.CodeGenerator
}

////////////////////////////////////////////////////////
/////////////////////////////////////////
//              ENTER PASS

func (sl *scriptyListener) EnterFunction_def(ctx *parsergen.Function_defContext) {
	fnCtx := ctx.Function_name().(*parsergen.Function_nameContext)
	funcName := fnCtx.NAME().GetText()
	paramsCtx := ctx.Parameter_defs().(*parsergen.Parameter_defsContext)
	params := paramsCtx.AllNAME()

	funcAst := &ast.Function{
		Proto: &ast.FunctionPrototype{
			Name:   funcName,
			Params: params,
		},
	}

	funcSig := funcAst.SignatureHash()
	_, ok := sl.functions[funcSig]
	if ok {
		logrus.Errorf("Function with signature '%s' exists already", funcSig)
		return
	}

	sl.functions[funcSig] = funcAst
}

func (sl *scriptyListener) EnterVariable(ctx *parsergen.VariableContext) {
	n := ctx.NAME().GetText()
	sl.expressions[ctx] = &ast.Variable{
		Name: n,
	}
}

func (sl *scriptyListener) EnterLiteral(ctx *parsergen.LiteralContext) {
	if ctx.NUMBER() != nil {
		numAsString := ctx.NUMBER().GetText()
		f, err := strconv.ParseFloat(numAsString, 64)
		if err != nil {
			logrus.Errorf("Can't convert string to number: %s", err.Error())
		}
		sl.expressions[ctx] = &ast.NumberLiteralExpression{
			Num: f,
		}
	} else if ctx.STRING() != nil {
		str := ctx.STRING().GetText()
		sl.expressions[ctx] = &ast.StringLiteralExpression{
			Str: str,
		}
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
		logrus.Errorf("Can't find variable or literal: %s", key.GetText())
		return
	}
	sl.expressions[ctx] = &ast.VarOrLiteral{
		Child: gen,
	}
}

func (sl *scriptyListener) ExitMath_expression(ctx *parsergen.Math_expressionContext) {
	lhs, okl := sl.expressions[ctx.Var_or_literal(0)]
	if !okl {
		logrus.Errorf("Can't find var or literal: %s", ctx.Var_or_literal(0).GetText())
		return
	}

	rhs, okr := sl.expressions[ctx.Var_or_literal(1)]
	if !okr {
		logrus.Errorf("Can't find var or literal: %s", ctx.Var_or_literal(1).GetText())
		return
	}

	operand := ctx.ARITHMETIC_OP(0).GetText()
	sl.expressions[ctx] = &ast.MathExpression{
		LeftHandSide:  lhs,
		RightHandSide: rhs,
		Operand:       operand,
	}
}

func (sl *scriptyListener) ExitExpr(ctx *parsergen.ExprContext) {
	var key antlr.ParserRuleContext
	if ctx.Math_expression() != nil {
		key = ctx.Math_expression()
	} // TODO -- add other possibilities here

	if key == nil {
		logrus.Errorf("Can't find a good key in expr")
		return
	}

	gen, ok := sl.expressions[key]
	if !ok {
		logrus.Errorf("Can't find expr: %s", key.String(nil, nil))
		return
	}
	sl.expressions[ctx] = &ast.Expr{
		Child: gen,
	}
}

func (sl *scriptyListener) ExitExpression(ctx *parsergen.ExpressionContext) {
	var key antlr.ParserRuleContext
	if ctx.Expr() != nil {
		key = ctx.Expr()
	} // TODO -- add other possibilities here

	if key == nil {
		logrus.Errorf("Can't find a good key")
		return
	}

	gen, ok := sl.expressions[key]
	if !ok {
		logrus.Errorf("Can't find expr: %s", key.GetText())
		return
	}
	sl.expressions[ctx] = &ast.Expression{
		Child: gen,
	}
}

func (sl *scriptyListener) ExitFunction_def(ctx *parsergen.Function_defContext) {
	fnCtx := ctx.Function_name().(*parsergen.Function_nameContext)
	funcName := fnCtx.NAME().GetText()
	paramsCtx := ctx.Parameter_defs().(*parsergen.Parameter_defsContext)
	params := paramsCtx.AllNAME()

	proto := &ast.FunctionPrototype{
		Name:   funcName,
		Params: params,
	}

	f, ok := sl.functions[proto.SignatureHash()]
	if !ok {
		logrus.Errorf("Can't find function with signature '%s'", proto.SignatureHash())
	}

	funcExprCtx := ctx.Function_expressions().(*parsergen.Function_expressionsContext)
	logrus.Infof("funcExprCtx: %s", funcExprCtx.GetText())
	codegen, ok := sl.expressions[funcExprCtx.Expression(0)]
	if !ok {
		logrus.Errorf("Can't find expression for '%s'", funcExprCtx.Expression(0).GetText())
	}

	expr := codegen.(*ast.Expression)
	f.Body = &ast.FunctionBody{
		Expression: expr,
	}
}

func (sl *scriptyListener) ExitProgram(ctx *parsergen.ProgramContext) {
	numExpressions := len(ctx.AllExpression())
	logrus.Infof("found %d top-level expressions", numExpressions)
	tle := make([]*ast.Expression, numExpressions)
	for i := 0; i < numExpressions; i++ {
		e, ok := sl.expressions[ctx.Expression(i)]
		if !ok {
			logrus.Errorf("Can't find expression: %s", ctx.Expression(i).GetText())
			return
		}
		tle[i] = e.(*ast.Expression)
	}

	logrus.Infof("found %d functions", len(ctx.AllFunction_def()))
	fs := make(map[string]*ast.Function)
	for _, feTmp := range ctx.AllFunction_def() {
		fe := feTmp.(*parsergen.Function_defContext)
		fnCtx := fe.Function_name().(*parsergen.Function_nameContext)
		funcName := fnCtx.NAME().GetText()
		paramsCtx := fe.Parameter_defs().(*parsergen.Parameter_defsContext)
		params := paramsCtx.AllNAME()

		proto := &ast.FunctionPrototype{
			Name:   funcName,
			Params: params,
		}
		fn := sl.functions[proto.SignatureHash()]

		fs[funcName] = fn
	}

	sl.theProgram = &ast.Program{
		Functions:           fs,
		TopLevelExpressions: tle,
	}
}
