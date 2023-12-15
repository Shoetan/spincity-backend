package types

type LoginPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type ResponseUser struct {
	ID string `json:"id"`
	Name string  `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt string `json:"delete_at"`
}