package factorial

func GetFactorial(in int) int {
	result := 1
	for i := 1; i <= in; i++ {
		result *= i
	}

	return result

}
