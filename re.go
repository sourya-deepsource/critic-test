package checker_test

import (
	"fmt"
	"regexp"
)

func reCheck() {
	re := *regexp.MustCompile(`.`)
	fmt.Printf("%+v\n", re.SubexpIndex("something"))
}
// la lee lo
