# go-collections

[![BuildStatus](https://github.com/koh789/go-collections/actions/workflows/test.yml/badge.svg)](https://github.com/koh789/go-collections/actions/workflows/test.yml)


functional library for slice

## Install

```
go get -u github.com/koh789/go-collections
```

## Requirements

go collections library requires Go version `>=1.18`


## Usage

### Map

```go:サンプルコード
func exampleMap() {
	inputs := []int{1, 2, 3}
	results := col.Map(inputs, func(num int) string {
		return fmt.Sprintf("result_%d", num)
	})
	fmt.Printf("map input slice  :%v\n", inputs)
	fmt.Printf("map results      :%v\n", results)
}
//map input slice  :[1 2 3]
//map results      :[result_1 result_2 result_3]
```
### MapE

```go:サンプルコード
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
//mapE input slice   :[1 3 5 6 7 11]
//mapE results       :[]
//mapE error         :even error
```

### FlatMap

```go:サンプルコード
func exampleFlatMap() {
	type user struct {
		pets []string
	}
	inputs := []user{{pets: []string{"dog", "cat"}}, {pets: []string{"dog", "pig"}}, {pets: []string{"dog"}}}
	results := col.FlatMap(inputs, func(u user) []string { return u.pets })
	fmt.Printf("flatMap input slice   :%+v\n", inputs)
	fmt.Printf("flatMap results       :%v\n", results)
}
//flatMap input slice   :[{pets:[dog cat]} {pets:[dog pig]} {pets:[dog]}]
//flatMap results       :[dog cat dog pig dog]
```



### MapWithIndex

```go:サンプルコード
func exampleMapWithIndex() {
	type user struct {
		name string
	}
	users := []user{{name: "AAA"}, {name: "BBB"}}
	results := col.MapWithIndex(users, func(i int, u user) string { return fmt.Sprintf("%d_%s", i, u.name) })
	fmt.Printf("mapWithIndex input slice   :%+v\n", users)
	fmt.Printf("mapWithIndex results       :%v\n", results)
}
//mapWithIndex input slice   :[{name:AAA} {name:BBB}]
//mapWithIndex results       :[0_AAA 1_BBB]
```

### Filter

```go:サンプルコード
func exampleFilter() {
	inputs := []int{1, 2, 3, 4, 5}
	results := col.Filter(inputs, func(num int) bool {
		return num%2 == 0
	})
	fmt.Printf("filter input slice   :%+v\n", inputs)
	fmt.Printf("filter results       :%v\n", results)
}
//filter input slice   :[1 2 3 4 5]
//filter results       :[2 4]

```

### Uniq

```go:サンプルコード
func exampleUniq() {
	inputs := []string{"a", "b", "b", "c", "c", "c"}
	results := col.Uniq(inputs)
	fmt.Printf("filter input slice   :%+v\n", inputs)
	fmt.Printf("filter results       :%v\n", results)
}
//filter input slice   :[a b b c c c]
//filter results       :[a b c]
```


### GroupByUniq


```go:サンプルコード
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

//groupByUniq input slice   :[{ID:1 Name:one} {ID:3 Name:three} {ID:10 Name:ten}]
//groupByUniq results       :map[1:[{1 one}] 3:[{3 three}] 10:[{10 ten}]]

```

### GroupBy
```go:サンプルコード
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
//groupBy input slice   :[{ID:1 Country:JP} {ID:3 Country:JP} {ID:10 Country:US} {ID:50 Country:JP}]
//groupBy results       :map[JP:[{1 JP} {3 JP} {50 JP}] US:[{10 US}]]

```

### Chunk

```go:サンプルコード
func exampleChunk() {
	inputs := []string{"a", "b", "c", "d", "e", "f", "g"}
	results := col.Chunk(inputs, 3)
	fmt.Printf("chunk input slice   :%+v\n", inputs)
	fmt.Printf("chunk results       :%v\n", results)
}
//chunk input slice   :[a b c d e f g]
//chunk results       :[[a b c] [d e f] [g]]

```
