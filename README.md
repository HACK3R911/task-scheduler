#Планировщик задач
Ограничение времени	30 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Напишите функцию CreateTaskSchedule(tasks []string, tasksPerDay int) [][]string на языке Go, которая помогает распределить задачи по дням недели.

Функция принимает список задач tasks и количество задач на день tasksPerDay. Необходимо создать расписание, где каждая строка представляет один день. Если в последнем дне не хватает задач, они должны быть заполнены пустыми строками.

Формат ввода
tasks := []string{"Совещание", "Отчет", "Письма", "Созвон", "Проект", "Встреча", "Презентация"}
tasksPerDay := 3

result := CreateTaskSchedule(tasks, tasksPerDay)
Формат вывода
результат должен быть:

[
    ["Совещание", "Отчет", "Письма"],       // Первый день
    ["Созвон", "Проект", "Встреча"],        // Второй день
    ["Презентация", "", ""]                 // Третий день (неполный)
]
Примечания
Все импорты и package main писать нужно! Функцию main писать не нужно Если список задач пустой или tasksPerDay <= 0, вернуть пустой массив [][]string{}