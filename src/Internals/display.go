package RED

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	tsize "github.com/kopoli/go-terminal-size"
)

// Ajoute un texte personalisé
func DisplayText(textToPrint string, isCentered bool, offset int) {
	var width int
	width, _ = sizeTest()
	if isCentered {
		fmt.Print(strings.Repeat(" ", (width-len(textToPrint)+offset)/2))
	}
	fmt.Println(textToPrint)
}

// Ajoute un titre
func DisplayTitle(textToPrint string) {
	var width int
	width, _ = sizeTest()
	DisplayLine()
	fmt.Print(strings.Repeat(" ", (width-len(textToPrint))/2))
	fmt.Println(textToPrint)
	DisplayLine()
}

// Ajoute une ligne de -
func DisplayLine() {
	var width int
	width, _ = sizeTest()
	fmt.Println(strings.Repeat("-", width))
}

// Saute une ligne
func NewLine(x int) {
	fmt.Print(strings.Repeat("\n", x))
}

func sizeTest() (Width int, Height int) {
	var s tsize.Size

	s, _ = tsize.GetSize()
	Width, Height = s.Width, s.Height
	return
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
