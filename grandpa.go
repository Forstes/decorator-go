package main

type Grandpa struct {
	name string
}

func (s *Grandpa) getAncestry() string {
	return s.name
}
