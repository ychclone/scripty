// Generated from scripty.g4 by ANTLR 4.7.

package parsergen // scripty
import "github.com/antlr/antlr4/runtime/Go/antlr"

// scriptyListener is a complete listener for a parse tree produced by scriptyParser.
type scriptyListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFunction_def is called when entering the function_def production.
	EnterFunction_def(c *Function_defContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterParameter_defs is called when entering the parameter_defs production.
	EnterParameter_defs(c *Parameter_defsContext)

	// EnterFunction_expressions is called when entering the function_expressions production.
	EnterFunction_expressions(c *Function_expressionsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterBoolean_expression is called when entering the boolean_expression production.
	EnterBoolean_expression(c *Boolean_expressionContext)

	// EnterMath_expression is called when entering the math_expression production.
	EnterMath_expression(c *Math_expressionContext)

	// EnterVar_or_literal is called when entering the var_or_literal production.
	EnterVar_or_literal(c *Var_or_literalContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFunction_def is called when exiting the function_def production.
	ExitFunction_def(c *Function_defContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitParameter_defs is called when exiting the parameter_defs production.
	ExitParameter_defs(c *Parameter_defsContext)

	// ExitFunction_expressions is called when exiting the function_expressions production.
	ExitFunction_expressions(c *Function_expressionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitBoolean_expression is called when exiting the boolean_expression production.
	ExitBoolean_expression(c *Boolean_expressionContext)

	// ExitMath_expression is called when exiting the math_expression production.
	ExitMath_expression(c *Math_expressionContext)

	// ExitVar_or_literal is called when exiting the var_or_literal production.
	ExitVar_or_literal(c *Var_or_literalContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
