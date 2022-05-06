/*
P
	I:
	O:
E

[1, 3, 5, 10, 11, 12, 13, 14], 9
D
A

If the value at the midpoint is the target value
	return midpoint index
If the value at the midpoint is less than the target value
	If the value to the right of the midpoint is greater than the target value
		return midpoint index plus 1
	Else, move lower bound to midpoint plus 1
If the value at the midpoint is greater than the target value
	If the value to the left of the midpoint is less than the target value
		return midpoint index minus 1
	Else, move upper bound to midpoint minus 1

*/

package main

func main() {

}
