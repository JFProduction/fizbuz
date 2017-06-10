package fizbuz

import (
	"bytes"
)

func fizbuz(i int) string {
	var str bytes.Buffer
	if i%3 == 0 {
		str.WriteString("fiz")
	}

	if i%5 == 0 {
		if str.Len() == 0 {
			str.WriteString("buz")
		} else {
			str.WriteString(" buz")
		}
	}

	return str.String()
}
