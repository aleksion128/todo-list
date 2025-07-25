package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"todoList/todoNote"
)

var ID int

func main() {
	ID = todoNote.CheckLastId() + 1
	doAll()
}

func viewNotes() {
	for _, tk := range todoNote.ReadJson() {
		fmt.Print(tk.ToString())
	}
}
func doAll() {
	for {
		fmt.Println("Введите команду: д - добавить, п - посмотреть, у - удалить, в - выйти")
		var cmnd string
		fmt.Scanln(&cmnd)
		switch cmnd {
		case "д":
			fmt.Println("Введите название заметки")
			title, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			todoNote.SetNote(&ID, title)
		case "п":
			viewNotes()
		case "у":
			fmt.Println("Введите ID заметки, которую хотите удалить")
			var toDelist int
			fmt.Scanln(&toDelist)
			todoNote.GetDone(toDelist, &todoNote.Tasks)
		case "в":
			fmt.Println("Завершаю работу и сохраняю файл")
			todoNote.DoJson()
			time.Sleep(time.Second)
			os.Exit(0)
		default:
			fmt.Println("Неизвестная команда. Повторите попытку")
		}
	}
}
