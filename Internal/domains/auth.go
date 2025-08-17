package domains

type AuthService interface {
	LoginService(email, password string) (string, error)
}
