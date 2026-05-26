// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"regexp"

	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/terror"
)

// Error instances.
var (
	ErrSyntax = terror.ClassParser.New(CodeSyntaxErr, "syntax error")
)

// Error codes.
const (
	CodeSyntaxErr terror.ErrCode = 1
)

var (
	specCodePattern = regexp.MustCompile(`\/\*!(M?[0-9]{5,6})?([^*]|\*+[^*/])*\*+\/`)
	specCodeStart   = regexp.MustCompile(`^\/\*!(M?[0-9]{5,6} )?[ \t]*`)
	specCodeEnd     = regexp.MustCompile(`[ \t]*\*\/$`)
)

func trimComment(txt string) string { _ = "STUB: not implemented"; return "" }

// Parser represents a parser instance. Some temporary objects are stored in it to reduce object allocation during Parse function.
type Parser struct {
	charset   string
	collation string
	result    []ast.StmtNode
	src       string
	lexer     Scanner

	// the following fields are used by yyParse to reduce allocation.
	cache  []yySymType
	yylval yySymType
	yyVAL  yySymType
}

type stmtTexter interface {
	stmtText() string
}

// New returns a Parser object.
func New() *Parser { _ = "STUB: not implemented"; return nil }

// Parse parses a query string to raw ast.StmtNode.
// If charset or collation is "", default charset and collation will be used.
func (parser *Parser) Parse(sql, charset, collation string) ([]ast.StmtNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseOneStmt parses a query and returns an ast.StmtNode.
// The query must have one statement, otherwise ErrSyntax is returned.
func (parser *Parser) ParseOneStmt(sql, charset, collation string) (ast.StmtNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.StmtNode), nil
}

// The select statement is not at the end of the whole statement, if the last
// field text was set from its offset to the end of the src string, update
// the last field text.
func (parser *Parser) setLastSelectFieldText(st *ast.SelectStmt, lastEnd int) {
	_ = "STUB: not implemented"
	return
}

func (parser *Parser) startOffset(v *yySymType) int { _ = "STUB: not implemented"; return 0 }

func (parser *Parser) endOffset(v *yySymType) int { _ = "STUB: not implemented"; return 0 }

func toInt(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

func toDecimal(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

func toFloat(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

// See https://dev.mysql.com/doc/refman/5.7/en/hexadecimal-literals.html
func toHex(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

// If parse hexadecimal literal to numerical value error, we should treat it as a string.

// See https://dev.mysql.com/doc/refman/5.7/en/bit-type.html
func toBit(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }
