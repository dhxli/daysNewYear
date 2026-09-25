package daysleft

import (
	"testing"
	"time"
)

func TestUntil(t *testing.T) {
	cases := []struct {
		name string
		in   time.Time
		want int
	}{
		{"1 января високосного", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 366},
		{"31 декабря", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 1},
		{"29 февраля високосного", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), 307},
		{"28 февраля невисокосного", time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC), 307},
		{"1 марта невисокосного", time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), 306},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Until(tc.in)
			if got != tc.want {
				t.Errorf("Until(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}