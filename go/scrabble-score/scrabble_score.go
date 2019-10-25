package scrabble

import "unicode"

//LetterToValueMap is a map of letter to values
// var LetterToValueMap = map[rune]int{
// 	'A': 1,
// 	'E': 1,
// 	'I': 1,
// 	'O': 1,
// 	'U': 1,
// 	'L': 1,
// 	'N': 1,
// 	'R': 1,
// 	'S': 1,
// 	'T': 1,
// 	'D': 2,
// 	'G': 2,
// 	'B': 3,
// 	'C': 3,
// 	'M': 3,
// 	'P': 3,
// 	'F': 4,
// 	'H': 4,
// 	'V': 4,
// 	'W': 4,
// 	'Y': 4,
// 	'K': 5,
// 	'J': 8,
// 	'X': 8,
// 	'Q': 10,
// 	'Z': 10,
// }

//Score function compute the scrabble score for a given word
func Score(input string) int {
	score := 0
	for _, v := range input {
		v = unicode.ToUpper(v)
		switch v {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score++
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}
	return score
}
