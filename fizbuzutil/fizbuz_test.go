package fizbuz

import "testing"

func TestFizBuz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, ""},
		{3, "fizz"},
		{5, "buzz"},
		{15, "fizzbuzz"},
	}

	for _, c := range cases {
		got := fizbuz(c.in)
		if got != c.want {
			t.Errorf("FizBuz(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
