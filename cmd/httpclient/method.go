package main

import "net/http"

type Method string

const (
	MethodGet     Method = http.MethodGet
	MethodPost    Method = http.MethodPost
	MethodPut     Method = http.MethodPut
	MethodDelete  Method = http.MethodDelete
	MethodPatch   Method = http.MethodPatch
	MethodHead    Method = http.MethodHead
	MethodOptions Method = http.MethodOptions
)

func (m Method) String() string {
	return string(m)
}
