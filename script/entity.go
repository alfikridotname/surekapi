package script

type Script struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Template string `json:"template"`
	IsActive bool   `json:"is_active"`
}
