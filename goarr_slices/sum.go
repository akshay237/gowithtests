package goarr_slices

import "fmt"

// Arrays are of fixed length and slices are of dynamic length.
func Sum(num [5]int) int {
	total := 0
	for _, i := range num {
		total += i
	}
	return total
}

//Sum on slices
func SumSlice(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func SumAll(numsToSum ...[]int) []int {
	//length := len(numsToSum)
	//sum := make([]int, length)
	sum := []int{}

	// iterate over all the arrays can on each index call sum slice
	for _, nums := range numsToSum {
		//sum[i] = SumSlice(nums)
		sum = append(sum, SumSlice(nums))
	}
	return sum
}

// sum of slices by removing head
func SumAllTails(numsToSum ...[]int) []int {
	sum := []int{}
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, SumSlice(nums[1:]))
		}
	}
	return sum
}

//Compile time errors are our friend because they help us write software that works,
//runtime errors are our enemies because they affect our users
func main() {
	// call sum
	arr := [5]int{1, 3, 5, 7, 9}
	sum := Sum(arr)
	fmt.Println("Sum is : ", sum)

	// call sum slice
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum1 := SumSlice(arr1)
	fmt.Println("Sum of slice is: ", sum1)
}
