package delivery

type RegisterReq struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Team     string `json:"team" form:"team"`
	Status   string `json:"status" form:"status"`
}

type LoginReq struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateReq struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Team     string `json:"team" form:"team"`
	Status   string `json:"status" form:"status"`
}

type ChangePass struct {
	OldPass string `json:"old_pass" form:"old_pass"`
	NewPass string `json:"new_pass" form:"new_pass"`
}