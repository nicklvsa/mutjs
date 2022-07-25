package main

import (
	"bufio"
	"os"
	"strings"
)

func readFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func writeFile(filePath string, contents []string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	for _, data := range contents {
		writer.WriteString(data + "\n")
	}

	writer.Flush()
	f.Close()

	return nil
}

func peekable(index int, slice []string) string {
	if index < len(slice) {
		return slice[index]
	}

	return ""
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines, err := readFile("example.gjs")
	handleErr(err)

	for i, line := range lines {
		directives := strings.Split(line, " ")
		for j, directive := range directives {
			directive = strings.TrimSpace(directive)

			if strings.HasPrefix(directive, "//") {
				break
			}

			switch directive {
			case "let":
				poke := peekable(j+1, directives)
				if poke != "mut" {
					directives[j] = "const"
				} else {
					directives[j+1] = ""
				}
			case "var":
				poke := peekable(j+1, directives)
				if poke != "mut" {
					directives[j] = "const"
				} else {
					directives[j+1] = ""
				}
			case "const":
				continue
			}
		}

		lines[i] = strings.TrimSpace(strings.Join(directives, " "))
	}

	handleErr(writeFile("out.js", lines))
}
