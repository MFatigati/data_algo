/*
P
	I:
	O:
E
D
A
*/

package main

func leftHalf(x float64) []int {
	midpoint := len(nums) / 2
	return nums[0:midpoint]
}

func rightHalf(x float64) []int {
	midpoint := len(nums) / 2
	return nums[midpoint:]
}

func combine() {
	// some custom function to process the results of the halved sections, possibly along with something about the current section, and pass that result back up to the next outer function
}

func myPow(x float64, n int) float64 {
	// base cases; very different for every problem

	// divide
	left := leftHalf(x)
	right := rightHalf(x)

	// conquer (recurse)
	leftResult := myPow(left)
	rightResult := myPow(right)

	// combine; very different for every problem
	return combine(leftResult, rightResult)
}

func main() {

}
