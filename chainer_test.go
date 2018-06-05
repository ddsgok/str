package str_test

import (
	"testing"

	"github.com/ddspog/bdd"
	"github.com/ddspog/str"
)

// Feature Chainer creation.
// - As a developer,
// - I want to create Chainer object with various properties,
// - So that I can use to print information in various ways.
func Test_Chainer_creation(t *testing.T) {
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

			it("should have s.Print() prints %[golden.out]q", func(assert bdd.Assert) {
				tw := newTestWriter()
				s.Print(tw)
				assert.Equal(gold.Out, tw.String())
			})
		})
	})

	input2, gold2 := &struct {
		Formatter string        `yaml:"fmt"`
		Args      []interface{} `yaml:"args"`
	}{}, &struct {
		Out string `yaml:"out"`
	}{}

	given(t, "a format that equal %[input.fmt]q and args equal to %[input.args]v", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input2, gold2)
		format, args := input2.Formatter, input2.Args

		when("s := str.New(format, args...).String() is called", func(it bdd.It) {
			s := str.New(format, args...).String()

			golden.Update(func() interface{} {
				gold2.Out = s
				return gold2
			})

			it("should have s equal to %[golden.out]q", func(assert bdd.Assert) {
				assert.Equal(gold2.Out, s)
			})
		})
	})
}

// Feature Chainer Split.
// - As a developer,
// - I want to split Chainer object into a string array,
// - So that I can use to other operations.
func Test_Chainer_Split(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Text string `yaml:"text"`
		Sep  string `yaml:"sep"`
	}{}, &struct {
		Arr     []string      `yaml:"arr"`
		AnomArr []interface{} `yaml:"anom_arr"`
	}{}

	given(t, "a text that equal %[input.text]q", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)
		text := input.Text

		when("sa := str.New(text).Split(%[input.sep]q) is called", func(it bdd.It) {
			sa := str.New(text).Split(input.Sep)

			arr := sa.Array()
			anomArr := sa.AnomArray()

			golden.Update(func() interface{} {
				gold.Arr = arr
				gold.AnomArr = anomArr
				return gold
			})

			it("should have sa.Array() equal to %[golden.arr]q", func(assert bdd.Assert) {
				assert.Equal(gold.Arr, arr)
			})

			it("should have s.AnomArray() equal to %[golden.anom_arr]q", func(assert bdd.Assert) {
				assert.Equal(gold.AnomArr, anomArr)
			})
		})
	})
}
