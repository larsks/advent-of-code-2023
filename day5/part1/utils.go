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
