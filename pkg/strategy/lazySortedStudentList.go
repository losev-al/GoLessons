package strategy

import "fmt"

// LazySortedStudentList describe student list with support of lazy sorted
type LazySortedStudentList struct {
	students []Student
	strategy SortStrategy
	needSort bool
}

// Add adds student into list
func (l *LazySortedStudentList) Add(s Student) {
	l.students = append(l.students, s)
	l.needSort = true
}

// SetSort set strategy for sort list
func (l *LazySortedStudentList) SetSort(s SortStrategy) {
	l.strategy = s
	l.needSort = true
}

// PrintNames print list with names only
func (l *LazySortedStudentList) PrintNames() {
	l.sortIfNeeded()
	for _, s := range l.students {
		fmt.Println(s.Name)
	}
}

// PrintFullInfo print list with full student info
func (l *LazySortedStudentList) PrintFullInfo() {
	l.sortIfNeeded()
	for _, s := range l.students {
		fmt.Printf("%20s\t%4.2f\n", s.Name, s.AverageScore)
	}
}

func (l *LazySortedStudentList) sortIfNeeded() {
	if l.needSort && l.strategy != nil {
		l.strategy.Sort(l.students)
	}
}
