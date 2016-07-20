package model

type Client struct {
	UID       *string `json:"uid" gorm:"primary_key"`
	Secret    *string `json:"secret" validate:"required"`
	CreatedAt *int64  `json:"created_at"`
	UpdatedAt *int64  `json:"updated_at"`
}

type AccessToken struct {
	Id        *int64  `gorm:"primary_key"`
	Token     *string `json:"token"`
	CreatedAt *int64  `json:"created_at"`
	UpdatedAt *int64  `json:"updated_at"`
}
