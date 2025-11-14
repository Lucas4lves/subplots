package models

type SubPlot struct {
	ID        int    `json:"subplot_id"`
	Title     string `json:"subplot_title"`
	Content   string `json:"subplot_content"`
	Status    string `json:"subplot_status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewSubplot(id int, title, content, status, createdAt, updatedAt string) *SubPlot {
	return &SubPlot{
		ID:        id,
		Title:     title,
		Content:   content,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
