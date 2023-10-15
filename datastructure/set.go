package datastructure

type Set struct {
	items map[any]struct{}
}

func MakeSet() *Set {
	return &Set{
		items: make(map[any]struct{}),
	}
}

func (s *Set) Add(item any) {
	s.items[item] = struct{}{}
}

func (s *Set) Remove(item any) {
	_, exists := s.items[item]
	if !exists {
		return
	}

	delete(s.items, item)
}

func (s *Set) Contains(item any) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) IsEmpty() bool {
	return len(s.items) == 0
}
