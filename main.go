package main

import (
	"todoList/handlers"

	"github.com/gin-gonic/gin"
)

var ID int

func main() {
	r := gin.Default()
	r.GET("/todos", handlers.GetTasks)
	r.POST("/todos", handlers.CreateTask)
	r.PUT("/todos/:id", handlers.UpdateTask)
	r.DELETE("/todos/:id", handlers.DeleteTask)

	r.Run()
	// ID = todoNote.CheckLastId() + 1
	// doAll()
}

// func viewNotes() {
// 	for _, tk := range todoNote.ReadJson() {
// 		fmt.Print(tk.ToString())
// 	}
// }
// func doAll() {
// 	for {
// 		fmt.Println("Введите команду: д - добавить, п - посмотреть, у - удалить, в - выйти")
// 		var cmnd string
// 		fmt.Scanln(&cmnd)
// 		switch cmnd {
// 		case "д":
// 			fmt.Println("Введите название заметки")
// 			title, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// 			todoNote.SetNote(&ID, title)
// 		case "п":
// 			viewNotes()
// 		case "у":
// 			fmt.Println("Введите ID заметки, которую хотите удалить")
// 			var toDelist int
// 			fmt.Scanln(&toDelist)
// 			todoNote.GetDone(toDelist, &todoNote.Tasks)
// 		case "в":
// 			fmt.Println("Завершаю работу и сохраняю файл")
// 			todoNote.DoJson()
// 			time.Sleep(time.Second)
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Неизвестная команда. Повторите попытку")
// 		}
// 	}
// }
