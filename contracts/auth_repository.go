package contracts

type AuthRepository interface {
	GetItems() ([]string, error)
}
