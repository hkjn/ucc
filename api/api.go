// Package api defines the ucc API.
package api // import "hkjn.me/ucc/api"

import (
	"crypto/rand"
	"fmt"
)

type (
	Uuid []byte
	// Inputs defines the inputs for the computation, e.g. input data,
	// source code, or context for the environment of the computation.
	Inputs interface {
		// Id returns a unique id for the input.
		Id() Uuid
	}
	// Output defines the output produced by the computation.
	Output interface {
		// Id returns a unique id for the output.
		Id() Uuid
	}
	// Computation defines a ucc computation.
	Computation interface {
		// Get retrieves data from the computation.
		//
		// Get retrieves the result of the input from cache if it exists,
		// from peers otherwise, computing it if needed.
		Get(Inputs) (*Output, error)
	}
)

// NewUuid returns a random Uuid value.
func NewUuid() (*Uuid, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("couldn't get random bytes for Uuid: %v", err)
	}
	uuid := Uuid([]byte(fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])))
	return &uuid, nil
}
