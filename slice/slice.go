package slice

type T interface{}
type mapFun func(t T) T
type consumeFun func(t T)
type reduceFun func(t1, t2 T) T
type filterFun func(t T) bool
type compareFun func(t1, t2 T) bool

func foreach(slice []T, fun consumeFun) {
	for _, temp := range slice {
		fun(temp)
	}
}

func reduce(slice []T, init T, fun reduceFun) T {

	result := init
	for i, size := 0, len(slice); i < size; i++ {
		result = fun(result, slice[i])
	}
	return result

}

func mapper(slice []T, fun mapFun) []T {
	mappers := make([]T, 0, len(slice))
	for _, temp := range slice {
		mappers = append(mappers, fun(temp))
	}
	return mappers
}

func filter(slice []T, fun filterFun) []T {
	ts := make([]T, 0)
	for _, temp := range slice {
		if fun(temp) {
			ts = append(ts, temp)
		}
	}
	return ts
}

func anyMatch(slice []T, fun filterFun) bool {
	for _, temp := range slice {
		if fun(temp) {
			return true
		}
	}
	return false
}

func allMatch(slice []T, fun filterFun) bool {
	result := true
	for _, temp := range slice {
		result = result && fun(temp)
	}
	return result
}

func first(slice []T, filter filterFun) T {
	for _, temp := range slice {
		if filter(temp) {
			return temp
		}
	}
	return nil

}

func last(slice []T, filter filterFun) T {
	for i := len(slice) - 1; i >= 0; i-- {
		if filter(slice[i]) {
			return slice[i]
		}
	}
	return nil
}

func sort(slice []T, compare compareFun) []T {
	for i := len(slice); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			if compare(slice[j-1], slice[j]) {
				intermediate := slice[j]
				slice[j] = slice[j-1]
				slice[j-1] = intermediate
			}
		}
	}
	return slice
}

func compare(slice []T, compare compareFun) T {
	result := slice[0]
	for _, temp := range slice {
		if compare(temp, result) {
			result = temp
		}
	}
	return result
}
