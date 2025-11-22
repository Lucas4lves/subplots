package models

type Plot struct {
	ID        int    `json:"plot_id"`
	Title     string `json:"plot_title"`
	Story     string `json:"plot_story"`
	Status    string `json:"plot_status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewPlot(id int, title, story, status, createdAt, updatedAt string) *Plot {
	return &Plot{
		ID:        id,
		Title:     title,
		Story:     story,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
