package ast

type Root struct {
	Tree interface{}
}

type Ident struct {
	Name string
}

type Import struct {
	Package *Ident
	From    Ident
	As      *Ident
}

type Number struct{}

type Float struct{}

type Op struct{}

type If struct{}

type Else struct{}

type Elseif struct{}

type Unless struct{}

type Whitespace struct{}

type EOF struct{}
