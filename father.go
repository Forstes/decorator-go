package main

type Father struct {
	name     string
	ancestor Ancestor
}

func (s *Father) getAncestry() string {
	return s.name + "-" + s.ancestor.getAncestry()
}
