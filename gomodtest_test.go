package gomodtest

import "testing"

func TestHelloWorld(t *testing.T) {
	res := HelloWord()
	if res != "Hello gomodtest" {
		t.Errorf(`fail to HelloWord %v`, res)
	}
}
