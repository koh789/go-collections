package gol

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	for _, table := range []struct {
		title  string
		inputs []int
		expect []string
	}{
		{
			title:  "empty as an argument, empty is returned",
			inputs: []int{},
			expect: []string{},
		},
		{
			title:  "nil as an argument, empty is returned",
			inputs: nil,
			expect: []string{},
		},
		{
			title:  "if numeric slice as an argument, the result processed by the function is returned",
			inputs: []int{1, 4, 100},
			expect: []string{"output_1", "output_4", "output_100"},
		},
		{
			title:  "if numeric slice as an argument, the result processed by the function is returned",
			inputs: []int{1, 4, 100},
			expect: []string{"output_1", "output_4", "output_100"},
		},
	} {
		t.Run(table.title, func(t *testing.T) {
			actual := Map(table.inputs, func(input int) string { return fmt.Sprintf("output_%d", input) })
			assert.ElementsMatch(t, table.expect, actual)
		})
	}
}

type Dummy struct {
	Nums []int
}

func TestFlatMap(t *testing.T) {
	for _, table := range []struct {
		title  string
		inputs []Dummy
		expect []int
	}{
		{
			title:  "emptyを渡すと,empty返却",
			inputs: []Dummy{},
			expect: []int{},
		},
		{
			title:  "nilを渡すと,empty返却",
			inputs: nil,
			expect: []int{},
		},
		{
			title: "sliceを持つstructのsliceを渡すと, flatten処理した結果が返却",
			inputs: []Dummy{
				{Nums: []int{1, 2, 3}},
				{Nums: []int{1, 2, 3}},
				{Nums: []int{3, 4, 5, 6}},
			},
			expect: []int{1, 2, 3, 1, 2, 3, 3, 4, 5, 6},
		},
		{
			title: "sliceを持つstructを, flatten処理した結果が返却",
			inputs: []Dummy{
				{Nums: []int{1, 2, 3}},
			},
			expect: []int{1, 2, 3},
		},
	} {
		t.Run(table.title, func(t *testing.T) {
			actual := FlatMap(table.inputs, func(d Dummy) []int { return d.Nums })
			assert.ElementsMatch(t, table.expect, actual)
		})
	}
}

func TestMapE(t *testing.T) {
	errEven := errors.New("even error")
	fn := func(input int) (string, error) {
		if input%2 == 0 {
			return "", errEven
		}
		return fmt.Sprintf("output_%d", input), nil
	}

	for _, table := range []struct {
		title       string
		inputs      []int
		expect      []string
		errorsOccur bool
		err         error
	}{
		{
			title:  "empty as an argument, empty is returned",
			inputs: []int{},
			expect: []string{},
		},
		{
			title:  "nil as an argument, empty is returned",
			inputs: nil,
			expect: []string{},
		},
		{
			title:  "if odd numeric slice as an argument, the result processed by the function is returned(no err)",
			inputs: []int{1, 5, 7, 13},
			expect: []string{"output_1", "output_5", "output_7", "output_13"},
		},
		{
			title:       "if odd numeric slice as an argument, err returned",
			inputs:      []int{1, 5, 100},
			errorsOccur: true,
			err:         errEven,
			expect:      []string{},
		},
	} {
		t.Run(table.title, func(t *testing.T) {

			actual, err := MapE(table.inputs, fn)
			if table.errorsOccur {
				assert.NotNil(t, err)
				assert.Equal(t, table.err, err)
			}
			assert.ElementsMatch(t, table.expect, actual)
		})
	}
}

func TestMapWithIndexE(t *testing.T) {
	errEven := errors.New("even error")
	fn := func(index, input int) (string, error) {
		if input%2 == 0 {
			return "", errEven
		}
		return fmt.Sprintf("i_%d_output_%d", index, input), nil
	}

	for _, table := range []struct {
		title       string
		inputs      []int
		expect      []string
		errorsOccur bool
		err         error
	}{
		{
			title:  "empty as an argument, empty is returned",
			inputs: []int{},
			expect: []string{},
		},
		{
			title:  "nil as an argument, empty is returned",
			inputs: nil,
			expect: []string{},
		},
		{
			title:  "if odd numeric slice as an argument, the result processed by the function is returned(no err)",
			inputs: []int{1, 5, 7, 13},
			expect: []string{"i_0_output_1", "i_1_output_5", "i_2_output_7", "i_3_output_13"},
		},
		{
			title:       "if odd numeric slice as an argument, err returned",
			inputs:      []int{1, 5, 100},
			errorsOccur: true,
			err:         errEven,
			expect:      []string{},
		},
	} {
		t.Run(table.title, func(t *testing.T) {

			actual, err := MapWithIndexE(table.inputs, fn)
			if table.errorsOccur {
				assert.NotNil(t, err)
				assert.Equal(t, table.err, err)
			}
			assert.ElementsMatch(t, table.expect, actual)
		})
	}
}

