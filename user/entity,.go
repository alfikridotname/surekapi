package user

import "time"

type Tabler interface {
	TableName() string
}

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	IsLock    bool      `json:"is_lock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy int       `json:"created_by"`
}

func (User) TableName() string {
	return "_users"
}
