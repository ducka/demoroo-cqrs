package cqrs

type Repository[TEntity any] interface {
	GetAll() []TEntity
	GetById(id string) TEntity
	Commit(entity TEntity)
}
