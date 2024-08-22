package main

import "testing"

func TestSayHello(t *testing.T) {
	msg := sayHello("Test")

	if msg != "Hello Test!" {
		t.Errorf("sayHello() = %s; want Hello Test!", msg)
	}

}
