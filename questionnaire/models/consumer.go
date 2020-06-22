package models

type Consumer struct {
	Id              uint   `json:"Id"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	TelephoneNumber string `json:"telephoneNumber"`
}

func UserList() Consumer {
	user := Consumer{Id: 1, Name: "张三", Gender: "男", TelephoneNumber: "12345678912"}
	return user
}

func UserAdd(u Consumer) int {
	println("add user ", u.Name)
	return 1
}




