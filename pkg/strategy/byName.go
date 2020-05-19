package strategy

import "sort"

type byName []Student

func (s byName) Len() int           { return len(s) }
func (s byName) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// ByNameStrategy describe strategy for sort by name
type ByNameStrategy struct{}

// Sort sort by Name
func (b ByNameStrategy) Sort(s []Student) {
	sort.Sort(byName(s))
}
