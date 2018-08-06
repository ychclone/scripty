// Generated from scripty.g4 by ANTLR 4.7.

package parsergen // scripty
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasescriptyListener is a complete listener for a parse tree produced by scriptyParser.
type BasescriptyListener struct{}

var _ scriptyListener = &BasescriptyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasescriptyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasescriptyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasescriptyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasescriptyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasescriptyListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasescriptyListener) ExitProgram(ctx *ProgramContext) {}

// EnterFunction_def is called when production function_def is entered.
func (s *BasescriptyListener) EnterFunction_def(ctx *Function_defContext) {}

// ExitFunction_def is called when production function_def is exited.
func (s *BasescriptyListener) ExitFunction_def(ctx *Function_defContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BasescriptyListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BasescriptyListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterParameter_defs is called when production parameter_defs is entered.
func (s *BasescriptyListener) EnterParameter_defs(ctx *Parameter_defsContext) {}

// ExitParameter_defs is called when production parameter_defs is exited.
func (s *BasescriptyListener) ExitParameter_defs(ctx *Parameter_defsContext) {}

// EnterFunction_expressions is called when production function_expressions is entered.
func (s *BasescriptyListener) EnterFunction_expressions(ctx *Function_expressionsContext) {}

// ExitFunction_expressions is called when production function_expressions is exited.
func (s *BasescriptyListener) ExitFunction_expressions(ctx *Function_expressionsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasescriptyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasescriptyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasescriptyListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasescriptyListener) ExitExpr(ctx *ExprContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasescriptyListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasescriptyListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BasescriptyListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BasescriptyListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterBoolean_expression is called when production boolean_expression is entered.
func (s *BasescriptyListener) EnterBoolean_expression(ctx *Boolean_expressionContext) {}

// ExitBoolean_expression is called when production boolean_expression is exited.
func (s *BasescriptyListener) ExitBoolean_expression(ctx *Boolean_expressionContext) {}

// EnterMath_expression is called when production math_expression is entered.
func (s *BasescriptyListener) EnterMath_expression(ctx *Math_expressionContext) {}

// ExitMath_expression is called when production math_expression is exited.
func (s *BasescriptyListener) ExitMath_expression(ctx *Math_expressionContext) {}

// EnterVar_or_literal is called when production var_or_literal is entered.
func (s *BasescriptyListener) EnterVar_or_literal(ctx *Var_or_literalContext) {}

// ExitVar_or_literal is called when production var_or_literal is exited.
func (s *BasescriptyListener) ExitVar_or_literal(ctx *Var_or_literalContext) {}

// EnterVar is called when production var is entered.
func (s *BasescriptyListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BasescriptyListener) ExitVar(ctx *VarContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasescriptyListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasescriptyListener) ExitLiteral(ctx *LiteralContext) {}
