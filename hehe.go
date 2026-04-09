package hehe

func Add(a int, b int) int {
	result := a

	for i := 0; i < b; i++ {
		result = result + 1
	}

	return result
}
