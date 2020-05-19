package main

import (
	"fmt"

	"github.com/losev-al/GoLessons/pkg/strategy"
)

func main() {
	list := strategy.LazySortedStudentList{}
	list.Add(strategy.Student{Name: "Petrov", AverageScore: 4})
	list.Add(strategy.Student{Name: "Sidorov", AverageScore: 3})
	list.Add(strategy.Student{Name: "Ivanov", AverageScore: 5})
	fmt.Println("--- in add order ---")
	list.PrintFullInfo()
	fmt.Println("--- by name order ---")
	list.SetSort(strategy.ByNameStrategy{})
	list.PrintFullInfo()
	fmt.Println("--- by average score order ---")
	list.SetSort(strategy.ByAverageScoreStrategy{})
	list.PrintFullInfo()
	fmt.Println("--- after add and not change strategy ---")
	list.Add(strategy.Student{Name: "Fedorov", AverageScore: 4.5})
	list.PrintFullInfo()
	fmt.Println("--- after add and set nil strategy (added student must be last) ---")
	list.Add(strategy.Student{Name: "Alexandrov", AverageScore: 3.5})
	list.SetSort(nil)
	list.PrintFullInfo()
}
