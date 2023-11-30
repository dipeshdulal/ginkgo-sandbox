package contracts

type AuthService interface {
	GetItems() ([]string, error)
}
