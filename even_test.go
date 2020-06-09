package even

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "zero",
			n:    0,
			want: true,
		},
		{
			name: "one",
			n:    1,
			want: false,
		},
		{
			name: "three",
			n:    3,
			want: false,
		},
		{
			name: "ten",
			n:    2,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.n); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
