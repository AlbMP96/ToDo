package urls

const (
	BaseUrl       = "/api"
	AuthUrl       = "/auth"
	UsersUrl      = BaseUrl + "/users"
	CreateUserUrl = AuthUrl + "/register"
	LoginUrl      = AuthUrl + "/login"
	LogoutUrl     = AuthUrl + "/logout"
	MeUrl         = AuthUrl + "/me"
)
