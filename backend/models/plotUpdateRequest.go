package models

type PlotUpdateRequest struct {
	Title      *string `json:"plot_title"`
	Story      *string `json:"plot_story"`
	Status     *bool   `json:"plot_status"`
	Updated_at string  `json:"updated_at"`
}

func NewPlotUpdateRequest(title, story *string, status *bool) *PlotUpdateRequest {
	if title != nil {
		title = title
	}

	if story != nil {
		story = story
	}

	if status != nil {
		status = status
	}

	return &PlotUpdateRequest{
		Title:      title,
		Story:      story,
		Status:     status,
		Updated_at: "",
	}
}
