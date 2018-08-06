// Generated from scripty.g4 by ANTLR 4.7.

package parsergen // scripty
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 130,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 6, 2, 33, 10, 2, 13, 2,
	14, 2, 34, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 51, 10, 5, 12, 5, 14, 5, 54, 11, 5, 5, 5,
	56, 10, 5, 3, 6, 7, 6, 59, 10, 6, 12, 6, 14, 6, 62, 11, 6, 3, 7, 3, 7,
	3, 7, 5, 7, 67, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 72, 10, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 81, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 86,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 93, 10, 10, 12, 10, 14,
	10, 96, 11, 10, 5, 10, 98, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 7, 11, 107, 10, 11, 12, 11, 14, 11, 110, 11, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 117, 10, 12, 12, 12, 14, 12, 120, 11, 12, 3,
	13, 3, 13, 5, 13, 124, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 2, 2,
	16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 3, 3, 2, 14,
	15, 2, 132, 2, 32, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 45, 3, 2, 2, 2, 8,
	55, 3, 2, 2, 2, 10, 60, 3, 2, 2, 2, 12, 66, 3, 2, 2, 2, 14, 71, 3, 2, 2,
	2, 16, 85, 3, 2, 2, 2, 18, 87, 3, 2, 2, 2, 20, 101, 3, 2, 2, 2, 22, 111,
	3, 2, 2, 2, 24, 123, 3, 2, 2, 2, 26, 125, 3, 2, 2, 2, 28, 127, 3, 2, 2,
	2, 30, 33, 5, 4, 3, 2, 31, 33, 5, 12, 7, 2, 32, 30, 3, 2, 2, 2, 32, 31,
	3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2,
	35, 3, 3, 2, 2, 2, 36, 37, 7, 3, 2, 2, 37, 38, 5, 6, 4, 2, 38, 39, 7, 4,
	2, 2, 39, 40, 5, 8, 5, 2, 40, 41, 7, 5, 2, 2, 41, 42, 7, 6, 2, 2, 42, 43,
	5, 10, 6, 2, 43, 44, 7, 7, 2, 2, 44, 5, 3, 2, 2, 2, 45, 46, 7, 13, 2, 2,
	46, 7, 3, 2, 2, 2, 47, 52, 7, 13, 2, 2, 48, 49, 7, 8, 2, 2, 49, 51, 7,
	13, 2, 2, 50, 48, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52,
	53, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 47, 3, 2, 2,
	2, 55, 56, 3, 2, 2, 2, 56, 9, 3, 2, 2, 2, 57, 59, 5, 12, 7, 2, 58, 57,
	3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2,
	61, 11, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 67, 5, 16, 9, 2, 64, 67, 5,
	18, 10, 2, 65, 67, 5, 14, 8, 2, 66, 63, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	66, 65, 3, 2, 2, 2, 67, 13, 3, 2, 2, 2, 68, 72, 5, 22, 12, 2, 69, 72, 5,
	20, 11, 2, 70, 72, 5, 24, 13, 2, 71, 68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2,
	71, 70, 3, 2, 2, 2, 72, 15, 3, 2, 2, 2, 73, 74, 7, 13, 2, 2, 74, 75, 7,
	9, 2, 2, 75, 86, 7, 13, 2, 2, 76, 77, 7, 13, 2, 2, 77, 80, 7, 9, 2, 2,
	78, 81, 5, 20, 11, 2, 79, 81, 5, 22, 12, 2, 80, 78, 3, 2, 2, 2, 80, 79,
	3, 2, 2, 2, 81, 86, 3, 2, 2, 2, 82, 83, 7, 13, 2, 2, 83, 84, 7, 9, 2, 2,
	84, 86, 5, 18, 10, 2, 85, 73, 3, 2, 2, 2, 85, 76, 3, 2, 2, 2, 85, 82, 3,
	2, 2, 2, 86, 17, 3, 2, 2, 2, 87, 88, 7, 13, 2, 2, 88, 97, 7, 4, 2, 2, 89,
	94, 5, 14, 8, 2, 90, 91, 7, 8, 2, 2, 91, 93, 5, 14, 8, 2, 92, 90, 3, 2,
	2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 98,
	3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 89, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2,
	98, 99, 3, 2, 2, 2, 99, 100, 7, 5, 2, 2, 100, 19, 3, 2, 2, 2, 101, 102,
	5, 24, 13, 2, 102, 103, 7, 11, 2, 2, 103, 108, 5, 24, 13, 2, 104, 105,
	7, 11, 2, 2, 105, 107, 5, 24, 13, 2, 106, 104, 3, 2, 2, 2, 107, 110, 3,
	2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 21, 3, 2, 2,
	2, 110, 108, 3, 2, 2, 2, 111, 112, 5, 24, 13, 2, 112, 113, 7, 12, 2, 2,
	113, 118, 5, 24, 13, 2, 114, 115, 7, 12, 2, 2, 115, 117, 5, 24, 13, 2,
	116, 114, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118,
	119, 3, 2, 2, 2, 119, 23, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 124, 5,
	26, 14, 2, 122, 124, 5, 28, 15, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2,
	2, 2, 124, 25, 3, 2, 2, 2, 125, 126, 7, 13, 2, 2, 126, 27, 3, 2, 2, 2,
	127, 128, 9, 2, 2, 2, 128, 29, 3, 2, 2, 2, 16, 32, 34, 52, 55, 60, 66,
	71, 80, 85, 94, 97, 108, 118, 123,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'def'", "'('", "')'", "'{'", "'}'", "','", "'='", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "ASSIGNMENT_OP", "STMT_END", "BOOLEAN_OP",
	"ARITHMETIC_OP", "NAME", "NUMBER", "STRING", "WHITESPACE", "COMMENT", "UNKNOWN",
	"ANY", "INDENT", "DEDENT", "NEWLINE", "ENDMARKER",
}

