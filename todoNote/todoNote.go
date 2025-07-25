package todoNote

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Note struct {
	ID    int
	Title string
}

var Tasks = make([]Note, 0, 10000)

func CheckLastId() int {
	var tsks []Note
	tks, err := os.ReadFile("tasks.json")
	if err != nil {
		return 0
	} else {
		json.Unmarshal(tks, &tsks)
	}
	if len(tsks) == 0 {
		return 0
	}
	return tsks[len(tsks)-1].ID
}

func ReadJson() []Note {
	var tsks []Note
	tks, err := os.ReadFile("tasks.json")
	if err != nil {
		Tasks = make([]Note, 0, 10000)
		return Tasks
	} else {
		json.Unmarshal(tks, &tsks)
		Tasks = tsks
	}
	return Tasks
}

// Создание заметки
func SetNote(idNew *int, t string) *Note {
	*idNew++
	Task := &Note{ID: *idNew, Title: t}
	Tasks = append(Tasks, *Task)
	DoJson()
	return Task
}
func DoJson() {
	jsonInf, err := json.Marshal(Tasks)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("tasks.json", jsonInf, 0666)
}
func (note Note) GetId() int {
	return note.ID
}

func (note Note) GetTitle() string {
	return note.Title
}

// Удаление выполненной заметки
func GetDone(ind int, tsks *[]Note) []Note {
	res := make([]Note, 0, len(*tsks)-1)
	for _, tk := range *tsks {
		if tk.GetId() != ind {
			res = append(res, tk)
		}
	}
	*tsks = res
	DoJson()
	return res
}

// Посмотреть все заметки
func ViewAllNotes() []Note {

	return Tasks
}
func (note *Note) ToString() string {
	return "ID: " + strconv.Itoa(note.ID) + ", Название: " + note.Title
}
