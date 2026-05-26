package util

import (
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, httpStatus int, obj string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func Json(w http.ResponseWriter, r *http.Request, httpStatus int, obj interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}
