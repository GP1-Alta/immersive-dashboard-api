package delivery

type MenteeResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
	Status   string `json:"status"`
	ClassID  uint   `json:"class_id"`
	Category string `json:"category"`
}
