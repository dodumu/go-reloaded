package helper

import "strings"

func EditText(s string) string {
	var result strings.Builder
	inQuote := false

	for i := 0; i < len(s); i++ {
		char := s[i]

		
		if char == '\'' {
			if !inQuote {
				inQuote = true
				result.WriteByte(char)

				if i+1 < len(s) && s[i+1] == ' ' {
					i++
				}
			} else {
				inQuote = false

				current := result.String()
				if len(current) > 0 && current[len(current)-1] == ' ' {
					result.Reset()
					result.WriteString(current[:len(current)-1])
				}
				result.WriteByte(char)
			}
		} else {
			result.WriteByte(char)
		}
	}
	return result.String()
}
