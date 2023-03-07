package delivery

type AddLogReq struct {
	MenteeID uint
	UserID   uint `json:"user_id"`
	StatusID uint `json:"status_id"`
	Feedback string `json:"feedback"`
}