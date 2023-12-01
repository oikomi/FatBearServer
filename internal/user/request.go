package user

type Request struct {
	Name string `json:"name" form:"name"`
}

type LoginReq struct {
	UserName string `json:"user_name" `
	PassWord string `json:"password" `
	Role     string `json:"role" `
}

type AddTokenReq struct {
	Name  string `json:"name" example:"充值的用户名称" format:"string"`
	Token int    `json:"token" example:"1000" format:"int"`
}

type GetTokenReq struct {
	Name string `json:"name" form:"name"`
}
