package iteration

func Repeat(character string, number int) (repeated string) {
	if number == 0 {
		repeated = character
		return
	}

	for i := 0; i < number; i++ {
		repeated += character
	}
	return
}
