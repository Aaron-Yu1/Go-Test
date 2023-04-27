package users

type User struct {
	ID          int64  `json:"id"`
	FirestName  string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_create"`
}
