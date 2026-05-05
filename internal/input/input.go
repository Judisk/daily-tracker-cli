package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AskInt(prompt string, min, max int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		num, err := strconv.Atoi(text)
		if err != nil || num < min || num > max {
			fmt.Println("Enter number between", min, "and", max)
			continue
		}

		return num
	}
}
