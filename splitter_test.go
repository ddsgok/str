package str_test

import (
	"github.com/ddsgok/bdd"
	"github.com/ddsgok/str"
	"testing"
)

// Feature Splitter creation.
// - As a developer,
// - I want to create Splitter object with various properties,
// - So that I can use to print information in various ways.
func Test_Splitter_creation(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Set []string `yaml:"set"`
	}{}, &struct {
		Arr []string `yaml:"arr"`
		Str string   `yaml:"str"`
	}{}

	given(t, "a set that equal %[input.set]v", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		set := input.Set

		when("sa := str.With(set) is called", func(it bdd.It) {
			sa := str.With(set)

			golden.Update(func() interface{} {
				gold.Arr = sa.Array()
				gold.Str = sa.String()
				return gold
			})

			it("should have sa.Array() equal to %[golden.arr]q", func(assert bdd.Assert) {
				assert.Equal(gold.Arr, sa.Array())
			})

			it("should have sa.String() equal to %[golden.str]q", func(assert bdd.Assert) {
				assert.Equal(gold.Str, sa.String())
			})

			it("should have sa.Error() return an error with %[golden.str]q message", func(assert bdd.Assert) {
				err := sa.Error()
				assert.Error(err)
				assert.Equal(gold.Str, err.Error())
			})

			it("should have sa.Print() prints %[golden.str]q", func(assert bdd.Assert) {
				tw := newTestWriter()
				sa.Print(tw)
				assert.Equal(gold.Str, tw.String())
			})
		})
	})
}

// Feature Splitter Join.
// - As a developer,
// - I want to join Splitter object into a string,
// - So that I can use to other operations.
func Test_Splitter_Join(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Arr  []string `yaml:"arr"`
		Conn string   `yaml:"conn"`
	}{}, &struct {
		Str string `yaml:"str"`
	}{}

	given(t, "a set that equal %[input.arr]v", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		set := input.Arr

		when("sa := str.With(set) is called", func(it bdd.It) {
			sa := str.With(set)
			text := sa.Join(input.Conn).String()

			golden.Update(func() interface{} {
				gold.Str = text
				return gold
			})

			it("should have sa.Join(%[input.conn]q) equal to %[golden.str]q", func(assert bdd.Assert) {
				assert.Equal(gold.Str, text)
			})
		})
	})
}
