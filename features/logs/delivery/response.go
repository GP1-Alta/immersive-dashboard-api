package delivery

import "time"

type LogResponse struct {
	ID         uint      `json:"id"`
	UserName   string    `json:"mentor"`
	StatusName string    `json:"status"`
	Feedback   string    `json:"feedback"`
	CreatedAt  time.Time `json:"created_at"`
	Image      string    `json:"image"`
}

type ListLogResponse []LogResponse
