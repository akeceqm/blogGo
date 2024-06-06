package middlewares

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString() string {
	var textScan string
	reader := bufio.NewReader(os.Stdin)

	fmt.Scan(&textScan)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		return err.Error()
	}
	result := strings.ReplaceAll(textScan+" "+text, "\r\n", "")
	result = strings.ReplaceAll(result, "\n", "")
	return result
}
