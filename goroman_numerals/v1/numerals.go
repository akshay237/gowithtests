package v1

import "strings"

func convertToRoman(num int) string {

	var result strings.Builder

	// for i := num; i > 0; i-- {
	// 	if i == 5 {
	// 		result.WriteString("V")
	// 		break
	// 	}
	// 	if i == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}
	// 	result.WriteString("I")
	// }

	for num > 0 {
		switch {
		case num > 4:
			result.WriteString("V")
			num -= 5
		case num > 3:
			result.WriteString("IV")
			num -= 4
		default:
			result.WriteString("I")
			num -= 1
		}
	}

	return result.String()
}
