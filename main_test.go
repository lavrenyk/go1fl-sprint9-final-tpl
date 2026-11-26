package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100} {
		data := generateRandomElements(size)
		if len(data) != max(size, 0) {
			t.Fatalf("size %d: unexpected length %d", size, len(data))
		}
		for _, value := range data {
			if value <= 0 {
				t.Fatalf("expected positive value, got %d", value)
			}
		}
	}
}

func TestMaximums(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"negative", []int{-9, -2, -5}, -2},
		{"equal", []int{4, 4, 4}, 4},
		{"first", []int{9, 2, 3, 4, 5, 6, 7, 8}, 9},
		{"middle", []int{1, 2, 3, 9, 5, 6, 7, 8}, 9},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum: got %d, want %d", got, test.want)
			}
			if got := maxChunks(test.data); got != test.want {
				t.Errorf("maxChunks: got %d, want %d", got, test.want)
			}
		})
	}
}
