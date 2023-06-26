package goiterations

const num = 4

func Repeat(ch string) string {
	repeat := ""
	for i := 0; i < num; i++ {
		repeat += ch
	}
	return repeat
}
