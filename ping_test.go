package easyping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ping(t *testing.T) {
	_, err := Ping("127.0.0.1")
	assert.NoError(t, err)

	_, err = Ping("localhost")
	assert.NoError(t, err)
}
