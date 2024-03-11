package taskbase

import "fmt"

type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// Метод получения всех задач
func List() map[string]Task {
	return tasks
}

// Метод получения задачи по ID
func ById(id string) (Task, error) {

	task, ok := tasks[id]

	if !ok {
		return Task{}, fmt.Errorf("element was not found")
	}

	return task, nil
}

// Метод добавления задачи
func Add(t Task) {

	tasks[t.ID] = t
}

// Метод удаления задачи по ID
func Delete(id string) error {

	if _, ok := tasks[id]; !ok {
		return fmt.Errorf("element was not found")
	}

	delete(tasks, id)

	return nil
}
