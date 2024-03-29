package user

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	IsActive  bool      `json:"is_active"`
	IsLock    bool      `json:"is_lock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
