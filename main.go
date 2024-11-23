package main

import (
	"data_str_and_algo/basics"
	"fmt"
)

func main() {
	bs := basics.BasicProblems{}
	var fiboSeries []int

	for i := 1; i <=20; i++{
		fiboSeries = append(fiboSeries, bs.FibonaciSeries(i))
	}
	// num := bs.FibonaciSeries(2) 
	fmt.Println(fiboSeries)

}
 
