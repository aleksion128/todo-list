package storage

import "todoList/models"

var tasks = make(map[int]models.Task)

func GetAll() []models.Task {
	res := make([]models.Task, 0, len(tasks))
	for _, v := range tasks {
		res = append(res, v)
	}
	return res
}

func AddTask(t models.Task) {
	tasks[t.Id] = t
}

func ChangeTask(t models.Task) {
	tasks[t.Id] = t
}

func DeleteTask(id int) {
	delete(tasks, id)
}
