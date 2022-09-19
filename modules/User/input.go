package user

type RegisterUserInput struct {
	Name     string `json:"name" form:"name" binding:"required"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginInput struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type FormUpdateUserInput struct {
	ID       uint64
	Name     string `json:"name"  form:"name" binding:"required"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password"`
	Error    error
}
