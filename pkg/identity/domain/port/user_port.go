package port

type IUserPort[DTO any, U any] interface {
	GetByID(id uint) (U, error)
	GetByEmail(email string) (U, error)
	Create(user DTO) error
}
