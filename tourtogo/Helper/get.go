package Helper

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func GetString(prompt string)string{
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	cleanedInput := strings.TrimSpace(input)

	return cleanedInput
}