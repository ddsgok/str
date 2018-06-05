package str_test

import (
	"github.com/ddspog/bdd"
	"github.com/ddspog/str"
	"testing"
)

// Feature Collecter creation.
// - As a developer,
// - I want to create Splitter object with various properties,
// - So that I can use to print information in various ways.
func Test_Collecter_creation(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Set interface{} `yaml:"set"`
	}{}, &struct {
		Arr     [][]string      `yaml:"arr"`
		AnomArr [][]interface{} `yaml:"anom_arr"`
		Str     string          `yaml:"str"`
	}{}

	given(t, "a set that equal %[input.set]v", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		set := input.Set

		when("sa := str.Collect(set) is called", func(it bdd.It) {
			sa := str.Collect(set)

			golden.Update(func() interface{} {
				gold.Arr = sa.Array()
				gold.AnomArr = sa.AnomArray()
				gold.Str = sa.String()
				return gold
			})

			it("should have sa.Array() equal to %[golden.arr]q", func(assert bdd.Assert) {
				assert.Equal(gold.Arr, sa.Array())
			})

			it("should have sa.AnomArray() equal to %[golden.anom_arr]q", func(assert bdd.Assert) {
				assert.Equal(gold.AnomArr, sa.AnomArray())
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

// Feature Collecter FmtAll.
// - As a developer,
// - I want to fmt elements on Collecter to a splitter,
// - So that I can use to other operations.
func Test_Collecter_FmtAll(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Arr    [][]string    `yaml:"arr"`
		Format string        `yaml:"format"`
		Args   []interface{} `yaml:"args"`
	}{}, &struct {
		Arr []string `yaml:"arr"`
		Str string   `yaml:"str"`
	}{}

	given(t, "a set that equal %[input.arr]v, a format string %[input.format]q", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		set, format := input.Arr, input.Format

		when("split := str.Collect(set).FmtAll(format) is called", func(it bdd.It) {
			split := str.Collect(set).FmtAll(format)

			golden.Update(func() interface{} {
				gold.Arr = split.Array()
				gold.Str = split.String()
				return gold
			})

			it("should have split.Array() equal to %[golden.arr]v", func(assert bdd.Assert) {
				assert.Equal(gold.Arr, split.Array())
			})

			it("should have split.String() equal to %[golden.str]v", func(assert bdd.Assert) {
				assert.Equal(gold.Str, split.String())
			})
		})
	})

	given(t, "a set that equal %[input.arr]v, a format string %[input.format]q and args equal to %[input.args]v", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		set, format, args := input.Arr, input.Format, input.Args

		when("split := str.Collect(set).FmtAll(format, args...) is called", func(it bdd.It) {
			split := str.Collect(set).FmtAll(format, args...)

			golden.Update(func() interface{} {
				gold.Arr = split.Array()
				gold.Str = split.String()
				return gold
			})

			it("should have split.Array() equal to %[golden.arr]v", func(assert bdd.Assert) {
				assert.Equal(gold.Arr, split.Array())
			})

			it("should have split.String() equal to %[golden.str]v", func(assert bdd.Assert) {
				assert.Equal(gold.Str, split.String())
			})
		})
	})
}

// Feature Collecter JoinAll.
// - As a developer,
// - I want to join Collecter object into a splitter,
// - So that I can use to other operations.
func Test_Collecter_JoinAll(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Arr  [][]string `yaml:"arr"`
		Conn string     `yaml:"conn"`
	}{}, &struct {
		Arr []string `yaml:"arr"`
		Str string   `yaml:"str"`
	}{}

	given(t, "a set that equal %[input.arr]v", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		set := input.Arr

		when("split := str.Collect(set).JoinAll(%[input.conn]q) is called", func(it bdd.It) {
			split := str.Collect(set).JoinAll(input.Conn)

			golden.Update(func() interface{} {
				gold.Arr = split.Array()
				gold.Str = split.String()
				return gold
			})

			it("should have split.Array() equal to %[golden.arr]v", func(assert bdd.Assert) {
				assert.Equal(gold.Arr, split.Array())
			})

			it("should have split.String() equal to %[golden.str]v", func(assert bdd.Assert) {
				assert.Equal(gold.Str, split.String())
			})
		})
	})
}
