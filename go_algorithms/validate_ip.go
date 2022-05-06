/*
P
	I: a value of letters, digits, and ./:
	O: string "IPv6", "IPv4", or "Neither"
E
D
A
Attempt to split the string at periods
	if the result is length 4, check if valid IPv4
		For each element
		If valid, return "IPv4"
Attempt to split the string at colons
	if the result is length 8, check if valid IPv6
		For each element
		If valid, return "IPv4"
Return "Neither"

Greedy option:
	- Proceed through the strings, front to back, checking rules as you go, breaking when a rule is broken

*/

package main

func main() {

}
