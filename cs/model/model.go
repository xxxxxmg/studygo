package model

type EmailLogin struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

type EmailRegist struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}
