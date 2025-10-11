package overlay

import (
	// "fmt"
	"strings"
	// "testing"
)

func createTestModel(width, height int, fillChar string) string {
	var sb strings.Builder

	for range height {
		sb.WriteString(strings.Repeat(fillChar, width))
		sb.WriteString("\n")
	}

	return sb.String()
}


//TODO: create tests for overlay function

// func TestOverlay(t *testing.T) {
// 	forground := createTestModel(4, 6, " ")
// 	background := createTestModel(6, 6, "#")
//
// 	fmt.Print(Overlay(forground, background, 0, 1))
//
// }
