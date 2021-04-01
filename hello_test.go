package gomodtest

import "testing"

func TestHelloWorld(t *testing.T) {
	res := HelloWord()
	if res != "Hello go_mod_test" {
		t.Errorf(`fail to HelloWord %v`, res)
	}
}
