package Grouter

import "testing"

func TestNewRule(t *testing.T) {
	rule := NewRule("POST", "/liubang/@name:([a-z]+)/@age:([0-9])", nil)
	rule = rule
}
