package users

type Users struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Tokens struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}
