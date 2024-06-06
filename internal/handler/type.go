package handler

import (
	"bufio"
	"os"
)

var (
	title       string
	description string
	err         error
)

var read = bufio.NewReader(os.Stdin)
