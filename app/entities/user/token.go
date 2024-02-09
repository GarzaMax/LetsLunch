package user

type AuthToken struct {
	UserName string
	Password string
}

func (a *AuthToken) Claims() map[string]interface{} {
	return map[string]interface{}{
		"userName": a.UserName,
		"password": a.Password,
	}
}
