package main

// // getMaxTime 获取需要等待的最大时间
// func getMaxTime(takeWater []int16) int16 {
// 	var maxTime int16 = 0
// 	for i, e := range takeWater {
// 		if i == 0 || e > maxTime {
// 			maxTime = e
// 		}
// 	}
// 	return maxTime
// }

// // getMinTime 获取需要等待的最小时间
// func getMinTime(takeWater []int16) int16 {
// 	var minTime int16 = 0
// 	for i, e := range takeWater {
// 		if i == 0 || e < minTime {
// 			minTime = e
// 		}
// 	}
// 	return minTime
// }

// // monsterDrink 检测喝水需要的时间
// // people -> 前面的人数
// // pipe -> 水管
// // times -> 每个人需要的时间
// // 考虑了每个人接水的时间可能不同的情况
// func monsterDrink(people, pipe int16, times []int16) int16 {
// 	// 过滤掉可能多余的时间
// 	times = times[:people]
// 	// 如果水管的数量大于人的数量，则不需要等待就可以喝水
// 	if people < pipe {
// 		return 1
// 	}
// 	var maxTime int16 = 1
// 	var takeWater []int16
// 	for len(times) >= int(pipe) {
// 		// 每次获取满足水管数量的人数
// 		takeWater = times[:pipe]
// 		times = times[pipe:]

// 		maxTime += getMaxTime(takeWater)
// 	}

// 	// 返回需要等待的最大的时间
// 	return maxTime
// }

// monsterDrink 检测喝水需要的时间
// people -> 前面的人数
// pipe -> 水管
// times -> 每个人需要的时间
// 考虑了每个人接水的时间可能不同的情况
func monsterDrink(people, pipe int16, times []int16) int16 {
	// 过滤掉可能多余的时间
	times = times[:people]
	// 如果水管的数量大于人的数量，则可以直接接水，然后喝水
	if people < pipe {
		return 1
	}

	var time int16
	for i := range times {
		time += int16(i)
	}
	//
	return (pipe / time) + (pipe % time) + 1
}

func main() {
	monsterDrink(4, 2, []int16{1, 1, 1, 1})
}
