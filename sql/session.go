// Copyright 2013 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

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

package sql

import (
	"fmt"
	"sync"

	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/infoschema"
	"github.com/chrislusf/gleam/sql/parser"
	"github.com/chrislusf/gleam/sql/plan"
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
)

// Session context
type Session interface {
	context.Context
	Status() uint16 // Flag of current status, such as autocommit.
	String() string // For debug
	Close() error
}

var (
	_         Session = (*session)(nil)
	sessionMu sync.Mutex
)

type stmtRecord struct {
	stmtID uint32
	st     ast.Statement
	params []interface{}
}

type stmtHistory struct {
	history []*stmtRecord
}

func (h *stmtHistory) add(stmtID uint32, st ast.Statement, params ...interface{}) {
	_ = "STUB: not implemented"
	return
}

type session struct {
	values      map[fmt.Stringer]interface{}
	parser      *parser.Parser
	sessionVars *variable.SessionVars
}

func (s *session) Status() uint16 { _ = "STUB: not implemented"; return 0 }

func (s *session) String() string {
	_ = "STUB: not implemented"
	// TODO: how to print binded context in values appropriately?
	return ""
}

func (s *session) ParseSQL(sql, charset, collation string) ([]ast.StmtNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkArgs makes sure all the arguments' types are known and can be handled.
// integer types are converted to int64 and uint64, time.Time is converted to types.Time.
// time.Duration is converted to types.Duration, other known types are leaved as it is.
func checkArgs(args ...interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *session) SetValue(key fmt.Stringer, value interface{}) { _ = "STUB: not implemented"; return }

func (s *session) Value(key fmt.Stringer) interface{} { _ = "STUB: not implemented"; return nil }

func (s *session) ClearValue(key fmt.Stringer) { _ = "STUB: not implemented"; return }

// Close function does some clean work when session end.
func (s *session) Close() error {
	_ = "STUB: not implemented"

	// GetSessionVars implements the context.Context interface.
	return nil
}

func (s *session) GetSessionVars() *variable.SessionVars { _ = "STUB: not implemented"; return nil }

// Some vars name for debug.
const (
	retryEmptyHistoryList = "RetryEmptyHistoryList"
)

// CreateSession creates a new session environment.
func CreateSession(info infoschema.InfoSchema) (Session, error) {
	_ = "STUB: not implemented"
	return *new(Session), nil
}

func createSession(info infoschema.InfoSchema) (*session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compile is safe for concurrent use by multiple goroutines.
func Compile(ctx context.Context, rawStmt ast.StmtNode) (plan.Plan, error) {
	_ = "STUB: not implemented"
	return *new(plan.Plan), nil
}

// runStmt executes the ast.Statement and commit or rollback the current transaction.
func runStmt(ctx context.Context, s ast.Statement) (ast.RecordSet, error) {
	_ = "STUB: not implemented"
	return *new(ast.RecordSet), nil
}
