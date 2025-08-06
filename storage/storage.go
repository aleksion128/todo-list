package storage

import (
	"database/sql"
	"fmt"
	"log"
	"todoList/models"

	"golang.org/x/crypto/bcrypt"
)

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

// TASKS
func CreateNote(db *sql.DB, userID int, title, content string) int {
	var id int
	err := db.QueryRow(`INSERT INTO notes (user_id, title, content)
		VALUES ($1, $2, $3)
		RETURNING id`, userID, title, content).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id
}

func GetUserNotes(db *sql.DB, userID int) []models.Task {
	rows, err := db.Query(`SELECT id, title, content 
		FROM notes 
		WHERE user_id = $1`, userID)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var Tasks []models.Task
	for rows.Next() {
		var tsk models.Task
		err := rows.Scan(&tsk.Id, &tsk.Title, &tsk.Content)
		tsk.UserID = userID
		if err != nil {
			log.Fatal(err)
		}
		Tasks = append(Tasks, tsk)
	}
	return Tasks
}

func UpdateNote(db *sql.DB, noteID int, title, content string) {
	_, err := db.Exec(`UPDATE notes
		SET title = $1, content = $2, updated_at = NOW()
		WHERE id = $3`, title, content, noteID)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteNote(db *sql.DB, noteID int) {
	_, err := db.Exec(`DELETE FROM notes WHERE id = $1`, noteID)
	if err != nil {
		log.Fatal(err)
	}
}

func ChangeTask(t models.Task) {
	tasks[t.Id] = t
}

func DeleteTask(id int) {
	delete(tasks, id)
}

// USER
func HashPass(pass string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(password)
}

func CreateUser(db *sql.DB, username, email, password string) int {
	var id int
	err := db.QueryRow(`INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`, username, email, HashPass(password)).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func DeleteUser(db *sql.DB, id int) {
	err := db.QueryRow(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		fmt.Println("Проблемы с удалением пользователя из базы данных")
	}
}

func UpdateUsername(db *sql.DB, id int, username string) {
	err := db.QueryRow(`UPDATE users SET username = $1 WHERE id = $2`, username, id)
	if err != nil {
		fmt.Println("Проблемы с обновлением пользователя в базе данных")
	}
}

func UpdateEmail(db *sql.DB, id int, email string) {
	err := db.QueryRow(`UPDATE users SET email=$1 WHERE id = $2`, email, id)
	if err != nil {
		fmt.Println("Проблемы с обновлением пользователя в базе данных")
	}
}

func UpdatePassword(db *sql.DB, id int, password string) {
	err := db.QueryRow(`UPDATE users SET password_hash=$1 WHERE id = $2`, HashPass(password), id)
	if err != nil {
		fmt.Println("Проблемы с обновлением пользователя в базе данных")
	}
}