var ruleNames = []string{
	"program", "function_def", "function_name", "parameter_defs", "function_expressions",
	"expression", "expr", "assignment", "function_call", "boolean_expression",
	"math_expression", "var_or_literal", "variable", "literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type scriptyParser struct {
	*antlr.BaseParser
}

func NewscriptyParser(input antlr.TokenStream) *scriptyParser {
	this := new(scriptyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "scripty.g4"

	return this
}

// scriptyParser tokens.
const (
	scriptyParserEOF           = antlr.TokenEOF
	scriptyParserT__0          = 1
	scriptyParserT__1          = 2
	scriptyParserT__2          = 3
	scriptyParserT__3          = 4
	scriptyParserT__4          = 5
	scriptyParserT__5          = 6
	scriptyParserASSIGNMENT_OP = 7
	scriptyParserSTMT_END      = 8
	scriptyParserBOOLEAN_OP    = 9
	scriptyParserARITHMETIC_OP = 10
	scriptyParserNAME          = 11
	scriptyParserNUMBER        = 12
	scriptyParserSTRING        = 13
	scriptyParserWHITESPACE    = 14
	scriptyParserCOMMENT       = 15
	scriptyParserUNKNOWN       = 16
	scriptyParserANY           = 17
	scriptyParserINDENT        = 18
	scriptyParserDEDENT        = 19
	scriptyParserNEWLINE       = 20
	scriptyParserENDMARKER     = 21
)

// scriptyParser rules.
const (
	scriptyParserRULE_program              = 0
	scriptyParserRULE_function_def         = 1
	scriptyParserRULE_function_name        = 2
	scriptyParserRULE_parameter_defs       = 3
	scriptyParserRULE_function_expressions = 4
	scriptyParserRULE_expression           = 5
	scriptyParserRULE_expr                 = 6
	scriptyParserRULE_assignment           = 7
	scriptyParserRULE_function_call        = 8
	scriptyParserRULE_boolean_expression   = 9
	scriptyParserRULE_math_expression      = 10
	scriptyParserRULE_var_or_literal       = 11
	scriptyParserRULE_variable             = 12
	scriptyParserRULE_literal              = 13
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllFunction_def() []IFunction_defContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunction_defContext)(nil)).Elem())
	var tst = make([]IFunction_defContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunction_defContext)
		}
	}

	return tst
}

