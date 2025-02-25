package services

type HashService interface {
	HashPassword(string) (string, error)
	ComparePassword(hashPassword, password string) bool
}
