package main

import (
    "fmt"
    "testing"
)

func TestMainFunction(t *testing.T) {
    expected := "Hello, World!"
    result := fmt.Sprint("Hello, World!")
    if result != expected {
        t.Errorf("expected %s but got %s", expected, result)
    }
}
