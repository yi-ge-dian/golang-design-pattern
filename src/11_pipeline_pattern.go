package main

import "fmt"

func CallPipelinePattern() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}
	result := pipeline(data)
	for res := range result {
		fmt.Println(res)
	}
}

func pipeline(data []int) <-chan int {
	squareChan := square(data)
	doubleChan := double(squareChan)
	return doubleChan
}

func square(data []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range data {
			out <- v * v
		}
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * 2
		}
	}()
	return out
}
