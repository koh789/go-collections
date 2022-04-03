package col

func Map[T, V any](elms []T, fn func(T) V) []V {
	outputs := make([]V, len(elms), cap(elms))
	for i, elm := range elms {
		outputs[i] = fn(elm)
	}
	return outputs
}

func MapWithIndex[T, V any](elms []T, fn func(int, T) V) []V {
	outputs := make([]V, len(elms), cap(elms))
	for i, elm := range elms {
		outputs[i] = fn(i, elm)
	}
	return outputs
}

func FlatMap[T, V any](elms []T, fn func(T) []V) []V {
	outputs := make([]V, 0)
	for _, elm := range elms {
		outputs = append(outputs, fn(elm)...)
	}
	return outputs
}

func MapE[T, V any](elms []T, fn func(T) (V, error)) ([]V, error) {

	outputs := make([]V, len(elms), cap(elms))
	for i, elm := range elms {
		o, err := fn(elm)
		if err != nil {
			return nil, err
		}
		outputs[i] = o
	}
	return outputs, nil
}

func MapWithIndexE[T, V any](elms []T, fn func(int, T) (V, error)) ([]V, error) {
	outputs := make([]V, len(elms), cap(elms))
	for i, elm := range elms {
		o, err := fn(i, elm)
		if err != nil {
			return nil, err
		}
		outputs[i] = o
	}
	return outputs, nil
}

func Filter[T any](elms []T, fn func(T) bool) []T {
	outputs := make([]T, 0)
	for _, elm := range elms {
		if fn(elm) {
			outputs = append(outputs, elm)
		}
	}
	return outputs
}

func Foreach[T any](elms []T, fn func(T)) {
	for _, elm := range elms {
		fn(elm)
	}
}

func Uniq[T comparable](elms []T) []T {
	outputs := make([]T, 0, len(elms))
	m := make(map[T]bool)
	for _, elm := range elms {
		if _, ok := m[elm]; !ok {
			m[elm] = true
			outputs = append(outputs, elm)
		}
	}
	return outputs
}

func GroupByUniq[T any, V comparable](elms []T, fn func(T) V) map[V]T {
	outputs := make(map[V]T, 0)
	for _, elm := range elms {
		key := fn(elm)
		if _, ok := outputs[key]; ok {
			continue
		}
		outputs[key] = elm
	}
	return outputs
}

func GroupBy[T any, V comparable](elms []T, fn func(T) V) map[V][]T {
	outputs := make(map[V][]T, 0)
	for _, elm := range elms {
		key := fn(elm)
		if values, ok := outputs[key]; ok {
			outputs[key] = append(values, elm)
			continue
		}
		outputs[key] = []T{elm}
	}
	return outputs
}

func Chunk[T any](elms []T, size int) [][]T {
	total := len(elms)
	if size <= 0 {
		return [][]T{elms}
	}
	if total <= 0 {
		return [][]T{}
	}
	if total <= size {
		return [][]T{elms}
	}
	end := size
	outputs := make([][]T, 0)
	for start := 0; start < total && end <= total; end += min(size, total-end) {
		outputs = append(outputs, elms[start:end])
		start += size
	}
	return outputs
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
