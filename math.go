package github.com/AOEIUVBPMFDTNL/common

import (
	"math"
	"math/big"
)

// 阶乘
func Factorial(n int) (res int) {
	if n < 1 {
		res = 1
	} else {
		res = n
		for n > 1 {
			n--
			res *= n
		}
	}
	return
}

// 大数阶乘
func FactorialBig(n int64) (res *big.Int) {
	if n < 1 {
		res = big.NewInt(1)
	} else {
		res = big.NewInt(n)
		for n > 1 {
			n--
			res.Mul(res, big.NewInt(n))
		}
	}
	return
}

// 计算排列数
func PermutationCount(n int, m int, repeat bool) int {
	if repeat {
		return int(math.Pow(float64(n), float64(m)))
	}
	return Factorial(n) / Factorial(n-m)
}

// 计算组合数
func CombinationCount(n int, m int, repeat bool) int {
	if repeat {
		return Factorial(n+m-1) / (Factorial(n-1) * Factorial(m))
	}
	return Factorial(n) / (Factorial(n-m) * Factorial(m))
}

// 排列数解析
func PermutationParse[T comparable](arr []T, r int, repeat bool) (result [][]T) {
	n := len(arr)
	if repeat {
		t := PermutationCount(n, r, true)
		if t == 0 {
			return
		}
		for i := 0; i < t; i++ {
			v := make([]T, r)
			j := i
			for k := 0; k < r; k++ {
				x := j % n
				j = int(j / n)
				v[k] = arr[x]
			}
			result = append(result, v)
		}
	} else {
		if r > n {
			return
		}
		idxs := make([]int, n)
		for i := range idxs {
			idxs[i] = i
		}
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		cmb := make([]T, r)
		res := make([]T, r)
		for i, el := range idxs[:r] {
			cmb[i] = arr[el]
		}
		copy(res, cmb)
		result = append(result, res)
		for n > 0 {
			i := r - 1
			for ; i >= 0; i -= 1 {
				cycles[i] -= 1
				if cycles[i] == 0 {
					index := idxs[i]
					for j := i; j < n-1; j += 1 {
						idxs[j] = idxs[j+1]
					}
					idxs[n-1] = index
					cycles[i] = n - i
				} else {
					j := cycles[i]
					idxs[i], idxs[n-j] = idxs[n-j], idxs[i]
					for k := i; k < r; k += 1 {
						cmb[k] = arr[idxs[k]]
					}
					rc := make([]T, r)
					copy(rc, cmb)
					result = append(result, rc)
					break
				}
			}
			if i < 0 {
				return
			}
		}
	}
	return
}

// 组合数解析
func CombinationParse[T comparable](arr []T, r int, repeat bool) (result [][]T) {
	sendIndex := func(base []T, index []int) []T {
		res := make([]T, len(index))
		for i, idx := range index {
			res[i] = base[idx]
		}
		return res
	}
	n := len(arr)
	t := CombinationCount(n, r, repeat)
	if repeat {
		idxs := make([]int, r)
		result = append(result, sendIndex(arr, idxs))
		for i, j := 1, r-1; i < t; i++ {
			if idxs[j] == n-1 {
				for idxs[j] == n-1 {
					j--
				}
				v := idxs[j] + 1
				for i := j; i < r; i++ {
					idxs[i] = v
				}
				j = r - 1
			} else {
				idxs[j] = idxs[j] + 1
			}
			result = append(result, sendIndex(arr, idxs))
		}
	} else {
		if t == 0 {
			return
		}
		idxs := make([]int, r)
		for i := range idxs {
			idxs[i] = i
		}
		result = append(result, sendIndex(arr, idxs))
		for i, j := 1, r-1; i < t; i++ {
			if idxs[j] == j+n-r {
				for idxs[j] == j+n-r {
					j--
				}
				v := idxs[j] + 1
				for i := j; i < r; i++ {
					idxs[i] = v
					v++
				}
				j = r - 1
			} else {
				idxs[j] = idxs[j] + 1
			}
			result = append(result, sendIndex(arr, idxs))
		}
	}
	return
}

// 多切片组合计算
func CombinationManyCount[T comparable](arr [][]T, repeat bool) (total int) {
	// 循环次数
	var count = 1
	for i := range arr {
		count *= len(arr[i])
	}
	if repeat {
		total = count
		return
	}
	for i := 0; i < count; i++ {
		t := i
		a := make([]T, len(arr))
		for m := 0; m < len(arr); m++ {
			if t/len(arr) >= 0 {
				a[m] = arr[m][t%len(arr[m])]
				t /= len(arr[m])
			}
		}
		m := make(map[T]struct{})
		for v := range a {
			if _, ok := m[a[v]]; ok {
				break
			}
			m[a[v]] = struct{}{}
		}
		if len(m) == len(a) {
			total++
		}
	}
	return
}

// 多切片组合解析
func CombinationManyParse[T comparable](arr [][]T, repeat bool) (result [][]T) {
	// 循环次数
	var count = 1
	for i := range arr {
		count *= len(arr[i])
	}
	for i := 0; i < count; i++ {
		t := i
		a := make([]T, len(arr))
		for m := 0; m < len(arr); m++ {
			if t/len(arr) >= 0 {
				a[m] = arr[m][t%len(arr[m])]
				t /= len(arr[m])
			}
		}
		if repeat {
			result = append(result, a)
		} else {
			m := make(map[T]struct{})
			for v := range a {
				if _, ok := m[a[v]]; ok {
					break
				}
				m[a[v]] = struct{}{}
			}
			if len(m) == len(a) {
				result = append(result, a)
			}
		}
	}
	return
}
