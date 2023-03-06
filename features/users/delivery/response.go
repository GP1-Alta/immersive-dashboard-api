package delivery

type UserResponse struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

type ListUserResponse []UserResponse