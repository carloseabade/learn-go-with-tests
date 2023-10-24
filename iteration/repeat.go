package iteration

import (
	"fmt"
	"strings"
)

func Repeat(character string, n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		fmt.Fprintf(&b, "%s", character)
	}
	return b.String()
}
