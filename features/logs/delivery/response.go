package delivery

type LogResponse struct {
	ID         uint   `json:"id"`
	UserName   string `json:"mentor"`
	StatusName string `json:"name"`
	Feedback   string `json:"fedback"`
	CreatedAt  string `json:"date"`
	Image      string `json:"image"`
}

type ListLogResponse []LogResponse
