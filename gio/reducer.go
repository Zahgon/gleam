package gio

import (
	"context"
)

func (runner *gleamRunner) processReducer(ctx context.Context, f Reducer, keyPositions []int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (runner *gleamRunner) doProcessReducer(f Reducer) (err error) {
	_ = "STUB: not implemented"
	// get the first row
	return nil
}

// fmt.Fprintf(os.Stderr, "lastKeys:%v\n", lastKeys)

func (runner *gleamRunner) doProcessReducerByKeys(f Reducer, keyPositions []int) (err error) {
	_ = "STUB: not implemented"

	// get the first row
	return nil
}

func reduce(f Reducer, x, y []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
