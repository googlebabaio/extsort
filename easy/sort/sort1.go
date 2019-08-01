package main

import "sort"

func main() {
	a:=[]int{9,2,4,5,3,8,3,2,1}
	sort.Ints(a)
	for _,v:=range a{
		println(v)
	}
}
