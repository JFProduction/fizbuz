package fizbuz

import (
	"bytes"
)

func fizbuz(i int) string {
	var str bytes.Buffer
	if i%3 == 0 {
		str.WriteString("fizz")
	}

	if i%5 == 0 {
		str.WriteString("buzz")
	}

	return str.String()
}
