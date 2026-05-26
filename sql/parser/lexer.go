// Copyright 2016 PingCAP, Inc.
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
	"bytes"
)

var _ = yyLexer(&Scanner{})

// Pos represents the position of a token.
type Pos struct {
	Line   int
	Col    int
	Offset int
}

// Scanner implements the yyLexer interface.
type Scanner struct {
	r   reader
	buf bytes.Buffer

	errs         []error
	stmtStartPos int

	// for scanning such kind of comment: /*! MySQL-specific code */
	specialComment *specialCommentScanner
}

type specialCommentScanner struct {
	*Scanner
	Pos
}

// Errors returns the errors during a scan.
func (s *Scanner) Errors() []error {
	_ = "STUB: not implemented"

	// reset resets the sql string to be scanned.
	return nil
}

func (s *Scanner) reset(sql string) { _ = "STUB: not implemented"; return }

func (s *Scanner) stmtText() string { _ = "STUB: not implemented"; return "" }

// trim new line

// Errorf tells scanner something is wrong.
// Scanner satisfies yyLexer interface which need this function.
func (s *Scanner) Errorf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Lex returns a token and store the token value in v.
// Scanner satisfies yyLexer interface.
// 0 and invalid are special token id this function would return:
// return 0 tells parser that scanner meets EOF,
// return invalid tells parser that scanner meets illegal character.
func (s *Scanner) Lex(v *yySymType) int { _ = "STUB: not implemented"; return 0 }

// NewScanner returns a new scanner object.
func NewScanner(s string) *Scanner { _ = "STUB: not implemented"; return nil }

func (s *Scanner) skipEmptySpace() rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) scan() (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// enter specialComment scan mode.
// for scanning such kind of comment: /*! MySQL-specific code */

// return the specialComment scan result as the result

// leave specialComment scan mode after all stream consumed.

// when scanner meets EOF, the returned token should be 0,
// because 0 is a special token id to remind the parser that stream is end.

// search a trie to get a token.

func startWithXx(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithb(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithSharp(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithDash(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithSlash(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// See http://dev.mysql.com/doc/refman/5.7/en/comments.html
// Convert "/*!VersionNumber MySQL-specific-code */" to "MySQL-specific-code".

func sqlOffsetInComment(comment string) int {
	_ = "STUB: not implemented"
	// find the first SQL token offset in pattern like "/*!40101 mysql specific code */"
	return 0
}

func startWithAt(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func scanIdentifier(s *Scanner) (int, Pos, string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

var (
	quotedIdentifier = -identifier
)

func scanQuotedIdent(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// don't return identifier in case that it's interpreted as keyword token later.

func startString(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// Quoted strings placed next to each other are concatenated to a single string.
// See http://dev.mysql.com/doc/refman/5.7/en/string-literals.html

// lazyBuf is used to avoid allocation if possible.
// it has a useBuf field indicates whether bytes.Buffer is necessary. if
// useBuf is false, we can avoid calling bytes.Buffer.String(), which
// make a copy of data and cause allocation.
type lazyBuf struct {
	useBuf bool
	r      *reader
	b      *bytes.Buffer
	p      *Pos
}

func (mb *lazyBuf) setUseBuf(str string) { _ = "STUB: not implemented"; return }

func (mb *lazyBuf) writeRune(r rune, w int) { _ = "STUB: not implemented"; return }

func (mb *lazyBuf) data() string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) scanString() (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// handleEscape handles the case in scanString when previous char is '\'.
func handleEscape(s *Scanner) rune { _ = "STUB: not implemented"; return 0 }

/*
	\" \' \\ \n \0 \b \Z \r \t ==> escape to one char
	\% \_ ==> preserve both char
	other ==> remove \
*/

func startWithNumber(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// Identifiers may begin with a digit but unless quoted may not consist solely of digits.

func startWithDot(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// Fail to parse a float, reset to dot.

func (s *Scanner) scanOct() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanHex() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanBit() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanFloat(beg *Pos) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"

	// float = D1 . D2 e D3
	return 0, *new(Pos), ""
}

func (s *Scanner) scanDigits() string { _ = "STUB: not implemented"; return "" }

type reader struct {
	s string
	p Pos
	w int
}

var eof = Pos{-1, -1, -1}

func (r *reader) eof() bool { _ = "STUB: not implemented"; return false }

// peek() peeks a rune from underlying reader.
// if reader meets EOF, it will return unicode.ReplacementChar. to distinguish from
// the real unicode.ReplacementChar, the caller should call r.eof() again to check.
func (r *reader) peek() rune { _ = "STUB: not implemented"; return 0 }

// illegal UTF-8 encoding

// illegal UTF-8 encoding

// inc increase the position offset of the reader.
// peek must be called before calling inc!
func (r *reader) inc() { _ = "STUB: not implemented"; return }

func (r *reader) incN(n int) { _ = "STUB: not implemented"; return }

func (r *reader) readByte() (ch rune) { _ = "STUB: not implemented"; return 0 }

func (r *reader) pos() Pos { _ = "STUB: not implemented"; return *new(Pos) }

func (r *reader) data(from *Pos) string { _ = "STUB: not implemented"; return "" }

func (r *reader) incAsLongAs(fn func(rune) bool) rune { _ = "STUB: not implemented"; return 0 }
