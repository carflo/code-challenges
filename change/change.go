package change

import (
	"fmt"
	"sort"
)

// GetChange returns the amount of each coin denomination given the
// amount of change needed and the coin denominations available
func GetChange(amount int, denominations []int) (map[int]int, error) {
	// Sort the denominations you have on hand
	sort.Sort(sort.Reverse(sort.IntSlice(denominations)))

	// Check for edge cases
	if len(denominations) == 0 {
		return nil, fmt.Errorf("No change to give out")
	}

	if amount != 0 && denominations[len(denominations)-1] > amount {
		return nil, fmt.Errorf("Not enough change")
	}

	// Get change!
	result := make(map[int]int)
	for _, denom := range denominations {
		result[denom] = 0
	}
	remainder := amount
	for _, denom := range denominations {
		amountOfCoinDenom := remainder / denom
		result[denom] = amountOfCoinDenom
		remainder = remainder - (amountOfCoinDenom * denom)
		if remainder == 0 {
			break
		}
	}

	return result, nil
}
