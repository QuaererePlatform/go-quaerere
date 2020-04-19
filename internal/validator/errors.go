package validator

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/hashicorp/go-multierror"
)

type (
	Error interface {
		Add(string)
		Error() string
		Errors() []error
	}

	err struct {
		errors []error
		lock *sync.RWMutex
	}
)

func NewError() Error {
	return &err{
		lock: new(sync.RWMutex),
	}
}

func (e *err) Add(s string) {
	e.lock.Lock()
	e.errors = append(e.errors, errors.New(s))
	e.lock.Unlock()
}

func (e *err) Error() string {
	var result *multierror.Error
	for _, err := range e.errors {
		result = multierror.Append(result, err)
	}

	result.ErrorFormat = func(es []error) string {
		if len(es) == 1 {
			return es[0].Error()
		}

		points := make([]string, len(es))
		for i, err := range es {
			points[i] = fmt.Sprintf("* %s", err)
		}

		return fmt.Sprintf("%d errors occured:\n\t%s\n\t", len(es), strings.Join(points, "\n\t"))
	}

	return result.Error()
}

func (e *err) Errors() []error {
	return e.errors
}