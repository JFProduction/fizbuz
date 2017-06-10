package fizbuz

import "testing"

func TestFizBuz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, ""},
		{3, "fiz"},
		{5, "buz"},
		{15, "fiz buz"},
	}

	for _, c := range cases {
		got := fizbuz(c.in)
		if got != c.want {
			t.Errorf("FizBuz(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
