package array

//数字1～1000放在含有1001个元素的数组中，其中只有唯一的一个元素值重复，其他数字均只出现一次。设计一个算法，将重复元素找出来，要求每个数组元素只能访问一次。
//如果不使用辅助存储空间，能否设计一个算法实现？

// 累加求和法
// 存在的问题：累加数值过大会溢出
func FindDupByHash(arr []int) int {
	l := len(arr)
	suma := 0
	sumb := 0
	for _, v := range arr {
		suma += v
	}
	for i := 1; i < l; i++ {
		sumb += i
	}
	return suma-sumb
}
