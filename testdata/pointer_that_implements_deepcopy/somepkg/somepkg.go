package somepkg

import (
	"github.com/slobodyanyuk/deep-copy/testdata/pointer_that_implements_deepcopy/anotherpkg"
)

type SomeStruct struct {
	AnotherStruct *anotherpkg.AnotherStruct
}