func TestUniq(t *testing.T) {
	for _, table := range []struct {
		title  string
		inputs []string
		expect []string
	}{
		{
			title:  "empty as an argument, empty is returned",
			inputs: []string{},
			expect: []string{},
		},
		{
			title:  "nil as an argument, empty is returned",
			inputs: nil,
			expect: []string{},
		},
		{
			title:  "if a slice containing duplicates is passed as an argument, it is de-duplicated and returned",
			inputs: []string{"a", "b", "c", "c", "b", "aa", "aa", "b"},
			expect: []string{"a", "b", "c", "aa"},
		},
		{
			title:  "If a slice that does not contain duplicates is passed as an argument, it is returned as is",
			inputs: []string{"a", "aa", "b", "c"},
			expect: []string{"a", "aa", "b", "c"},
		},
	} {
		t.Run(table.title, func(t *testing.T) {
			actual := Uniq(table.inputs)
			assert.ElementsMatch(t, table.expect, actual)
		})
	}
}

func TestGroupByUniq(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}

	for _, table := range []struct {
		title  string
		inputs []user
		expect map[int]user
	}{
		{
			title:  "empty as an argument, empty is returned",
			inputs: []user{},
			expect: map[int]user{},
		},
		{
			title:  "nil as an argument, empty is returned",
			inputs: nil,
			expect: map[int]user{},
		},
		{
			title: "if a slice containing duplicates is passed as an argument, it is de-duplicated and returned",
			inputs: []user{
				{ID: 1, Name: "one"},
				{ID: 3, Name: "three"},
				{ID: 10, Name: "ten"},
				{ID: 50, Name: "fifty"},
			},
			expect: map[int]user{
				1:  {ID: 1, Name: "one"},
				3:  {ID: 3, Name: "three"},
				10: {ID: 10, Name: "ten"},
				50: {ID: 50, Name: "fifty"},
			},
		},
		{
			title: "If a slice that does not contain duplicates is passed as an argument, it is returned as is",
			inputs: []user{
				{ID: 1, Name: "one"},
				{ID: 3, Name: "three"},
				{ID: 3, Name: "three_duplicated"},
				{ID: 10, Name: "ten"},
				{ID: 10, Name: "ten_duplicated"},
				{ID: 50, Name: "fifty"},
			},
			expect: map[int]user{
				1:  {ID: 1, Name: "one"},
				3:  {ID: 3, Name: "three"},
				10: {ID: 10, Name: "ten"},
				50: {ID: 50, Name: "fifty"},
			},
		},
	} {
		t.Run(table.title, func(t *testing.T) {
			actual := GroupByUniq(table.inputs, func(u user) int { return u.ID })
			assert.Equal(t, table.expect, actual)
		})
	}
}

func TestChunk(t *testing.T) {
	for _, table := range []struct {
		title  string
		inputs []string
		size   int
		expect [][]string
	}{
		{
			title:  "7の要素に対してsize:3でchunk,chunkされた要素が3つ返却",
			inputs: []string{"a", "b", "c", "d", "e", "f", "g"},
			size:   3,
			expect: [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g"}},
		},
		{
			title:  "9の要素に対してsize:3でchunk,chunkされた要素が3つ返却",
			inputs: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
			size:   3,
			expect: [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}},
		},
		{
			title:  "7の要素に対してsize:8でchunk,chunkされた要素が1つ返却",
			inputs: []string{"a", "b", "c", "d", "e", "f", "g"},
			size:   8,
			expect: [][]string{{"a", "b", "c", "d", "e", "f", "g"}},
		},
		{
			title:  "7の要素に対してsize:1でchunk,chunkされた要素が7つ返却",
			inputs: []string{"a", "b", "c", "d", "e", "f", "g"},
			size:   1,
			expect: [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}, {"g"}},
		},
		{
			title:  "7の要素に対してsize:0でchunk,元の要素が二次元sliceとなり返却",
			inputs: []string{"a", "b", "c", "d", "e", "f", "g"},
			size:   0,
			expect: [][]string{{"a", "b", "c", "d", "e", "f", "g"}},
		},
		{
			title:  "0の要素に対してsize:3でchunk,empty返却",
			inputs: []string{},
			size:   3,
			expect: [][]string{},
		},
	} {
		t.Run(table.title, func(t *testing.T) {
			actual := Chunk(table.inputs, table.size)
			assert.ElementsMatch(t, table.expect, actual)
		})
	}
}
