package interfaces

type IHashService interface {
	Hash(value string) (hashedValue string, err error)
	ValidateHash(value, providedValue string) (err error)
}
