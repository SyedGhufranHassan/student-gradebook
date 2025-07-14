package main

import (
	"bufio"
	"os"
	"strings"
)

var Reader = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	print(prompt)
	input, _ := Reader.ReadString('\n')
	return strings.TrimSpace(input)

}