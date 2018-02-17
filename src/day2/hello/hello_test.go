package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Log("TestHello")

	hello := Hello()
	assert.Equal(t, "Hello", hello)
}
