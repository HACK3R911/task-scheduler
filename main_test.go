package main

import (
	"reflect"
	"testing"
)

func TestCreateTaskSchedule(t *testing.T) {
	tests := []struct {
		name        string
		tasks       []string
		tasksPerDay int
		want        [][]string
	}{
		{
			name:        "Basic test",
			tasks:       []string{"Совещание", "Отчет", "Письма", "Созвон", "Проект", "Встреча", "Презентация"},
			tasksPerDay: 3,
			want: [][]string{
				{"Совещание", "Отчет", "Письма"},
				{"Созвон", "Проект", "Встреча"},
				{"Презентация", "", ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateTaskSchedule(tt.tasks, tt.tasksPerDay)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTaskSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
