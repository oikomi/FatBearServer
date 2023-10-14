package auth

type LoginReq struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
}

type RegisterReq struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email"`
}
