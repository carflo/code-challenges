package main

import "sort"

// Implements the https://golang.org/pkg/sort/ interface
type sortedMap struct {
	unsortedMap map[string]int
	keys        []string
}

func (sm *sortedMap) Len() int {
	return len(sm.unsortedMap)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.unsortedMap[sm.keys[i]] > sm.unsortedMap[sm.keys[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.keys[i], sm.keys[j] = sm.keys[j], sm.keys[i]
}

// sortedKeys returns a sorted slice of the keys in the map
// can leverage this against the unsorted map to print "in order"
func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.unsortedMap = m
	sm.keys = make([]string, len(m))
	i := 0
	for key := range m {
		sm.keys[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.keys
}
