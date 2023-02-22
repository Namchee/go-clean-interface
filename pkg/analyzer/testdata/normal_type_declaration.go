package testdata

type NotAnInterface struct {
	aNumber int
}

type IntegerArray []int

func (ia IntegerArray) IsEmpty() bool {
	return len(ia) == 0
}
