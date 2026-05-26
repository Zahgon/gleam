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

package expression

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/util/distinct"
	"github.com/chrislusf/gleam/sql/util/types"
)

// AggregationFunction stands for aggregate functions.
type AggregationFunction interface {
	fmt.Stringer
	json.Marshaler
	// Update during executing.
	Update(row []types.Datum, groupKey []byte, ctx context.Context) error

	// StreamUpdate updates data using streaming algo.
	StreamUpdate(row []types.Datum, ctx context.Context) error

	// SetMode sets aggFunctionMode for aggregate function.
	SetMode(mode AggFunctionMode)

	// GetMode gets aggFunctionMode from aggregate function.
	GetMode() AggFunctionMode

	// GetGroupResult will be called when all data have been processed.
	GetGroupResult(groupKey []byte) types.Datum

	// GetStreamResult gets a result using streaming agg.
	GetStreamResult() types.Datum

	// GetArgs stands for getting all arguments.
	GetArgs() []Expression

	// GetName gets the aggregation function name.
	GetName() string

	// SetArgs sets argument by index.
	SetArgs(args []Expression)

	// Clear collects the mapper's memory.
	Clear()

	// IsDistinct indicates if the aggregate function contains distinct attribute.
	IsDistinct() bool

	// SetContext sets the aggregate evaluation context.
	SetContext(ctx map[string](*aggEvaluateContext))

	// Equal checks whether two aggregation functions are equal.
	Equal(agg AggregationFunction, ctx context.Context) bool

	// Clone copies an aggregate function totally.
	Clone() AggregationFunction

	// GetType gets field type of aggregate function.
	GetType() *types.FieldType

	// CalculateDefaultValue gets the default value when the aggregate function's input is null.
	// The input stands for the schema of Aggregation's child. If the function can't produce a default value, the second
	// return value will be false.
	CalculateDefaultValue(schema Schema, ctx context.Context) (types.Datum, bool)
}

// aggEvaluateContext is used to store intermediate result when calculating aggregate functions.
type aggEvaluateContext struct {
	DistinctChecker *distinct.Checker
	Count           int64
	Value           types.Datum
	Buffer          *bytes.Buffer // Buffer is used for group_concat.
	GotFirstRow     bool          // It will check if the agg has met the first row key.
}

// NewAggFunction creates a new AggregationFunction.
func NewAggFunction(funcType string, funcArgs []Expression, distinct bool) AggregationFunction {
	_ = "STUB: not implemented"
	return *new(AggregationFunction)
}

type aggCtxMapper map[string]*aggEvaluateContext

// AggFunctionMode stands for the aggregation function's mode.
type AggFunctionMode int

const (
	// CompleteMode function accepts origin data.
	CompleteMode AggFunctionMode = iota
	// FinalMode function accepts partial data.
	FinalMode
)

type aggFunction struct {
	name         string
	mode         AggFunctionMode
	Args         []Expression
	Distinct     bool
	resultMapper aggCtxMapper
	streamCtx    *aggEvaluateContext
}

