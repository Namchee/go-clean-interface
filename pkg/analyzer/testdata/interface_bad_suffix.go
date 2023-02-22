package testdata

type BadSuffixedInterfaceItf interface { // want "Interface BadSuffixedInterfaceItf has unnecessary suffix `Itf`"
	Bar(a int) int
}
