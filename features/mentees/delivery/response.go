package delivery

type MenteeResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Status   string `json:"status"`
	Class    string `json:"class"`
	Category string `json:"category"`
}
