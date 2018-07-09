package main

import "testing"

func DefaultTest(t *testing.T) {
	if !true {
		t.Fatal()
	}
}
