package utils

type Customer struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type LoginInput struct {
	UserName string `json:"userNameForLogin"`
	Password string `json:"password"`
}
