package RED

import (
	"fmt"
	"os"
	"os/exec"
	"runtime" // Add this line to import the strconv package
	"strings"

	"github.com/fatih/color"
	tsize "github.com/kopoli/go-terminal-size"
)

// Ajoute un texte personalisé
func DisplayText(passedParams DisplayTextOptions) {
	Params := DisplayTextOptions{
		TextToPrint: "Text",
		IsCentered:  false,
		Offset:      0,
		FgColor:     color.FgWhite,
		BgColor:     color.BgBlack,
		Bold:        false,
		Underline:   false,
	}

	col := color.New(Params.FgColor, Params.BgColor)

	if passedParams.TextToPrint != "" {
		Params.TextToPrint = passedParams.TextToPrint
	}
	Params.IsCentered = passedParams.IsCentered
	if passedParams.Offset != 0 {
		Params.Offset = passedParams.Offset
	}
	if passedParams.FgColor != 0 {
		Params.FgColor = passedParams.FgColor
		col.Add(passedParams.FgColor)
	}

	if passedParams.BgColor != 0 {
		Params.BgColor = passedParams.BgColor
		col.Add(passedParams.BgColor)
	}
	Params.Bold = passedParams.Bold
	Params.Underline = passedParams.Underline

	if Params.Bold {
		col.Add(color.Bold)
	}

	if Params.Underline {
		col.Add(color.Underline)
	}

	var width int

	width, _ = SizeTest()
	if Params.IsCentered {
		col.Print(strings.Repeat(" ", (width-len(Params.TextToPrint)+Params.Offset)/2))
	}
	col.Println(Params.TextToPrint)
}

// Ajoute un titre
func DisplayTitle(textToPrint string) {
	var width int
	width, _ = SizeTest()
	DisplayLine()
	fmt.Print(strings.Repeat(" ", (width-len(textToPrint))/2))
	fmt.Println(textToPrint)
	DisplayLine()
}

// Ajoute une ligne de -
func DisplayLine() {
	var width int
	width, _ = SizeTest()
	fmt.Println(strings.Repeat("-", width))
}

// Saute une ligne
func NewLine(x int) {
	fmt.Print(strings.Repeat("\n", x))
}

func SizeTest() (Width int, Height int) {
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

func BoxStrings(lines []string) {
	// Trouver la longueur de la plus longue chaîne
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	// Créer la bordure du haut et du bas
	topBottomBorder := "+" + strings.Repeat("-", maxLength+2) + "+"

	// Afficher la bordure du haut
	fmt.Println(topBottomBorder)

	// Afficher chaque ligne entourée de barres verticales et ajustée à la longueur maximale
	for _, line := range lines {
		fmt.Printf("| %-*s |\n", maxLength, line) // %-*s permet d'aligner les chaînes à gauche
	}

	// Afficher la bordure du bas
	fmt.Println(topBottomBorder)
}
