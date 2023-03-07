package delivery

type AddLogReq struct {
	MenteeID uint
	UserID   uint
	StatusID uint   `json:"status_id" form:"status_id"`
	Feedback string `json:"feedback" form:"feedback"`
	Image    string `json:"image" form:"image"`
}