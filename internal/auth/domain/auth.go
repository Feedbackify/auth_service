package domain

type AuthUseCase interface {
	Login()
	Register()
	ChangePassword()
}

type AuthRepository interface {
	StoreToken()
	RetrieveToken()
	DeleteToken()
}
