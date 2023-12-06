package main

import "strconv"

func s2i(s []string) []int {
	i := make([]int, len(s))

	for index, sval := range s {
		ival, err := strconv.Atoi(sval)
		if err != nil {
			panic(err)
		}
		i[index] = ival
	}

	return i
}

func rangelist2i(s []string) [][]int {
	rangelist := s2i(s)
	i := [][]int{}

	for index := 0; index < len(s); index += 2 {
		i = append(i, []int{rangelist[index], rangelist[index+1]})
	}

	return i
}
