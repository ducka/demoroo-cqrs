package data

import (
	"github.com/ducka/demoroo/internal/domain/entities"
)

var store = map[string]entities.Branch{}

type BranchRepository struct{}

func NewBranchRepository() BranchRepository {
	repo := BranchRepository{}
	repo.Commit(entities.Branch{
		ID:   "1",
		Name: "KFC",
	})
	repo.Commit(entities.Branch{
		ID:   "2",
		Name: "Mc Donalds",
	})
	repo.Commit(entities.Branch{
		ID:   "3",
		Name: "Sainsbury",
	})
	repo.Commit(entities.Branch{
		ID:   "4",
		Name: "Tesco",
	})
	repo.Commit(entities.Branch{
		ID:   "5",
		Name: "Waitroes",
	})
	repo.Commit(entities.Branch{
		ID:   "6",
		Name: "Woolworths",
	})
	repo.Commit(entities.Branch{
		ID:   "7",
		Name: "Woolworths",
	})
	repo.Commit(entities.Branch{
		ID:   "8",
		Name: "Kaki Katsu",
	})
	return repo
}

func (t BranchRepository) GetAll() []entities.Branch {
	items := make([]entities.Branch, 0, len(store))
	for _, v := range store {
		items = append(items, v)
	}
	return items
}

func (t BranchRepository) GetById(id string) entities.Branch {
	return store[id]
}

func (t BranchRepository) Commit(branch entities.Branch) {
	store[branch.ID] = branch
}
