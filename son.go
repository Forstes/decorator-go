package main

type Son struct {
	name     string
	ancestor Ancestor
}

func (s *Son) getAncestry() string {
	return s.name + "-" + s.ancestor.getAncestry()
}
