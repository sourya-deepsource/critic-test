package checker_test

func sliceArrayMultipleTimes() {
	var xs [3]int

	_ = xs[:][:]    // want `autofix required for unnecessary slice`
	_ = xs[:][:][:] // want `autofix required for unnecessary slice`
}

func dullStringSlicing() {
	var s string

	_ = s[:] // want `autofix required for unnecessary slice`

	_ = s[:][:] // want `autofix required for unnecessary slice`

	_ = s[:][:][:] // want `autofix required for unnecessary slice`
}

func dullSlicing() {
	{
		var xs []byte
		var ys []byte
		copy(xs[:], ys[:]) // want `autofix required for unnecessary slice`
	}
	{
		var xs []int
		_ = xs[:] // want `autofix required for unnecessary slice`
	}
	{
		var xs [][]int
		_ = xs[0][:] // want `autofix required for unnecessary slice`
	}
	{
		var xs []string
		_ = xs[:] // want `autofix required for unnecessary slice`
	}
	{
		var xs []struct{}
		_ = xs[:] // want `autofix required for unnecessary slice`
	}
	{
		var xs map[string][][]int
		_ = xs["0"][0][:] // want `autofix required for unnecessary slice`
	}
}
