package testrepo0955

type MyInt int

const (
	A MyInt = iota
	B
	C
	D

	// another way

	E = MyInt(iota)
	F
	G
	H
)
