package main

import (
	"errors"
	"fmt"

	"github.com/koh789/go-collections/pkg/col"
)

func main() {
	exampleMap()
	exampleMapE()
	exampleFlatMap()
	exampleMapWithIndex()
	exampleFilter()
	exampleUniq()
	exampleGroupBy()
	exampleGroupByUniq()
	exampleChunk()
}

func exampleMap() {
	inputs := []int{1, 2, 3}
	results := col.Map(inputs, func(num int) string {
		return fmt.Sprintf("result_%d", num)
	})
	fmt.Printf("map input slice  :%v\n", inputs)
	fmt.Printf("map results      :%v\n", results)
}

func exampleMapE() {
	errEven := errors.New("even error")
	inputs := []int{1, 3, 5, 6, 7, 11}
	results, err := col.MapE(inputs, func(input int) (string, error) {
		if input%2 == 0 {
			return "", errEven
		}
		return fmt.Sprintf("output_%d", input), nil
	})
	fmt.Printf("mapE input slice   :%v\n", inputs)
	fmt.Printf("mapE results       :%v\n", results)
	fmt.Printf("mapE error         :%v\n", err)
}

func exampleFlatMap() {
	type user struct {
		pets []string
	}
	inputs := []user{{pets: []string{"dog", "cat"}}, {pets: []string{"dog", "pig"}}, {pets: []string{"dog"}}}
	results := col.FlatMap(inputs, func(u user) []string { return u.pets })
	fmt.Printf("flatMap input slice   :%+v\n", inputs)
	fmt.Printf("flatMap results       :%v\n", results)
}

func exampleMapWithIndex() {
	type user struct {
		name string
	}
	users := []user{{name: "AAA"}, {name: "BBB"}}
	results := col.MapWithIndex(users, func(i int, u user) string { return fmt.Sprintf("%d_%s", i, u.name) })
	fmt.Printf("mapWithIndex input slice   :%+v\n", users)
	fmt.Printf("mapWithIndex results       :%v\n", results)
}

func exampleFilter() {
	inputs := []int{1, 2, 3, 4, 5}
	results := col.Filter(inputs, func(num int) bool {
		return num%2 == 0
	})
	fmt.Printf("filter input slice   :%+v\n", inputs)
	fmt.Printf("filter results       :%v\n", results)
}

func exampleUniq() {
	inputs := []string{"a", "b", "b", "c", "c", "c"}
	results := col.Uniq(inputs)
	fmt.Printf("filter input slice   :%+v\n", inputs)
	fmt.Printf("filter results       :%v\n", results)
}
func exampleGroupByUniq() {
	type user struct {
		ID   int
		Name string
	}
	inputs := []user{{ID: 1, Name: "one"}, {ID: 3, Name: "three"}, {ID: 10, Name: "ten"}}
	results := col.GroupBy(inputs, func(u user) int { return u.ID })
	fmt.Printf("groupByUniq input slice   :%+v\n", inputs)
	fmt.Printf("groupByUniq results       :%v\n", results)
}

func exampleGroupBy() {
	type user struct {
		ID      int
		Country string
	}
	inputs := []user{
		{ID: 1, Country: "JP"},
		{ID: 3, Country: "JP"},
		{ID: 10, Country: "US"},
		{ID: 50, Country: "JP"},
	}
	results := col.GroupBy(inputs, func(u user) string { return u.Country })
	fmt.Printf("groupBy input slice   :%+v\n", inputs)
	fmt.Printf("groupBy results       :%v\n", results)
}

func exampleChunk() {
	inputs := []string{"a", "b", "c", "d", "e", "f", "g"}
	results := col.Chunk(inputs, 3)
	fmt.Printf("chunk input slice   :%+v\n", inputs)
	fmt.Printf("chunk results       :%v\n", results)
}
