package servicesH

type UserProfileHandler interface {
	MyProfile()
}

type PasswordChanger interface {
	ChangePassword() error
}
