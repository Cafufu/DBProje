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

type Bill struct {
	UserId   int    `json:"userId"`
	TypeName string `json:"billType"`
	BillName string `json:"billName"`
	Year     string `json:"year"`
	Month    string `json:"month"`
	Amount   string `json:"amount"`
}

type BillInfo struct {
	UserId int `json:"userId"`
	TypeId int `json:"billtype"`
}