func (s *ProgramContext) Function_def(i int) IFunction_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_defContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunction_defContext)
}

func (s *ProgramContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *scriptyParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, scriptyParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<scriptyParserT__0)|(1<<scriptyParserNAME)|(1<<scriptyParserNUMBER)|(1<<scriptyParserSTRING))) != 0) {
		p.SetState(30)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case scriptyParserT__0:
			{
				p.SetState(28)
				p.Function_def()
			}

		case scriptyParserNAME, scriptyParserNUMBER, scriptyParserSTRING:
			{
				p.SetState(29)
				p.Expression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunction_defContext is an interface to support dynamic dispatch.
type IFunction_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_defContext differentiates from other interfaces.
	IsFunction_defContext()
}

type Function_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_defContext() *Function_defContext {
	var p = new(Function_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_function_def
	return p
}

func (*Function_defContext) IsFunction_defContext() {}

func NewFunction_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_defContext {
	var p = new(Function_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_function_def

	return p
}

func (s *Function_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_defContext) Function_name() IFunction_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *Function_defContext) Parameter_defs() IParameter_defsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_defsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_defsContext)
}

func (s *Function_defContext) Function_expressions() IFunction_expressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_expressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_expressionsContext)
}

func (s *Function_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterFunction_def(s)
	}
}

func (s *Function_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitFunction_def(s)
	}
}

func (p *scriptyParser) Function_def() (localctx IFunction_defContext) {
	localctx = NewFunction_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, scriptyParserRULE_function_def)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(scriptyParserT__0)
	}
	{
		p.SetState(35)
		p.Function_name()
	}
	{
		p.SetState(36)
		p.Match(scriptyParserT__1)
	}
	{
		p.SetState(37)
		p.Parameter_defs()
	}
	{
		p.SetState(38)
		p.Match(scriptyParserT__2)
	}
	{
		p.SetState(39)
		p.Match(scriptyParserT__3)
	}
	{
		p.SetState(40)
		p.Function_expressions()
	}
	{
		p.SetState(41)
		p.Match(scriptyParserT__4)
	}

	return localctx
}

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_function_name
	return p
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(scriptyParserNAME, 0)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterFunction_name(s)
	}
}

func (s *Function_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitFunction_name(s)
	}
}

func (p *scriptyParser) Function_name() (localctx IFunction_nameContext) {
	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, scriptyParserRULE_function_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(scriptyParserNAME)
	}

	return localctx
}

// IParameter_defsContext is an interface to support dynamic dispatch.
type IParameter_defsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_defsContext differentiates from other interfaces.
	IsParameter_defsContext()
}

type Parameter_defsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_defsContext() *Parameter_defsContext {
	var p = new(Parameter_defsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_parameter_defs
	return p
}

func (*Parameter_defsContext) IsParameter_defsContext() {}

func NewParameter_defsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_defsContext {
	var p = new(Parameter_defsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_parameter_defs

	return p
}

func (s *Parameter_defsContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_defsContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(scriptyParserNAME)
}

func (s *Parameter_defsContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(scriptyParserNAME, i)
}

func (s *Parameter_defsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_defsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_defsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterParameter_defs(s)
	}
}

func (s *Parameter_defsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitParameter_defs(s)
	}
}

