package strategy

import "sort"

type byAverageScore []Student

func (s byAverageScore) Len() int           { return len(s) }
func (s byAverageScore) Less(i, j int) bool { return s[i].AverageScore < s[j].AverageScore }
func (s byAverageScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// ByAverageScoreStrategy describe strategy for sort by name
type ByAverageScoreStrategy struct{}

// Sort sort by Average Score
func (b ByAverageScoreStrategy) Sort(s []Student) {
	sort.Sort(byAverageScore(s))
}
