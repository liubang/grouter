package grouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRule(t *testing.T) {
	rule := NewRule("POST", "/liubang/@name:([a-z]+)/@age:([0-9])", nil)
	keys := rule.ParamKeys
	assert.Equal(t, 2, len(keys))
	assert.Equal(t, "name", keys[0])
	assert.Equal(t, "age", keys[1])
}
