package main

import (
	"testing"
)

// Test for the FindCommandPath function
func TestFindCommandPath(t *testing.T) {
  _, err := FindCommandPath("go")
  if err != nil {
    t.Errorf("Expected no error, got %v", err)
  }

  _, err = FindCommandPath("nonexistentcommand")
  if err == nil {
    t.Errorf("Expected error, got nil")
  }
}
