package area

type CreateAreaInput struct {
	Param1   int32  `json:"param_1" form:"param_1" binding:"required"`
	Param2   int64  `json:"param_2" form:"param_2" binding:"required"`
	TypeArea string `json:"type_area" form:"type_area" binding:"required"`
}
