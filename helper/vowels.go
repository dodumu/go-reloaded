package helper

import "strings"

func Replace(s string) string {
	wordcheck := strings.Fields(s)

	for i := 0; i < len(wordcheck)-1; i++ {

		switch wordcheck[i] {
		case "a":
			if i+1 < len(wordcheck) {
				nextword := wordcheck[i+1]

				if len(nextword) > 0 &&
					(nextword[0] == 'e' ||
						nextword[0] == 'i' ||
						nextword[0] == 'o' ||
						nextword[0] == 'u' ||
						nextword[0] == 'h' ||
						nextword[0] == 'a') {
					wordcheck[i] = "an"
				}
			}
		}
	}
	return strings.Join(wordcheck, " ")
}
