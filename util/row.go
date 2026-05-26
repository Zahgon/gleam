//go:generate truepack -fast-strings -o row_codec.go
package util

type Row struct {
	K []interface{} `msg:"K"`
	V []interface{} `msg:"V"`
	T int64         `msg:"T"`
}

func NewRow(timestamp int64, objects ...interface{}) *Row { _ = "STUB: not implemented"; return nil }

func (row *Row) AppendKey(objects ...interface{}) *Row { _ = "STUB: not implemented"; return nil }

func (row *Row) AppendValue(objects ...interface{}) *Row { _ = "STUB: not implemented"; return nil }

// UseKeys use the indexes[] specified fields as key fields
// and the rest of fields as value fields
func (row *Row) UseKeys(indexes []int) (err error) { _ = "STUB: not implemented"; return nil }
