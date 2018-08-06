grammar scripty;

tokens { INDENT, DEDENT, NEWLINE, ENDMARKER }

program: (function_def | expression)+ ;

function_def: 'def' function_name '(' parameter_defs ')' '{' function_expressions '}' ;

function_name : NAME ;

parameter_defs: (NAME (',' NAME)*)? ;

function_expressions: expression* ;

expression: (assignment | function_call | expr) ;

expr: math_expression | boolean_expression | var_or_literal ;

assignment:
      NAME ASSIGNMENT_OP NAME STMT_END
    | NAME ASSIGNMENT_OP (boolean_expression | math_expression) STMT_END
    | NAME ASSIGNMENT_OP function_call STMT_END
    ;

function_call: NAME '(' (expr (',' expr)*)? ')' STMT_END ;

boolean_expression: var_or_literal BOOLEAN_OP var_or_literal (BOOLEAN_OP var_or_literal)* ;

math_expression: var_or_literal ARITHMETIC_OP var_or_literal (ARITHMETIC_OP var_or_literal)* ;

var_or_literal: var | literal ;

var: NAME;

literal: NUMBER | STRING;

ASSIGNMENT_OP: '=' ;

STMT_END: ';' ;

BOOLEAN_OP: '&&' | '||' ;

ARITHMETIC_OP: '+' | '-' | '*' | '/' ;

NAME: [a-zA-Z_] [a-zA-Z0-9_]*
    ;

NUMBER
    :   '0' ([xX] [0-9a-fA-F]+         ([lL]  | [eE] [+-]? [0-9]+)?
    |        [oO] [0-7]+                [lL]?
    |        [bB] [01]+                 [lL]?)
    | ([0-9]+ '.' [0-9]* | '.' [0-9]+)         ([eE] [+-]? [0-9]+)?       [jJ]?
    |  [0-9]+                          ([lL]  | [eE] [+-]? [0-9]+ [jJ]? | [jJ])?
    ;

STRING
    : ([uUbB]? [rR]? | [rR]? [uUbB]?)
    ( '\''     ('\\' (([ \t]+ ('\r'? '\n')?)|.) | ~[\\\r\n'])*  '\''
    | '"'      ('\\' (([ \t]+ ('\r'? '\n')?)|.) | ~[\\\r\n"])*  '"'
    | '"""'    ('\\' .                          | ~'\\'     )*? '"""'
    | '\'\'\'' ('\\' .                          | ~'\\'     )*? '\'\'\''
    )
    ;

WHITESPACE: ('\t' | ' ')+ -> channel(HIDDEN);
COMMENT: '#' ~[\r\n]* -> skip;
UNKNOWN: . -> skip;
ANY: . ;