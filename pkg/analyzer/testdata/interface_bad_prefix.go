package testdata

type IBadPrefixedInterface interface { // want "Interface IBadPrefixedInterface has unnecessary prefix `I`"
	Foo(b int) int
}
