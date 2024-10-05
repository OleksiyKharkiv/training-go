package ports

import "training-go/sorting_structures/internal/domain/entity"

type Sorter interface {
	Sort(people []entity.Person) []entity.Person
}
