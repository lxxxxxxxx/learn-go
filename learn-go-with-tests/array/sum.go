package array

//Sum : 输入一个切片，返回切片中所有int的和
func Sum(slice []int) (sum int) {
	for _, num := range slice {
		sum += num
	}
	return
}

//AllSum : 使用命名变量返回切片
func AllSum(numbersToSum ...[]int) (sums []int) {
	len := len(numbersToSum)
	sums = make([]int, len)

	for i, nums := range numbersToSum {
		sums[i] = Sum(nums)
	}
	return
}

//AllSum1 : 使用append函数在切片末尾添加元素
func AllSum1(numsToSum ...[]int) []int {
	var sums []int

	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}

//SumAllTial : 计算所有切片除第一个元素外的和
func SumAllTial(numsToSum ...[]int) []int {
	var sums []int

	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}
	return sums
}
