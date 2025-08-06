package storage

import (
	"database/sql"
	"fmt"
	"log"
	"todoList/models"

	"golang.org/x/crypto/bcrypt"
)

// TASKS
func GetAll(db *sql.DB, id int) []models.Task {
	rows, err := db.Query(`SELECT id, user_id, title, content FROM notes WHERE user_id=$1`, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var Tasks []models.Task
	for rows.Next() {
		var tsk models.Task
		err := rows.Scan(&tsk.Id, &tsk.UserID, &tsk.Title, &tsk.Content)
		if err != nil {
			log.Println(err)
		}
		Tasks = append(Tasks, tsk)
	}
	return Tasks
}

func AddTask(db *sql.DB, t models.Task) int {
	var id int
	err := db.QueryRow(`INSERT INTO notes (user_id, title, content) VALUES ($1,$2,$3) RETURNING id`, t.UserID, t.Title, t.Content).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id
}

func UpdateTask(db *sql.DB, t models.Task) {
	fmt.Println(t)
	err := db.QueryRow(`UPDATE notes
		SET title = $1, content = $2
		WHERE id = $3`, t.Title, t.Content, t.Id)
	if err != nil {
		log.Println(err)
	}
}

func DeleteTask(db *sql.DB, noteID int) {
	err := db.QueryRow(`DELETE FROM notes WHERE id = $1`, noteID)
	if err != nil {
		log.Println(err)
	}
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
