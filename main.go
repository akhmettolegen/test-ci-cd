// main.go
package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("Hello, CI/CD!")
}

func TestMain(t *testing.T) {
	expected := "Hello, CI/CD!"
	if expected != "Hello, CI/CD!" {
		t.Errorf("Expected %s but got %s", expected, "Hello, CI/CD!")
	}
}
