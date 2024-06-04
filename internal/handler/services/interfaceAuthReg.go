package servicesH

type AuthHandler interface {
	AuthorizationUser() error
}

type RegistrationHandler interface {
	RegistrationUser() error
}

type MyProfileHandler interface {
	MyProfile() error
}