// Equal implements AggregationFunction interface.
func (af *aggFunction) Equal(b AggregationFunction, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// String implements fmt.Stringer interface.
func (af *aggFunction) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implements json.Marshaler interface.
func (af *aggFunction) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func newAggFunc(name string, args []Expression, dist bool) aggFunction {
	_ = "STUB: not implemented"
	return *new(aggFunction)
}

// CalculateDefaultValue implements AggregationFunction interface.
func (af *aggFunction) CalculateDefaultValue(schema Schema, ctx context.Context) (types.Datum, bool) {
	_ = "STUB: not implemented"
	return *new(types.Datum), false
}

// IsDistinct implements AggregationFunction interface.
func (af *aggFunction) IsDistinct() bool {
	_ = "STUB: not implemented"

	// Clear implements AggregationFunction interface.
	return false
}

func (af *aggFunction) Clear() { _ = "STUB: not implemented"; return }

// GetName implements AggregationFunction interface.
func (af *aggFunction) GetName() string {
	_ = "STUB: not implemented"

	// SetMode implements AggregationFunction interface.
	return ""
}

func (af *aggFunction) SetMode(mode AggFunctionMode) {
	_ = "STUB: not implemented"

	// GetMode implements AggregationFunction interface.
	return
}

func (af *aggFunction) GetMode() AggFunctionMode {
	_ = "STUB: not implemented"

	// GetArgs implements AggregationFunction interface.
	return *new(AggFunctionMode)
}

func (af *aggFunction) GetArgs() []Expression {
	_ = "STUB: not implemented"

	// SetArgs implements AggregationFunction interface.
	return nil
}

func (af *aggFunction) SetArgs(args []Expression) { _ = "STUB: not implemented"; return }

func (af *aggFunction) getContext(groupKey []byte) *aggEvaluateContext {
	_ = "STUB: not implemented"
	return nil
}

func (af *aggFunction) getStreamedContext() *aggEvaluateContext {
	_ = "STUB: not implemented"
	return nil
}

// SetContext implements AggregationFunction interface.
func (af *aggFunction) SetContext(ctx map[string](*aggEvaluateContext)) {
	_ = "STUB: not implemented"
	return
}

func (af *aggFunction) updateSum(row []types.Datum, groupKey []byte, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (af *aggFunction) streamUpdateSum(row []types.Datum, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type sumFunction struct {
	aggFunction
}

// Clone implements AggregationFunction interface.
func (sf *sumFunction) Clone() AggregationFunction {
	_ = "STUB: not implemented"
	return *new(AggregationFunction)
}

// Update implements AggregationFunction interface.
func (sf *sumFunction) Update(row []types.Datum, groupKey []byte, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// StreamUpdate implements AggregationFunction interface.
func (sf *sumFunction) StreamUpdate(row []types.Datum, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetGroupResult implements AggregationFunction interface.
func (sf *sumFunction) GetGroupResult(groupKey []byte) (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// GetStreamResult implements AggregationFunction interface.
func (sf *sumFunction) GetStreamResult() (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// CalculateDefaultValue implements AggregationFunction interface.
func (sf *sumFunction) CalculateDefaultValue(schema Schema, ctx context.Context) (d types.Datum, valid bool) {
	_ = "STUB: not implemented"
	return *new(types.Datum), false
}

// GetType implements AggregationFunction interface.
func (sf *sumFunction) GetType() *types.FieldType { _ = "STUB: not implemented"; return nil }

type countFunction struct {
	aggFunction
}

// Clone implements AggregationFunction interface.
func (cf *countFunction) Clone() AggregationFunction {
	_ = "STUB: not implemented"
	return *new(AggregationFunction)
}

// CalculateDefaultValue implements AggregationFunction interface.
func (cf *countFunction) CalculateDefaultValue(schema Schema, ctx context.Context) (d types.Datum, valid bool) {
	_ = "STUB: not implemented"
	return *new(types.Datum), false
}

// GetType implements AggregationFunction interface.
func (cf *countFunction) GetType() *types.FieldType { _ = "STUB: not implemented"; return nil }

// Update implements AggregationFunction interface.
func (cf *countFunction) Update(row []types.Datum, groupKey []byte, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// StreamUpdate implements AggregationFunction interface.
func (cf *countFunction) StreamUpdate(row []types.Datum, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetGroupResult implements AggregationFunction interface.
func (cf *countFunction) GetGroupResult(groupKey []byte) (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// GetStreamResult implements AggregationFunction interface.
func (cf *countFunction) GetStreamResult() (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

type avgFunction struct {
	aggFunction
}

// Clone implements AggregationFunction interface.
func (af *avgFunction) Clone() AggregationFunction {
	_ = "STUB: not implemented"
	return *new(AggregationFunction)
}

// GetType implements AggregationFunction interface.
func (af *avgFunction) GetType() *types.FieldType { _ = "STUB: not implemented"; return nil }

func (af *avgFunction) updateAvg(row []types.Datum, groupKey []byte, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Update implements AggregationFunction interface.
func (af *avgFunction) Update(row []types.Datum, groupKey []byte, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// StreamUpdate implements AggregationFunction interface.
func (af *avgFunction) StreamUpdate(row []types.Datum, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (af *avgFunction) calculateResult(ctx *aggEvaluateContext) (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// GetGroupResult implements AggregationFunction interface.
func (af *avgFunction) GetGroupResult(groupKey []byte) types.Datum {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// GetStreamResult implements AggregationFunction interface.
func (af *avgFunction) GetStreamResult() (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

type concatFunction struct {
	aggFunction
}

// Clone implements AggregationFunction interface.
func (cf *concatFunction) Clone() AggregationFunction {
	_ = "STUB: not implemented"
	return *new(AggregationFunction)
}

// GetType implements AggregationFunction interface.
func (cf *concatFunction) GetType() *types.FieldType { _ = "STUB: not implemented"; return nil }

// Update implements AggregationFunction interface.
func (cf *concatFunction) Update(row []types.Datum, groupKey []byte, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// now use comma separator

// TODO: if total length is greater than global var group_concat_max_len, truncate it.

// StreamUpdate implements AggregationFunction interface.
func (cf *concatFunction) StreamUpdate(row []types.Datum, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// now use comma separator

// TODO: if total length is greater than global var group_concat_max_len, truncate it.

// GetGroupResult implements AggregationFunction interface.
func (cf *concatFunction) GetGroupResult(groupKey []byte) (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// GetStreamResult implements AggregationFunction interface.
func (cf *concatFunction) GetStreamResult() (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

type maxMinFunction struct {
	aggFunction
	isMax bool
}

// Clone implements AggregationFunction interface.
func (mmf *maxMinFunction) Clone() AggregationFunction {
	_ = "STUB: not implemented"
	return *new(AggregationFunction)
}

// CalculateDefaultValue implements AggregationFunction interface.
func (mmf *maxMinFunction) CalculateDefaultValue(schema Schema, ctx context.Context) (d types.Datum, valid bool) {
	_ = "STUB: not implemented"
	return *new(types.Datum), false
}

// GetType implements AggregationFunction interface.
func (mmf *maxMinFunction) GetType() *types.FieldType { _ = "STUB: not implemented"; return nil }

// GetGroupResult implements AggregationFunction interface.
func (mmf *maxMinFunction) GetGroupResult(groupKey []byte) (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// GetStreamResult implements AggregationFunction interface.
func (mmf *maxMinFunction) GetStreamResult() (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// Update implements AggregationFunction interface.
func (mmf *maxMinFunction) Update(row []types.Datum, groupKey []byte, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// StreamUpdate implements AggregationFunction interface.
func (mmf *maxMinFunction) StreamUpdate(row []types.Datum, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type firstRowFunction struct {
	aggFunction
}

// Clone implements AggregationFunction interface.
func (ff *firstRowFunction) Clone() AggregationFunction {
	_ = "STUB: not implemented"
	return *new(AggregationFunction)
}

// GetType implements AggregationFunction interface.
func (ff *firstRowFunction) GetType() *types.FieldType { _ = "STUB: not implemented"; return nil }

// Update implements AggregationFunction interface.
func (ff *firstRowFunction) Update(row []types.Datum, groupKey []byte, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// StreamUpdate implements AggregationFunction interface.
func (ff *firstRowFunction) StreamUpdate(row []types.Datum, ectx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetGroupResult implements AggregationFunction interface.
func (ff *firstRowFunction) GetGroupResult(groupKey []byte) types.Datum {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// GetStreamResult implements AggregationFunction interface.
func (ff *firstRowFunction) GetStreamResult() (d types.Datum) {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}

// CalculateDefaultValue implements AggregationFunction interface.
func (ff *firstRowFunction) CalculateDefaultValue(schema Schema, ctx context.Context) (d types.Datum, valid bool) {
	_ = "STUB: not implemented"
	return *new(types.Datum), false
}
