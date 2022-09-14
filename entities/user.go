package entities

type User struct {
	UserId       string `json:"user_id"`
	UserPassword string `json:"user_password"`
	//Msg          string
}

type ResponseUser struct {
	Status   bool
	Messagae string
	Data     User
}
