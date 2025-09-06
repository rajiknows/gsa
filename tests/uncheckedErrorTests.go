package tests

func getsum(a int, b int) (int, error) {
	return a + b, nil
}

func uncheckederror() {
	sum, _ := getsum(1, 2)
	println("sum %d", sum)
}