func (p *scriptyParser) Parameter_defs() (localctx IParameter_defsContext) {
	localctx = NewParameter_defsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, scriptyParserRULE_parameter_defs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == scriptyParserNAME {
		{
			p.SetState(45)
			p.Match(scriptyParserNAME)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == scriptyParserT__5 {
			{
				p.SetState(46)
				p.Match(scriptyParserT__5)
			}
			{
				p.SetState(47)
				p.Match(scriptyParserNAME)
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IFunction_expressionsContext is an interface to support dynamic dispatch.
type IFunction_expressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_expressionsContext differentiates from other interfaces.
	IsFunction_expressionsContext()
}

type Function_expressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_expressionsContext() *Function_expressionsContext {
	var p = new(Function_expressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_function_expressions
	return p
}

func (*Function_expressionsContext) IsFunction_expressionsContext() {}

func NewFunction_expressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_expressionsContext {
	var p = new(Function_expressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_function_expressions

	return p
}

func (s *Function_expressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_expressionsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Function_expressionsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Function_expressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_expressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_expressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterFunction_expressions(s)
	}
}

func (s *Function_expressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitFunction_expressions(s)
	}
}

func (p *scriptyParser) Function_expressions() (localctx IFunction_expressionsContext) {
	localctx = NewFunction_expressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, scriptyParserRULE_function_expressions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<scriptyParserNAME)|(1<<scriptyParserNUMBER)|(1<<scriptyParserSTRING))) != 0 {
		{
			p.SetState(55)
			p.Expression()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ExpressionContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExpressionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *scriptyParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, scriptyParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(61)
			p.Assignment()
		}

	case 2:
		{
			p.SetState(62)
			p.Function_call()
		}

	case 3:
		{
			p.SetState(63)
			p.Expr()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Math_expression() IMath_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMath_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *ExprContext) Boolean_expression() IBoolean_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *ExprContext) Var_or_literal() IVar_or_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_or_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_or_literalContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *scriptyParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, scriptyParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Math_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Boolean_expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Var_or_literal()
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(scriptyParserNAME)
}

func (s *AssignmentContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(scriptyParserNAME, i)
}

func (s *AssignmentContext) ASSIGNMENT_OP() antlr.TerminalNode {
	return s.GetToken(scriptyParserASSIGNMENT_OP, 0)
}

func (s *AssignmentContext) Boolean_expression() IBoolean_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *AssignmentContext) Math_expression() IMath_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMath_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *AssignmentContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *scriptyParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, scriptyParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(scriptyParserNAME)
		}
		{
			p.SetState(72)
			p.Match(scriptyParserASSIGNMENT_OP)
		}
		{
			p.SetState(73)
			p.Match(scriptyParserNAME)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(scriptyParserNAME)
		}
		{
			p.SetState(75)
			p.Match(scriptyParserASSIGNMENT_OP)
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(76)
				p.Boolean_expression()
			}

		case 2:
			{
				p.SetState(77)
				p.Math_expression()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Match(scriptyParserNAME)
		}
		{
			p.SetState(81)
			p.Match(scriptyParserASSIGNMENT_OP)
		}
		{
			p.SetState(82)
			p.Function_call()
		}

	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) NAME() antlr.TerminalNode {
	return s.GetToken(scriptyParserNAME, 0)
}

func (s *Function_callContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Function_callContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *scriptyParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, scriptyParserRULE_function_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(scriptyParserNAME)
	}
	{
		p.SetState(86)
		p.Match(scriptyParserT__1)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<scriptyParserNAME)|(1<<scriptyParserNUMBER)|(1<<scriptyParserSTRING))) != 0 {
		{
			p.SetState(87)
			p.Expr()
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == scriptyParserT__5 {
			{
				p.SetState(88)
				p.Match(scriptyParserT__5)
			}
			{
				p.SetState(89)
				p.Expr()
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(97)
		p.Match(scriptyParserT__2)
	}

	return localctx
}

// IBoolean_expressionContext is an interface to support dynamic dispatch.
type IBoolean_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolean_expressionContext differentiates from other interfaces.
	IsBoolean_expressionContext()
}

type Boolean_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_expressionContext() *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_boolean_expression
	return p
}

func (*Boolean_expressionContext) IsBoolean_expressionContext() {}

func NewBoolean_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_boolean_expression

	return p
}

func (s *Boolean_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_expressionContext) AllVar_or_literal() []IVar_or_literalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_or_literalContext)(nil)).Elem())
	var tst = make([]IVar_or_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_or_literalContext)
		}
	}

	return tst
}

func (s *Boolean_expressionContext) Var_or_literal(i int) IVar_or_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_or_literalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_or_literalContext)
}

func (s *Boolean_expressionContext) AllBOOLEAN_OP() []antlr.TerminalNode {
	return s.GetTokens(scriptyParserBOOLEAN_OP)
}

