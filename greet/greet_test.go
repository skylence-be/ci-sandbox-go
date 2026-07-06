package greet

import "testing"

func TestGreeting(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"named", "skylence", "hello, skylence"},
		{"empty falls back", "", "hello, world"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := Greeting(tc.in); got != tc.want {
				t.Fatalf("Greeting(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
