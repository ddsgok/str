package str_test

import (
	"testing"

	"github.com/ddspog/bdd"
	"github.com/ddspog/str"
)

// Feature Chained creation.
// - As a developer,
// - I want to create Chained object with various properties,
// - So that I can use to print information in various ways.
func Test_Chained_creation(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Val interface{} `yaml:"val"`
	}{}, &struct {
		Out interface{} `yaml:"out"`
	}{}

	given(t, "a val that equal %[input.val]T(%[input.val]v)", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		val := input.Val

		when("s := str.New(val) is called", func(it bdd.It) {
			s := str.New(val)

			golden.Update(func() interface{} {
				gold.Out = s
				return gold
			})

			it("should have s equal to %[golden.out]q", func(assert bdd.Assert) {
				assert.Equal(gold.Out, s)
			})

			it("should have s.String() equal to %[golden.out]q", func(assert bdd.Assert) {
				assert.Equal(gold.Out, s.String())
			})

			it("should have s.Error() return an error with %[golden.out]q message", func(assert bdd.Assert) {
				err := s.Error()
				assert.Error(err)
				assert.Equal(gold.Out, err.Error())
			})
		})
	})
}