func (s *Boolean_expressionContext) BOOLEAN_OP(i int) antlr.TerminalNode {
	return s.GetToken(scriptyParserBOOLEAN_OP, i)
}

func (s *Boolean_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterBoolean_expression(s)
	}
}

func (s *Boolean_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitBoolean_expression(s)
	}
}

func (p *scriptyParser) Boolean_expression() (localctx IBoolean_expressionContext) {
	localctx = NewBoolean_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, scriptyParserRULE_boolean_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Var_or_literal()
	}
	{
		p.SetState(100)
		p.Match(scriptyParserBOOLEAN_OP)
	}
	{
		p.SetState(101)
		p.Var_or_literal()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == scriptyParserBOOLEAN_OP {
		{
			p.SetState(102)
			p.Match(scriptyParserBOOLEAN_OP)
		}
		{
			p.SetState(103)
			p.Var_or_literal()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMath_expressionContext is an interface to support dynamic dispatch.
type IMath_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMath_expressionContext differentiates from other interfaces.
	IsMath_expressionContext()
}

type Math_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_expressionContext() *Math_expressionContext {
	var p = new(Math_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_math_expression
	return p
}

func (*Math_expressionContext) IsMath_expressionContext() {}

func NewMath_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_expressionContext {
	var p = new(Math_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_math_expression

	return p
}

func (s *Math_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_expressionContext) AllVar_or_literal() []IVar_or_literalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_or_literalContext)(nil)).Elem())
	var tst = make([]IVar_or_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_or_literalContext)
		}
	}

	return tst
}

func (s *Math_expressionContext) Var_or_literal(i int) IVar_or_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_or_literalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_or_literalContext)
}

func (s *Math_expressionContext) AllARITHMETIC_OP() []antlr.TerminalNode {
	return s.GetTokens(scriptyParserARITHMETIC_OP)
}

func (s *Math_expressionContext) ARITHMETIC_OP(i int) antlr.TerminalNode {
	return s.GetToken(scriptyParserARITHMETIC_OP, i)
}

func (s *Math_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterMath_expression(s)
	}
}

func (s *Math_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitMath_expression(s)
	}
}

func (p *scriptyParser) Math_expression() (localctx IMath_expressionContext) {
	localctx = NewMath_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, scriptyParserRULE_math_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Var_or_literal()
	}
	{
		p.SetState(110)
		p.Match(scriptyParserARITHMETIC_OP)
	}
	{
		p.SetState(111)
		p.Var_or_literal()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == scriptyParserARITHMETIC_OP {
		{
			p.SetState(112)
			p.Match(scriptyParserARITHMETIC_OP)
		}
		{
			p.SetState(113)
			p.Var_or_literal()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVar_or_literalContext is an interface to support dynamic dispatch.
type IVar_or_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_or_literalContext differentiates from other interfaces.
	IsVar_or_literalContext()
}

type Var_or_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_or_literalContext() *Var_or_literalContext {
	var p = new(Var_or_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_var_or_literal
	return p
}

func (*Var_or_literalContext) IsVar_or_literalContext() {}

func NewVar_or_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_or_literalContext {
	var p = new(Var_or_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_var_or_literal

	return p
}

func (s *Var_or_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_or_literalContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Var_or_literalContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Var_or_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_or_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_or_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterVar_or_literal(s)
	}
}

func (s *Var_or_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitVar_or_literal(s)
	}
}

func (p *scriptyParser) Var_or_literal() (localctx IVar_or_literalContext) {
	localctx = NewVar_or_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, scriptyParserRULE_var_or_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case scriptyParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Variable()
		}

	case scriptyParserNUMBER, scriptyParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) NAME() antlr.TerminalNode {
	return s.GetToken(scriptyParserNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *scriptyParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, scriptyParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(scriptyParserNAME)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scriptyParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scriptyParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(scriptyParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(scriptyParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scriptyListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *scriptyParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, scriptyParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(125)
	_la = p.GetTokenStream().LA(1)

	if !(_la == scriptyParserNUMBER || _la == scriptyParserSTRING) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
