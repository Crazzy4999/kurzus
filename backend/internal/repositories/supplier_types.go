package repository

type SupplierTypesRepositoryI interface {
	GetTypeByID(int) (*string, error)
}
