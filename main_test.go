package main

import "testing"

func TestDefaultText(t *testing.T) {
	got := defaultText()
	need := "Hello World"
	if got != need {
		t.Errorf("defaultText() = %s; want %s", got, need)
	}
}

func TestCustomText(t *testing.T) {
	got := customText("Coleton")
	need := "Hello, Coleton"
	if got != need {
		t.Errorf("defaultText() = %s; want %s", got, need)
	}
}