package change

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetChange(t *testing.T) {
	tests := []struct {
		name          string
		in            int
		denominations []int
		out           map[int]int
		err           error
	}{
		{"$0.00", 0, []int{10, 25, 1}, map[int]int{25: 0, 10: 0, 1: 0}, nil},
		{"$0.97", 97, []int{10, 25, 5, 1}, map[int]int{25: 3, 10: 2, 5: 0, 1: 2}, nil},
		{"$1.18", 118, []int{10, 25, 5, 1}, map[int]int{25: 4, 10: 1, 5: 1, 1: 3}, nil},
		{"$0.50", 50, []int{1}, map[int]int{1: 50}, nil},
		{"$0.05 - expect err", 5, []int{10}, nil, fmt.Errorf("Not enough change")},
		{"$0.05 - expect err", 5, []int{}, nil, fmt.Errorf("No change to give out")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetChange(tt.in, tt.denominations)
			if !reflect.DeepEqual(tt.err, err) {
				t.Errorf("Failed test: %v.\n Got: %v, want: %v\n", tt.name, err, tt.err)
			}
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("Failed test: %v.\n Got: %v, want: %v\n", tt.name, result, tt.out)
			}
		})
	}
}
