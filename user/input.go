package user

type RegisterUserInput struct {
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckUsernameInput struct {
	Username string `json:"username" binding:"required"`
}
