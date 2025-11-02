package libs

import (
	"fmt"

	"github.com/fatih/color"
)

// ColorPrintln prints text in the specified color
func ColorPrintln(c color.Attribute, text string) {
	color.Set(c)
	fmt.Println(text)
	color.Unset()
}

// ColorPrintf prints formatted text in the specified color
func ColorPrintf(c color.Attribute, format string, a ...interface{}) {
	color.Set(c)
	fmt.Printf(format, a...)
	color.Unset()
}

// ColorSprint returns a string formatted in the specified color
func ColorSprint(c color.Attribute, text string) string {
	color.Set(c)
	coloredText := fmt.Sprint(text)
	color.Unset()
	return coloredText
}

// ColorPrint prints text in the specified color without a newline
func ColorPrint(c color.Attribute, text string) {
	color.Set(c)
	fmt.Print(text)
	color.Unset()
}
