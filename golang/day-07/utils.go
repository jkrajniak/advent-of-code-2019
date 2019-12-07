package main

import (
	"strconv"
	"strings"
)

func convertToIntSlice(s []string) ([]int, error) {
	var ints []int
	for _, s := range s {
		i, err := strconv.Atoi(strings.Trim(s, "\n"))
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func getPermutations(seqToPermute []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permute(c, seqToPermute)
	}(c)
	return c
}
func permute(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp

		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output

		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}
