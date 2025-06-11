package user

type UserDB struct {
	UUID           string `json:"uuid"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashedpassword"`
}

type UserMini struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
