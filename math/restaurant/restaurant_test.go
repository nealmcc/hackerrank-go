package restaurant

import "testing"

func TestRestaurant(t *testing.T) {
	tt := []struct {
		name string
		l, b int32
		want int32
	}{
		{
			name: "example 1",
			l:    2,
			b:    2,
			want: 1,
		},
		{
			name: "example 2",
			l:    6,
			b:    9,
			want: 6,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

		})
	}
}
