package main

import "testing"

func TestCountActiveModes(t *testing.T) {
	tests := []struct {
		name  string
		modes []bool
		want  int
	}{
		{
			name:  "no active modes",
			modes: []bool{false, false, false, false},
			want:  0,
		},
		{
			name:  "one active mode",
			modes: []bool{true, false, false, false},
			want:  1,
		},
		{
			name:  "two active modes",
			modes: []bool{true, true, false, false},
			want:  2,
		},
		{
			name:  "all modes active",
			modes: []bool{true, true, true, true},
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countActiveModes(tt.modes...)
			if got != tt.want {
				t.Fatalf("countActiveModes() = %d, want %d", got, tt.want)
			}
		})
	}
}
