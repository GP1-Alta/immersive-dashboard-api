package delivery

type LogResponse struct {
	ID         uint   `json:"id"`
	UserName   string `json:"mentor"`
	StatusName string `json:"status"`
	Feedback   string `json:"feedback"`
	CreatedAt  string `json:"created_at"`
	Image      string `json:"image"`
}

type ListLogResponse []LogResponse
