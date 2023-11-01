package user

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpOutput struct {
	Id   int        `json:"id"`
	Name string     `json:"name"`
	Info InfoOutput `json:"info"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type UpdateUserOutput struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
