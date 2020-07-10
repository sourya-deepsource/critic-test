package checker_test

import (
	"fmt"
	"time"
)

func f() {
	_ = fmt.Sprintf("%d %d %d", 1, 2, time.Now().Nanosecond)
}
