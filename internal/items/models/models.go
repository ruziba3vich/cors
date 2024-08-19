package models

type (
	User struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	LoginUserResponse struct {
		User  *User  `json:"user"`
		Token string `json:"token"`
	}
)
