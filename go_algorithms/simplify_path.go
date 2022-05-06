package main

import (
	"fmt"
	"regexp"
)

func simplifyPath(path string) string {
	temp := []string{}

	re := regexp.MustCompile("/+")
	splitPath := re.Split(path, -1)

	for _, elem := range splitPath {
		if elem == ".." {
			temp = temp[:len(temp)-1]
		} else if elem == "." {
			continue
		} else {
			temp = append(temp, elem)
		}
	}

	returnString := ""

	fmt.Println(temp)
	for _, elem := range temp {
		if elem == " " || elem == "" {
			continue
		} else {
			returnString = returnString + "/" + elem
		}
	}

	if len(returnString) == 0 {
		return "/"
	}

	return returnString
}

func main() {

	// fmt.Println(simplifyPath("/home//foo/../baz")) // "/home/baz"
	// fmt.Println(simplifyPath("/home/"))            // "/home"
	// fmt.Println(simplifyPath("/../"))              // "/"
	fmt.Println(simplifyPath("/home/../../..")) // "/"
}
