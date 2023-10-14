package user

type Request struct {
	Name string `json:"name" form:"name"`
}

type LoginReq struct {
	UserName string `json:"user_name" `
	PassWord string `json:"password" `
	Role     string `json:"role" `
}
