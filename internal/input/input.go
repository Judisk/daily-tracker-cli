package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AskInt(prompt string, min, max int) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)

		text, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		text = strings.TrimSpace(text)

		num, err := strconv.Atoi(text)
		if err != nil || num < min || num > max {
			fmt.Println("Enter number between", min, "and", max)
			continue
		}

		return num, nil
	}
}
