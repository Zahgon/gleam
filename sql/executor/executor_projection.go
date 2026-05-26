package executor

import (
	"regexp"

	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/expression"
)

type ProjectionExec struct {
	Src      Executor
	schema   expression.Schema
	executed bool
	ctx      context.Context
	exprs    []expression.Expression
}

// Schema implements the Executor Schema interface.
func (e *ProjectionExec) Schema() expression.Schema {
	_ = "STUB: not implemented"

	// Next implements the Executor Next interface.
	return *new(expression.Schema)
}

func (e *ProjectionExec) Exec() *flow.Dataset {
	_ = "STUB: not implemented"

	// fmt.Printf("=================\n")
	// fmt.Printf("context: %+v\n", e.ctx)
	return nil
}

// fmt.Printf("input: %s TblName:%s DBName:%s FromID:%s %d\n", col, col.TblName, col.DBName, col.FromID, col.Position)

/*
	for _, col := range e.Schema().Columns {
		fmt.Printf("output:%s TblName:%s DBName:%s FromID:%s %d\n", col, col.TblName, col.DBName, col.FromID, col.Position)
	}
*/

// fmt.Printf("expression: %s %+v\n", expr.String(), reflect.TypeOf(expr))

/*
		sqlText := fmt.Sprintf(`
	        function(%s)
	          return %s
	        end
	    `, inputParams, outputParams)
*/

// println(sqlText)

// ret := d.Map("map", sqlText)

var re = regexp.MustCompile(`([a-z]+\w*\.)(\w+)`)

func removeTableName(sqlText string) string { _ = "STUB: not implemented"; return "" }
