package models

import "fmt"

type RequestError struct {
	StatusCode int    `json:"code"`
	Err        string `json:"error"`
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
