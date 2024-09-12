package RED

import (
	"fmt"
	"strings"

	tsize "github.com/kopoli/go-terminal-size"
)

func DisplayText(textToPrint string, center bool, offset int) {
	var width int
	width, _ = sizeTest()
	if center {
		fmt.Print(strings.Repeat(" ", (width-len(textToPrint)+offset)/2))
	}
	fmt.Println(textToPrint)
}

func DisplayTitle(textToPrint string) {
	var width int
	width, _ = sizeTest()
	DisplayLine()
	fmt.Print(strings.Repeat(" ", (width-len(textToPrint))/2))
	fmt.Println(textToPrint)
	DisplayLine()
}

func DisplayLine() {
	var width int
	width, _ = sizeTest()
	fmt.Println(strings.Repeat("-", width))
}

func sizeTest() (Width int, Height int) {
	var s tsize.Size

	s, _ = tsize.GetSize()
	Width, Height = s.Width, s.Height
	return
}
