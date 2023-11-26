package user

type CreateReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   int    `json:"gender" binding:"required"`
}

type UpdateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   int    `json:"gender"`
}

type Detail struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Gender   int    `json:"gender"`
}
