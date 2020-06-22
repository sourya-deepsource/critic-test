package checker_test

func warnings1() {
	s := "hello"

	// case: edit fits inside another edit
	if len(s[:]) == 0 {
		// do nothing
	}

	// case: edits partially overlap -> Issue 1: (1, 10) Issue 2: (9, 13)
	// TODO: Add example of partial overlap
}
