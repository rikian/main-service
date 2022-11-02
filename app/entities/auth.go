package entities

type Identity struct {
	Id       string
	JwtToken string
}

type ResponseSession struct {
	UserSession string
	RememberMe  string
}

type RequestLogin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type RequestRegister struct {
	UserName      string `json:"user_name"`
	UserEmail     string `json:"user_email"`
	UserPassword1 string `json:"user_password_1"`
	UserPassword2 string `json:"user_password_2"`
	UserTerm      bool   `json:"user_term"`
}

type ResponseLogin struct {
	Message string
	Status  string
	Code    uint32
	User    *UserAuth
}

type ResponseRegister struct {
	Message string
	Status  string
	Code    uint32
	User    *UserAuth
}

type UserAuth struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

type ErrorSession struct {
	Status  uint
	Message string
}
