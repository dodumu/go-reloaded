package helper

import (
	"strconv"
	"strings"
)

func ConvertCases(s string) string {
	text := strings.Fields(s)

	for i := 0; i < len(text); i++ {

		if strings.HasPrefix(text[i], "(") {

			// Join possible split modifier
			modifier := text[i]

			// If it doesn't end with ')', combine with next word
			if !strings.HasSuffix(modifier, ")") && i+1 < len(text) {
				modifier = modifier + " " + text[i+1]
			}

			// Remove parentheses
			modifier = strings.Trim(modifier, "()")

			parts := strings.Split(modifier, ",")
			command := strings.TrimSpace(parts[0])

			count := 1
			if len(parts) == 2 {
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					count = n
				}
			}

			
			// Apply to previous words
			for j := 1; j <= count && i-j >= 0; j++ {
				switch command {
				case "up":
					text[i-j] = strings.ToUpper(text[i-j])
				case "low":
					text[i-j] = strings.ToLower(text[i-j])
				case "cap":
					word := text[i-j]
					if len(word) > 0 {
						text[i-j] = strings.ToUpper(word[:1]) + word[1:]
					}
				}
			}

			// Remove modifier words
			if strings.Contains(text[i], ",") {
				text = append(text[:i], text[i+2:]...)
			} else {
				text = append(text[:i], text[i+1:]...)
			}

			i--
		}
	}

	return strings.Join(text, " ")
}
