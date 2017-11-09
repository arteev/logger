package logger

import (
	"fmt"
	"testing"
)

func TestLevelNames(t *testing.T) {
	l := Level(100)
	if got := l.Name(); got != "" {
		t.Errorf("Expected %q,got %q", "", got)
	}
	want := "Error"
	if got := LevelError.Name(); got != want {
		t.Errorf("Expected %q,got %q", want, got)
	}
}

func TestLogStringer(t *testing.T) {
	want := fmt.Sprintf("%s (%d)", "Debug", int(LevelDebug))
	got := fmt.Sprintf("%s", Debug)
	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
