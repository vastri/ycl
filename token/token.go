// Copyright 2016 Vastri. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

// Package token defines constants representing the lexical tokens of YCL.
package token

type Token int

// The list of tokens.
const (
	token_beg Token = iota

	// Special tokens.
	ILLEGAL
	EOF
	COMMENT

	// Identifiers and basic type literals.
	literal_beg

	IDENT
	BOOL
	INT
	FLOAT
	STRING

	literal_end

	// Operators and delimiters.
	operator_beg

	NOT // !
	AND // &&
	OR  // ||

	EQL // ==
	NEQ // !=
	LSS // <
	LEQ // <=
	GTR // >
	GEQ // >=

	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	LPAREN // (
	LBRACK // [
	LBRACE // {
	RPAREN // )
	RBRACK // ]
	RBRACE // }

	COMMA  // ,
	COLON  // :
	PERIOD // .
	ASSIGN // =

	operator_end

	token_end = operator_end
)

// String returns the string corresponding to the token tok.
func (tok Token) String() string {
	switch tok {
	case EOF:
		return "EOF"
	case COMMENT:
		return "COMMENT"
	case IDENT:
		return "IDENT"
	case BOOL:
		return "BOOL"
	case INT:
		return "INT"
	case FLOAT:
		return "FLOAT"
	case STRING:
		return "STRING"
	case NOT:
		return "!"
	case AND:
		return "&&"
	case OR:
		return "||"
	case EQL:
		return "=="
	case NEQ:
		return "!="
	case LSS:
		return "<"
	case LEQ:
		return "<="
	case GTR:
		return ">"
	case GEQ:
		return ">="
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case QUO:
		return "/"
	case REM:
		return "%"
	case LPAREN:
		return "("
	case LBRACK:
		return "["
	case LBRACE:
		return "{"
	case RPAREN:
		return ")"
	case RBRACK:
		return "]"
	case RBRACE:
		return "}"
	case COMMA:
		return ","
	case COLON:
		return ":"
	case PERIOD:
		return "."
	case ASSIGN:
		return "="
	default:
		return "ILLEGAL"
	}
}

// Precedence returns the operator precedence of the operator op.
func (op Token) Precedence() int {
	switch op {
	case OR:
		return 1
	case AND:
		return 2
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 3
	case ADD, SUB:
		return 4
	case MUL, QUO, REM:
		return 5
	default:
		return 0
	}
}

// IsLiteral returns true for tokens corresponding to identifiers and literals.
func (tok Token) IsLiteral() bool {
	return literal_beg < tok && tok < literal_end
}

// IsOperator returns true for tokens corresponding to operators and delimiters.
func (tok Token) IsOperator() bool {
	return operator_beg < tok && tok < operator_end
}
