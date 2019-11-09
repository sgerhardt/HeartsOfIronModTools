package Parse

import "strings"

func For(label string, data string) (string, int) {
	start := strings.Index(data, label)
	if start == -1 {
		return "", -1
	}
	labeledData := data[start:]

	parenStack := []string{}
	for i, char := range labeledData {
		if string(char) == "{" {
			parenStack = append(parenStack, "{")
			continue
		}
		if string(char) == "}" {
			n := len(parenStack) - 1
			if n <= 0 {
				// reached end of label
				return labeledData[:i+1], start + i
			}
			parenStack = parenStack[:n]
		}
	}
	return "", -1
}
