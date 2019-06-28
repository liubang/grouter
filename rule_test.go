package Grouter

import (
	"fmt"
	"testing"
)

func TestNewRule(t *testing.T) {
	rule := NewRule("POST", "/liubang/@name:([a-z]+)/@age:([0-9])", nil)
	fmt.Println(rule.ParamKeys)
}
