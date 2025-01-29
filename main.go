package main

import (
	"bufio"
	"fmt"
	"os"
)

func CreateTaskSchedule(tasks []string, tasksPerDay int) [][]string {
	if len(tasks) == 0 || tasksPerDay <= 0 {
		return [][]string{}
	}

	totalDays := (len(tasks) + tasksPerDay - 1) / tasksPerDay
	schedule := make([][]string, totalDays)

	for day := 0; day < totalDays; day++ {
		schedule[day] = make([]string, tasksPerDay)
		start := day * tasksPerDay
		end := start + tasksPerDay
		if end > len(tasks) {
			end = len(tasks)
		}
		copy(schedule[day], tasks[start:end])
	}

	return schedule
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024), 1024)

	var n int
	fmt.Scan(&n)

	tasks := make([]string, n)

	for i := 0; i < n && scanner.Scan(); i++ {
		tasks[i] = scanner.Text()
	}

	var tasksPerDay int
	fmt.Scan(&tasksPerDay)

	result := CreateTaskSchedule(tasks, tasksPerDay)
	fmt.Println(result)
}
