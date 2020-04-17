## Change

Returns amount of each coin denomination given an amount of change to return and a slice of coin denominations available

## Example

Change: 50 cents
Denominations: 25, 10, 5, 1

`GetChange(50, []int{25, 10, 5, 1})` returns a map `25: 2, 10: 0, 5: 0, 1: 0`

`GetChange(97, []int{25, 10, 5, 1})` returns a map `25: 3, 10: 2, 5: 0, 1: 2`
