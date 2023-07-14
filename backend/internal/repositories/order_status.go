package repository

type OrderStatusRepositoryI interface {
	GetStatusByID(int) (*string, error)
}
