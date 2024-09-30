package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type KEYVALUE struct {
	TypeKey   string
	TypeValue string
}

func ScanENVFile(filePath string) []KEYVALUE {

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var secrets []KEYVALUE

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "=")

		if len(parts) != 2 {
			fmt.Println("Invalid Format in line: ", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.TrimLeft(parts[1], " ")

		secrets = append(secrets, KEYVALUE{key, value})
	}

	return secrets
}
