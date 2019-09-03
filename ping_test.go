package easyping

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Ping(t *testing.T) {
	_, err := Ping("127.0.0.1")
	assert.NoError(t, err)

	_, err = Ping("localhost")
	assert.NoError(t, err)
}

func Test_AdvancedPing(t *testing.T) {
	_, err := AdvancedPing(&Options{Address: "192.168.1.100", Timeout: time.Second * 1})
	assert.NoError(t, err)
}
