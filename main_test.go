package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// To run the tests you need to use the command go test
// If you want a more detailed and verbose result you need to use go test -v
func TestAddSuccess(t *testing.T) {
	result := Add(20, 2)
	expect := 22

	if result != expect {
		t.Errorf("Expecting %d, Received %d", expect, result)
	}
}

func TestAddSuccessTestify(t *testing.T) {
	c := require.New(t)

	result := Add(20, 2)
	expect := 22

	c.Equal(expect, result)
}
