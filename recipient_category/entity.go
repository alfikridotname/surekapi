package recipient_category

type RecipientCategory struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
