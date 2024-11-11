package utils

import (
	"fmt"
	"strconv"
)

func ConvertStringsToInts(strings []string) ([]int, error) {
    ints := make([]int, len(strings)) // Create a slice of ints with the same length as the input slice

    for i, s := range strings {
        num, err := strconv.Atoi(s)
        if err != nil {
            return nil, fmt.Errorf("failed to convert %q to int: %w", s, err)
        }
        ints[i] = num
    }

    return ints, nil
}
