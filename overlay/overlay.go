package overlay

import (
	"errors"
	"fmt"
	"strings"

	"github.com/suryaaprakassh/comprosaur/shared"
	ansi "github.com/charmbracelet/x/ansi"
)

var OutOfBoundsError error = errors.New("Overlay out of bounds")

// returns the width and height of the string
// returns format - lines, width
func getLines(content string) ([]string, int) {
	var lines = strings.Split(content, "\n")
	
	var widest int = 0

	for _,line := range lines {
		widest = max(widest,ansi.StringWidth(line))
	}

	return lines,widest
}

// takes in the forground and background model and the position of the
// foreground model in the background model
// returns a error if something breaks
func Overlay(ctx shared.CTX, foreground, background string, pos_x, pos_y int) (string, error) {
	fLines, fWidth:= getLines(foreground)
	bLines, bWidth := getLines(background)

	fHeight := len(fLines)
	bHeight := len(bLines)

	if pos_x+fWidth > bWidth || pos_y+fHeight > bHeight {
		fmt.Println(fWidth, pos_x, pos_x+fWidth)
		fmt.Println(fHeight, pos_y, pos_y+fHeight)
		return "", OutOfBoundsError
	}

	var builder strings.Builder

	modelIdx := 0

	for idx, line := range bLines {
		if idx < pos_y || idx > pos_y+fHeight-1 {
			builder.WriteString(line)
			builder.WriteByte(byte('\n'))
			continue
		}
		//left
			left := line[:pos_x]
			builder.WriteString(left)
		//center
			builder.WriteString(fLines[modelIdx])
			modelIdx++
		//right
			right := line[pos_x+fWidth:]
			builder.WriteString(right)
			builder.WriteByte(byte('\n'))
	}

	return builder.String(), nil
}
