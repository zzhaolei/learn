package main

// arrayOperation 计算能被3整除的元素数量
// 元素对`3`取余，只会余：`1，2，0`，这三种可能。
// 余1的均是奇数，余2的均是不能被3整除的偶数。
// 遍历完元素之后，则直接对比余数为`1`和余数为`2`的情况
// `余1和余2`的原始`元素`相加，可以得一个被整除的数，此时会消耗两个数
func arrayOperation(num int16, list []int16) int16 {
	if len(list) == 0 {
		return 0
	}
	// result: 能被整除的数量
	// remainder1: 余1的数量
	// remainder2: 余2的数量
	var result, remainder1, remainder2 int16
	var remainder int16
	var i int16
	for i = 0; i < num; i++ {
		remainder = list[i] % 3
		if remainder == 0 {
			result += 1
		} else if remainder == 1 {
			// 如果余2的次数有至少一次，则直接两数相加获得一次可以被整除的数
			if remainder2-1 >= 0 {
				remainder2 -= 1
				result += 1
			} else {
				remainder1 += 1
			}
		} else if remainder == 2 {
			// 如果余1的次数有至少一次，则直接两数相加获得一次可以被整除的数
			if remainder1-1 >= 0 {
				remainder1 -= 1
				result += 1
			} else {
				remainder2 += 1
			}
		}
	}
	// `3`个奇数相加得一个可以被`3`整除的数
	if remainder1 != 0 && remainder1%3 == 0 {
		result += 1
	}
	// `2`个偶数相加得一个可以被`3`除后得1的数
	// 2 + 2 ==> 4 ==> 4 % 3 ==> 1
	if remainder2 != 0 && remainder2%3 == 1 {
		result += 1
	}
	return result
}

func main() {
	arrayOperation(2, []int16{})
	arrayOperation(5, []int16{})
	arrayOperation(5, []int16{3, 1, 2, 3, 1})
	arrayOperation(7, []int16{1, 1, 1, 1, 1, 2, 2})
}
