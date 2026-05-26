package flow

import (
	"io"

	"github.com/chrislusf/gleam/util"
)

// Output concurrently collects outputs from previous step to the driver.
func (d *Dataset) Output(f func(io.Reader) error) *Dataset { _ = "STUB: not implemented"; return nil }

// Fprintf formats using the format for each row and writes to writer.
func (d *Dataset) Fprintf(writer io.Writer, format string) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// Fprintlnf add "\n" at the end of each format
func (d *Dataset) Fprintlnf(writer io.Writer, format string) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// Printf prints to os.Stdout in the specified format
func (d *Dataset) Printf(format string) *Dataset { _ = "STUB: not implemented"; return nil }

// Printlnf prints to os.Stdout in the specified format,
// adding an "\n" at the end of each format
func (d *Dataset) Printlnf(format string) *Dataset { _ = "STUB: not implemented"; return nil }

// SaveFirstRowTo saves the first row's values into the operands.
func (d *Dataset) SaveFirstRowTo(decodedObjects ...interface{}) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) OutputRow(f func(*util.Row) error) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func setValueTo(src, dst interface{}) error { _ = "STUB: not implemented"; return nil }
