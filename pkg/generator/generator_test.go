package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCooldownKey(t *testing.T) {
	channel := "test-channel"
	cmd := "test-cmd"
	expected := "test-channel:test-cmd"

	actual := NewCooldownKey(channel, cmd)

	assert.Equal(t, expected, actual, "generated cooldown key should match expected value")
}
