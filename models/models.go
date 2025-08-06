package models

type Task struct {
	Id      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatedUsernameUser struct {
	Username string `json:"username"`
}

type UpdatedEmailUser struct {
	Email string `json:"email"`
}
