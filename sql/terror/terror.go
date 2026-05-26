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

package terror

import (
	"github.com/chrislusf/gleam/sql/mysql"
)

// Common base error instances.
var (
	CommitNotInTransaction   = ClassExecutor.New(CodeCommitNotInTransaction, "commit not in transaction")
	RollbackNotInTransaction = ClassExecutor.New(CodeRollbackNotInTransaction, "rollback not in transaction")
	ExecResultIsEmpty        = ClassExecutor.New(CodeExecResultIsEmpty, "exec result is empty")

	MissConnectionID = ClassExpression.New(CodeMissConnectionID, "miss connection id information")
)

// ErrCode represents a specific error type in a error class.
// Same error code can be used in different error classes.
type ErrCode int

// Executor error codes.
const (
	CodeUnknown                  ErrCode = -1
	CodeCommitNotInTransaction           = 1
	CodeRollbackNotInTransaction         = 2
	CodeExecResultIsEmpty                = 3
)

// Expression error codes.
const (
	CodeMissConnectionID ErrCode = iota + 1
)

// ErrClass represents a class of errors.
type ErrClass int

// Error classes.
const (
	ClassAutoid ErrClass = iota + 1
	ClassDDL
	ClassDomain
	ClassEvaluator
	ClassExecutor
	ClassExpression
	ClassInspectkv
	ClassKV
	ClassMeta
	ClassOptimizer
	ClassOptimizerPlan
	ClassParser
	ClassPerfSchema
	ClassPrivilege
	ClassSchema
	ClassServer
	ClassStructure
	ClassVariable
	ClassXEval
	ClassTable
	ClassTypes
	// Add more as needed.
)

// String implements fmt.Stringer interface.
func (ec ErrClass) String() string { _ = "STUB: not implemented"; return "" }

// EqualClass returns true if err is *Error with the same class.
func (ec ErrClass) EqualClass(err error) bool { _ = "STUB: not implemented"; return false }

// NotEqualClass returns true if err is not *Error with the same class.
func (ec ErrClass) NotEqualClass(err error) bool { _ = "STUB: not implemented"; return false }

// New creates an *Error with an error code and an error message.
// Usually used to create base *Error.
func (ec ErrClass) New(code ErrCode, message string) *Error { _ = "STUB: not implemented"; return nil }

// Error implements error interface and adds integer Class and Code, so
// errors with different message can be compared.
type Error struct {
	class   ErrClass
	code    ErrCode
	message string
	args    []interface{}
	file    string
	line    int
}

// Class returns ErrClass
func (e *Error) Class() ErrClass {
	_ = "STUB: not implemented"

	// Code returns ErrCode
	return *new(ErrClass)
}

func (e *Error) Code() ErrCode {
	_ = "STUB: not implemented"

	// MarshalJSON implements json.Marshaler interface.
	return *new(ErrCode)
}

func (e *Error) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler interface.
func (e *Error) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Location returns the location where the error is created,
// implements juju/errors locationer interface.
func (e *Error) Location() (file string, line int) {
	_ = "STUB: not implemented"
	return "",

		// Error implements error interface.
		0
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) getMsg() string { _ = "STUB: not implemented"; return "" }

// Gen generates a new *Error with the same class and code, and a new formatted message.
func (e *Error) Gen(format string, args ...interface{}) *Error {
	_ = "STUB: not implemented"
	return nil
}

// GenByArgs generates a new *Error with the same class and code, and new arguments.
func (e *Error) GenByArgs(args ...interface{}) *Error { _ = "STUB: not implemented"; return nil }

// FastGen generates a new *Error with the same class and code, and a new formatted message.
// This will not call runtime.Caller to get file and line.
func (e *Error) FastGen(format string, args ...interface{}) *Error {
	_ = "STUB: not implemented"
	return nil
}

// Equal checks if err is equal to e.
func (e *Error) Equal(err error) bool { _ = "STUB: not implemented"; return false }

// NotEqual checks if err is not equal to e.
func (e *Error) NotEqual(err error) bool { _ = "STUB: not implemented"; return false }

// ToSQLError convert Error to mysql.SQLError.
func (e *Error) ToSQLError() *mysql.SQLError { _ = "STUB: not implemented"; return nil }

var defaultMySQLErrorCode uint16

func (e *Error) getMySQLErrorCode() uint16 { _ = "STUB: not implemented"; return 0 }

var (
	// ErrClassToMySQLCodes is the map of ErrClass to code-map.
	ErrClassToMySQLCodes map[ErrClass](map[ErrCode]uint16)
)

func init() {
	ErrClassToMySQLCodes = make(map[ErrClass](map[ErrCode]uint16))
	defaultMySQLErrorCode = mysql.ErrUnknown
}

// ErrorEqual returns a boolean indicating whether err1 is equal to err2.
func ErrorEqual(err1, err2 error) bool { _ = "STUB: not implemented"; return false }

// ErrorNotEqual returns a boolean indicating whether err1 isn't equal to err2.
func ErrorNotEqual(err1, err2 error) bool { _ = "STUB: not implemented"; return false }
