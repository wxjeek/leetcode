package main

import (
	`fmt`
	`strconv`
)

/**

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21

 */

func reverse(x int) int {
	if x<0 {
		r := rever(strconv.Itoa(-x))
		i,_:=strconv.Atoi(r)
		if i>1<<31{
			return 0
		}
		return -i
	}else {
		r := rever(strconv.Itoa(x))
		i,_:=strconv.Atoi(r)
		if i>1<<31{
			return 0
		}
		return i
	}
}

func rever(s string) string {
	b := []byte(s)
	for i:=0;i<len(b)/2;i++{
		b[i],b[len(b)-1-i] = s[len(s)-1-i],s[i]
	}
	return string(b)
}

func main()  {
	i := -321
	fmt.Println(reverse(i))
}
