package RED

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	fmt.Println(input)
}
