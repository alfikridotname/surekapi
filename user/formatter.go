package user

type UserFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return formatter
}
