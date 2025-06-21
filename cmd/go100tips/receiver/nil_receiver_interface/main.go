package main

import (
	"log"
	"strings"

	"github.com/cockroachdb/errors"
)

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var m *MultiError
	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}
	return m
}

func main() {
	customer := Customer{Age: 33, Name: "John"}
	if err := customer.Validate(); err != nil {
		log.Fatalf("customer is invalid: %v", err)
	}
}
