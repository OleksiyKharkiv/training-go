package service

import (
	"training-go/sorting_structures/internal/domain/entity"
	"training-go/sorting_structures/internal/ports"
)

type SorterServiceImpl struct{}

func NewSorterService() ports.Sorter {
	return &SorterServiceImpl{}
}
func (s *SorterServiceImpl) Sort(people []entity.Person) []entity.Person {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	return people
}
