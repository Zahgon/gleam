package gio

// Emit encode and write a row of data to os.Stdout
func Emit(anyObject ...interface{}) error { _ = "STUB: not implemented"; return nil }

// TsEmit encode and write a row of data to os.Stdout
// with ts in milliseconds epoch time
func TsEmit(ts int64, anyObject ...interface{}) error { _ = "STUB: not implemented"; return nil }

func TsEmitKV(ts int64, keys, values []interface{}) error { _ = "STUB: not implemented"; return nil }
